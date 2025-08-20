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
	"fmt"
	"reflect"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
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
			// Request object not found, could have been deleted after reconcile request.
			// Return and don't requeue
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return ctrl.Result{}, err
	}

	// Check if the resource is being deleted
	if !deployment.DeletionTimestamp.IsZero() {
		// Resource is being deleted, handle finalization
		if err := r.finalizeDeployment(ctx, deployment); err != nil {
			logger.Error(err, "Failed to finalize deployment")
			return ctrl.Result{RequeueAfter: time.Minute}, err
		}
		return ctrl.Result{}, nil
	}

	// Add finalizer if it doesn't exist
	if !containsFinalizer(deployment.Finalizers, FinalizerName) {
		deployment.Finalizers = append(deployment.Finalizers, FinalizerName)
		if err := r.Update(ctx, deployment); err != nil {
			logger.Error(err, "Failed to add finalizer")
			return ctrl.Result{RequeueAfter: time.Minute}, err
		}
		return ctrl.Result{Requeue: true}, nil
	}

	// Set default values if not specified
	r.setDefaults(deployment)

	// Reconcile Ollama deployment
	if err := r.reconcileOllama(ctx, deployment); err != nil {
		logger.Error(err, "Failed to reconcile Ollama")
		return ctrl.Result{RequeueAfter: time.Minute}, err
	}

	// Reconcile OpenWebUI deployment if enabled
	if deployment.Spec.OpenWebUI.Enabled {
		if err := r.reconcileOpenWebUI(ctx, deployment); err != nil {
			logger.Error(err, "Failed to reconcile OpenWebUI")
			return ctrl.Result{RequeueAfter: time.Minute}, err
		}
	}

	// Reconcile Tabby deployment if enabled
	if deployment.Spec.Tabby.Enabled {
		if err := r.reconcileTabby(ctx, deployment); err != nil {
			logger.Error(err, "Failed to reconcile Tabby")
			return ctrl.Result{RequeueAfter: time.Minute}, err
		}
	}

	// Update status
	if err := r.updateStatus(ctx, deployment); err != nil {
		logger.Error(err, "Failed to update status")
		return ctrl.Result{RequeueAfter: time.Minute}, err
	}

	return ctrl.Result{RequeueAfter: time.Minute * 5}, nil
}

