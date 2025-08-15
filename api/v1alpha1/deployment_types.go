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

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OllamaModel defines a model to be deployed with Ollama
// Format: "modelname:tag" (e.g., "llama2:7b", "mistral:7b", "codellama:13b")
type OllamaModel string

// OllamaSpec defines the desired state of Ollama deployment
type OllamaSpec struct {
	// Replicas is the number of Ollama pods to run
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10
	Replicas int32 `json:"replicas,omitempty"`

	// Image is the Ollama container image to use
	Image string `json:"image,omitempty"`

	// ImageTag is the Ollama image tag to use
	ImageTag string `json:"imageTag,omitempty"`

	// Resources defines the resource requirements for Ollama pods
	Resources ResourceRequirements `json:"resources,omitempty"`

	// Models is the list of models to deploy with Ollama
	Models []OllamaModel `json:"models"`

	// Service defines the service configuration for Ollama
	Service ServiceSpec `json:"service,omitempty"`
}

// ServiceSpec defines service configuration
type ServiceSpec struct {
	// Type is the type of service to expose
	// +kubebuilder:validation:Enum=ClusterIP;NodePort;LoadBalancer
	Type string `json:"type,omitempty"`

	// Port is the port to expose the service
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port,omitempty"`
}

// IngressSpec defines ingress configuration
type IngressSpec struct {
	// Enabled determines if an Ingress should be created
	Enabled bool `json:"enabled,omitempty"`

	// Host is the hostname for the Ingress
	Host string `json:"host,omitempty"`

	// Annotations are custom annotations for the Ingress
	Annotations map[string]string `json:"annotations,omitempty"`

	// TLS configuration for the Ingress
	TLS *networkingv1.IngressTLS `json:"tls,omitempty"`
}

// OpenWebUISpec defines the desired state of OpenWebUI deployment
type OpenWebUISpec struct {
	// Enabled determines if OpenWebUI should be deployed
	Enabled bool `json:"enabled,omitempty"`

	// Replicas is the number of OpenWebUI pods to run
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5
	Replicas int32 `json:"replicas,omitempty"`

	// Image is the OpenWebUI container image to use
	Image string `json:"image,omitempty"`

	// ImageTag is the OpenWebUI image tag to use
	ImageTag string `json:"imageTag,omitempty"`

	// Resources defines the resource requirements for OpenWebUI pods
	Resources ResourceRequirements `json:"resources,omitempty"`

	// Service defines the service configuration for OpenWebUI
	Service ServiceSpec `json:"service,omitempty"`

	// Ingress defines the ingress configuration for OpenWebUI
	Ingress IngressSpec `json:"ingress,omitempty"`

	// Plugins defines the list of plugins to deploy and configure
	Plugins []OpenWebUIPlugin `json:"plugins,omitempty"`
}

