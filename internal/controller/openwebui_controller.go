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
	"encoding/json"
	"fmt"
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/cluster-api/util/patch"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/log"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// ensurePipelineSecret ensures the pipelines secret exists and returns the API key
func (r *LMDeploymentReconciler) ensurePipelineSecret(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) (string, error) {
	secretName := fmt.Sprintf("%s-pipelines-secret", deployment.Name)
	existingSecret := &corev1.Secret{}

	err := r.Get(ctx, client.ObjectKey{
		Name:      secretName,
		Namespace: deployment.Namespace,
	}, existingSecret)

	if err == nil {
		// Secret exists, check if it has PIPELINES_API_KEY
		if apiKeyBytes, exists := existingSecret.Data["PIPELINES_API_KEY"]; exists {
			apiKey := string(apiKeyBytes)
			if apiKey != "" {
				return apiKey, nil
			}
		}
	} else if !errors.IsNotFound(err) {
		// Error other than not found
		return "", fmt.Errorf("failed to get pipelines secret: %w", err)
	}

	// Generate new API key and create secret
	apiKey, err := generateSecureSecret(32)
	if err != nil {
		return "", fmt.Errorf("failed to generate API key: %w", err)
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      secretName,
			Namespace: deployment.Namespace,
			Labels: map[string]string{
				"app":            "pipelines",
				"llm-deployment": deployment.Name,
			},
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			"PIPELINES_API_KEY": []byte(apiKey),
		},
	}

	_ = controllerutil.SetControllerReference(deployment, secret, r.Scheme)

	if err := r.createOrUpdateSecret(ctx, secret); err != nil {
		return "", fmt.Errorf("failed to create pipelines secret: %w", err)
	}

	return apiKey, nil
}

// reconcileOpenWebUI reconciles the OpenWebUI deployment
func (r *LMDeploymentReconciler) reconcileOpenWebUI(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
	// Reconcile Redis if needed for OpenWebUI
	if err := r.reconcileRedis(ctx, deployment); err != nil {
		return err
	}

	// Reconcile Pipelines if enabled
	if deployment.Spec.OpenWebUI.Pipelines != nil && deployment.Spec.OpenWebUI.Pipelines.Enabled {
		if err := r.reconcilePipelines(ctx, deployment); err != nil {
			return err
		}
	}

	// Create or update OpenWebUI PVC if persistence is enabled
	if deployment.Spec.OpenWebUI.Persistence != nil && deployment.Spec.OpenWebUI.Persistence.Enabled {
		pvc := r.buildOpenWebUIPVC(deployment)
		if err := r.ensurePVC(ctx, pvc); err != nil {
			return err
		}
	}

	// Create or update OpenWebUI secret for WEBUI_SECRET_KEY
	openwebuiSecret, err := r.buildOpenWebUISecret(deployment)
	if err != nil {
		return err
	}
	if err := r.createSecretIfNotExists(ctx, openwebuiSecret); err != nil {
		return err
	}

	// Create or update OpenWebUI config Secret
	openwebuiConfig, err := r.buildOpenWebUIConfig(ctx, deployment)
	if err != nil {
		return fmt.Errorf("failed to build OpenWebUI config: %w", err)
	}

	if err := r.createOrUpdateSecret(ctx, openwebuiConfig); err != nil {
		return fmt.Errorf("failed to create or update OpenWebUI config secret: %w", err)
	}

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
	if deployment.Spec.OpenWebUI.Ingress.Host != "" {
		ingress := r.buildOpenWebUIIngress(deployment)
		if err := r.createOrUpdateIngress(ctx, ingress); err != nil {
			return err
		}
	}

	return nil
}

// reconcilePipelines reconciles the OpenWebUI Pipelines deployment
func (r *LMDeploymentReconciler) reconcilePipelines(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
	// Ensure pipeline secret exists and get API key
	_, err := r.ensurePipelineSecret(ctx, deployment)
	if err != nil {
		return fmt.Errorf("failed to ensure pipeline secret: %w", err)
	}

	// Create or update Pipelines PVC if persistence is enabled
	if deployment.Spec.OpenWebUI.Pipelines.Persistence.Enabled {
		pvc := r.buildPipelinesPVC(deployment)
		if err := r.ensurePVC(ctx, pvc); err != nil {
			return err
		}
	}

	// Create or update Pipelines deployment
	pipelinesDeployment := r.buildPipelinesDeployment(deployment)
	if err := r.createOrUpdateDeployment(ctx, pipelinesDeployment); err != nil {
		return err
	}

	// Create or update Pipelines service
	pipelinesService := r.buildPipelinesService(deployment)
	if err := r.createOrUpdateService(ctx, pipelinesService); err != nil {
		return err
	}

	return nil
}