// setDefaults sets default values for the Deployment
func (r *LMDeploymentReconciler) setDefaults(deployment *llmgeeperiov1alpha1.LMDeployment) {
	// Set Ollama defaults
	if deployment.Spec.Ollama.Image == "" {
		deployment.Spec.Ollama.Image = "ollama/ollama:latest"
	}
	if deployment.Spec.Ollama.Replicas == 0 {
		deployment.Spec.Ollama.Replicas = 1
	}
	if deployment.Spec.Ollama.Service.Type == "" {
		deployment.Spec.Ollama.Service.Type = "ClusterIP"
	}
	if deployment.Spec.Ollama.Service.Port == 0 {
		deployment.Spec.Ollama.Service.Port = 11434
	}

	// Set OpenWebUI defaults
	if deployment.Spec.OpenWebUI.Image == "" {
		deployment.Spec.OpenWebUI.Image = "ghcr.io/open-webui/open-webui:main"
	}
	if deployment.Spec.OpenWebUI.Replicas == 0 {
		deployment.Spec.OpenWebUI.Replicas = 1
	}
	if deployment.Spec.OpenWebUI.Service.Type == "" {
		deployment.Spec.OpenWebUI.Service.Type = "ClusterIP"
	}
	if deployment.Spec.OpenWebUI.Service.Port == 0 {
		deployment.Spec.OpenWebUI.Service.Port = 8080
	}

	// Set OpenWebUI Redis defaults
	if deployment.Spec.OpenWebUI.Redis.Image == "" {
		deployment.Spec.OpenWebUI.Redis.Image = "redis:7-alpine"
	}
	if deployment.Spec.OpenWebUI.Redis.Service.Port == 0 {
		deployment.Spec.OpenWebUI.Redis.Service.Port = 6379
	}
	if deployment.Spec.OpenWebUI.Redis.Service.Type == "" {
		deployment.Spec.OpenWebUI.Redis.Service.Type = "ClusterIP"
	}
	if deployment.Spec.OpenWebUI.Redis.Persistence.Size == "" {
		deployment.Spec.OpenWebUI.Redis.Persistence.Size = "1Gi"
	}
	if deployment.Spec.OpenWebUI.Redis.Password == "" {
		// Generate a default Redis password if not provided
		deployment.Spec.OpenWebUI.Redis.Password = "redis-password-123"
	}

	// Set OpenWebUI Langfuse defaults and auto-enable pipelines if needed
	if deployment.Spec.OpenWebUI.Langfuse != nil && deployment.Spec.OpenWebUI.Langfuse.Enabled {
		// Auto-enable pipelines if Langfuse is enabled
		if deployment.Spec.OpenWebUI.Pipelines == nil {
			deployment.Spec.OpenWebUI.Pipelines = &llmgeeperiov1alpha1.PipelinesSpec{}
		}
		deployment.Spec.OpenWebUI.Pipelines.Enabled = true

		// Set Langfuse defaults
		if deployment.Spec.OpenWebUI.Langfuse.ProjectName == "" {
			deployment.Spec.OpenWebUI.Langfuse.ProjectName = deployment.Name
		}
		if deployment.Spec.OpenWebUI.Langfuse.Environment == "" {
			deployment.Spec.OpenWebUI.Langfuse.Environment = "production"
		}

		// If no URL is provided, set up self-hosted Langfuse defaults
		if deployment.Spec.OpenWebUI.Langfuse.URL == "" {
			if deployment.Spec.OpenWebUI.Langfuse.Deploy == nil {
				deployment.Spec.OpenWebUI.Langfuse.Deploy = &llmgeeperiov1alpha1.LangfuseDeploySpec{}
			}

			// Set self-hosted defaults
			if deployment.Spec.OpenWebUI.Langfuse.Deploy.Image == "" {
				deployment.Spec.OpenWebUI.Langfuse.Deploy.Image = "langfuse/langfuse:latest"
			}
			if deployment.Spec.OpenWebUI.Langfuse.Deploy.Replicas == 0 {
				deployment.Spec.OpenWebUI.Langfuse.Deploy.Replicas = 1
			}
			if deployment.Spec.OpenWebUI.Langfuse.Deploy.Port == 0 {
				deployment.Spec.OpenWebUI.Langfuse.Deploy.Port = 3000
			}
			if deployment.Spec.OpenWebUI.Langfuse.Deploy.ServiceType == "" {
				deployment.Spec.OpenWebUI.Langfuse.Deploy.ServiceType = "ClusterIP"
			}

			// Set persistence defaults
			if deployment.Spec.OpenWebUI.Langfuse.Deploy.Persistence == nil {
				deployment.Spec.OpenWebUI.Langfuse.Deploy.Persistence = &llmgeeperiov1alpha1.LangfusePersistenceSpec{}
			}
			if deployment.Spec.OpenWebUI.Langfuse.Deploy.Persistence.Size == "" {
				deployment.Spec.OpenWebUI.Langfuse.Deploy.Persistence.Size = "10Gi"
			}

		}
	}

	// Set OpenWebUI Pipelines defaults (for both manual and auto-enabled)
	if deployment.Spec.OpenWebUI.Pipelines != nil && deployment.Spec.OpenWebUI.Pipelines.Enabled {
		if deployment.Spec.OpenWebUI.Pipelines.Image == "" {
			deployment.Spec.OpenWebUI.Pipelines.Image = "ghcr.io/open-webui/pipelines:main"
		}
		if deployment.Spec.OpenWebUI.Pipelines.Replicas == 0 {
			deployment.Spec.OpenWebUI.Pipelines.Replicas = 1
		}
		if deployment.Spec.OpenWebUI.Pipelines.Port == 0 {
			deployment.Spec.OpenWebUI.Pipelines.Port = 9099
		}
		if deployment.Spec.OpenWebUI.Pipelines.ServiceType == "" {
			deployment.Spec.OpenWebUI.Pipelines.ServiceType = "ClusterIP"
		}
		if deployment.Spec.OpenWebUI.Pipelines.PipelinesDir == "" {
			deployment.Spec.OpenWebUI.Pipelines.PipelinesDir = "/app/pipelines"
		}
	}

	// Set Tabby defaults
	if deployment.Spec.Tabby.Image == "" {
		deployment.Spec.Tabby.Image = "tabbyml/tabby:latest"
	}
	if deployment.Spec.Tabby.Replicas == 0 {
		deployment.Spec.Tabby.Replicas = 1
	}
	if deployment.Spec.Tabby.Service.Type == "" {
		deployment.Spec.Tabby.Service.Type = "ClusterIP"
	}
	if deployment.Spec.Tabby.Service.Port == 0 {
		deployment.Spec.Tabby.Service.Port = 8080
	}
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
	req := corev1.ResourceRequirements{}

	if resources.Requests != nil {
		if resources.Requests.CPU != "" {
			if req.Requests == nil {
				req.Requests = corev1.ResourceList{}
			}
			req.Requests[corev1.ResourceCPU] = resource.MustParse(resources.Requests.CPU)
		}
		if resources.Requests.Memory != "" {
			if req.Requests == nil {
				req.Requests = corev1.ResourceList{}
			}
			req.Requests[corev1.ResourceMemory] = resource.MustParse(resources.Requests.Memory)
		}
	}

	if resources.Limits != nil {
		if resources.Limits.CPU != "" {
			if req.Limits == nil {
				req.Limits = corev1.ResourceList{}
			}
			req.Limits[corev1.ResourceCPU] = resource.MustParse(resources.Limits.CPU)
		}
		if resources.Limits.Memory != "" {
			if req.Limits == nil {
				req.Limits = corev1.ResourceList{}
			}
			req.Limits[corev1.ResourceMemory] = resource.MustParse(resources.Limits.Memory)
		}
	}

	return req
}

