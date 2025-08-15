package controller

import (
	"context"
	"fmt"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	"github.com/BurntSushi/toml"
	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

// reconcileTabby reconciles the Tabby deployment
func (r *OllamaDeploymentReconciler) reconcileTabby(ctx context.Context, deployment *llmgeeperiov1alpha1.Deployment) error {
	// Create or update Tabby ConfigMap
	tabbyConfigMap := r.buildTabbyConfigMap(deployment)
	if err := r.createOrUpdateConfigMap(ctx, tabbyConfigMap); err != nil {
		return err
	}

	// Create or update Tabby deployment
	tabbyDeployment := r.buildTabbyDeployment(deployment)
	if err := r.createOrUpdateDeployment(ctx, tabbyDeployment); err != nil {
		return err
	}

	// Create or update Tabby service
	tabbyService := r.buildTabbyService(deployment)
	if err := r.createOrUpdateService(ctx, tabbyService); err != nil {
		return err
	}

	// Create or update Tabby ingress if enabled
	if deployment.Spec.Tabby.Ingress.Enabled && deployment.Spec.Tabby.Ingress.Host != "" {
		tabbyIngress := r.buildTabbyIngress(deployment)
		if err := r.createOrUpdateIngress(ctx, tabbyIngress); err != nil {
			return err
		}
	}

	return nil
}

// buildTabbyDeployment builds the Tabby deployment object
func (r *OllamaDeploymentReconciler) buildTabbyDeployment(deployment *llmgeeperiov1alpha1.Deployment) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "tabby",
		"llm-deployment": deployment.Name,
	}

	// Set default values
	image := deployment.Spec.Tabby.Image
	if image == "" {
		image = "tabbyml/tabby"
	}
	imageTag := deployment.Spec.Tabby.ImageTag
	if imageTag == "" {
		imageTag = "latest"
	}
	replicas := deployment.Spec.Tabby.Replicas
	if replicas == 0 {
		replicas = 1
	}
	servicePort := deployment.Spec.Tabby.Service.Port
	if servicePort == 0 {
		servicePort = 8080
	}

	// Note: Model configuration is now handled via config.toml file

	// Build environment variables (only custom ones, no Ollama-specific ones)
	envVars := []corev1.EnvVar{}
	if deployment.Spec.Tabby.EnvVars != nil {
		envVars = append(envVars, deployment.Spec.Tabby.EnvVars...)
	}

	// Build volume mounts
	volumeMounts := []corev1.VolumeMount{
		{
			Name:      "tabby-data",
			MountPath: "/data",
		},
		{
			Name:      "tabby-config",
			MountPath: "/root/.tabby",
		},
	}

	// Add custom volume mounts
	if deployment.Spec.Tabby.VolumeMounts != nil {
		volumeMounts = append(volumeMounts, deployment.Spec.Tabby.VolumeMounts...)
	}

	// Build volumes
	volumes := []corev1.Volume{
		{
			Name: "tabby-data",
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		},
		{
			Name: "tabby-config",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: deployment.GetTabbyConfigMapName(),
					},
					Items: []corev1.KeyToPath{
						{
							Key:  "config.toml",
							Path: "config.toml",
						},
					},
				},
			},
		},
	}

	// Add custom volumes
	if deployment.Spec.Tabby.Volumes != nil {
		volumes = append(volumes, deployment.Spec.Tabby.Volumes...)
	}

	tabbyDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetTabbyDeploymentName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &replicas,
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
							Name:  "tabby",
							Image: fmt.Sprintf("%s:%s", image, imageTag),
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: servicePort,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Resources:    r.buildResourceRequirements(deployment.Spec.Tabby.Resources),
							Env:          envVars,
							VolumeMounts: volumeMounts,
						},
					},
					Volumes: volumes,
				},
			},
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, tabbyDeployment, r.Scheme)
	return tabbyDeployment
}

