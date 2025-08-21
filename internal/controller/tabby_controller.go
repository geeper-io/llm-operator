package controller

import (
	"bytes"
	"context"
	"fmt"

	"github.com/BurntSushi/toml"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// reconcileTabby reconciles the Tabby deployment
func (r *LMDeploymentReconciler) reconcileTabby(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
	// Create or update Tabby PVC if persistence is enabled
	if deployment.Spec.Tabby.Persistence.Enabled {
		tabbyPVC := r.buildTabbyPVC(deployment)
		if err := r.ensurePVC(ctx, tabbyPVC); err != nil {
			return err
		}
	}

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
	if deployment.Spec.Tabby.Ingress.Host != "" {
		tabbyIngress := r.buildTabbyIngress(deployment)
		if err := r.createOrUpdateIngress(ctx, tabbyIngress); err != nil {
			return err
		}
	}

	return nil
}

// buildTabbyDeployment builds the Tabby deployment object
func (r *LMDeploymentReconciler) buildTabbyDeployment(deployment *llmgeeperiov1alpha1.LMDeployment) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "tabby",
		"llm-deployment": deployment.Name,
	}

	// Set default values
	// Build Tabby deployment
	image := deployment.Spec.Tabby.Image
	if image == "" {
		image = "tabbyml/tabby:latest"
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
			Name:      "tabby-config",
			MountPath: "/tmp/config",
		},
		{
			Name:      "tabby-data",
			MountPath: "/data",
		},
	}

	// Add custom volume mounts
	if deployment.Spec.Tabby.VolumeMounts != nil {
		volumeMounts = append(volumeMounts, deployment.Spec.Tabby.VolumeMounts...)
	}

	// Build volumes
	volumes := []corev1.Volume{
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

	// Add tabby-data volume based on persistence configuration
	if deployment.Spec.Tabby.Persistence.Enabled {
		// Use PVC for persistence
		volumes = append(volumes, corev1.Volume{
			Name: "tabby-data",
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: deployment.GetTabbyPVCName(),
				},
			},
		})
	} else {
		// Use EmptyDir for non-persistent storage
		volumes = append(volumes, corev1.Volume{
			Name: "tabby-data",
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		})
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
					InitContainers: []corev1.Container{
						{
							Name:    "tabby-config-init",
							Image:   "busybox:1.35",
							Command: []string{"/bin/sh"},
							Args: []string{
								"-c",
								"cp /tmp/config/config.toml /data/config.toml && echo 'Config file copied successfully'",
							},
							VolumeMounts: []corev1.VolumeMount{
								{
									Name:      "tabby-config",
									MountPath: "/tmp/config",
								},
							},
						},
					},
					Containers: []corev1.Container{
						{
							Name:  "tabby",
							Image: image,
							Args:  []string{"serve", "--device", deployment.Spec.Tabby.Device},
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
	_ = controllerutil.SetControllerReference(deployment, tabbyDeployment, r.Scheme)
	return tabbyDeployment
}

// buildTabbyService builds the Tabby service object
func (r *LMDeploymentReconciler) buildTabbyService(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.Service {
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
		serviceType = corev1.ServiceTypeClusterIP
	}

	tabbyService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetTabbyServiceName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: serviceType,
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
	_ = controllerutil.SetControllerReference(deployment, tabbyService, r.Scheme)
	return tabbyService
}

// buildTabbyIngress builds the Tabby ingress object
func (r *LMDeploymentReconciler) buildTabbyIngress(deployment *llmgeeperiov1alpha1.LMDeployment) *networkingv1.Ingress {
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
	_ = controllerutil.SetControllerReference(deployment, tabbyIngress, r.Scheme)
	return tabbyIngress
}

// buildTabbyConfigMap builds the Tabby configuration ConfigMap
func (r *LMDeploymentReconciler) buildTabbyConfigMap(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.ConfigMap {
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

// buildTabbyPVC builds the Tabby PVC object
func (r *LMDeploymentReconciler) buildTabbyPVC(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.PersistentVolumeClaim {
	labels := map[string]string{
		"app":            "tabby",
		"llm-deployment": deployment.Name,
	}

	// Set default values
	var storageClass *string
	if deployment.Spec.Tabby.Persistence.StorageClass != "" {
		storageClass = &deployment.Spec.Tabby.Persistence.StorageClass
	}
	size := deployment.Spec.Tabby.Persistence.Size
	if size == "" {
		size = "10Gi" // Default size
	}

	return &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetTabbyPVCName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteOnce,
			},
			StorageClassName: storageClass,
			Resources: corev1.VolumeResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse(size),
				},
			},
		},
	}
}

// generateTabbyConfig generates the Tabby TOML configuration
func (r *LMDeploymentReconciler) generateTabbyConfig(deployment *llmgeeperiov1alpha1.LMDeployment) (string, error) {
	// Build Ollama service host
	ollamaHost := fmt.Sprintf("%s.%s:%d",
		deployment.GetOllamaServiceName(),
		deployment.Namespace,
		deployment.GetOllamaServicePort())

	// Create Tabby configuration
	config := &TabbyConfig{
		Model: TabbyModelConfig{
			Completion: TabbyCompletionConfig{
				HTTP: TabbyHTTPConfig{
					Kind:           "ollama/completion",
					ModelName:      deployment.Spec.Tabby.CompletionModel,
					APIEndpoint:    fmt.Sprintf("http://%s", ollamaHost),
					PromptTemplate: "<PRE> {prefix} <SUF>{suffix} <MID>",
				},
			},
			Chat: TabbyChatConfig{
				HTTP: TabbyHTTPConfig{
					Kind:        "ollama/chat",
					ModelName:   deployment.Spec.Tabby.ChatModel,
					APIEndpoint: fmt.Sprintf("http://%s", ollamaHost),
				},
			},
			Embedding: TabbyEmbeddingConfig{
				Local: TabbyLocalConfig{
					ModelID: "Nomic-Embed-Text",
				},
			},
		},
	}

	// Encode to TOML
	var buf bytes.Buffer
	encoder := toml.NewEncoder(&buf)
	if err := encoder.Encode(config); err != nil {
		return "", fmt.Errorf("failed to encode Tabby config: %w", err)
	}

	return buf.String(), nil
}

// TabbyConfig represents the Tabby configuration structure
type TabbyConfig struct {
	Model TabbyModelConfig `toml:"model"`
}

type TabbyModelConfig struct {
	Completion TabbyCompletionConfig `toml:"completion"`
	Chat       TabbyChatConfig       `toml:"chat"`
	Embedding  TabbyEmbeddingConfig  `toml:"embedding"`
}

type TabbyCompletionConfig struct {
	HTTP TabbyHTTPConfig `toml:"http"`
}

type TabbyChatConfig struct {
	HTTP TabbyHTTPConfig `toml:"http"`
}

type TabbyEmbeddingConfig struct {
	Local TabbyLocalConfig `toml:"local"`
}

type TabbyHTTPConfig struct {
	Kind            string   `toml:"kind"`
	ModelName       string   `toml:"model_name"`
	APIEndpoint     string   `toml:"api_endpoint"`
	SupportedModels []string `toml:"supported_models,omitempty"`
	PromptTemplate  string   `toml:"prompt_template,omitempty"`
}

type TabbyLocalConfig struct {
	ModelID string `toml:"model_id"`
}