// TabbySpec defines the desired state of Tabby deployment
type TabbySpec struct {
	// Enabled determines if Tabby should be deployed
	Enabled bool `json:"enabled,omitempty"`

	// Replicas is the number of Tabby pods to run
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5
	Replicas int32 `json:"replicas,omitempty"`

	// Image is the Tabby container image to use
	Image string `json:"image,omitempty"`

	// ImageTag is the Tabby image tag to use
	ImageTag string `json:"imageTag,omitempty"`

	// Resources defines the resource requirements for Tabby pods
	Resources ResourceRequirements `json:"resources,omitempty"`

	// Service defines the service configuration for Tabby
	Service ServiceSpec `json:"service,omitempty"`

	// Ingress defines the ingress configuration for Tabby
	Ingress IngressSpec `json:"ingress,omitempty"`

	// OllamaServiceName is the name of the Ollama service to connect to
	// If not specified, it will default to the deployment's Ollama service
	OllamaServiceName string `json:"ollamaServiceName,omitempty"`

	// OllamaServicePort is the port of the Ollama service to connect to
	// If not specified, it will default to the deployment's Ollama service port
	OllamaServicePort int32 `json:"ollamaServicePort,omitempty"`

	// ModelName is the name of the Ollama model to use for code completion
	ModelName string `json:"modelName,omitempty"`

	// EnvVars defines environment variables for Tabby
	EnvVars []corev1.EnvVar `json:"envVars,omitempty"`

	// VolumeMounts defines volume mounts for Tabby
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`

	// Volumes defines volumes for Tabby
	Volumes []corev1.Volume `json:"volumes,omitempty"`
}

// OpenWebUIPlugin defines a plugin for OpenWebUI
type OpenWebUIPlugin struct {
	// Name is the unique name of the plugin
	Name string `json:"name"`

	// Enabled determines if this plugin should be deployed
	Enabled bool `json:"enabled,omitempty"`

	// Type is the type of plugin (e.g., "openapi", "custom")
	// +kubebuilder:validation:Enum=openapi;custom
	Type string `json:"type"`

	// Image is the container image for the plugin
	Image string `json:"image"`

	// ImageTag is the image tag for the plugin
	ImageTag string `json:"imageTag,omitempty"`

	// Replicas is the number of plugin pods to run
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3
	Replicas int32 `json:"replicas,omitempty"`

	// Port is the port the plugin service exposes
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port"`

	// Resources defines the resource requirements for plugin pods
	Resources ResourceRequirements `json:"resources,omitempty"`

	// ServiceType is the type of service to expose the plugin
	// +kubebuilder:validation:Enum=ClusterIP;NodePort;LoadBalancer
	ServiceType string `json:"serviceType,omitempty"`

	// EnvVars defines environment variables for the plugin
	EnvVars []corev1.EnvVar `json:"envVars,omitempty"`

	// VolumeMounts defines volume mounts for the plugin
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`

	// Volumes defines volumes for the plugin
	Volumes []corev1.Volume `json:"volumes,omitempty"`

	// ConfigMapName is the name of the ConfigMap containing plugin configuration
	ConfigMapName string `json:"configMapName,omitempty"`

	// SecretName is the name of the Secret containing plugin credentials
	SecretName string `json:"secretName,omitempty"`
}

// ResourceRequirements describes the compute resource requirements
type ResourceRequirements struct {
	// Limits describes the maximum amount of compute resources allowed
	Limits ResourceList `json:"limits,omitempty"`

	// Requests describes the minimum amount of compute resources required
	Requests ResourceList `json:"requests,omitempty"`
}

// ResourceList is a set of (resource name, quantity) pairs
type ResourceList struct {
	// CPU is the CPU resource (e.g., "100m", "2")
	CPU string `json:"cpu,omitempty"`

	// Memory is the memory resource (e.g., "128Mi", "2Gi")
	Memory string `json:"memory,omitempty"`

	// Storage is the storage resource (e.g., "1Gi", "100Gi")
	Storage string `json:"storage,omitempty"`
}

// DeploymentSpec defines the desired state of Deployment
type DeploymentSpec struct {
	// Ollama defines the Ollama deployment configuration
	Ollama OllamaSpec `json:"ollama"`

	// OpenWebUI defines the OpenWebUI deployment configuration
	OpenWebUI OpenWebUISpec `json:"openwebui,omitempty"`

	// Tabby defines the Tabby deployment configuration
	Tabby TabbySpec `json:"tabby,omitempty"`
}

// DeploymentStatus defines the observed state of Deployment
type DeploymentStatus struct {
	// Phase represents the current phase of the deployment
	Phase string `json:"phase,omitempty"`

	// Conditions represent the latest available observations of the deployment's current state
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// OllamaStatus represents the status of Ollama deployment
	OllamaStatus DeploymentComponentStatus `json:"ollamaStatus,omitempty"`

	// OpenWebUIStatus represents the status of OpenWebUI deployment
	OpenWebUIStatus DeploymentComponentStatus `json:"openwebuiStatus,omitempty"`

	// TabbyStatus represents the status of Tabby deployment
	TabbyStatus DeploymentComponentStatus `json:"tabbyStatus,omitempty"`

	// ReadyReplicas is the number of ready replicas
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// TotalReplicas is the total number of replicas
	TotalReplicas int32 `json:"totalReplicas,omitempty"`
}

// DeploymentComponentStatus represents the status of a deployment component
type DeploymentComponentStatus struct {
	// AvailableReplicas is the number of available replicas
	AvailableReplicas int32 `json:"availableReplicas,omitempty"`

	// ReadyReplicas is the number of ready replicas
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// UpdatedReplicas is the number of updated replicas
	UpdatedReplicas int32 `json:"updatedReplicas,omitempty"`

	// Conditions represent the latest available observations of the component's current state
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type="string",JSONPath=".status.phase"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.readyReplicas"
// +kubebuilder:printcolumn:name="Total",type="string",JSONPath=".status.totalReplicas"
// +kubebuilder:printcolumn:name="Age",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Namespaced,shortName=ollamadep

// Deployment is the Schema for the ollamadeployments API
type Deployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeploymentSpec   `json:"spec,omitempty"`
	Status DeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeploymentList contains a list of Deployment
type DeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Deployment `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Deployment{}, &DeploymentList{})
}
