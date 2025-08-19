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
	"crypto/rand"
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

// Essential Langfuse environment variables
const (
	ENV_NEXTAUTH_URL        = "NEXTAUTH_URL"
	ENV_NEXT_PUBLIC_API_URL = "NEXT_PUBLIC_API_URL"
	ENV_NEXTAUTH_SECRET     = "NEXTAUTH_SECRET"
	ENV_SALT                = "SALT"
	ENV_ENCRYPTION_KEY      = "ENCRYPTION_KEY"
	ENV_NODE_ENV            = "NODE_ENV"
	ENV_PORT                = "PORT"
	ENV_HOSTNAME            = "HOSTNAME"
)

// Default values
const (
	DefaultLangfuseImage       = "langfuse/langfuse:latest"
	DefaultLangfuseWorkerImage = "langfuse/langfuse-worker:latest"
	DefaultLangfusePort        = 3000
	DEFAULT_NODE_ENV           = "production"
	DEFAULT_HOSTNAME           = "0.0.0.0"
)

// generateSecureSecret generates a cryptographically secure random secret
func generateSecureSecret(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}
	return string(bytes), nil
}

// reconcileLangfuse reconciles the Langfuse deployment
func (r *LMDeploymentReconciler) reconcileLangfuse(ctx context.Context, deployment *llmgeeperiov1alpha1.LMDeployment) error {
	langfuseSpec := deployment.Spec.OpenWebUI.Langfuse

	// If URL is provided, no need to deploy self-hosted Langfuse
	if langfuseSpec.URL != "" {
		return nil
	}

	// Create or update Langfuse secrets
	langfuseSecret, err := r.buildLangfuseSecrets(deployment)
	if err != nil {
		return fmt.Errorf("failed to build Langfuse secrets: %w", err)
	}

	if err := r.createSecretIfNotExists(ctx, langfuseSecret); err != nil {
		return fmt.Errorf("failed to create Langfuse secrets: %w", err)
	}

	// Create or update PVC if persistence is enabled
	if langfuseSpec.Deploy != nil &&
		langfuseSpec.Deploy.Persistence != nil &&
		langfuseSpec.Deploy.Persistence.Enabled {
		pvc := r.buildLangfusePVC(deployment)
		if err := r.createOrUpdatePVC(ctx, pvc); err != nil {
			return err
		}
	}

	// Create or update Langfuse web deployment
	webDeployment := r.buildLangfuseWebDeployment(deployment)
	if err := r.createOrUpdateDeployment(ctx, webDeployment); err != nil {
		return err
	}

	// Create or update Langfuse worker deployment
	workerDeployment := r.buildLangfuseWorkerDeployment(deployment)
	if err := r.createOrUpdateDeployment(ctx, workerDeployment); err != nil {
		return err
	}

	// Create or update Langfuse web service
	webService := r.buildLangfuseWebService(deployment)
	if err := r.createOrUpdateService(ctx, webService); err != nil {
		return err
	}

	// Create or update Ingress if enabled
	if langfuseSpec.Deploy != nil &&
		langfuseSpec.Deploy.Ingress != nil &&
		langfuseSpec.Deploy.Ingress.Host != "" {
		ingress := r.buildLangfuseIngress(deployment)
		if err := r.createOrUpdateIngress(ctx, ingress); err != nil {
			return err
		}
	}

	return nil
}

// buildLangfuseSecrets creates the necessary Kubernetes secrets for Langfuse
func (r *LMDeploymentReconciler) buildLangfuseSecrets(deployment *llmgeeperiov1alpha1.LMDeployment) (*corev1.Secret, error) {
	// Generate cryptographically secure secrets
	nextAuthSecret, err := generateSecureSecret(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate next auth secret: %w", err)
	}

	salt, err := generateSecureSecret(16)
	if err != nil {
		return nil, fmt.Errorf("failed to generate salt: %w", err)
	}

	encryptionKey, err := generateSecureSecret(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate encryption key: %w", err)
	}

	// Main Langfuse secret containing authentication keys
	langfuseSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetLangfuseSecretName(),
			Namespace: deployment.Namespace,
			Labels: map[string]string{
				"app":            "langfuse",
				"llm-deployment": deployment.Name,
			},
		},
		Type: corev1.SecretTypeOpaque,
		StringData: map[string]string{
			"nextauth-secret": nextAuthSecret,
			"salt":            salt,
			"encryption-key":  encryptionKey,
		},
	}
	controllerutil.SetControllerReference(deployment, langfuseSecret, r.Scheme)
	return langfuseSecret, nil
}

