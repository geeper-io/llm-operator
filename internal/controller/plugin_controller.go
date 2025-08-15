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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

// reconcilePlugins reconciles all OpenWebUI plugins
func (r *OllamaDeploymentReconciler) reconcilePlugins(ctx context.Context, deployment *llmgeeperiov1alpha1.Deployment) error {
	for _, plugin := range deployment.Spec.OpenWebUI.Plugins {
		if !plugin.Enabled {
			continue
		}

		// Create or update plugin deployment
		pluginDeployment := r.buildPluginDeployment(deployment, &plugin)
		if err := r.createOrUpdateDeployment(ctx, pluginDeployment); err != nil {
			return err
		}

		// Create or update plugin service
		pluginService := r.buildPluginService(deployment, &plugin)
		if err := r.createOrUpdateService(ctx, pluginService); err != nil {
			return err
		}
	}
	return nil
}

// buildPluginDeployment builds a plugin deployment object
func (r *OllamaDeploymentReconciler) buildPluginDeployment(deployment *llmgeeperiov1alpha1.Deployment, plugin *llmgeeperiov1alpha1.OpenWebUIPlugin) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "openwebui-plugin",
		"plugin-name":    plugin.Name,
		"llm-deployment": deployment.Name,
	}

	// Set default image tag if not specified
	imageTag := plugin.ImageTag
	if imageTag == "" {
		imageTag = "latest"
	}

	// Set default replicas if not specified
	replicas := plugin.Replicas
	if replicas == 0 {
		replicas = 1
	}

	// Set default service type if not specified
	serviceType := plugin.ServiceType
	if serviceType == "" {
		serviceType = "ClusterIP"
	}

	// Build environment variables
	envVars := plugin.EnvVars

	// Add ConfigMap reference if specified
	if plugin.ConfigMapName != "" {
		envVars = append(envVars, corev1.EnvVar{
			Name: "PLUGIN_CONFIG",
			ValueFrom: &corev1.EnvVarSource{
				ConfigMapKeyRef: &corev1.ConfigMapKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: plugin.ConfigMapName,
					},
					Key: "config",
				},
			},
		})
	}

	// Add Secret reference if specified
	if plugin.SecretName != "" {
		envVars = append(envVars, corev1.EnvVar{
			Name: "PLUGIN_CREDENTIALS",
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: plugin.SecretName,
					},
					Key: "credentials",
				},
			},
		})
	}

	pluginDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-plugin-%s", deployment.Name, plugin.Name),
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
							Name:  plugin.Name,
							Image: fmt.Sprintf("%s:%s", plugin.Image, imageTag),
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: plugin.Port,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Resources:    r.buildResourceRequirements(plugin.Resources),
							Env:          envVars,
							VolumeMounts: plugin.VolumeMounts,
						},
					},
					Volumes: plugin.Volumes,
				},
			},
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, pluginDeployment, r.Scheme)
	return pluginDeployment
}

// buildPluginService builds a plugin service object
func (r *OllamaDeploymentReconciler) buildPluginService(deployment *llmgeeperiov1alpha1.Deployment, plugin *llmgeeperiov1alpha1.OpenWebUIPlugin) *corev1.Service {
	labels := map[string]string{
		"app":            "openwebui-plugin",
		"plugin-name":    plugin.Name,
		"llm-deployment": deployment.Name,
	}

	// Set default service type if not specified
	serviceType := plugin.ServiceType
	if serviceType == "" {
		serviceType = "ClusterIP"
	}

	pluginService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      fmt.Sprintf("%s-plugin-%s", deployment.Name, plugin.Name),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType(serviceType),
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       plugin.Port,
					TargetPort: intstr.FromInt32(plugin.Port),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Selector: labels,
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, pluginService, r.Scheme)
	return pluginService
}
