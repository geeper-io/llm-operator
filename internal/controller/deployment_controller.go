/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"reflect"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/cluster-api/util/patch"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

const (
	// FinalizerName is the name of the finalizer
	FinalizerName = "llm.geeper.io/finalizer"
)

// LMDeploymentReconciler reconciles a LMDeployment object
type LMDeploymentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=llm.geeper.io,resources=lmdeployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=llm.geeper.io,resources=lmdeployments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=llm.geeper.io,resources=lmdeployments/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=configmaps,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=secrets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=persistentvolumes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=persistentvolumeclaims,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *LMDeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("Reconciling Deployment")

	// Fetch the Deployment instance
	deployment := &llmgeeperiov1alpha1.LMDeployment{}
	err := r.Get(ctx, req.NamespacedName, deployment)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	// Check if the resource is being deleted
	if !deployment.DeletionTimestamp.IsZero() {
		// Resource is being deleted, handle finalization
		if err := r.finalizeDeployment(ctx, deployment); err != nil {
			return ctrl.Result{}, fmt.Errorf("failed to finalize deployment: %w", err)
		}
		return ctrl.Result{}, nil
	}

	// Add finalizer if it doesn't exist
	if !containsFinalizer(deployment.Finalizers, FinalizerName) {
		deployment.Finalizers = append(deployment.Finalizers, FinalizerName)
		if err := r.Update(ctx, deployment); err != nil {
			return ctrl.Result{}, fmt.Errorf("failed to add finalizer: %w", err)
		}
		return ctrl.Result{}, nil
	}

	// Reconcile model serving deployment (Ollama or vLLM)
	if deployment.Spec.VLLM.Enabled {
		// Reconcile vLLM deployment
		if err := r.reconcileVLLM(ctx, deployment); err != nil {
			return ctrl.Result{}, fmt.Errorf("failed to reconcile vLLM: %w", err)
		}
	}

	if deployment.Spec.Ollama.Enabled {
		// Reconcile Ollama deployment (default)
		if err := r.reconcileOllama(ctx, deployment); err != nil {
			return ctrl.Result{}, fmt.Errorf("failed to reconcile Ollama: %w", err)
		}
	}

	if deployment.Spec.OpenWebUI.Enabled {
		if err := r.reconcileOpenWebUI(ctx, deployment); err != nil {
			return ctrl.Result{}, fmt.Errorf("failed to reconcile OpenWebUI: %w", err)
		}
	}

	if deployment.Spec.Tabby.Enabled {
		if err := r.reconcileTabby(ctx, deployment); err != nil {
			return ctrl.Result{}, fmt.Errorf("failed to reconcile Tabby: %w", err)
		}
	}

	// Update status
	if err := r.updateStatus(ctx, deployment); err != nil {
		return ctrl.Result{}, fmt.Errorf("failed to update deployment status: %w", err)
	}

	// Only requeue if there are actual changes that need monitoring
	// If everything is stable, don't requeue unnecessarily
	return ctrl.Result{}, nil
}

// containsFinalizer checks if a slice contains a specific finalizer
func containsFinalizer(finalizers []string, finalizer string) bool {
	for _, f := range finalizers {
		if f == finalizer {
			return true
		}
	}
	return false
}

// removeFinalizer removes a finalizer from a slice
func removeFinalizer(finalizers []string, finalizer string) []string {
	var result []string
	for _, f := range finalizers {
		if f != finalizer {
			result = append(result, f)
		}
	}
	return result
}

// finalizeDeployment handles the finalization of a deployment
func (r *LMDeploymentReconciler) finalizeDeployment(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
	logger := log.FromContext(ctx)
	logger.Info("Finalizing deployment", "name", deployment.Name)

	// Remove finalizer
	deployment.Finalizers = removeFinalizer(deployment.Finalizers, FinalizerName)
	if err := r.Update(ctx, deployment); err != nil {
		return err
	}

	logger.Info("Finalization completed", "name", deployment.Name)
	return nil
}

