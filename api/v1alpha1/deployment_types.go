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
	"fmt"

	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// OllamaSpec defines the desired state of Ollama deployment
type OllamaSpec struct {
	// Replicas is the number of Ollama pods to run
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=10
	Replicas int32 `json:"replicas,omitempty"`

	// Image is the Ollama container image to use (including tag)
	Image string `json:"image,omitempty"`

	// Resources defines the resource requirements for Ollama pods
	Resources ResourceRequirements `json:"resources,omitempty"`

	// Models is the list of models to deploy with Ollama
	Models []string `json:"models"`

	// Service defines the service configuration for Ollama
	Service ServiceSpec `json:"service,omitempty"`
}

// ServiceSpec defines service configuration
type ServiceSpec struct {
	// Type is the type of service to expose
	// +kubebuilder:validation:Enum=ClusterIP;NodePort;LoadBalancer
	Type corev1.ServiceType `json:"type,omitempty"`

	// Port is the port to expose the service
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port,omitempty"`
}

// IngressSpec defines ingress configuration
type IngressSpec struct {
	// Host is the hostname for the Ingress
	Host string `json:"host,omitempty"`

	// Annotations are custom annotations for the Ingress
	Annotations map[string]string `json:"annotations,omitempty"`
}

// OpenWebUISpec defines the desired state of OpenWebUI deployment
type OpenWebUISpec struct {
	// Enabled determines if OpenWebUI should be deployed
	Enabled bool `json:"enabled,omitempty"`

	// Replicas is the number of OpenWebUI pods to run
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5
	Replicas int32 `json:"replicas,omitempty"`

	// Image is the OpenWebUI container image to use (including tag)
	Image string `json:"image,omitempty"`

	// EnvVars defines environment variables for the Pipelines
	EnvVars []corev1.EnvVar `json:"envVars,omitempty"`

	// Resources defines the resource requirements for OpenWebUI pods
	Resources ResourceRequirements `json:"resources,omitempty"`

	// Service defines the service configuration for OpenWebUI
	Service ServiceSpec `json:"service,omitempty"`

	// Ingress defines the ingress configuration for OpenWebUI
	Ingress IngressSpec `json:"ingress,omitempty"`

	// Redis defines the Redis configuration for OpenWebUI
	Redis RedisSpec `json:"redis,omitempty"`

	// Pipelines defines the OpenWebUI Pipelines configuration
	Pipelines *PipelinesSpec `json:"pipelines,omitempty"`

	// Langfuse defines the Langfuse monitoring configuration for OpenWebUI
	Langfuse *LangfuseSpec `json:"langfuse,omitempty"`

	// Persistence defines OpenWebUI persistence configuration
	Persistence *OpenWebUIPersistenceSpec `json:"persistence,omitempty"`
}