// createOrUpdateDeployment creates or updates a deployment
func (r *LMDeploymentReconciler) createOrUpdateDeployment(ctx context.Context, deployment *appsv1.Deployment) error {
	existing := &appsv1.Deployment{}
	err := r.Get(ctx, types.NamespacedName{Name: deployment.Name, Namespace: deployment.Namespace}, existing)
	if err != nil && errors.IsNotFound(err) {
		// Create new deployment
		if err := r.Create(ctx, deployment); err != nil {
			return err
		}
	} else if err == nil {
		// Update existing deployment
		if !reflect.DeepEqual(existing.Spec, deployment.Spec) {
			existing.Spec = deployment.Spec
			if err := r.Update(ctx, existing); err != nil {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}

// createOrUpdateService creates or updates a service
func (r *LMDeploymentReconciler) createOrUpdateService(ctx context.Context, service *corev1.Service) error {
	existing := &corev1.Service{}
	err := r.Get(ctx, types.NamespacedName{Name: service.Name, Namespace: service.Namespace}, existing)
	if err != nil && errors.IsNotFound(err) {
		// Create new service
		if err := r.Create(ctx, service); err != nil {
			return err
		}
	} else if err == nil {
		// Update existing service
		if !reflect.DeepEqual(existing.Spec, service.Spec) {
			existing.Spec = service.Spec
			if err := r.Update(ctx, existing); err != nil {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}

// createOrUpdateIngress creates or updates an ingress
func (r *LMDeploymentReconciler) createOrUpdateIngress(ctx context.Context, ingress *networkingv1.Ingress) error {
	existing := &networkingv1.Ingress{}
	err := r.Get(ctx, types.NamespacedName{Name: ingress.Name, Namespace: ingress.Namespace}, existing)
	if err != nil && errors.IsNotFound(err) {
		// Create new ingress
		if err := r.Create(ctx, ingress); err != nil {
			return err
		}
	} else if err == nil {
		// Update existing ingress
		if !reflect.DeepEqual(existing.Spec, ingress.Spec) {
			existing.Spec = ingress.Spec
			if err := r.Update(ctx, existing); err != nil {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}

// createOrUpdateConfigMap creates or updates a ConfigMap
func (r *LMDeploymentReconciler) createOrUpdateConfigMap(ctx context.Context, configMap *corev1.ConfigMap) error {
	existing := &corev1.ConfigMap{}
	err := r.Get(ctx, types.NamespacedName{Name: configMap.Name, Namespace: configMap.Namespace}, existing)
	if err != nil && errors.IsNotFound(err) {
		// Create new ConfigMap
		if err := r.Create(ctx, configMap); err != nil {
			return err
		}
	} else if err == nil {
		// Update existing ConfigMap
		if !reflect.DeepEqual(existing.Data, configMap.Data) {
			existing.Data = configMap.Data
			if err := r.Update(ctx, existing); err != nil {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}

// updateStatus updates the status of the Deployment
func (r *LMDeploymentReconciler) updateStatus(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
	// Get Ollama deployment status
	ollamaDeployment := &appsv1.Deployment{}
	err := r.Get(ctx, types.NamespacedName{
		Name:      deployment.GetOllamaDeploymentName(),
		Namespace: deployment.Namespace,
	}, ollamaDeployment)

	if err == nil {
		deployment.Status.OllamaStatus.AvailableReplicas = ollamaDeployment.Status.AvailableReplicas
		deployment.Status.OllamaStatus.ReadyReplicas = ollamaDeployment.Status.ReadyReplicas
		deployment.Status.OllamaStatus.UpdatedReplicas = ollamaDeployment.Status.UpdatedReplicas
		// Convert DeploymentCondition to metav1.Condition
		for _, cond := range ollamaDeployment.Status.Conditions {
			metaCond := metav1.Condition{
				Type:               string(cond.Type),
				Status:             metav1.ConditionStatus(cond.Status),
				LastTransitionTime: cond.LastTransitionTime,
				Reason:             cond.Reason,
				Message:            cond.Message,
			}
			deployment.Status.OllamaStatus.Conditions = append(deployment.Status.OllamaStatus.Conditions, metaCond)
		}
	}

	// Get OpenWebUI deployment status if enabled
	if deployment.Spec.OpenWebUI.Enabled {
		openwebuiDeployment := &appsv1.Deployment{}
		err := r.Get(ctx, types.NamespacedName{
			Name:      deployment.GetOpenWebUIDeploymentName(),
			Namespace: deployment.Namespace,
		}, openwebuiDeployment)

		if err == nil {
			deployment.Status.OpenWebUIStatus.AvailableReplicas = openwebuiDeployment.Status.AvailableReplicas
			deployment.Status.OpenWebUIStatus.ReadyReplicas = openwebuiDeployment.Status.ReadyReplicas
			deployment.Status.OpenWebUIStatus.UpdatedReplicas = openwebuiDeployment.Status.UpdatedReplicas
			// Convert DeploymentCondition to metav1.Condition
			for _, cond := range openwebuiDeployment.Status.Conditions {
				metaCond := metav1.Condition{
					Type:               string(cond.Type),
					Status:             metav1.ConditionStatus(cond.Status),
					LastTransitionTime: cond.LastTransitionTime,
					Reason:             cond.Reason,
					Message:            cond.Message,
				}
				deployment.Status.OpenWebUIStatus.Conditions = append(deployment.Status.OpenWebUIStatus.Conditions, metaCond)
			}
		}
	}

	// Get Tabby deployment status if enabled
	if deployment.Spec.Tabby.Enabled {
		tabbyDeployment := &appsv1.Deployment{}
		err := r.Get(ctx, types.NamespacedName{
			Name:      deployment.GetTabbyDeploymentName(),
			Namespace: deployment.Namespace,
		}, tabbyDeployment)

		if err == nil {
			deployment.Status.TabbyStatus.AvailableReplicas = tabbyDeployment.Status.AvailableReplicas
			deployment.Status.TabbyStatus.ReadyReplicas = tabbyDeployment.Status.ReadyReplicas
			deployment.Status.TabbyStatus.UpdatedReplicas = tabbyDeployment.Status.UpdatedReplicas
			// Convert DeploymentCondition to metav1.Condition
			for _, cond := range tabbyDeployment.Status.Conditions {
				metaCond := metav1.Condition{
					Type:               string(cond.Type),
					Status:             metav1.ConditionStatus(cond.Status),
					LastTransitionTime: cond.LastTransitionTime,
					Reason:             cond.Reason,
					Message:            cond.Message,
				}
				deployment.Status.TabbyStatus.Conditions = append(deployment.Status.TabbyStatus.Conditions, metaCond)
			}
		}
	}

	// Calculate overall status
	deployment.Status.TotalReplicas = deployment.Spec.Ollama.Replicas
	if deployment.Spec.OpenWebUI.Enabled {
		deployment.Status.TotalReplicas += deployment.Spec.OpenWebUI.Replicas
	}

	// Add Tabby replicas
	if deployment.Spec.Tabby.Enabled {
		deployment.Status.TotalReplicas += deployment.Spec.Tabby.Replicas
	}

	deployment.Status.ReadyReplicas = deployment.Status.OllamaStatus.ReadyReplicas
	if deployment.Spec.OpenWebUI.Enabled {
		deployment.Status.ReadyReplicas += deployment.Status.OpenWebUIStatus.ReadyReplicas
	}

	// Add Tabby ready replicas
	if deployment.Spec.Tabby.Enabled {
		deployment.Status.ReadyReplicas += deployment.Status.TabbyStatus.ReadyReplicas
	}

	// Set phase
	if deployment.Status.ReadyReplicas == 0 {
		deployment.Status.Phase = "Pending"
	} else if deployment.Status.ReadyReplicas == deployment.Status.TotalReplicas {
		deployment.Status.Phase = "Ready"
	} else {
		deployment.Status.Phase = "Progressing"
	}

	// Update status
	return r.Status().Update(ctx, deployment)
}

// SetupWithManager sets up the controller with the Manager.
func (r *LMDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	// Initialize specialized controllers
	return ctrl.NewControllerManagedBy(mgr).
		For(&llmgeeperiov1alpha1.LMDeployment{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&corev1.Secret{}).
		Owns(&networkingv1.Ingress{}).
		Owns(&corev1.PersistentVolumeClaim{}).
		Named("lmdeployment").
		Complete(r)
}

// ensurePVC creates a PersistentVolumeClaim only if it doesn't exist
func (r *LMDeploymentReconciler) ensurePVC(ctx context.Context, pvc *corev1.PersistentVolumeClaim) error {
	existing := &corev1.PersistentVolumeClaim{}
	err := r.Get(ctx, types.NamespacedName{Name: pvc.Name, Namespace: pvc.Namespace}, existing)
	if err != nil && errors.IsNotFound(err) {
		// Create new PVC only if it doesn't exist
		if err := r.Create(ctx, pvc); err != nil {
			return err
		}
	} else if err != nil {
		// Return error if it's not a "not found" error
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
