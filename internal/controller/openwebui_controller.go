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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
	"k8s.io/apimachinery/pkg/api/resource"
)

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

	// Reconcile Langfuse if enabled
	if deployment.Spec.OpenWebUI.Langfuse != nil && deployment.Spec.OpenWebUI.Langfuse.Enabled {
		if err := r.reconcileLangfuse(ctx, deployment); err != nil {
			return err
		}
	}

	// Create or update OpenWebUI configuration ConfigMap if plugins are defined
	if len(deployment.Spec.OpenWebUI.Plugins) > 0 {
		openwebuiConfig := r.buildOpenWebUIConfigMap(deployment)
		if err := r.createOrUpdateConfigMap(ctx, openwebuiConfig); err != nil {
			return err
		}
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
	if deployment.Spec.OpenWebUI.Ingress.Enabled && deployment.Spec.OpenWebUI.Ingress.Host != "" {
		ingress := r.buildOpenWebUIIngress(deployment)
		if err := r.createOrUpdateIngress(ctx, ingress); err != nil {
			return err
		}
	}

	return nil
}

// reconcilePipelines reconciles the OpenWebUI Pipelines deployment
func (r *LMDeploymentReconciler) reconcilePipelines(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
	// Create or update PVC if persistence is enabled
	if deployment.Spec.OpenWebUI.Pipelines.Persistence != nil && deployment.Spec.OpenWebUI.Pipelines.Persistence.Enabled {
		pvc := r.buildPipelinesPVC(deployment)
		if err := r.createOrUpdatePVC(ctx, pvc); err != nil {
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

	// Create or update OpenWebUI configuration to include Pipelines connection
	openwebuiConfig := r.buildOpenWebUIConfigMap(deployment)
	if err := r.createOrUpdateConfigMap(ctx, openwebuiConfig); err != nil {
		return err
	}

	return nil
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
		{
			Name:  "OLLAMA_BASE_URL",
			Value: fmt.Sprintf("http://%s:%d", ollamaServiceName, deployment.GetOllamaServicePort()),
		},
		{
			Name:  "WEBUI_SECRET_KEY",
			Value: "your-secret-key-here",
		},
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
			{
				Name:  "REDIS_URL",
				Value: redisURL,
			},
		}...)
	}

	// Add Langfuse environment variables if enabled
	if deployment.Spec.OpenWebUI.Langfuse != nil && deployment.Spec.OpenWebUI.Langfuse.Enabled {
		langfuseSpec := deployment.Spec.OpenWebUI.Langfuse

		// Determine Langfuse URL - use self-hosted service if no external URL
		langfuseURL := langfuseSpec.URL
		if langfuseURL == "" && langfuseSpec.Deploy != nil {
			langfuseURL = fmt.Sprintf("http://%s:%d", deployment.GetLangfuseServiceName(), langfuseSpec.Deploy.Port)
		}

		// Add Langfuse environment variables
		if langfuseURL != "" {
			envVars = append(envVars, corev1.EnvVar{
				Name:  "LANGFUSE_HOST",
				Value: langfuseURL,
			})
		}

		if langfuseSpec.PublicKey != "" {
			envVars = append(envVars, corev1.EnvVar{
				Name:  "LANGFUSE_PUBLIC_KEY",
				Value: langfuseSpec.PublicKey,
			})
		}

		if langfuseSpec.SecretKey != "" {
			envVars = append(envVars, corev1.EnvVar{
				Name:  "LANGFUSE_SECRET_KEY",
				Value: langfuseSpec.SecretKey,
			})
		}

		if langfuseSpec.ProjectName != "" {
			envVars = append(envVars, corev1.EnvVar{
				Name:  "LANGFUSE_PROJECT",
				Value: langfuseSpec.ProjectName,
			})
		}

		if langfuseSpec.Environment != "" {
			envVars = append(envVars, corev1.EnvVar{
				Name:  "LANGFUSE_ENVIRONMENT",
				Value: langfuseSpec.Environment,
			})
		}

		if langfuseSpec.Debug {
			envVars = append(envVars, corev1.EnvVar{
				Name:  "LANGFUSE_DEBUG",
				Value: "true",
			})
		}
	}

	// Add custom environment variables if specified
	if len(deployment.Spec.OpenWebUI.EnvVars) > 0 {
		envVars = append(envVars, deployment.Spec.OpenWebUI.EnvVars...)
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
		Resources: r.buildResourceRequirements(deployment.Spec.OpenWebUI.Resources),
		Env:       envVars,
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      "openwebui-config",
				MountPath: "/app/backend/data",
				SubPath:   "config.json",
			},
		},
	}

	// Build volumes and volume mounts
	volumes := []corev1.Volume{}
	volumeMounts := []corev1.VolumeMount{}

	// Add plugin configuration volume if plugins are defined
	if len(deployment.Spec.OpenWebUI.Plugins) > 0 {
		// Add config volume
		volumes = append(volumes, corev1.Volume{
			Name: "openwebui-config",
			VolumeSource: corev1.VolumeSource{
				ConfigMap: &corev1.ConfigMapVolumeSource{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: fmt.Sprintf("%s-openwebui-config", deployment.Name),
					},
				},
			},
		})

		// Add config volume mount
		volumeMounts = append(volumeMounts, corev1.VolumeMount{
			Name:      "openwebui-config",
			MountPath: "/app/backend/data",
			SubPath:   "config.json",
		})
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
					Containers: []corev1.Container{container},
					Volumes:    volumes,
				},
			},
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, openwebuiDeployment, r.Scheme)
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

	serviceType := pipelinesSpec.ServiceType
	if serviceType == "" {
		serviceType = "ClusterIP"
	}

	// Build environment variables
	envVars := []corev1.EnvVar{
		{
			Name:  "PIPELINES_DIR",
			Value: pipelinesSpec.PipelinesDir,
		},
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
		langfusePipelineURL := "https://github.com/open-webui/pipelines/blob/main/examples/monitoring/langfuse_monitor_pipeline.py"

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
			Requests: corev1.ResourceList{},
			Limits:   corev1.ResourceList{},
		}

		if pipelinesSpec.Resources.Requests.CPU != "" {
			container.Resources.Requests[corev1.ResourceCPU] = resource.MustParse(pipelinesSpec.Resources.Requests.CPU)
		}
		if pipelinesSpec.Resources.Requests.Memory != "" {
			container.Resources.Requests[corev1.ResourceMemory] = resource.MustParse(pipelinesSpec.Resources.Requests.Memory)
		}
		if pipelinesSpec.Resources.Limits.CPU != "" {
			container.Resources.Limits[corev1.ResourceCPU] = resource.MustParse(pipelinesSpec.Resources.Limits.CPU)
		}
		if pipelinesSpec.Resources.Limits.Memory != "" {
			container.Resources.Limits[corev1.ResourceMemory] = resource.MustParse(pipelinesSpec.Resources.Limits.Memory)
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
	controllerutil.SetControllerReference(deployment, deploymentObj, r.Scheme)
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

	serviceType := pipelinesSpec.ServiceType
	if serviceType == "" {
		serviceType = "ClusterIP"
	}

	service := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetPipelinesServiceName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type:     corev1.ServiceType(serviceType),
			Selector: labels,
			Ports: []corev1.ServicePort{
				{
					Port:       port,
					TargetPort: intstr.FromInt(int(port)),
					Protocol:   corev1.ProtocolTCP,
				},
			},
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, service, r.Scheme)
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

	// Set owner reference
	controllerutil.SetControllerReference(deployment, pvc, r.Scheme)
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
			Type: corev1.ServiceType(deployment.Spec.OpenWebUI.Service.Type),
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
	controllerutil.SetControllerReference(deployment, openwebuiService, r.Scheme)
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
	controllerutil.SetControllerReference(deployment, ingress, r.Scheme)
	return ingress
}

