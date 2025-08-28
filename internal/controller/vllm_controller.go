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

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/utils/ptr"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

// ensureVLLMApiKeySecret ensures the vLLM API key secret exists and returns the API key
func (r *LMDeploymentReconciler) ensureVLLMApiKeySecret(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) (string, error) {
	secretName := deployment.GetVLLMApiKeySecretName()
	existingSecret := &corev1.Secret{}
	err := r.Get(ctx, client.ObjectKey{
		Name:      secretName,
		Namespace: deployment.Namespace,
	}, existingSecret)

	key := deployment.Spec.VLLM.ApiKey.Key
	if err == nil {
		if apiKeyBytes, exists := existingSecret.Data[key]; exists {
			apiKey := string(apiKeyBytes)
			if apiKey != "" {
				return apiKey, nil
			}
		}
	} else if !errors.IsNotFound(err) {
		// Error other than not found
		return "", fmt.Errorf("failed to get vLLM API key secret: %w", err)
	}

	// Generate new API key and create secret
	apiKey, err := generateSecureSecret(32)
	if err != nil {
		return "", fmt.Errorf("failed to generate vLLM API key: %w", err)
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: deployment.Namespace,
			Labels: map[string]string{
				"app":            "vllm",
				"llm-deployment": deployment.Name,
			},
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			key: []byte(apiKey),
		},
	}

	_ = controllerutil.SetControllerReference(deployment, secret, r.Scheme)

	if err := r.createOrUpdateSecret(ctx, secret); err != nil {
		return "", fmt.Errorf("failed to create vLLM API key secret: %w", err)
	}

	return apiKey, nil
}

// reconcileVLLM reconciles the vLLM deployment
func (r *LMDeploymentReconciler) reconcileVLLM(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
	// Ensure vLLM API key secret exists if enabled
	_, err := r.ensureVLLMApiKeySecret(ctx, deployment)
	if err != nil {
		return fmt.Errorf("failed to ensure vLLM API key secret: %w", err)
	}

	// Create or update vLLM model deployments
	for _, modelSpec := range deployment.Spec.VLLM.Models {
		// Create or update model deployment
		vllmDeployment := r.buildVLLMModelDeployment(deployment, modelSpec)
		if err := r.createOrUpdateDeployment(ctx, vllmDeployment); err != nil {
			return err
		}

		// Create or update model service
		vllmService := r.buildVLLMModelService(deployment, modelSpec)
		if err := r.createOrUpdateService(ctx, vllmService); err != nil {
			return err
		}

		// Create or update model PVC if persistence is enabled
		if modelSpec.Persistence != nil && modelSpec.Persistence.Enabled {
			vllmPVC := r.buildVLLMModelPVC(deployment, modelSpec)
			if err := r.ensurePVC(ctx, vllmPVC); err != nil {
				return err
			}
		}
	}

	// Create or update vLLM router
	routerDeployment := r.buildVLLMRouterDeployment(deployment)
	if err := r.createOrUpdateDeployment(ctx, routerDeployment); err != nil {
		return err
	}

	routerService := r.buildVLLMRouterService(deployment)
	if err := r.createOrUpdateService(ctx, routerService); err != nil {
		return err
	}

	return nil
}