// buildResourceRequirements builds resource requirements from the spec
func (r *LMDeploymentReconciler) buildResourceRequirements(resources llmgeeperiov1alpha1.ResourceRequirements) corev1.ResourceRequirements {
	return corev1.ResourceRequirements{
		Requests: resources.Requests,
		Limits:   resources.Limits,
	}
}

// createOrUpdateDeployment creates or updates a deployment using patch helper to avoid unnecessary reconciliations
func (r *LMDeploymentReconciler) createOrUpdateDeployment(ctx context.Context, deployment *appsv1.Deployment) error {
	existing := &appsv1.Deployment{}
	err := r.Get(ctx, types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, existing)
	if err != nil && errors.IsNotFound(err) {
		// Create new deployment
		if err := r.Create(ctx, deployment); err != nil {
			return err
		}
	} else if err == nil {
		// Update existing deployment using patch helper
		if !reflect.DeepEqual(existing.Spec, deployment.Spec) {
			patchHelper, err := patch.NewHelper(existing, r.Client)
			if err != nil {
				return fmt.Errorf("failed to create patch helper for deployment %s: %w", deployment.Name, err)
			}

			existing.Spec = deployment.Spec
			if err := patchHelper.Patch(ctx, existing); err != nil {
				return fmt.Errorf("failed to patch deployment %s: %w", deployment.Name, err)
			}
		}
	} else {
		return err
	}
	return nil
}

// createOrUpdateService creates or updates a service using patch helper to avoid unnecessary reconciliations
func (r *LMDeploymentReconciler) createOrUpdateService(ctx context.Context, service *corev1.Service) error {
	existing := &corev1.Service{}
	err := r.Get(ctx, types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, existing)
	if err != nil && errors.IsNotFound(err) {
		// Create new service
		if err := r.Create(ctx, service); err != nil {
			return err
		}
	} else if err == nil {
		// Update existing service using patch helper
		if !reflect.DeepEqual(existing.Spec, service.Spec) {
			patchHelper, err := patch.NewHelper(existing, r.Client)
			if err != nil {
				return fmt.Errorf("failed to create patch helper for service %s: %w", service.Name, err)
			}

			existing.Spec = service.Spec
			if err := patchHelper.Patch(ctx, existing); err != nil {
				return fmt.Errorf("failed to patch service %s: %w", service.Name, err)
			}
		}
	} else {
		return err
	}
	return nil
}

// createOrUpdateIngress creates or updates an ingress using patch helper to avoid unnecessary reconciliations
func (r *LMDeploymentReconciler) createOrUpdateIngress(ctx context.Context, ingress *networkingv1.Ingress) error {
	existing := &networkingv1.Ingress{}
	err := r.Get(ctx, types.NamespacedName{Name: ingress.Name, Namespace: ingress.Namespace}, existing)
	if err != nil && errors.IsNotFound(err) {
		// Create new ingress
		if err := r.Create(ctx, ingress); err != nil {
			return err
		}
	} else if err == nil {
		// Update existing ingress using patch helper
		if !reflect.DeepEqual(existing.Spec, ingress.Spec) {
			patchHelper, err := patch.NewHelper(existing, r.Client)
			if err != nil {
				return fmt.Errorf("failed to create patch helper for ingress %s: %w", ingress.Name, err)
			}

			existing.Spec = ingress.Spec
			if err := patchHelper.Patch(ctx, existing); err != nil {
				return fmt.Errorf("failed to patch ingress %s: %w", ingress.Name, err)
			}
		}
	} else {
		return err
	}
	return nil
}