// buildOpenWebUIPVC builds the OpenWebUI PVC object
func (r *LMDeploymentReconciler) buildOpenWebUIPVC(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.PersistentVolumeClaim {
	labels := map[string]string{
		"app":            "openwebui",
		"llm-deployment": deployment.Name,
	}

	// Set default size if not specified
	size := "1Gi"
	if deployment.Spec.OpenWebUI.Persistence != nil && deployment.Spec.OpenWebUI.Persistence.Size != "" {
		size = deployment.Spec.OpenWebUI.Persistence.Size
	}

	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetOpenWebUIPVCName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteOnce,
			},
			Resources: corev1.VolumeResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse(size),
				},
			},
		},
	}

	// Add storage class if specified
	if deployment.Spec.OpenWebUI.Persistence != nil && deployment.Spec.OpenWebUI.Persistence.StorageClass != "" {
		pvc.Spec.StorageClassName = &deployment.Spec.OpenWebUI.Persistence.StorageClass
	}

	// Note: We don't set controller reference on PVCs because they should persist
	// even if the LMDeployment is deleted to preserve user data
	return pvc
}

// buildOpenWebUIDeployment builds the OpenWebUI deployment object
func (r *LMDeploymentReconciler) buildOpenWebUIDeployment(deployment *llmgeeperiov1alpha1.LMDeployment) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "openwebui",
		"llm-deployment": deployment.Name,
	}

	ollamaServiceName := deployment.GetOllamaServiceName()

	// Build environment variables
	envVars := []corev1.EnvVar{
		{Name: "OLLAMA_BASE_URL", Value: fmt.Sprintf("http://%s:%d", ollamaServiceName, deployment.GetOllamaServicePort())},
		{Name: "ENABLE_VERSION_UPDATE_CHECK", Value: "False"},
		{
			Name: "WEBUI_SECRET_KEY",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: fmt.Sprintf("%s-openwebui-secret", deployment.Name),
					},
					Key: "WEBUI_SECRET_KEY",
				},
			},
		},
	}
	if deployment.Spec.OpenWebUI.Ingress.Host != "" {
		envVars = append(envVars, []corev1.EnvVar{
			// TODO: Use the correct protocol based on deployment type
			{Name: "WEBUI_HOST", Value: fmt.Sprintf("http://%s", deployment.Spec.OpenWebUI.Ingress.Host)},
			{Name: "CORS_ALLOW_ORIGIN", Value: fmt.Sprintf("http://%s;https://%s", deployment.Spec.OpenWebUI.Ingress.Host, deployment.Spec.OpenWebUI.Ingress.Host)},
		}...)
	}

	// Add Redis environment variables if Redis is enabled
	if deployment.Spec.OpenWebUI.Redis.Enabled {
		var redisURL string
		if deployment.Spec.OpenWebUI.Redis.RedisURL != "" {
			// Use external Redis URL
			redisURL = deployment.Spec.OpenWebUI.Redis.RedisURL
		} else {
			// Use internal Redis service
			redisURL = fmt.Sprintf("redis://:%s@%s:%d/0",
				deployment.Spec.OpenWebUI.Redis.Password,
				deployment.GetRedisServiceName(),
				deployment.Spec.OpenWebUI.Redis.Service.Port)
		}

		// Add Redis environment variables for OpenWebUI
		envVars = append(envVars, []corev1.EnvVar{
			{Name: "REDIS_URL", Value: redisURL},
		}...)
	}

	// Add custom environment variables if specified
	if len(deployment.Spec.OpenWebUI.EnvVars) > 0 {
		envVars = append(envVars, deployment.Spec.OpenWebUI.EnvVars...)
	}

	// Build volume mounts and volumes for config
	var volumeMounts []corev1.VolumeMount
	var volumes []corev1.Volume

	// Mount config file to temporary directory
	volumeMounts = append(volumeMounts, []corev1.VolumeMount{
		{
			Name:      "openwebui-config",
			MountPath: "/tmp/config",
			ReadOnly:  true,
		},
		{
			Name:      "openwebui-data",
			MountPath: "/app/backend/data",
		},
	}...)
	volumes = append(volumes, corev1.Volume{
		Name: "openwebui-config",
		VolumeSource: corev1.VolumeSource{
			Secret: &corev1.SecretVolumeSource{
				SecretName: deployment.GetOpenWebUIConfigName(),
				Items: []corev1.KeyToPath{
					{
						Key:  "config.json",
						Path: "config.json",
					},
				},
			},
		},
	})

	// Always add data volume mount and volume
	// If persistence is enabled, use PVC; otherwise use emptyDir
	if deployment.Spec.OpenWebUI.Persistence != nil && deployment.Spec.OpenWebUI.Persistence.Enabled {
		// Use PVC for persistent storage
		volumes = append(volumes, corev1.Volume{
			Name: "openwebui-data",
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: deployment.GetOpenWebUIPVCName(),
				},
			},
		})
	} else {
		// Use emptyDir for temporary storage when persistence is disabled
		volumes = append(volumes, corev1.Volume{
			Name: "openwebui-data",
			VolumeSource: corev1.VolumeSource{
				EmptyDir: &corev1.EmptyDirVolumeSource{},
			},
		})
	}

	// Build init container to copy config file
	initContainer := corev1.Container{
		Name:    "copy-config",
		Image:   "busybox:1.35",
		Command: []string{"/bin/sh", "-c"},
		Args: []string{
			"mkdir -p /app/backend/data && cp /tmp/config/config.json /app/backend/data/config.json && echo 'Config file copied successfully'",
		},
		VolumeMounts: volumeMounts,
	}

	// Build container
	container := corev1.Container{
		Name:  "openwebui",
		Image: deployment.Spec.OpenWebUI.Image,
		Ports: []corev1.ContainerPort{
			{
				Name:          "http",
				ContainerPort: deployment.Spec.OpenWebUI.Service.Port,
				Protocol:      corev1.ProtocolTCP,
			},
		},
		Resources:    r.buildResourceRequirements(deployment.Spec.OpenWebUI.Resources),
		Env:          envVars,
		VolumeMounts: volumeMounts,
	}

	openwebuiDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetOpenWebUIDeploymentName(),
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
					Affinity:       deployment.Spec.OpenWebUI.Affinity,
					InitContainers: []corev1.Container{initContainer},
					Containers:     []corev1.Container{container},
					Volumes:        volumes,
				},
			},
		},
	}

	// Set owner reference
	_ = controllerutil.SetControllerReference(deployment, openwebuiDeployment, r.Scheme)
	return openwebuiDeployment
}