// buildOpenWebUIConfigMap builds the OpenWebUI configuration ConfigMap
func (r *LMDeploymentReconciler) buildOpenWebUIConfigMap(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.ConfigMap {
	// Start with base configuration
	config := map[string]interface{}{
		"version": 0,
		"ui": map[string]interface{}{
			"enable_signup": false,
		},
		"openai": map[string]interface{}{
			"enable":        false,
			"api_base_urls": []string{"https://api.openai.com/v1"},
			"api_keys":      []string{""},
			"api_configs": map[string]interface{}{
				"0": map[string]interface{}{},
			},
		},
		"tool_server": map[string]interface{}{
			"connections": []map[string]interface{}{},
		},
	}

	// Add tool server connections for each plugin
	if len(deployment.Spec.OpenWebUI.Plugins) > 0 {
		connections := []map[string]interface{}{}

		for _, plugin := range deployment.Spec.OpenWebUI.Plugins {
			if !plugin.Enabled {
				continue
			}

			// Build connection configuration
			connection := map[string]interface{}{
				"url":       fmt.Sprintf("http://%s:%d", deployment.GetPluginServiceName(plugin.Name), plugin.Port),
				"path":      "openapi.json", // Default OpenAPI path
				"auth_type": "none",         // Default auth type
				"key":       "",             // Default empty key
				"config": map[string]interface{}{
					"enable": true,
					"access_control": map[string]interface{}{
						"read": map[string]interface{}{
							"group_ids": []string{},
							"user_ids":  []string{},
						},
						"write": map[string]interface{}{
							"group_ids": []string{},
							"user_ids":  []string{},
						},
					},
				},
				"info": map[string]interface{}{
					"name":        plugin.Name,
					"description": fmt.Sprintf("%s tool for %s", plugin.Type, deployment.Name),
				},
			}

			// Override defaults if ConfigMap is specified
			if plugin.ConfigMapName != "" {
				connection["path"] = "openapi.json" // Could be made configurable
			}

			// Override auth type if Secret is specified
			if plugin.SecretName != "" {
				connection["auth_type"] = "bearer" // Default to bearer when credentials are provided
			}

			connections = append(connections, connection)
		}

		config["tool_server"].(map[string]interface{})["connections"] = connections
	}

	// Add Pipelines connection if enabled
	if deployment.Spec.OpenWebUI.Pipelines != nil && deployment.Spec.OpenWebUI.Pipelines.Enabled {
		pipelinesConnections := []map[string]interface{}{}

		// Add Pipelines connection
		pipelinesConnection := map[string]interface{}{
			"url":       fmt.Sprintf("http://%s:%d", deployment.GetPipelinesServiceName(), deployment.Spec.OpenWebUI.Pipelines.Port),
			"path":      "openapi.json",
			"auth_type": "none",
			"key":       "",
			"config": map[string]interface{}{
				"enable": true,
				"access_control": map[string]interface{}{
					"read": map[string]interface{}{
						"group_ids": []string{},
						"user_ids":  []string{},
					},
					"write": map[string]interface{}{
						"group_ids": []string{},
						"user_ids":  []string{},
					},
				},
			},
			"info": map[string]interface{}{
				"name":        "pipelines",
				"description": fmt.Sprintf("OpenWebUI Pipelines for %s", deployment.Name),
			},
		}

		pipelinesConnections = append(pipelinesConnections, pipelinesConnection)

		// Add existing tool server connections if any
		if existingConnections, exists := config["tool_server"].(map[string]interface{})["connections"]; exists {
			if connections, ok := existingConnections.([]map[string]interface{}); ok {
				pipelinesConnections = append(pipelinesConnections, connections...)
			}
		}

		config["tool_server"].(map[string]interface{})["connections"] = pipelinesConnections
	}

	// Convert to JSON
	configJSON, _ := json.MarshalIndent(config, "", "  ")

	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-openwebui-config", deployment.Name),
			Namespace: deployment.Namespace,
			Labels: map[string]string{
				"app":            "openwebui",
				"llm-deployment": deployment.Name,
			},
		},
		Data: map[string]string{
			"config.json": string(configJSON),
		},
	}
}