// updateStatus updates the status of the Deployment using patch helper to avoid unnecessary reconciliations
func (r *LMDeploymentReconciler) updateStatus(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
	// Create patch helper before making any changes
	patchHelper, err := patch.NewHelper(deployment, r.Client)
	if err != nil {
		return fmt.Errorf("failed to create patch helper: %w", err)
	}

	// Get model serving deployment status (Ollama or vLLM)
	if deployment.Spec.VLLM.Enabled {
		// Get vLLM model deployment statuses
		var totalVLLMReplicas int32
		var totalVLLMReadyReplicas int32
		var totalVLLMAvailableReplicas int32
		var totalVLLMUpdatedReplicas int32

		for _, modelSpec := range deployment.Spec.VLLM.Models {
			replicas := modelSpec.Replicas
			if replicas == 0 {
				replicas = 1
			}
			totalVLLMReplicas += replicas

			// Get individual model deployment status
			vllmDeployment := &appsv1.Deployment{}
			err = r.Get(ctx, types.NamespacedName{
				Name:      deployment.GetVLLMModelDeploymentName(modelSpec.Name),
				Namespace: deployment.Namespace,
			}, vllmDeployment)

			if err == nil {
				totalVLLMReadyReplicas += vllmDeployment.Status.ReadyReplicas
				totalVLLMAvailableReplicas += vllmDeployment.Status.AvailableReplicas
				totalVLLMUpdatedReplicas += vllmDeployment.Status.UpdatedReplicas
			}
		}

		// Update vLLM status
		deployment.Status.VLLMStatus.ReadyReplicas = totalVLLMReadyReplicas
		deployment.Status.VLLMStatus.AvailableReplicas = totalVLLMAvailableReplicas
		deployment.Status.VLLMStatus.UpdatedReplicas = totalVLLMUpdatedReplicas

		routerReplicas := deployment.Spec.VLLM.Router.Replicas
		if routerReplicas == 0 {
			routerReplicas = 1
		}
		totalVLLMReplicas += routerReplicas

		deployment.Status.TotalReplicas = totalVLLMReplicas
		deployment.Status.ReadyReplicas = totalVLLMReadyReplicas
	} else {
		// Get Ollama deployment status
		ollamaDeployment := &appsv1.Deployment{}
		err = r.Get(ctx, types.NamespacedName{
			Name:      deployment.GetOllamaDeploymentName(),
			Namespace: deployment.Namespace,
		}, ollamaDeployment)

		if err == nil {
			deployment.Status.OllamaStatus.AvailableReplicas = ollamaDeployment.Status.AvailableReplicas
			deployment.Status.OllamaStatus.ReadyReplicas = ollamaDeployment.Status.ReadyReplicas
			deployment.Status.OllamaStatus.UpdatedReplicas = ollamaDeployment.Status.UpdatedReplicas
		}

		deployment.Status.TotalReplicas = deployment.Spec.Ollama.Replicas
		deployment.Status.ReadyReplicas = deployment.Status.OllamaStatus.ReadyReplicas
	}

	// Get OpenWebUI deployment status if enabled
	if deployment.Spec.OpenWebUI.Enabled {
		openwebuiDeployment := &appsv1.Deployment{}
		err = r.Get(ctx, types.NamespacedName{
			Name:      deployment.GetOpenWebUIDeploymentName(),
			Namespace: deployment.Namespace,
		}, openwebuiDeployment)

		if err == nil {
			deployment.Status.OpenWebUIStatus.AvailableReplicas = openwebuiDeployment.Status.AvailableReplicas
			deployment.Status.OpenWebUIStatus.ReadyReplicas = openwebuiDeployment.Status.ReadyReplicas
			deployment.Status.OpenWebUIStatus.UpdatedReplicas = openwebuiDeployment.Status.UpdatedReplicas
		}

		deployment.Status.TotalReplicas += deployment.Spec.OpenWebUI.Replicas
		deployment.Status.ReadyReplicas += deployment.Status.OpenWebUIStatus.ReadyReplicas
	}

	// Get Tabby deployment status if enabled
	if deployment.Spec.Tabby.Enabled {
		tabbyDeployment := &appsv1.Deployment{}
		err = r.Get(ctx, types.NamespacedName{
			Name:      deployment.GetTabbyDeploymentName(),
			Namespace: deployment.Namespace,
		}, tabbyDeployment)

		if err == nil {
			deployment.Status.TabbyStatus.AvailableReplicas = tabbyDeployment.Status.AvailableReplicas
			deployment.Status.TabbyStatus.ReadyReplicas = tabbyDeployment.Status.ReadyReplicas
			deployment.Status.TabbyStatus.UpdatedReplicas = tabbyDeployment.Status.UpdatedReplicas
		}

		deployment.Status.TotalReplicas += deployment.Spec.Tabby.Replicas
		deployment.Status.ReadyReplicas += deployment.Status.TabbyStatus.ReadyReplicas
	}

	// Set phase
	// nolint:staticcheck
	if deployment.Status.ReadyReplicas == 0 {
		deployment.Status.Phase = "Pending"
	} else if deployment.Status.ReadyReplicas == deployment.Status.TotalReplicas {
		deployment.Status.Phase = "Ready"
	} else {
		deployment.Status.Phase = "Progressing"
	}

	// Use patch helper to update status - this only updates fields that actually changed
	return patchHelper.Patch(ctx, deployment)
}

