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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

// reconcileVLLM reconciles the vLLM deployment
func (r *LMDeploymentReconciler) reconcileVLLM(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
	// Create or update vLLM deployment
	vllmDeployment := r.buildVLLMDeployment(deployment)
	if err := r.createOrUpdateDeployment(ctx, vllmDeployment); err != nil {
		return err
	}

	// Create or update vLLM service
	vllmService := r.buildVLLMService(deployment)
	if err := r.createOrUpdateService(ctx, vllmService); err != nil {
		return err
	}

	// Create or update vLLM PVC if persistence is enabled
	if deployment.Spec.VLLM.Persistence != nil && deployment.Spec.VLLM.Persistence.Enabled {
		vllmPVC := r.buildVLLMPVC(deployment)
		if err := r.ensurePVC(ctx, vllmPVC); err != nil {
			return err
		}
	}

	return nil
}

// buildVLLMDeployment builds the vLLM deployment object
func (r *LMDeploymentReconciler) buildVLLMDeployment(deployment *llmgeeperiov1alpha1.LMDeployment) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "vllm",
		"llm-deployment": deployment.Name,
	}

	// Set default image if not specified
	image := deployment.Spec.VLLM.Image
	if image == "" {
		image = "vllm/vllm-openai:latest"
	}

	// Build container spec
	container := corev1.Container{
		Name:  "vllm",
		Image: image,
		Ports: []corev1.ContainerPort{
			{
				Name:          "http",
				ContainerPort: deployment.Spec.VLLM.Service.Port,
				Protocol:      corev1.ProtocolTCP,
			},
		},
		Command:   []string{"vllm", "serve", deployment.Spec.VLLM.Model},
		Resources: r.buildResourceRequirements(deployment.Spec.VLLM.Resources),
		Env: []corev1.EnvVar{
			{
				Name:  "HOST",
				Value: "0.0.0.0",
			},
			{
				Name:  "PORT",
				Value: fmt.Sprintf("%d", deployment.Spec.VLLM.Service.Port),
			},
		},
		VolumeMounts: []corev1.VolumeMount{
			{
				Name:      "vllm-data",
				MountPath: "/root/.cache/huggingface",
			},
		},
	}

	// Add custom environment variables
	if len(deployment.Spec.VLLM.EnvVars) > 0 {
		container.Env = append(container.Env, deployment.Spec.VLLM.EnvVars...)
	}

	// Add custom volume mounts
	if len(deployment.Spec.VLLM.VolumeMounts) > 0 {
		container.VolumeMounts = append(container.VolumeMounts, deployment.Spec.VLLM.VolumeMounts...)
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
	if len(deployment.Spec.VLLM.Volumes) > 0 {
		volumes = append(volumes, deployment.Spec.VLLM.Volumes...)
	}

	// Use PVC if persistence is enabled
	if deployment.Spec.VLLM.Persistence != nil && deployment.Spec.VLLM.Persistence.Enabled {
		volumes[0] = corev1.Volume{
			Name: "vllm-data",
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: deployment.GetVLLMPVCName(),
				},
			},
		}
	}

	vllmDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetVLLMDeploymentName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &deployment.Spec.VLLM.Replicas,
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
					Affinity:   deployment.Spec.VLLM.Affinity,
				},
			},
		},
	}

	// Set owner reference
	_ = controllerutil.SetControllerReference(deployment, vllmDeployment, r.Scheme)
	return vllmDeployment
}

// buildVLLMService builds the vLLM service object
func (r *LMDeploymentReconciler) buildVLLMService(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.Service {
	labels := map[string]string{
		"app":            "vllm",
		"llm-deployment": deployment.Name,
	}

	vllmService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetVLLMServiceName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: deployment.Spec.VLLM.Service.Type,
			Ports: []corev1.ServicePort{
				{
					Name:       "http",
					Port:       deployment.Spec.VLLM.Service.Port,
					TargetPort: intstr.FromInt32(deployment.Spec.VLLM.Service.Port),
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

// buildVLLMPVC builds the vLLM PVC object
func (r *LMDeploymentReconciler) buildVLLMPVC(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.PersistentVolumeClaim {
	labels := map[string]string{
		"app":            "vllm",
		"llm-deployment": deployment.Name,
	}

	storageClass := ""
	if deployment.Spec.VLLM.Persistence.StorageClass != "" {
		storageClass = deployment.Spec.VLLM.Persistence.StorageClass
	}

	size := "10Gi" // Default size
	if deployment.Spec.VLLM.Persistence.Size != "" {
		size = deployment.Spec.VLLM.Persistence.Size
	}

	vllmPVC := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetVLLMPVCName(),
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