// buildLangfuseWebDeployment builds the Langfuse web deployment
func (r *LMDeploymentReconciler) buildLangfuseWebDeployment(deployment *llmgeeperiov1alpha1.LMDeployment) *appsv1.Deployment {
	langfuseSpec := deployment.Spec.OpenWebUI.Langfuse
	labels := map[string]string{
		"app.kubernetes.io/name":     "langfuse",
		"app.kubernetes.io/instance": deployment.Name,
		"app":                        "web",
	}

	// Set default values
	image := DefaultLangfuseImage

	if langfuseSpec.Deploy != nil {
		if langfuseSpec.Deploy.Image != "" {
			image = langfuseSpec.Deploy.Image
		}
		if langfuseSpec.Deploy.Replicas > 0 {
			langfuseSpec.Deploy.Replicas = 1
		}
		if langfuseSpec.Deploy.Port > 0 {
			langfuseSpec.Deploy.Port = DefaultLangfusePort
		}
	}

	// Build resource requirements
	var resources corev1.ResourceRequirements
	if langfuseSpec.Deploy != nil {
		resources = r.buildResourceRequirements(langfuseSpec.Deploy.Resources)
	}

	// Build volume mounts and volumes for persistence
	var volumeMounts []corev1.VolumeMount
	var volumes []corev1.Volume

	if langfuseSpec.Deploy != nil &&
		langfuseSpec.Deploy.Persistence != nil &&
		langfuseSpec.Deploy.Persistence.Enabled {
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

	// Build essential environment variables
	envVars := r.buildLangfuseWebEnvVars(deployment)

	// Add custom environment variables if specified
	if langfuseSpec.Deploy != nil &&
		len(langfuseSpec.Deploy.EnvVars) > 0 {
		envVars = append(envVars, langfuseSpec.Deploy.EnvVars...)
	}

	langfuseDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetLangfuseDeploymentName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &langfuseSpec.Deploy.Replicas,
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
							Name:            "langfuse-web",
							Image:           image,
							SecurityContext: &corev1.SecurityContext{},
							Ports: []corev1.ContainerPort{
								{
									Name:          "http",
									ContainerPort: langfuseSpec.Deploy.Port,
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

// buildLangfuseWorkerDeployment builds the Langfuse worker deployment
func (r *LMDeploymentReconciler) buildLangfuseWorkerDeployment(deployment *llmgeeperiov1alpha1.LMDeployment) *appsv1.Deployment {
	langfuseSpec := deployment.Spec.OpenWebUI.Langfuse
	labels := map[string]string{
		"app.kubernetes.io/name":     "langfuse",
		"app.kubernetes.io/instance": deployment.Name,
		"app":                        "worker",
	}

	// Set default values
	image := DefaultLangfuseWorkerImage

	if langfuseSpec.Deploy != nil {
		if langfuseSpec.Deploy.Image != "" {
			image = langfuseSpec.Deploy.Image
		}
		if langfuseSpec.Deploy.Replicas > 0 {
			langfuseSpec.Deploy.Replicas = 1
		}
	}

	// Build resource requirements
	var resources corev1.ResourceRequirements
	if langfuseSpec.Deploy != nil {
		resources = r.buildResourceRequirements(langfuseSpec.Deploy.Resources)
	}

	// Build essential environment variables for worker
	envVars := r.buildLangfuseWorkerEnvVars(deployment)

	// Add custom environment variables if specified
	if langfuseSpec.Deploy != nil && len(langfuseSpec.Deploy.EnvVars) > 0 {
		envVars = append(envVars, langfuseSpec.Deploy.EnvVars...)
	}

	workerDeployment := &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetLangfuseWorkerDeploymentName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: &langfuseSpec.Deploy.Replicas,
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
							Name:            "langfuse-worker",
							Image:           image,
							SecurityContext: &corev1.SecurityContext{},
							Resources:       resources,
							Env:             envVars,
						},
					},
				},
			},
		},
	}
	controllerutil.SetControllerReference(deployment, workerDeployment, r.Scheme)
	return workerDeployment
}