// buildPipelinesDeployment builds the Pipelines deployment object
func (r *LMDeploymentReconciler) buildPipelinesDeployment(deployment *llmgeeperiov1alpha1.LMDeployment) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "pipelines",
		"llm-deployment": deployment.Name,
	}

	pipelinesSpec := deployment.Spec.OpenWebUI.Pipelines

	// Set default values
	image := pipelinesSpec.Image
	if image == "" {
		image = "ghcr.io/open-webui/pipelines:main"
	}

	port := pipelinesSpec.Port
	if port == 0 {
		port = 9099
	}

	replicas := pipelinesSpec.Replicas
	if replicas == 0 {
		replicas = 1
	}

	// Build environment variables
	envVars := []corev1.EnvVar{
		{
			Name:  "PIPELINES_DIR",
			Value: pipelinesSpec.PipelinesDir,
		},
	}

	// Add PIPELINES_API_KEY if pipelines are enabled
	if deployment.Spec.OpenWebUI.Pipelines != nil && deployment.Spec.OpenWebUI.Pipelines.Enabled {
		// Read PIPELINES_API_KEY from the pipelines secret
		secretName := fmt.Sprintf("%s-pipelines-secret", deployment.Name)
		envVars = append(envVars, corev1.EnvVar{
			Name: "PIPELINES_API_KEY",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: secretName,
					},
					Key: "PIPELINES_API_KEY",
				},
			},
		})
	}

	// Add Langfuse environment variables if Langfuse is enabled
	if deployment.Spec.OpenWebUI.Langfuse != nil && deployment.Spec.OpenWebUI.Langfuse.Enabled {
		langfuseSpec := deployment.Spec.OpenWebUI.Langfuse

		// Determine Langfuse URL - use self-hosted service if no external URL
		langfuseURL := langfuseSpec.URL

		// Add Langfuse environment variables
		if langfuseURL != "" {
			envVars = append(envVars, corev1.EnvVar{
				Name:  "LANGFUSE_HOST",
				Value: langfuseURL,
			})
		}

		// Add Langfuse credentials from SecretRef if provided
		if langfuseSpec.SecretRef != nil {
			envVars = append(envVars, corev1.EnvVar{
				Name: "LANGFUSE_PUBLIC_KEY",
				ValueFrom: &corev1.EnvVarSource{
					SecretKeyRef: &corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: langfuseSpec.SecretRef.Name,
						},
						Key: "LANGFUSE_PUBLIC_KEY",
					},
				},
			})

			envVars = append(envVars, corev1.EnvVar{
				Name: "LANGFUSE_SECRET_KEY",
				ValueFrom: &corev1.EnvVarSource{
					SecretKeyRef: &corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: langfuseSpec.SecretRef.Name,
						},
						Key: "LANGFUSE_SECRET_KEY",
					},
				},
			})
		}

		if langfuseSpec.Debug {
			envVars = append(envVars, corev1.EnvVar{
				Name:  "LANGFUSE_DEBUG",
				Value: "true",
			})
		}
	}

	// Add custom environment variables
	if len(pipelinesSpec.EnvVars) > 0 {
		envVars = append(envVars, pipelinesSpec.EnvVars...)
	}

	// Build pipeline URLs list
	var allPipelineURLs []string

	// Add existing pipeline URLs
	if len(pipelinesSpec.PipelineURLs) > 0 {
		allPipelineURLs = append(allPipelineURLs, pipelinesSpec.PipelineURLs...)
	}

	// Automatically add Langfuse monitoring pipeline if Langfuse is enabled
	if deployment.Spec.OpenWebUI.Langfuse != nil && deployment.Spec.OpenWebUI.Langfuse.Enabled {
		langfusePipelineURL := "https://github.com/open-webui/pipelines/blob/main/examples/filters/langfuse_filter_pipeline.py"

		// Check if Langfuse pipeline is already in the list
		langfusePipelineExists := false
		for _, url := range allPipelineURLs {
			if url == langfusePipelineURL {
				langfusePipelineExists = true
				break
			}
		}

		// Add Langfuse pipeline if not already present
		if !langfusePipelineExists {
			allPipelineURLs = append(allPipelineURLs, langfusePipelineURL)
		}
	}

	// Add pipeline URLs environment variable if we have any
	if len(allPipelineURLs) > 0 {
		pipelineURLs := strings.Join(allPipelineURLs, ",")
		envVars = append(envVars, corev1.EnvVar{
			Name:  "PIPELINES_URLS",
			Value: pipelineURLs,
		})
	}

	// Build container
	container := corev1.Container{
		Name:  "pipelines",
		Image: image,
		Ports: []corev1.ContainerPort{
			{
				ContainerPort: port,
				Protocol:      corev1.ProtocolTCP,
			},
		},
		Env: envVars,
	}

	// Add resource requirements if specified
	if pipelinesSpec.Resources.Requests != nil || pipelinesSpec.Resources.Limits != nil {
		container.Resources = corev1.ResourceRequirements{
			Requests: pipelinesSpec.Resources.Requests,
			Limits:   pipelinesSpec.Resources.Limits,
		}
	}

	// Add volume mounts and volumes if specified
	if len(pipelinesSpec.VolumeMounts) > 0 {
		container.VolumeMounts = pipelinesSpec.VolumeMounts
	}

	// Build volumes
	volumes := []corev1.Volume{}

	// Add custom volumes if specified
	if len(pipelinesSpec.Volumes) > 0 {
		volumes = append(volumes, pipelinesSpec.Volumes...)
	}

	// Add persistence volume if enabled
	if pipelinesSpec.Persistence != nil && pipelinesSpec.Persistence.Enabled {
		volumeName := fmt.Sprintf("%s-pipelines-data", deployment.Name)
		volumes = append(volumes, corev1.Volume{
			Name: volumeName,
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: volumeName,
				},
			},
		})

		// Add volume mount for persistence
		container.VolumeMounts = append(container.VolumeMounts, corev1.VolumeMount{
			Name:      volumeName,
			MountPath: "/app/pipelines",
		})
	}

	// Build deployment
	deploymentObj := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetPipelinesDeploymentName(),
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
				},
			},
		},
	}

	// Set owner reference
	_ = controllerutil.SetControllerReference(deployment, deploymentObj, r.Scheme)
	return deploymentObj
}