// buildVLLMModelDeployment builds a vLLM model deployment object
func (r *LMDeploymentReconciler) buildVLLMModelDeployment(deployment *llmgeeperiov1alpha1.LMDeployment, modelSpec llmgeeperiov1alpha1.VLLMModelSpec) *appsv1.Deployment {
	labels := map[string]string{
		"app":             "vllm",
		"llm-deployment":  deployment.Name,
		"vllm-model":      modelSpec.Name,
		"vllm-model-name": modelSpec.Model,
	}

	// Use model-specific image or fall back to global default
	image := modelSpec.Image
	if image == "" && deployment.Spec.VLLM.GlobalConfig != nil {
		image = deployment.Spec.VLLM.GlobalConfig.Image
	}
	if image == "" {
		image = "vllm/vllm-openai:latest"
	}

	// Use model-specific replicas or default to 1
	replicas := modelSpec.Replicas
	if replicas == 0 {
		replicas = 1
	}

	// Use model-specific service port or fall back to global default
	servicePort := modelSpec.Service.Port
	if servicePort == 0 && deployment.Spec.VLLM.GlobalConfig != nil {
		servicePort = deployment.Spec.VLLM.GlobalConfig.Service.Port
	}
	if servicePort == 0 {
		servicePort = 8000
	}

	// Build container spec
	container := corev1.Container{
		Name:  "vllm",
		Image: image,
		Ports: []corev1.ContainerPort{
			{
				Name:          "http",
				ContainerPort: servicePort,
				Protocol:      corev1.ProtocolTCP,
			},
		},
		Command: []string{"vllm", "serve", modelSpec.Model},
		SecurityContext: &corev1.SecurityContext{
			RunAsGroup:     ptr.To(int64(44)),
			SeccompProfile: &corev1.SeccompProfile{Type: corev1.SeccompProfileTypeUnconfined},
			Capabilities:   &corev1.Capabilities{Add: []corev1.Capability{"SYS_PTRACE"}},
		},
		Resources: r.buildVLLMResourceRequirements(modelSpec.Resources, deployment.Spec.VLLM.GlobalConfig),
		Env: []corev1.EnvVar{
			{
				Name:  "HOST",
				Value: "0.0.0.0",
			},
			{
				Name:  "PORT",
				Value: fmt.Sprintf("%d", servicePort),
			},
		},
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      "vllm-data",
				MountPath: "/root/.cache/huggingface",
			},
		},
	}

	container.Env = append(container.Env, corev1.EnvVar{
		Name: "VLLM_API_KEY",
		ValueFrom: &corev1.EnvVarSource{
			SecretKeyRef: &corev1.SecretKeySelector{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: deployment.GetVLLMApiKeySecretName(),
				},
				Key: deployment.Spec.VLLM.ApiKey.Key,
			},
		},
	})

	// Add custom environment variables
	if len(modelSpec.EnvVars) > 0 {
		container.Env = append(container.Env, modelSpec.EnvVars...)
	}

	// Add custom volume mounts
	if len(modelSpec.VolumeMounts) > 0 {
		container.VolumeMounts = append(container.VolumeMounts, modelSpec.VolumeMounts...)
	}

	// Build volumes
	volumes := []corev1.Volume{
		{
			Name: "vllm-data",
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		},
	}

	// Add custom volumes
	if len(modelSpec.Volumes) > 0 {
		volumes = append(volumes, modelSpec.Volumes...)
	}

	// Use PVC if persistence is enabled
	if modelSpec.Persistence != nil && modelSpec.Persistence.Enabled {
		volumes[0] = corev1.Volume{
			Name: "vllm-data",
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: deployment.GetVLLMModelPVCName(modelSpec.Name),
				},
			},
		}
	}

	vllmDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetVLLMModelDeploymentName(modelSpec.Name),
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
					Containers: []corev1.Container{container},
					Volumes:    volumes,
					Affinity:   modelSpec.Affinity,
				},
			},
		},
	}

	// Set owner reference
	_ = controllerutil.SetControllerReference(deployment, vllmDeployment, r.Scheme)
	return vllmDeployment
}

// buildVLLMModelService builds a vLLM model service object
func (r *LMDeploymentReconciler) buildVLLMModelService(deployment *llmgeeperiov1alpha1.LMDeployment, modelSpec llmgeeperiov1alpha1.VLLMModelSpec) *corev1.Service {
	labels := map[string]string{
		"app":             "vllm",
		"llm-deployment":  deployment.Name,
		"vllm-model":      modelSpec.Name,
		"vllm-model-name": modelSpec.Model,
	}

	// Use model-specific service configuration or fall back to global default
	serviceType := modelSpec.Service.Type
	servicePort := modelSpec.Service.Port
	if serviceType == "" && deployment.Spec.VLLM.GlobalConfig != nil {
		serviceType = deployment.Spec.VLLM.GlobalConfig.Service.Type
	}
	if servicePort == 0 && deployment.Spec.VLLM.GlobalConfig != nil {
		servicePort = deployment.Spec.VLLM.GlobalConfig.Service.Port
	}
	if servicePort == 0 {
		servicePort = 8000
	}
	if serviceType == "" {
		serviceType = corev1.ServiceTypeClusterIP
	}

	vllmService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetVLLMModelServiceName(modelSpec.Name),
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
	_ = controllerutil.SetControllerReference(deployment, vllmService, r.Scheme)
	return vllmService
}

// buildVLLMModelPVC builds a vLLM model PVC object
func (r *LMDeploymentReconciler) buildVLLMModelPVC(deployment *llmgeeperiov1alpha1.LMDeployment, modelSpec llmgeeperiov1alpha1.VLLMModelSpec) *corev1.PersistentVolumeClaim {
	labels := map[string]string{
		"app":             "vllm",
		"llm-deployment":  deployment.Name,
		"vllm-model":      modelSpec.Name,
		"vllm-model-name": modelSpec.Model,
	}

	// Use model-specific persistence configuration or fall back to global default
	persistence := modelSpec.Persistence
	if persistence == nil && deployment.Spec.VLLM.GlobalConfig != nil {
		persistence = deployment.Spec.VLLM.GlobalConfig.Persistence
	}

	storageClass := ""
	if persistence != nil && persistence.StorageClass != "" {
		storageClass = persistence.StorageClass
	}

	size := "10Gi" // Default size
	if persistence != nil && persistence.Size != "" {
		size = persistence.Size
	}

	vllmPVC := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetVLLMModelPVCName(modelSpec.Name),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteOnce,
			},
			StorageClassName: &storageClass,
			Resources: corev1.VolumeResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse(size),
				},
			},
		},
	}

	// Set owner reference
	_ = controllerutil.SetControllerReference(deployment, vllmPVC, r.Scheme)
	return vllmPVC
}

