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
	"strings"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/util/intstr"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

const (
	// FinalizerName is the name of the finalizer
	FinalizerName = "llm.geeper.io/finalizer"
)

// OllamaDeploymentReconciler reconciles a Deployment object
type OllamaDeploymentReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=llm.geeper.io,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=llm.geeper.io,resources=deployments/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=llm.geeper.io,resources=deployments/finalizers,verbs=update
// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=networking.k8s.io,resources=ingresses,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=core,resources=pods,verbs=get;list;watch

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
func (r *OllamaDeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)
	logger.Info("Reconciling Deployment")

	// Fetch the Deployment instance
	deployment := &llmgeeperiov1alpha1.Deployment{}
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

	// Update status
	if err := r.updateStatus(ctx, deployment); err != nil {
		logger.Error(err, "Failed to update status")
		return ctrl.Result{RequeueAfter: time.Minute}, err
	}

	return ctrl.Result{RequeueAfter: time.Minute * 5}, nil
}

// setDefaults sets default values for the Deployment
func (r *OllamaDeploymentReconciler) setDefaults(deployment *llmgeeperiov1alpha1.Deployment) {
	// Set Ollama defaults
	if deployment.Spec.Ollama.Image == "" {
		deployment.Spec.Ollama.Image = "ollama/ollama"
	}
	if deployment.Spec.Ollama.ImageTag == "" {
		deployment.Spec.Ollama.ImageTag = "latest"
	}
	if deployment.Spec.Ollama.Replicas == 0 {
		deployment.Spec.Ollama.Replicas = 1
	}
	if deployment.Spec.Ollama.ServiceType == "" {
		deployment.Spec.Ollama.ServiceType = "ClusterIP"
	}
	if deployment.Spec.Ollama.ServicePort == 0 {
		deployment.Spec.Ollama.ServicePort = 11434
	}

	// Set OpenWebUI defaults
	if deployment.Spec.OpenWebUI.Image == "" {
		deployment.Spec.OpenWebUI.Image = "ghcr.io/open-webui/open-webui"
	}
	if deployment.Spec.OpenWebUI.ImageTag == "" {
		deployment.Spec.OpenWebUI.ImageTag = "main"
	}
	if deployment.Spec.OpenWebUI.Replicas == 0 {
		deployment.Spec.OpenWebUI.Replicas = 1
	}
	if deployment.Spec.OpenWebUI.ServiceType == "" {
		deployment.Spec.OpenWebUI.ServiceType = "ClusterIP"
	}
	if deployment.Spec.OpenWebUI.ServicePort == 0 {
		deployment.Spec.OpenWebUI.ServicePort = 8080
	}
}

// reconcileOllama reconciles the Ollama deployment
func (r *OllamaDeploymentReconciler) reconcileOllama(ctx context.Context, deployment *llmgeeperiov1alpha1.Deployment) error {
	// Create or update Ollama deployment
	ollamaDeployment := r.buildOllamaDeployment(deployment)
	if err := r.createOrUpdateDeployment(ctx, ollamaDeployment); err != nil {
		return err
	}

	// Create or update Ollama service
	ollamaService := r.buildOllamaService(deployment)
	if err := r.createOrUpdateService(ctx, ollamaService); err != nil {
		return err
	}

	return nil
}