// reconcileLangfuse reconciles the Langfuse deployment
func (r *LMDeploymentReconciler) reconcileLangfuse(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
	// If URL is provided, no need to deploy self-hosted Langfuse
	if deployment.Spec.OpenWebUI.Langfuse.URL != "" {
		return nil
	}

	// Create or update PVC if persistence is enabled
	if deployment.Spec.OpenWebUI.Langfuse.Deploy != nil &&
		deployment.Spec.OpenWebUI.Langfuse.Deploy.Persistence != nil &&
		deployment.Spec.OpenWebUI.Langfuse.Deploy.Persistence.Enabled {
		pvc := r.buildLangfusePVC(deployment)
		if err := r.createOrUpdatePVC(ctx, pvc); err != nil {
			return err
		}
	}

	// Create or update Langfuse deployment
	langfuseDeployment := r.buildLangfuseDeployment(deployment)
	if err := r.createOrUpdateDeployment(ctx, langfuseDeployment); err != nil {
		return err
	}

	// Create or update Langfuse service
	langfuseService := r.buildLangfuseService(deployment)
	if err := r.createOrUpdateService(ctx, langfuseService); err != nil {
		return err
	}

	// Create or update Ingress if enabled
	if deployment.Spec.OpenWebUI.Langfuse.Deploy != nil &&
		deployment.Spec.OpenWebUI.Langfuse.Deploy.Ingress != nil &&
		deployment.Spec.OpenWebUI.Langfuse.Deploy.Ingress.Enabled &&
		deployment.Spec.OpenWebUI.Langfuse.Deploy.Ingress.Host != "" {
		ingress := r.buildLangfuseIngress(deployment)
		if err := r.createOrUpdateIngress(ctx, ingress); err != nil {
			return err
		}
	}

	return nil
}