// buildPipelinesService builds the Pipelines service object
func (r *LMDeploymentReconciler) buildPipelinesService(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.Service {
	labels := map[string]string{
		"app":            "pipelines",
		"llm-deployment": deployment.Name,
	}

	pipelinesSpec := deployment.Spec.OpenWebUI.Pipelines

	port := pipelinesSpec.Port
	if port == 0 {
		port = 9099
	}

	serviceType := pipelinesSpec.Service.Type
	if serviceType == "" {
		serviceType = corev1.ServiceTypeClusterIP
	}

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetPipelinesServiceName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type:     serviceType,
			Selector: labels,
			Ports: []corev1.ServicePort{
				{
					Port:       port,
					TargetPort: intstr.FromInt32(port),
					Protocol:   corev1.ProtocolTCP,
				},
			},
		},
	}

	// Set owner reference
	_ = controllerutil.SetControllerReference(deployment, service, r.Scheme)
	return service
}

// buildPipelinesPVC builds the Pipelines PVC object
func (r *LMDeploymentReconciler) buildPipelinesPVC(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.PersistentVolumeClaim {
	labels := map[string]string{
		"app":            "pipelines",
		"llm-deployment": deployment.Name,
	}

	pipelinesSpec := deployment.Spec.OpenWebUI.Pipelines
	persistenceSpec := pipelinesSpec.Persistence

	// Set default values
	size := persistenceSpec.Size
	if size == "" {
		size = "10Gi"
	}

	storageClass := persistenceSpec.StorageClass

	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-pipelines-data", deployment.Name),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteOnce,
			},
			Resources: corev1.VolumeResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse(size),
				},
			},
		},
	}

	// Add storage class if specified
	if storageClass != "" {
		pvc.Spec.StorageClassName = &storageClass
	}

	// Note: We don't set controller reference on PVCs because they should persist
	// even if the LMDeployment is deleted to preserve user data
	return pvc
}