// reconcileOpenWebUI reconciles the OpenWebUI deployment
func (r *OllamaDeploymentReconciler) reconcileOpenWebUI(ctx context.Context, deployment *llmgeeperiov1alpha1.Deployment) error {
	// Create or update OpenWebUI deployment
	openwebuiDeployment := r.buildOpenWebUIDeployment(deployment)
	if err := r.createOrUpdateDeployment(ctx, openwebuiDeployment); err != nil {
		return err
	}

	// Create or update OpenWebUI service
	openwebuiService := r.buildOpenWebUIService(deployment)
	if err := r.createOrUpdateService(ctx, openwebuiService); err != nil {
		return err
	}

	// Create or update Ingress if enabled
	if deployment.Spec.OpenWebUI.IngressEnabled && deployment.Spec.OpenWebUI.IngressHost != "" {
		ingress := r.buildOpenWebUIIngress(deployment)
		if err := r.createOrUpdateIngress(ctx, ingress); err != nil {
			return err
		}
	}

	return nil
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
func (r *OllamaDeploymentReconciler) finalizeDeployment(ctx context.Context, deployment *llmgeeperiov1alpha1.Deployment) error {
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

// buildOllamaDeployment builds the Ollama deployment object
func (r *OllamaDeploymentReconciler) buildOllamaDeployment(deployment *llmgeeperiov1alpha1.Deployment) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "ollama",
		"llm-deployment": deployment.Name,
	}

	// Build postStart hook for model pulling
	var postStartCommands []string
	for _, model := range deployment.Spec.Ollama.Models {
		// Extract model name and tag from "modelname:tag" format
		modelStr := string(model)
		modelName := modelStr
		modelTag := "latest"

		// Check if tag is specified (format: "modelname:tag")
		if strings.Contains(modelStr, ":") {
			parts := strings.Split(modelStr, ":")
			if len(parts) == 2 {
				modelName = parts[0]
				modelTag = parts[1]
			}
		}

		// Add model pull command to postStart hook
		postStartCommands = append(postStartCommands,
			fmt.Sprintf("ollama pull %s:%s", modelName, modelTag))
	}

	ollamaDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-ollama", deployment.Name),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &deployment.Spec.Ollama.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "ollama",
							Image: fmt.Sprintf("%s:%s", deployment.Spec.Ollama.Image, deployment.Spec.Ollama.ImageTag),
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: deployment.Spec.Ollama.ServicePort,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Resources: r.buildResourceRequirements(deployment.Spec.Ollama.Resources),
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "ollama-data",
									MountPath: "/root/.ollama",
								},
							},
							Lifecycle: &corev1.Lifecycle{
								PostStart: &corev1.LifecycleHandler{
									Exec: &corev1.ExecAction{
										Command: []string{"/bin/sh", "-c", strings.Join(postStartCommands, " && ")},
									},
								},
							},
						},
					},
					Volumes: []corev1.Volume{
						{
							Name: "ollama-data",
							VolumeSource: corev1.VolumeSource{
								EmptyDir: &corev1.EmptyDirVolumeSource{},
							},
						},
					},
				},
			},
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, ollamaDeployment, r.Scheme)
	return ollamaDeployment
}

// buildOpenWebUIDeployment builds the OpenWebUI deployment object
func (r *OllamaDeploymentReconciler) buildOpenWebUIDeployment(deployment *llmgeeperiov1alpha1.Deployment) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "openwebui",
		"llm-deployment": deployment.Name,
	}

	ollamaServiceName := fmt.Sprintf("%s-ollama", deployment.Name)

	openwebuiDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-openwebui", deployment.Name),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &deployment.Spec.OpenWebUI.Replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: labels,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:  "openwebui",
							Image: fmt.Sprintf("%s:%s", deployment.Spec.OpenWebUI.Image, deployment.Spec.OpenWebUI.ImageTag),
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: deployment.Spec.OpenWebUI.ServicePort,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Resources: r.buildResourceRequirements(deployment.Spec.OpenWebUI.Resources),
							Env: []corev1.EnvVar{
								{
									Name:  "OLLAMA_BASE_URL",
									Value: fmt.Sprintf("http://%s:%d", ollamaServiceName, deployment.Spec.Ollama.ServicePort),
								},
								{
									Name:  "WEBUI_SECRET_KEY",
									Value: "your-secret-key-here",
								},
							},
						},
					},
				},
			},
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, openwebuiDeployment, r.Scheme)
	return openwebuiDeployment
}