// buildLangfuseDeployment builds the Langfuse deployment
func (r *LMDeploymentReconciler) buildLangfuseDeployment(deployment *llmgeeperiov1alpha1.LMDeployment) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "langfuse",
		"llm-deployment": deployment.Name,
	}

	// Set default values
	image := "langfuse/langfuse:latest"
	replicas := int32(1)
	port := int32(3000)

	if deployment.Spec.OpenWebUI.Langfuse.Deploy != nil {
		if deployment.Spec.OpenWebUI.Langfuse.Deploy.Image != "" {
			image = deployment.Spec.OpenWebUI.Langfuse.Deploy.Image
		}
		if deployment.Spec.OpenWebUI.Langfuse.Deploy.Replicas > 0 {
			replicas = deployment.Spec.OpenWebUI.Langfuse.Deploy.Replicas
		}
		if deployment.Spec.OpenWebUI.Langfuse.Deploy.Port > 0 {
			port = deployment.Spec.OpenWebUI.Langfuse.Deploy.Port
		}
	}

	// Build resource requirements
	var resources corev1.ResourceRequirements
	if deployment.Spec.OpenWebUI.Langfuse.Deploy != nil {
		resources = r.buildResourceRequirements(deployment.Spec.OpenWebUI.Langfuse.Deploy.Resources)
	}

	// Build volume mounts and volumes for persistence
	var volumeMounts []corev1.VolumeMount
	var volumes []corev1.Volume

	if deployment.Spec.OpenWebUI.Langfuse.Deploy != nil &&
		deployment.Spec.OpenWebUI.Langfuse.Deploy.Persistence != nil &&
		deployment.Spec.OpenWebUI.Langfuse.Deploy.Persistence.Enabled {
		volumeMounts = append(volumeMounts, corev1.VolumeMount{
			Name:      "langfuse-data",
			MountPath: "/app/data",
		})
		volumes = append(volumes, corev1.Volume{
			Name: "langfuse-data",
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: deployment.GetLangfusePVCName(),
				},
			},
		})
	}

	// Build environment variables
	envVars := []corev1.EnvVar{
		{
			Name:  "LANGFUSE_SECRET",
			Value: deployment.Spec.OpenWebUI.Langfuse.SecretKey,
		},
		{
			Name:  "LANGFUSE_PUBLIC_KEY",
			Value: deployment.Spec.OpenWebUI.Langfuse.PublicKey,
		},
		{
			Name:  "LANGFUSE_PROJECT_NAME",
			Value: deployment.Spec.OpenWebUI.Langfuse.ProjectName,
		},
		{
			Name:  "LANGFUSE_ENVIRONMENT",
			Value: deployment.Spec.OpenWebUI.Langfuse.Environment,
		},
	}

	// Add custom environment variables if specified
	if deployment.Spec.OpenWebUI.Langfuse.Deploy != nil &&
		len(deployment.Spec.OpenWebUI.Langfuse.Deploy.EnvVars) > 0 {
		envVars = append(envVars, deployment.Spec.OpenWebUI.Langfuse.Deploy.EnvVars...)
	}

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetLangfuseDeploymentName(),
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
							Name:  "langfuse",
							Image: image,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: port,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Resources:    resources,
							Env:          envVars,
							VolumeMounts: volumeMounts,
						},
					},
					Volumes: volumes,
				},
			},
		},
	}
}