// buildTabbyService builds the Tabby service object
func (r *OllamaDeploymentReconciler) buildTabbyService(deployment *llmgeeperiov1alpha1.Deployment) *corev1.Service {
	labels := map[string]string{
		"app":            "tabby",
		"llm-deployment": deployment.Name,
	}

	servicePort := deployment.Spec.Tabby.Service.Port
	if servicePort == 0 {
		servicePort = 8080
	}

	serviceType := deployment.Spec.Tabby.Service.Type
	if serviceType == "" {
		serviceType = "ClusterIP"
	}

	tabbyService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetTabbyServiceName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType(serviceType),
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       servicePort,
					TargetPort: intstr.FromInt32(servicePort),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Selector: labels,
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, tabbyService, r.Scheme)
	return tabbyService
}

// buildTabbyIngress builds the Tabby ingress object
func (r *OllamaDeploymentReconciler) buildTabbyIngress(deployment *llmgeeperiov1alpha1.Deployment) *networkingv1.Ingress {
	labels := map[string]string{
		"app":            "tabby",
		"llm-deployment": deployment.Name,
	}

	servicePort := deployment.Spec.Tabby.Service.Port
	if servicePort == 0 {
		servicePort = 8080
	}

	ingressHost := deployment.Spec.Tabby.Ingress.Host
	if ingressHost == "" {
		ingressHost = fmt.Sprintf("tabby-%s.localhost", deployment.Name)
	}

	tabbyIngress := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetTabbyIngressName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: ingressHost,
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &[]networkingv1.PathType{networkingv1.PathTypePrefix}[0],
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: deployment.GetTabbyServiceName(),
											Port: networkingv1.ServiceBackendPort{
												Number: servicePort,
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
	controllerutil.SetControllerReference(deployment, tabbyIngress, r.Scheme)
	return tabbyIngress
}

// buildTabbyConfigMap builds the Tabby configuration ConfigMap
func (r *OllamaDeploymentReconciler) buildTabbyConfigMap(deployment *llmgeeperiov1alpha1.Deployment) *corev1.ConfigMap {
	// Generate TOML configuration
	configTOML, err := r.generateTabbyConfig(deployment)
	if err != nil {
		// Log error but continue with empty config
		configTOML = "# Error generating configuration\n"
	}

	labels := map[string]string{
		"app":            "tabby",
		"llm-deployment": deployment.Name,
	}

	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetTabbyConfigMapName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Data: map[string]string{
			"config.toml": configTOML,
		},
	}
}

// generateTabbyConfig generates the Tabby TOML configuration
func (r *OllamaDeploymentReconciler) generateTabbyConfig(deployment *llmgeeperiov1alpha1.Deployment) (string, error) {
	// Determine Ollama service details
	ollamaServiceName := deployment.GetOllamaServiceName()
	ollamaServicePort := deployment.GetOllamaServicePort()

	// Set default model name if not specified
	modelName := deployment.Spec.Tabby.ModelName
	if modelName == "" && len(deployment.Spec.Ollama.Models) > 0 {
		// Use the first model from Ollama spec
		modelStr := string(deployment.Spec.Ollama.Models[0])
		if strings.Contains(modelStr, ":") {
			parts := strings.Split(modelStr, ":")
			if len(parts) == 2 {
				modelName = parts[0]
			} else {
				modelName = modelStr
			}
		} else {
			modelName = modelStr
		}
	}

	// Build TOML configuration
	config := map[string]interface{}{
		"chat": map[string]interface{}{
			"model": map[string]interface{}{
				"ollama": map[string]interface{}{
					"host":  fmt.Sprintf("%s.%s.svc.cluster.local:%d", ollamaServiceName, deployment.Namespace, ollamaServicePort),
					"model": modelName,
				},
			},
		},
		"completion": map[string]interface{}{
			"model": map[string]interface{}{
				"ollama": map[string]interface{}{
					"host":  fmt.Sprintf("%s.%s.svc.cluster.local:%d", ollamaServiceName, deployment.Namespace, ollamaServicePort),
					"model": modelName,
				},
			},
		},
		"model": map[string]interface{}{
			"embedding": map[string]interface{}{
				"local": map[string]interface{}{
					"model_id": "Nomic-Embed-Text",
				},
			},
		},
	}

	// Convert to TOML
	var buf strings.Builder
	if err := toml.NewEncoder(&buf).Encode(config); err != nil {
		return "", err
	}

	return buf.String(), nil
}
