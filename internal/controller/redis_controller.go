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
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

// reconcileRedis reconciles the Redis deployment for OpenWebUI
func (r *LMDeploymentReconciler) reconcileRedis(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
	// Only deploy Redis if OpenWebUI is enabled
	if !deployment.Spec.OpenWebUI.Enabled {
		return nil
	}

	// Auto-enable Redis if OpenWebUI has multiple replicas and Redis is not explicitly disabled
	if deployment.Spec.OpenWebUI.Replicas > 1 && !deployment.Spec.OpenWebUI.Redis.Enabled && deployment.Spec.OpenWebUI.Redis.RedisURL == "" {
		// Automatically enable Redis for multi-instance deployments
		deployment.Spec.OpenWebUI.Redis.Enabled = true
	}

	// Only deploy Redis if Redis is enabled
	if !deployment.Spec.OpenWebUI.Redis.Enabled {
		return nil
	}

	// If Redis URL is provided, don't deploy Redis
	if deployment.Spec.OpenWebUI.Redis.RedisURL != "" {
		return nil
	}

	// Create or update Redis PVC if persistence is enabled
	if deployment.Spec.OpenWebUI.Redis.Persistence.Enabled {
		redisPVC := r.buildRedisPVC(deployment)
		if err := r.createOrUpdatePVC(ctx, redisPVC); err != nil {
			return err
		}
	}

	// Create or update Redis deployment
	redisDeployment := r.buildRedisDeployment(deployment)
	if err := r.createOrUpdateDeployment(ctx, redisDeployment); err != nil {
		return err
	}

	// Create or update Redis service
	redisService := r.buildRedisService(deployment)
	if err := r.createOrUpdateService(ctx, redisService); err != nil {
		return err
	}

	return nil
}

// buildRedisDeployment builds the Redis deployment object
func (r *LMDeploymentReconciler) buildRedisDeployment(deployment *llmgeeperiov1alpha1.LMDeployment) *appsv1.Deployment {
	labels := map[string]string{
		"app":            "redis",
		"llm-deployment": deployment.Name,
	}

	// Build environment variables
	envVars := []corev1.EnvVar{
		{
			Name:  "REDIS_PASSWORD",
			Value: deployment.Spec.OpenWebUI.Redis.Password,
		},
	}

	// Build volumes and volume mounts
	volumes := []corev1.Volume{}
	volumeMounts := []corev1.VolumeMount{}

	// Add persistence volume if enabled
	if deployment.Spec.OpenWebUI.Redis.Persistence.Enabled {
		volumes = append(volumes, corev1.Volume{
			Name: "redis-data",
			VolumeSource: corev1.VolumeSource{
				PersistentVolumeClaim: &corev1.PersistentVolumeClaimVolumeSource{
					ClaimName: deployment.GetRedisPVCName(),
				},
			},
		})

		volumeMounts = append(volumeMounts, corev1.VolumeMount{
			Name:      "redis-data",
			MountPath: "/data",
		})
	}

	redisDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetRedisDeploymentName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: int32Ptr(1), // Redis should only have 1 replica
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
							Name:  "redis",
							Image: deployment.Spec.OpenWebUI.Redis.Image,
							Ports: []corev1.ContainerPort{
								{
									Name:          "redis",
									ContainerPort: deployment.Spec.OpenWebUI.Redis.Service.Port,
									Protocol:      corev1.ProtocolTCP,
								},
							},
							Resources:    r.buildResourceRequirements(deployment.Spec.OpenWebUI.Redis.Resources),
							Env:          envVars,
							VolumeMounts: volumeMounts,
							Command: []string{
								"redis-server",
								"--appendonly",
								"yes",
								"--requirepass",
								deployment.Spec.OpenWebUI.Redis.Password,
							},
						},
					},
					Volumes: volumes,
				},
			},
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, redisDeployment, r.Scheme)
	return redisDeployment
}

// buildRedisService builds the Redis service object
func (r *LMDeploymentReconciler) buildRedisService(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.Service {
	labels := map[string]string{
		"app":            "redis",
		"llm-deployment": deployment.Name,
	}

	redisService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetRedisServiceName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceType(deployment.Spec.OpenWebUI.Redis.Service.Type),
			Ports: []corev1.ServicePort{
				{
					Name:       "redis",
					Port:       deployment.Spec.OpenWebUI.Redis.Service.Port,
					TargetPort: intstr.FromInt32(deployment.Spec.OpenWebUI.Redis.Service.Port),
					Protocol:   corev1.ProtocolTCP,
				},
			},
			Selector: labels,
		},
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, redisService, r.Scheme)
	return redisService
}

// buildRedisPVC builds the Redis PVC object
func (r *LMDeploymentReconciler) buildRedisPVC(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.PersistentVolumeClaim {
	labels := map[string]string{
		"app":            "redis",
		"llm-deployment": deployment.Name,
	}

	storageClass := deployment.Spec.OpenWebUI.Redis.Persistence.StorageClass
	var pvcSpec corev1.PersistentVolumeClaimSpec

	if storageClass != "" {
		pvcSpec = corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteOnce,
			},
			StorageClassName: &storageClass,
			Resources: corev1.VolumeResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse(deployment.Spec.OpenWebUI.Redis.Persistence.Size),
				},
			},
		}
	} else {
		pvcSpec = corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteOnce,
			},
			Resources: corev1.VolumeResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse(deployment.Spec.OpenWebUI.Redis.Persistence.Size),
				},
			},
		}
	}

	redisPVC := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetRedisPVCName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: pvcSpec,
	}

	// Set owner reference
	controllerutil.SetControllerReference(deployment, redisPVC, r.Scheme)
	return redisPVC
}

// int32Ptr returns a pointer to an int32
func int32Ptr(i int32) *int32 {
	return &i
}