// buildOpenWebUIService builds the OpenWebUI service object
func (r *LMDeploymentReconciler) buildOpenWebUIService(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.Service {
	labels := map[string]string{
		"app":            "openwebui",
		"llm-deployment": deployment.Name,
	}

	openwebuiService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetOpenWebUIServiceName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: deployment.Spec.OpenWebUI.Service.Type,
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       deployment.Spec.OpenWebUI.Service.Port,
					TargetPort: intstr.FromInt32(deployment.Spec.OpenWebUI.Service.Port),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Selector: labels,
		},
	}

	// Set owner reference
	_ = controllerutil.SetControllerReference(deployment, openwebuiService, r.Scheme)
	return openwebuiService
}

// buildOpenWebUIIngress builds the OpenWebUI ingress object
func (r *LMDeploymentReconciler) buildOpenWebUIIngress(deployment *llmgeeperiov1alpha1.LMDeployment) *networkingv1.Ingress {
	labels := map[string]string{
		"app":            "openwebui",
		"llm-deployment": deployment.Name,
	}

	serviceName := deployment.GetOpenWebUIServiceName()
	pathType := networkingv1.PathTypePrefix

	ingress := &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetOpenWebUIIngressName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: deployment.Spec.OpenWebUI.Ingress.Host,
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
												Number: deployment.Spec.OpenWebUI.Service.Port,
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
	_ = controllerutil.SetControllerReference(deployment, ingress, r.Scheme)
	return ingress
}

