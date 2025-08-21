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
	"strings"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

// reconcileOllama reconciles the Ollama deployment
func (r *LMDeploymentReconciler) reconcileOllama(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
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

// buildOllamaDeployment builds the Ollama deployment object
func (r *LMDeploymentReconciler) buildOllamaDeployment(deployment *llmgeeperiov1alpha1.LMDeployment) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "ollama",
		"llm-deployment": deployment.Name,
	}

	// Build postStart hook for model pulling
	postStartCommands := make([]string, 0, len(deployment.Spec.Ollama.Models))
	for _, model := range deployment.Spec.Ollama.Models {
		// Extract model name and tag from "modelname:tag" format
		modelTag := "latest"
		modelName := model

		// Check if tag is specified (format: "modelname:tag")
		if strings.Contains(model, ":") {
			parts := strings.Split(model, ":")
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
			Name:      deployment.GetOllamaDeploymentName(),
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
							Image: deployment.Spec.Ollama.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: deployment.Spec.Ollama.Service.Port,
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
	_ = controllerutil.SetControllerReference(deployment, ollamaDeployment, r.Scheme)
	return ollamaDeployment
}

// buildOllamaService builds the Ollama service object
func (r *LMDeploymentReconciler) buildOllamaService(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.Service {
	labels := map[string]string{
		"app":            "ollama",
		"llm-deployment": deployment.Name,
	}

	ollamaService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetOllamaServiceName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: deployment.Spec.Ollama.Service.Type,
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       deployment.Spec.Ollama.Service.Port,
					TargetPort: intstr.FromInt32(deployment.Spec.Ollama.Service.Port),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Selector: labels,
		},
	}

	// Set owner reference
	_ = controllerutil.SetControllerReference(deployment, ollamaService, r.Scheme)
	return ollamaService
}
