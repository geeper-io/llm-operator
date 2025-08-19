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
	networkingv1 "k8s.io/api/networking/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

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

	langfuseDeployment := &appsv1.Deployment{
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
	controllerutil.SetControllerReference(deployment, langfuseDeployment, r.Scheme)
	return langfuseDeployment
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

	langfuseService := &corev1.Service{
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
	controllerutil.SetControllerReference(deployment, langfuseService, r.Scheme)
	return langfuseService
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

	controllerutil.SetControllerReference(deployment, pvc, r.Scheme)
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

	langfuseIngress := &networkingv1.Ingress{
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
	controllerutil.SetControllerReference(deployment, langfuseIngress, r.Scheme)
	return langfuseIngress
}