// TabbySpec defines the desired state of Tabby deployment
type TabbySpec struct {
	// Enabled determines if Tabby should be deployed
	Enabled bool `json:"enabled,omitempty"`

	// Replicas is the number of Tabby pods to run
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=5
	Replicas int32 `json:"replicas,omitempty"`

	// Image is the Tabby container image to use (including tag)
	Image string `json:"image,omitempty"`

	// Device specifies the device type for Tabby,
	// +kubebuilder:validation:Optional
	// +kubebuilder:validation:Enum=cpu;cuda;rocm;metal;vulkan
	// +kubebuilder:default=cpu
	Device string `json:"device,omitempty"`

	// Resources defines the resource requirements for Tabby pods
	Resources ResourceRequirements `json:"resources,omitempty"`

	// Service defines the service configuration for Tabby
	Service ServiceSpec `json:"service,omitempty"`

	// Ingress defines the ingress configuration for Tabby
	Ingress IngressSpec `json:"ingress,omitempty"`

	// ChatModel is the name of the Ollama model to use for chat functionality
	// Must be one of the models specified in spec.ollama.models
	ChatModel string `json:"chatModel,omitempty"`

	// CompletionModel is the name of the Ollama model to use for code completion
	// Must be one of the models specified in spec.ollama.models
	CompletionModel string `json:"completionModel,omitempty"`

	// EnvVars defines environment variables for Tabby
	EnvVars []corev1.EnvVar `json:"envVars,omitempty"`

	// VolumeMounts defines volume mounts for Tabby
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`

	// Volumes defines volumes for Tabby
	Volumes []corev1.Volume `json:"volumes,omitempty"`

	// ConfigMapName is the name of the ConfigMap containing Tabby configuration
	ConfigMapName string `json:"configMapName,omitempty"`

	// Persistence defines Tabby persistence configuration
	Persistence TabbyPersistenceSpec `json:"persistence,omitempty"`
}

// RedisSpec defines the Redis configuration for OpenWebUI
type RedisSpec struct {
	// Enabled determines if Redis should be deployed automatically
	// If false and RedisURL is not provided, Redis will not be deployed
	Enabled bool `json:"enabled,omitempty"`

	// RedisURL is the Redis connection URL
	// If not provided and Enabled is true, Redis will be deployed automatically
	// Format: redis://host:port/db or rediss://host:port/db for TLS
	RedisURL string `json:"redisUrl,omitempty"`

	// Image is the Redis container image to use (including tag)
	Image string `json:"image,omitempty"`

	// Resources defines the resource requirements for Redis pods
	Resources ResourceRequirements `json:"resources,omitempty"`

	// Service defines the service configuration for Redis
	Service ServiceSpec `json:"service,omitempty"`

	// Password is the Redis password (optional)
	Password string `json:"password,omitempty"`

	// Persistence defines Redis persistence configuration
	Persistence RedisPersistenceSpec `json:"persistence,omitempty"`
}

// RedisPersistenceSpec defines Redis persistence configuration
type RedisPersistenceSpec struct {
	// Enabled determines if Redis data should be persisted
	Enabled bool `json:"enabled,omitempty"`

	// StorageClass is the storage class to use for persistent volumes
	StorageClass string `json:"storageClass,omitempty"`

	// Size is the size of the persistent volume
	Size string `json:"size,omitempty"`
}

// PipelinesSpec defines the OpenWebUI Pipelines configuration
type PipelinesSpec struct {
	// Enabled determines if OpenWebUI Pipelines should be deployed
	Enabled bool `json:"enabled,omitempty"`

	// Image is the Pipelines container image to use (including tag)
	Image string `json:"image,omitempty"`

	// Replicas is the number of Pipelines pods to run
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=3
	Replicas int32 `json:"replicas,omitempty"`

	// Port is the port the Pipelines service exposes
	// +kubebuilder:validation:Minimum=1
	// +kubebuilder:validation:Maximum=65535
	Port int32 `json:"port,omitempty"`

	// Resources defines the resource requirements for Pipelines pods
	Resources ResourceRequirements `json:"resources,omitempty"`

	// Service describes service to expose the Pipelines
	Service ServiceSpec `json:"service,omitempty"`

	// PipelinesDir is the directory containing pipeline definitions
	// Default: /app/pipelines
	PipelinesDir string `json:"pipelinesDir,omitempty"`

	// PipelineURLs is a list of URLs to fetch pipeline definitions from
	// Format: https://github.com/open-webui/pipelines/blob/main/examples/filters/example.py
	PipelineURLs []string `json:"pipelineUrls,omitempty"`

	// EnvVars defines environment variables for the Pipelines
	EnvVars []corev1.EnvVar `json:"envVars,omitempty"`

	// VolumeMounts defines volume mounts for the Pipelines
	VolumeMounts []corev1.VolumeMount `json:"volumeMounts,omitempty"`

	// Volumes defines volumes for the Pipelines
	Volumes []corev1.Volume `json:"volumes,omitempty"`

	// Persistence defines Pipelines persistence configuration
	Persistence *PipelinesPersistenceSpec `json:"persistence,omitempty"`
}

// PipelinesPersistenceSpec defines Pipelines persistence configuration
type PipelinesPersistenceSpec struct {
	// Enabled determines if Pipelines data should be persisted
	Enabled bool `json:"enabled,omitempty"`

	// StorageClass is the storage class to use for persistent volumes
	StorageClass string `json:"storageClass,omitempty"`

	// Size is the size of the persistent volume
	Size string `json:"size,omitempty"`
}

// TabbyPersistenceSpec defines Tabby persistence configuration
type TabbyPersistenceSpec struct {
	// Enabled determines if Tabby data should be persisted
	Enabled bool `json:"enabled,omitempty"`

	// StorageClass is the storage class to use for persistent volumes
	StorageClass string `json:"storageClass,omitempty"`

	// Size is the size of the persistent volume
	Size string `json:"size,omitempty"`
}

// OpenWebUIPersistenceSpec defines OpenWebUI persistence configuration
type OpenWebUIPersistenceSpec struct {
	// Enabled determines if OpenWebUI data should be persisted
	Enabled bool `json:"enabled,omitempty"`

	// StorageClass is the storage class to use for persistent volumes
	StorageClass string `json:"storageClass,omitempty"`

	// Size is the size of the persistent volume
	Size string `json:"size,omitempty"`
}

// LangfuseSpec defines the Langfuse monitoring configuration
type LangfuseSpec struct {
	// Enabled determines if Langfuse monitoring should be enabled
	Enabled bool `json:"enabled,omitempty"`

	// URL is the Langfuse server URL
	// Format: https://cloud.langfuse.com or http://localhost:3000
	URL string `json:"url,omitempty"`

	// SecretRef is the reference to a Kubernetes secret containing Langfuse credentials
	// The secret should contain: LANGFUSE_PUBLIC_KEY, LANGFUSE_SECRET_KEY
	SecretRef *corev1.SecretReference `json:"secretRef,omitempty"`

	// ProjectName is the name of the project for Langfuse
	// If not provided, will default to deployment name
	ProjectName string `json:"projectName,omitempty"`

	// Environment is the environment name (e.g., "production", "staging", "development")
	Environment string `json:"environment,omitempty"`

	// Debug enables debug logging for Langfuse
	Debug bool `json:"debug,omitempty"`
}

// ResourceRequirements describes the compute resource requirements
type ResourceRequirements struct {
	// Limits describes the maximum amount of compute resources allowed
	Limits *ResourceList `json:"limits,omitempty"`

	// Requests describes the minimum amount of compute resources required
	Requests *ResourceList `json:"requests,omitempty"`
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

// LMDeploymentSpec defines the desired state of Deployment
type LMDeploymentSpec struct {
	// Ollama defines the Ollama deployment configuration
	Ollama OllamaSpec `json:"ollama"`

	// OpenWebUI defines the OpenWebUI deployment configuration
	// +kubebuilder:validation:Optional
	OpenWebUI OpenWebUISpec `json:"openwebui,omitempty"`

	// Tabby defines the Tabby deployment configuration
	// +kubebuilder:validation:Optional
	Tabby TabbySpec `json:"tabby,omitempty"`
}

// LMDeploymentStatus defines the observed state of Deployment
type LMDeploymentStatus struct {
	// Phase represents the current phase of the deployment
	Phase string `json:"phase,omitempty"`

	// Conditions represent the latest available observations of the deployment's current state
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// OllamaStatus represents the status of Ollama deployment
	OllamaStatus LMDeploymentComponentStatus `json:"ollamaStatus,omitempty"`

	// OpenWebUIStatus represents the status of OpenWebUI deployment
	OpenWebUIStatus LMDeploymentComponentStatus `json:"openwebuiStatus,omitempty"`

	// TabbyStatus represents the status of Tabby deployment
	TabbyStatus LMDeploymentComponentStatus `json:"tabbyStatus,omitempty"`

	// ReadyReplicas is the number of ready replicas
	ReadyReplicas int32 `json:"readyReplicas,omitempty"`

	// TotalReplicas is the total number of replicas
	TotalReplicas int32 `json:"totalReplicas,omitempty"`
}

// LMDeploymentComponentStatus represents the status of a deployment component
type LMDeploymentComponentStatus struct {
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

// LMDeployment is the Schema for the lmdeployments API
type LMDeployment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LMDeploymentSpec   `json:"spec,omitempty"`
	Status LMDeploymentStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LMDeploymentList contains a list of LMDeployment
type LMDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LMDeployment `json:"items"`
}

// GetOllamaServiceName returns the name of the Ollama service for this deployment
func (d *LMDeployment) GetOllamaServiceName() string {
	return fmt.Sprintf("%s-ollama", d.Name)
}

// GetOllamaServicePort returns the port of the Ollama service for this deployment
func (d *LMDeployment) GetOllamaServicePort() int32 {
	if d.Spec.Ollama.Service.Port == 0 {
		return 11434 // Default Ollama port
	}
	return d.Spec.Ollama.Service.Port
}

// GetOpenWebUIServiceName returns the name of the OpenWebUI service for this deployment
func (d *LMDeployment) GetOpenWebUIServiceName() string {
	return fmt.Sprintf("%s-openwebui", d.Name)
}

// GetTabbyServiceName returns the name of the Tabby service for this deployment
func (d *LMDeployment) GetTabbyServiceName() string {
	return fmt.Sprintf("%s-tabby", d.Name)
}

// GetRedisServiceName returns the name of the Redis service for this deployment
func (d *LMDeployment) GetRedisServiceName() string {
	return fmt.Sprintf("%s-redis", d.Name)
}

// GetRedisDeploymentName returns the name of the Redis deployment for this deployment
func (d *LMDeployment) GetRedisDeploymentName() string {
	return fmt.Sprintf("%s-redis", d.Name)
}

// GetRedisPVCName returns the name of the Redis PVC for this deployment
func (d *LMDeployment) GetRedisPVCName() string {
	return fmt.Sprintf("%s-redis", d.Name)
}

// GetOllamaDeploymentName returns the name of the Ollama deployment for this deployment
func (d *LMDeployment) GetOllamaDeploymentName() string {
	return fmt.Sprintf("%s-ollama", d.Name)
}

// GetOpenWebUIDeploymentName returns the name of the OpenWebUI deployment for this deployment
func (d *LMDeployment) GetOpenWebUIDeploymentName() string {
	return fmt.Sprintf("%s-openwebui", d.Name)
}

// GetTabbyDeploymentName returns the name of the Tabby deployment for this deployment
func (d *LMDeployment) GetTabbyDeploymentName() string {
	return fmt.Sprintf("%s-tabby", d.Name)
}

// GetOllamaIngressName returns the name of the Ollama ingress for this deployment
func (d *LMDeployment) GetOllamaIngressName() string {
	return fmt.Sprintf("%s-ollama-ingress", d.Name)
}

// GetOpenWebUIIngressName returns the name of the OpenWebUI ingress for this deployment
func (d *LMDeployment) GetOpenWebUIIngressName() string {
	return fmt.Sprintf("%s-openwebui-ingress", d.Name)
}

// GetOpenWebUIConfigName returns the name of the OpenWebUI config Secret for this deployment
func (d *LMDeployment) GetOpenWebUIConfigName() string {
	return fmt.Sprintf("%s-openwebui-config", d.Name)
}

// GetOpenWebUIPVCName returns the name of the OpenWebUI PVC for this deployment
func (d *LMDeployment) GetOpenWebUIPVCName() string {
	return fmt.Sprintf("%s-openwebui-data", d.Name)
}

// GetTabbyIngressName returns the name of the Tabby ingress for this deployment
func (d *LMDeployment) GetTabbyIngressName() string {
	return fmt.Sprintf("%s-tabby-ingress", d.Name)
}

// GetTabbyConfigMapName returns the name of the Tabby ConfigMap for this deployment
func (d *LMDeployment) GetTabbyConfigMapName() string {
	return fmt.Sprintf("%s-tabby-config", d.Name)
}

// GetTabbyPVCName returns the name of the Tabby PVC for this deployment
func (d *LMDeployment) GetTabbyPVCName() string {
	return fmt.Sprintf("%s-tabby-data", d.Name)
}

// GetPipelinesServiceName returns the name of the Pipelines service for this deployment
func (d *LMDeployment) GetPipelinesServiceName() string {
	return fmt.Sprintf("%s-pipelines", d.Name)
}

// GetPipelinesDeploymentName returns the name of the Pipelines deployment for this deployment
func (d *LMDeployment) GetPipelinesDeploymentName() string {
	return fmt.Sprintf("%s-pipelines", d.Name)
}

func init() {
	SchemeBuilder.Register(&LMDeployment{}, &LMDeploymentList{})
}