// buildLangfuseWebEnvVars builds essential environment variables for web deployment
func (r *LMDeploymentReconciler) buildLangfuseWebEnvVars(deployment *llmgeeperiov1alpha1.LMDeployment) []corev1.EnvVar {
	langfuseSpec := deployment.Spec.OpenWebUI.Langfuse

	// Get the host from ingress configuration
	host := "localhost"
	if langfuseSpec.Deploy != nil &&
		langfuseSpec.Deploy.Ingress != nil &&
		langfuseSpec.Deploy.Ingress.Host != "" {
		host = langfuseSpec.Deploy.Ingress.Host
	}

	envVars := []corev1.EnvVar{
		// Essential configuration
		{Name: ENV_NODE_ENV, Value: DEFAULT_NODE_ENV},
		{Name: ENV_HOSTNAME, Value: DEFAULT_HOSTNAME},

		// URLs from ingress host
		{Name: ENV_NEXTAUTH_URL, Value: fmt.Sprintf("http://%s", host)},
		{Name: ENV_NEXT_PUBLIC_API_URL, Value: fmt.Sprintf("http://%s", host)},

		// Secrets mounted from Kubernetes secrets
		{
			Name: ENV_NEXTAUTH_SECRET,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: deployment.GetLangfuseSecretName(),
					},
					Key: "nextauth-secret",
				},
			},
		},
		{
			Name: ENV_SALT,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: deployment.GetLangfuseSecretName(),
					},
					Key: "salt",
				},
			},
		},
		{
			Name: ENV_ENCRYPTION_KEY,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: deployment.GetLangfuseSecretName(),
					},
					Key: "encryption-key",
				},
			},
		},
	}

	return envVars
}

// buildLangfuseWorkerEnvVars builds essential environment variables for worker deployment
func (r *LMDeploymentReconciler) buildLangfuseWorkerEnvVars(deployment *llmgeeperiov1alpha1.LMDeployment) []corev1.EnvVar {
	langfuseSpec := deployment.Spec.OpenWebUI.Langfuse

	// Get the host from ingress configuration
	host := "localhost"
	if langfuseSpec.Deploy != nil &&
		langfuseSpec.Deploy.Ingress != nil &&
		langfuseSpec.Deploy.Ingress.Host != "" {
		host = langfuseSpec.Deploy.Ingress.Host
	}

	envVars := []corev1.EnvVar{
		// Essential configuration
		{Name: ENV_NODE_ENV, Value: DEFAULT_NODE_ENV},

		// URLs from ingress host
		{Name: ENV_NEXTAUTH_URL, Value: fmt.Sprintf("http://%s", host)},

		// Secrets mounted from Kubernetes secrets
		{
			Name: ENV_NEXTAUTH_SECRET,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: deployment.GetLangfuseSecretName(),
					},
					Key: "nextauth-secret",
				},
			},
		},
		{
			Name: ENV_SALT,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: deployment.GetLangfuseSecretName(),
					},
					Key: "salt",
				},
			},
		},
		{
			Name: ENV_ENCRYPTION_KEY,
			ValueFrom: &corev1.EnvVarSource{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: deployment.GetLangfuseSecretName(),
					},
					Key: "encryption-key",
				},
			},
		},
	}

	return envVars
}

// buildLangfuseWebService builds the Langfuse web service
func (r *LMDeploymentReconciler) buildLangfuseWebService(deployment *llmgeeperiov1alpha1.LMDeployment) *corev1.Service {
	labels := map[string]string{
		"helm.sh/chart":                "langfuse-1.4.1",
		"app.kubernetes.io/name":       "langfuse",
		"app.kubernetes.io/instance":   deployment.Name,
		"app.kubernetes.io/version":    "3.98.2",
		"app.kubernetes.io/managed-by": "Helm",
	}

	port := int32(DefaultLangfusePort)
	if deployment.Spec.OpenWebUI.Langfuse.Deploy != nil &&
		deployment.Spec.OpenWebUI.Langfuse.Deploy.Port > 0 {
		port = deployment.Spec.OpenWebUI.Langfuse.Deploy.Port
	}

	langfuseService := &corev1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      deployment.GetLangfuseServiceName(),
			Namespace: deployment.Namespace,
			Labels:    labels,
		},
		Spec: corev1.ServiceSpec{
			Type: corev1.ServiceTypeClusterIP,
			Ports: []corev1.ServicePort{
				{
					Port:       port,
					TargetPort: intstr.FromString("http"),
					Protocol:   corev1.ProtocolTCP,
					Name:       "http",
				},
			},
			Selector: map[string]string{
				"app.kubernetes.io/name":     "langfuse",
				"app.kubernetes.io/instance": deployment.Name,
				"app":                        "web",
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

	port := int32(DefaultLangfusePort)
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