// buildLangfuseService builds the Langfuse service
func (r *LMDeploymentReconciler) buildLangfuseService(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.Service {
	labels := map[string]string{
		"app":            "langfuse",
		"llm-deployment": deployment.Name,
	}

	// Set default values
	port := int32(3000)
	serviceType := corev1.ServiceTypeClusterIP

	if deployment.Spec.OpenWebUI.Langfuse.Deploy != nil {
		if deployment.Spec.OpenWebUI.Langfuse.Deploy.Port > 0 {
			port = deployment.Spec.OpenWebUI.Langfuse.Deploy.Port
		}
		if deployment.Spec.OpenWebUI.Langfuse.Deploy.ServiceType != "" {
			switch deployment.Spec.OpenWebUI.Langfuse.Deploy.ServiceType {
			case "NodePort":
				serviceType = corev1.ServiceTypeNodePort
			case "LoadBalancer":
				serviceType = corev1.ServiceTypeLoadBalancer
			default:
				serviceType = corev1.ServiceTypeClusterIP
			}
		}
	}

	return &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetLangfuseServiceName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type:     serviceType,
			Selector: labels,
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       port,
					TargetPort: intstr.FromInt32(port),
					Protocol:   corev1.ProtocolTCP,
				},
			},
		},
	}
}

// buildLangfusePVC builds the Langfuse PVC
func (r *LMDeploymentReconciler) buildLangfusePVC(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.PersistentVolumeClaim {
	labels := map[string]string{
		"app":            "langfuse",
		"llm-deployment": deployment.Name,
	}

	// Set default values
	size := "10Gi"
	storageClass := ""

	if deployment.Spec.OpenWebUI.Langfuse.Deploy != nil &&
		deployment.Spec.OpenWebUI.Langfuse.Deploy.Persistence != nil {
		if deployment.Spec.OpenWebUI.Langfuse.Deploy.Persistence.Size != "" {
			size = deployment.Spec.OpenWebUI.Langfuse.Deploy.Persistence.Size
		}
		if deployment.Spec.OpenWebUI.Langfuse.Deploy.Persistence.StorageClass != "" {
			storageClass = deployment.Spec.OpenWebUI.Langfuse.Deploy.Persistence.StorageClass
		}
	}

	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetLangfusePVCName(),
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

	if storageClass != "" {
		pvc.Spec.StorageClassName = &storageClass
	}

	return pvc
}

// buildLangfuseIngress builds the Langfuse ingress
func (r *LMDeploymentReconciler) buildLangfuseIngress(deployment *llmgeeperiov1alpha1.LMDeployment) *networkingv1.Ingress {
	labels := map[string]string{
		"app":            "langfuse",
		"llm-deployment": deployment.Name,
	}

	port := int32(3000)
	if deployment.Spec.OpenWebUI.Langfuse.Deploy != nil &&
		deployment.Spec.OpenWebUI.Langfuse.Deploy.Port > 0 {
		port = deployment.Spec.OpenWebUI.Langfuse.Deploy.Port
	}

	return &networkingv1.Ingress{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-langfuse-ingress", deployment.Name),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: networkingv1.IngressSpec{
			Rules: []networkingv1.IngressRule{
				{
					Host: deployment.Spec.OpenWebUI.Langfuse.Deploy.Ingress.Host,
					IngressRuleValue: networkingv1.IngressRuleValue{
						HTTP: &networkingv1.HTTPIngressRuleValue{
							Paths: []networkingv1.HTTPIngressPath{
								{
									Path:     "/",
									PathType: &[]networkingv1.PathType{networkingv1.PathTypePrefix}[0],
									Backend: networkingv1.IngressBackend{
										Service: &networkingv1.IngressServiceBackend{
											Name: deployment.GetLangfuseServiceName(),
											Port: networkingv1.ServiceBackendPort{
												Number: port,
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
}