// buildOllamaService builds the Ollama service object
func (r *OllamaDeploymentReconciler) buildOllamaService(deployment *llmgeeperiov1alpha1.Deployment) *corev1.Service {
	labels := map[string]string{
		"app":            "ollama",
		"llm-deployment": deployment.Name,
	}

	ollamaService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-ollama", deployment.Name),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType(deployment.Spec.Ollama.ServiceType),
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       deployment.Spec.Ollama.ServicePort,
					TargetPort: intstr.FromInt(int(deployment.Spec.Ollama.ServicePort)),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Selector: labels,
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, ollamaService, r.Scheme)
	return ollamaService
}

// buildOpenWebUIService builds the OpenWebUI service object
func (r *OllamaDeploymentReconciler) buildOpenWebUIService(deployment *llmgeeperiov1alpha1.Deployment) *corev1.Service {
	labels := map[string]string{
		"app":            "openwebui",
		"llm-deployment": deployment.Name,
	}

	openwebuiService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-openwebui", deployment.Name),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType(deployment.Spec.OpenWebUI.ServiceType),
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       deployment.Spec.OpenWebUI.ServicePort,
					TargetPort: intstr.FromInt(int(deployment.Spec.OpenWebUI.ServicePort)),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Selector: labels,
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, openwebuiService, r.Scheme)
	return openwebuiService
}

// buildOpenWebUIIngress builds the OpenWebUI ingress object
func (r *OllamaDeploymentReconciler) buildOpenWebUIIngress(deployment *llmgeeperiov1alpha1.Deployment) *networkingv1.Ingress {
	labels := map[string]string{
		"app":            "openwebui",
		"llm-deployment": deployment.Name,
	}

	serviceName := fmt.Sprintf("%s-openwebui", deployment.Name)
	pathType := networkingv1.PathTypePrefix

	ingress := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-openwebui", deployment.Name),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: deployment.Spec.OpenWebUI.IngressHost,
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &pathType,
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: serviceName,
											Port: networkingv1.ServiceBackendPort{
												Number: deployment.Spec.OpenWebUI.ServicePort,
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, ingress, r.Scheme)
	return ingress
}

// buildResourceRequirements builds resource requirements from the spec
func (r *OllamaDeploymentReconciler) buildResourceRequirements(resources llmgeeperiov1alpha1.ResourceRequirements) corev1.ResourceRequirements {
	req := corev1.ResourceRequirements{}

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

	return req
}

// createOrUpdateDeployment creates or updates a deployment
func (r *OllamaDeploymentReconciler) createOrUpdateDeployment(ctx context.Context, deployment *appsv1.Deployment) error {
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
func (r *OllamaDeploymentReconciler) createOrUpdateService(ctx context.Context, service *corev1.Service) error {
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
func (r *OllamaDeploymentReconciler) createOrUpdateIngress(ctx context.Context, ingress *networkingv1.Ingress) error {
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

// updateStatus updates the status of the Deployment
func (r *OllamaDeploymentReconciler) updateStatus(ctx context.Context, deployment *llmgeeperiov1alpha1.Deployment) error {
	// Get Ollama deployment status
	ollamaDeployment := &appsv1.Deployment{}
	err := r.Get(ctx, types.NamespacedName{
		Name:      fmt.Sprintf("%s-ollama", deployment.Name),
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
			Name:      fmt.Sprintf("%s-openwebui", deployment.Name),
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

	// Calculate overall status
	deployment.Status.TotalReplicas = deployment.Spec.Ollama.Replicas
	if deployment.Spec.OpenWebUI.Enabled {
		deployment.Status.TotalReplicas += deployment.Spec.OpenWebUI.Replicas
	}

	deployment.Status.ReadyReplicas = deployment.Status.OllamaStatus.ReadyReplicas
	if deployment.Spec.OpenWebUI.Enabled {
		deployment.Status.ReadyReplicas += deployment.Status.OpenWebUIStatus.ReadyReplicas
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
func (r *OllamaDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&llmgeeperiov1alpha1.Deployment{}).
		Owns(&appsv1.Deployment{}).
		Owns(&corev1.Service{}).
		Owns(&networkingv1.Ingress{}).
		Named("ollamadeployment").
		Complete(r)
}