// buildVLLMRouterDeployment builds the vLLM router deployment object
func (r *LMDeploymentReconciler) buildVLLMRouterDeployment(deployment *llmgeeperiov1alpha1.LMDeployment) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "vllm-router",
		"llm-deployment": deployment.Name,
	}

	// Use router-specific image or default
	image := deployment.Spec.VLLM.Router.Image
	if image == "" {
		image = "lmcache/lmstack-router:latest"
	}

	// Use router-specific replicas or default to 1
	replicas := deployment.Spec.VLLM.Router.Replicas
	if replicas == 0 {
		replicas = 1
	}

	// Use router-specific service port or default to 8000
	servicePort := deployment.Spec.VLLM.Router.Service.Port
	if servicePort == 0 {
		servicePort = 8000
	}

	// Build model endpoints for router configuration
	var modelEndpoints []string
	for _, modelSpec := range deployment.Spec.VLLM.Models {
		modelEndpoints = append(modelEndpoints, fmt.Sprintf("%s:%d", deployment.GetVLLMModelServiceName(modelSpec.Name), servicePort))
	}

	// Build container spec
	container := corev1.Container{
		Name:  "vllm-router",
		Image: image,
		Ports: []corev1.ContainerPort{
			{
				Name:          "http",
				ContainerPort: servicePort,
				Protocol:      corev1.ProtocolTCP,
			},
		},
		Args: []string{
			"--host", "0.0.0.0",
			"--port", fmt.Sprintf("%d", servicePort),
			"--service-discovery", "k8s",
			"--k8s-namespace", "k8s",
			"--k8s-label-selector", "app=vllm,llm-deployment=" + deployment.Name,
		},
		Resources: r.buildResourceRequirements(deployment.Spec.VLLM.Router.Resources),
		Env: []corev1.EnvVar{
			{
				Name:  "HOST",
				Value: "0.0.0.0",
			},
			{
				Name:  "PORT",
				Value: fmt.Sprintf("%d", servicePort),
			},
			{
				Name:  "MODEL_ENDPOINTS",
				Value: fmt.Sprintf("%s", modelEndpoints),
			},
		},
	}

	container.Env = append(container.Env, corev1.EnvVar{
		Name: "VLLM_API_KEY",
		ValueFrom: &corev1.EnvVarSource{
			SecretKeyRef: &corev1.SecretKeySelector{
				LocalObjectReference: corev1.LocalObjectReference{
					Name: deployment.GetVLLMApiKeySecretName(),
				},
				Key: deployment.Spec.VLLM.ApiKey.Key,
			},
		},
	})

	// Add custom environment variables
	if len(deployment.Spec.VLLM.Router.EnvVars) > 0 {
		container.Env = append(container.Env, deployment.Spec.VLLM.Router.EnvVars...)
	}

	routerDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetVLLMRouterDeploymentName(),
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
					Containers: []corev1.Container{container},
					Affinity:   deployment.Spec.VLLM.Router.Affinity,
				},
			},
		},
	}

	// Set owner reference
	_ = controllerutil.SetControllerReference(deployment, routerDeployment, r.Scheme)
	return routerDeployment
}

// buildVLLMRouterService builds the vLLM router service object
func (r *LMDeploymentReconciler) buildVLLMRouterService(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.Service {
	labels := map[string]string{
		"app":            "vllm-router",
		"llm-deployment": deployment.Name,
	}

	// Use router-specific service configuration or default
	serviceType := deployment.Spec.VLLM.Router.Service.Type
	servicePort := deployment.Spec.VLLM.Router.Service.Port
	if serviceType == "" {
		serviceType = corev1.ServiceTypeClusterIP
	}
	if servicePort == 0 {
		servicePort = 8000
	}

	routerService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetVLLMRouterServiceName(),
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
	_ = controllerutil.SetControllerReference(deployment, routerService, r.Scheme)
	return routerService
}

// buildVLLMResourceRequirements builds resource requirements with fallback to global defaults
func (r *LMDeploymentReconciler) buildVLLMResourceRequirements(modelResources llmgeeperiov1alpha1.ResourceRequirements, globalConfig *llmgeeperiov1alpha1.VLLMGlobalConfig) corev1.ResourceRequirements {
	// Start with model-specific resources
	requests := modelResources.Requests
	limits := modelResources.Limits

	// Fall back to global defaults if not specified
	if globalConfig != nil {
		if len(requests) == 0 && len(globalConfig.Resources.Requests) > 0 {
			requests = globalConfig.Resources.Requests
		}
		if len(limits) == 0 && len(globalConfig.Resources.Limits) > 0 {
			limits = globalConfig.Resources.Limits
		}
	}

	return corev1.ResourceRequirements{
		Requests: requests,
		Limits:   limits,
	}
}