// buildOpenWebUISecret builds the OpenWebUI secret for WEBUI_SECRET_KEY
func (r *LMDeploymentReconciler) buildOpenWebUISecret(deployment *llmgeeperiov1alpha1.LMDeployment) (*corev1.Secret, error) {
	// Generate a secure random secret key
	secretKey, err := generateSecureSecret(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate secure secret key: %w", err)
	}

	secret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-openwebui-secret", deployment.Name),
			Namespace: deployment.Namespace,
			Labels: map[string]string{
				"app":            "openwebui",
				"llm-deployment": deployment.Name,
			},
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			"WEBUI_SECRET_KEY": []byte(secretKey),
		},
	}

	_ = controllerutil.SetControllerReference(deployment, secret, r.Scheme)
	return secret, nil
}

// buildOpenWebUIConfig creates the config.json Secret for OpenWebUI
func (r *LMDeploymentReconciler) buildOpenWebUIConfig(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) (*corev1.Secret, error) {
	// Base configuration based on the provided config file
	config := map[string]interface{}{
		"version": 0,
		"ui": map[string]interface{}{
			"enable_signup": true,
		},
		"openai": map[string]interface{}{
			"enable":        true,
			"api_base_urls": []string{},
			"api_keys":      []string{},
			"api_configs":   map[string]interface{}{},
		},
	}

	// Add pipeline configuration if enabled
	if deployment.Spec.OpenWebUI.Pipelines != nil && deployment.Spec.OpenWebUI.Pipelines.Enabled {
		// Get API key from the existing pipelines secret
		apiKey, err := r.ensurePipelineSecret(ctx, deployment)
		if err != nil {
			return nil, fmt.Errorf("failed to ensure pipeline secret: %w", err)
		}

		// Add pipeline to OpenAI api_base_urls and api_configs
		openaiConfig := config["openai"].(map[string]interface{})
		openaiConfig["api_base_urls"] = []string{
			fmt.Sprintf("http://%s:9099", deployment.GetPipelinesServiceName()),
		}
		openaiConfig["api_configs"] = map[string]interface{}{
			"0": map[string]interface{}{
				"enable":          true,
				"tags":            []string{},
				"prefix_id":       "",
				"model_ids":       []string{},
				"connection_type": "local",
			},
		}
		openaiConfig["api_keys"] = []string{
			apiKey,
		}
		config["openai"] = openaiConfig
	}

	// Convert config to JSON
	configJSON, err := json.Marshal(config)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal config to JSON: %w", err)
	}

	openwebuiConfig := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetOpenWebUIConfigName(),
			Namespace: deployment.Namespace,
			Labels: map[string]string{
				"app":            "openwebui",
				"llm-deployment": deployment.Name,
			},
		},
		Type: corev1.SecretTypeOpaque,
		Data: map[string][]byte{
			"config.json": configJSON,
		},
	}
	_ = controllerutil.SetControllerReference(deployment, openwebuiConfig, r.Scheme)
	return openwebuiConfig, nil
}

// createOrUpdateSecret creates or updates a Kubernetes Secret using patch helper to avoid unnecessary reconciliations
func (r *LMDeploymentReconciler) createOrUpdateSecret(ctx context.Context, secret *corev1.Secret) error {
	logger := log.FromContext(ctx)
	existingSecret := &corev1.Secret{}
	err := r.Get(ctx, client.ObjectKeyFromObject(secret), existingSecret)
	if err != nil {
		if errors.IsNotFound(err) {
			// Secret doesn't exist, create it
			if err := r.Create(ctx, secret); err != nil {
				return fmt.Errorf("failed to create secret %s: %w", secret.Name, err)
			}
			logger.Info("Created OpenWebUI secret", "name", secret.Name, "namespace", secret.Namespace)
		} else {
			return fmt.Errorf("failed to get secret %s: %w", secret.Name, err)
		}
	} else {
		// Secret exists, use patch helper to only update if changes exist
		patchHelper, err := patch.NewHelper(existingSecret, r.Client)
		if err != nil {
			return fmt.Errorf("failed to create patch helper for secret %s: %w", secret.Name, err)
		}

		// Update the existing secret with new data and labels
		existingSecret.Data = secret.Data
		existingSecret.Labels = secret.Labels

		// Use patch helper to update - this only patches if changes exist
		if err := patchHelper.Patch(ctx, existingSecret); err != nil {
			return fmt.Errorf("failed to patch secret %s: %w", secret.Name, err)
		}
		// Note: PatchHelper only applies patches when there are actual changes,
		// so we don't log here to avoid spam when no changes are made
	}
	return nil
}