// SetupWithManager sets up the controller with the Manager.
func (r *LMDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// Initialize specialized controllers
	return ctrl.NewControllerManagedBy(mgr).
		For(&llmgeeperiov1alpha1.LMDeployment{}).
		Owns(&appsv1.Deployment{}).
		// PVCs are managed manually via ensurePVC to avoid immutable field issues
		// Services, ConfigMaps, Secrets, and Ingresses are managed by individual controllers
		Named("lmdeployment").
		Complete(r)
}

// ensurePVC creates a PersistentVolumeClaim only if it doesn't exist
func (r *LMDeploymentReconciler) ensurePVC(ctx context.Context, pvc *corev1.PersistentVolumeClaim) error {
	logger := log.FromContext(ctx)
	existing := &corev1.PersistentVolumeClaim{}
	err := r.Get(ctx, types.NamespacedName{Name: pvc.Name, Namespace: pvc.Namespace}, existing)
	if err != nil && errors.IsNotFound(err) {
		// Create new PVC only if it doesn't exist
		logger.Info("Creating new PVC", "name", pvc.Name, "namespace", pvc.Namespace, "storageClass", pvc.Spec.StorageClassName)
		if err := r.Create(ctx, pvc); err != nil {
			// Check if it's already exists error (race condition)
			return fmt.Errorf("failed to create PVC %s: %w", pvc.Name, err)
		}
		logger.Info("Successfully created PVC", "name", pvc.Name, "namespace", pvc.Namespace)
	} else if err != nil {
		// Return error if it's not a "not found" error
		logger.Error(err, "Failed to get PVC", "name", pvc.Name, "namespace", pvc.Namespace)
		return err
	}
	// If PVC exists, do nothing (don't try to update immutable fields)
	return nil
}

// createSecretIfNotExists creates a Kubernetes secret only if it doesn't already exist
func (r *LMDeploymentReconciler) createSecretIfNotExists(ctx context.Context, secret *corev1.Secret) error {
	logger := log.FromContext(ctx)
	existingSecret := &corev1.Secret{}
	err := r.Get(ctx, client.ObjectKeyFromObject(secret), existingSecret)
	if err != nil {
		if errors.IsNotFound(err) {
			// Secret doesn't exist, create it
			if err := r.Create(ctx, secret); err != nil {
				return fmt.Errorf("failed to create secret %s: %w", secret.Name, err)
			}
			logger.Info("Created Langfuse secret", "name", secret.Name, "namespace", secret.Namespace)
		} else {
			return fmt.Errorf("failed to get secret %s: %w", secret.Name, err)
		}
	}
	return nil
}

// generateSecureSecret generates a cryptographically secure random secret
func generateSecureSecret(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}
