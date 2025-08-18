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

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

// reconcileTools reconciles all OpenWebUI tools
func (r *LMDeploymentReconciler) reconcileTools(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
	for _, tool := range deployment.Spec.OpenWebUI.Tools {
		if !tool.Enabled {
			continue
		}

		// Create or update tool deployment
		toolDeployment := r.buildToolDeployment(deployment, &tool)
		if err := r.createOrUpdateDeployment(ctx, toolDeployment); err != nil {
			return err
		}

		// Create or update tool service
		toolService := r.buildToolService(deployment, &tool)
		if err := r.createOrUpdateService(ctx, toolService); err != nil {
			return err
		}
	}
	return nil
}

// buildToolDeployment builds a tool deployment object
func (r *LMDeploymentReconciler) buildToolDeployment(deployment *llmgeeperiov1alpha1.LMDeployment, tool *llmgeeperiov1alpha1.OpenWebUITool) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "openwebui-tool",
		"tool-name":      tool.Name,
		"llm-deployment": deployment.Name,
	}

	// Build tool deployment
	image := tool.Image
	if image == "" {
		image = "tool:latest"
	}

	toolDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetToolDeploymentName(tool.Name),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &tool.Replicas,
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
							Name:  "tool",
							Image: image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: tool.Port,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Resources:    r.buildResourceRequirements(tool.Resources),
							Env:          tool.EnvVars,
							VolumeMounts: tool.VolumeMounts,
						},
					},
					Volumes: tool.Volumes,
				},
			},
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, toolDeployment, r.Scheme)
	return toolDeployment
}

// buildToolService builds a tool service object
func (r *LMDeploymentReconciler) buildToolService(deployment *llmgeeperiov1alpha1.LMDeployment, tool *llmgeeperiov1alpha1.OpenWebUITool) *corev1.Service {
	labels := map[string]string{
		"app":            "openwebui-tool",
		"tool-name":      tool.Name,
		"llm-deployment": deployment.Name,
	}

	// Set default service type if not specified
	serviceType := tool.ServiceType
	if serviceType == "" {
		serviceType = "ClusterIP"
	}

	toolService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetToolServiceName(tool.Name),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType(serviceType),
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       tool.Port,
					TargetPort: intstr.FromInt32(tool.Port),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Selector: labels,
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, toolService, r.Scheme)
	return toolService
}
