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

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

// ReconcileOpenWebUI reconciles the OpenWebUI deployment
func (r *OllamaDeploymentReconciler) ReconcileOpenWebUI(ctx context.Context, deployment *llmgeeperiov1alpha1.Deployment) error {
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
	if deployment.Spec.OpenWebUI.IngressEnabled && deployment.Spec.OpenWebUI.IngressHost != "" {
		ingress := r.buildOpenWebUIIngress(deployment)
		if err := r.createOrUpdateIngress(ctx, ingress); err != nil {
			return err
		}
	}

	return nil
}

// buildOpenWebUIDeployment builds the OpenWebUI deployment object
func (r *OllamaDeploymentReconciler) buildOpenWebUIDeployment(deployment *llmgeeperiov1alpha1.Deployment) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "openwebui",
		"llm-deployment": deployment.Name,
	}

	ollamaServiceName := fmt.Sprintf("%s-ollama", deployment.Name)

	// Build environment variables
	envVars := []corev1.EnvVar{
		{
			Name:  "OLLAMA_BASE_URL",
			Value: fmt.Sprintf("http://%s:%d", ollamaServiceName, deployment.Spec.Ollama.ServicePort),
		},
		{
			Name:  "WEBUI_SECRET_KEY",
			Value: "your-secret-key-here",
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
							Resources:    r.buildResourceRequirements(deployment.Spec.OpenWebUI.Resources),
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
	controllerutil.SetControllerReference(deployment, openwebuiDeployment, r.Scheme)
	return openwebuiDeployment
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
					TargetPort: intstr.FromInt32(deployment.Spec.OpenWebUI.ServicePort),
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

// buildOpenWebUIConfigMap builds the OpenWebUI configuration ConfigMap
func (r *OllamaDeploymentReconciler) buildOpenWebUIConfigMap(deployment *llmgeeperiov1alpha1.Deployment) *corev1.ConfigMap {
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
				"url":       fmt.Sprintf("http://%s-plugin-%s:%d", deployment.Name, plugin.Name, plugin.Port),
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
