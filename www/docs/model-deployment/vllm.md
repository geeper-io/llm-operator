---
id: vllm-deployment
title: vLLM Deployment
sidebar_label: vLLM
sidebar_position: 1
description: Deploy high-performance vLLM model serving with GPU acceleration and model routing
---

# vLLM Deployment

vLLM (Very Large Language Model) is a high-performance inference engine that provides optimized model serving with GPU acceleration. The LLM Operator supports vLLM as an alternative to Ollama for production environments that require maximum performance and efficiency, with support for multiple models and intelligent routing.

## What is vLLM?

vLLM is a high-performance inference engine for large language models that provides:

- **🚀 Optimized Inference**: PagedAttention algorithm for efficient memory management
- **⚡ High Throughput**: Significantly faster inference compared to standard serving
- **🎯 GPU Optimization**: CUDA, ROCm, and Metal support for hardware acceleration
- **🔄 Continuous Batching**: Efficient handling of multiple concurrent requests
- **📊 Production Ready**: Built for high-scale production deployments
- **🛣️ Model Routing**: Intelligent routing between multiple models
- **📈 Multi-Model Support**: Serve multiple models simultaneously with individual scaling

## Key Features

### **Performance Optimizations**
- **PagedAttention**: Advanced memory management for large models
- **Continuous Batching**: Dynamic batching for optimal throughput
- **Tensor Parallelism**: Multi-GPU support for large models
- **Quantization**: Support for various precision formats (FP16, INT8, etc.)

### **Hardware Support**
- **NVIDIA GPUs**: Full CUDA support with optimized kernels
- **AMD GPUs**: ROCm support for AMD graphics cards
- **Apple Silicon**: Metal Performance Shaders (MPS) support
- **CPU Fallback**: Efficient CPU inference when GPUs unavailable

### **Model Management**
- **Multiple Models**: Deploy and serve multiple models simultaneously
- **Individual Scaling**: Scale each model independently based on demand
- **Model Routing**: Intelligent router for distributing requests across models
- **Global Configuration**: Shared settings that apply to all models

## Configuration

### Basic Multi-Model vLLM Deployment

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: vllm-basic
  namespace: default
spec:
  vllm:
    enabled: true
    
    # Global configuration for all models
    globalConfig:
      image: "vllm/vllm-openai:latest"
      service:
        type: ClusterIP
        port: 8000
      resources:
        requests:
          cpu: "2"
          memory: "8Gi"
        limits:
          cpu: "4"
          memory: "16Gi"
    
    # Multiple models with individual configurations
    models:
      - name: "llama2-7b"
        model: "meta-llama/Llama-2-7b-chat-hf"
        replicas: 1
      
      - name: "codellama-7b"
        model: "codellama/CodeLlama-7b-Instruct-hf"
        replicas: 1
        resources:
          requests:
            cpu: "3"
            memory: "12Gi"
          limits:
            cpu: "6"
            memory: "24Gi"
```

### vLLM with API Key Authentication

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: vllm-secure
  namespace: default
spec:
  vllm:
    enabled: true
    
    # Enable API key authentication using SecretReference
    apiKey:
      # Optional: specify custom secret name
      # name: "my-vllm-api-key"
      # Optional: specify custom key name in secret (defaults to VLLM_API_KEY)
      # key: "CUSTOM_API_KEY"
    
    globalConfig:
      image: "vllm/vllm-openai:latest"
      service:
        type: ClusterIP
        port: 8000
    
    models:
      - name: "llama2-7b"
        model: "meta-llama/Llama-2-7b-chat-hf"
        replicas: 1
      
      - name: "codellama-7b"
        model: "codellama/CodeLlama-7b-Instruct-hf"
        replicas: 1
    
    router:
      enabled: true
      replicas: 1
      image: "lmcache/lmstack-router:latest"
```

### GPU-Accelerated Multi-Model Deployment

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: vllm-gpu
  namespace: default
spec:
  vllm:
    enabled: true
    
    globalConfig:
      image: "vllm/vllm-openai:latest"
      resources:
        requests:
          cpu: "4"
          memory: "16Gi"
          nvidia.com/gpu: "1"
        limits:
          cpu: "8"
          memory: "32Gi"
          nvidia.com/gpu: "1"
    
    models:
      - name: "llama2-13b"
        model: "meta-llama/Llama-2-13b-chat-hf"
        replicas: 2
        affinity:
          nodeAffinity:
            requiredDuringSchedulingIgnoredDuringExecution:
              nodeSelectorTerms:
              - matchExpressions:
                - key: nvidia.com/gpu
                  operator: Exists
          podAntiAffinity:
            preferredDuringSchedulingIgnoredDuringExecution:
            - weight: 100
              podAffinityTerm:
                labelSelector:
                  matchExpressions:
                  - key: app
                    operator: In
                    values: ["vllm"]
                topologyKey: kubernetes.io/hostname
      
      - name: "codellama-34b"
        model: "codellama/CodeLlama-34b-Instruct-hf"
        replicas: 1
        resources:
          requests:
            cpu: "6"
            memory: "24Gi"
            nvidia.com/gpu: "1"
          limits:
            cpu: "12"
            memory: "48Gi"
            nvidia.com/gpu: "1"
```

### Multi-Model with Router and Persistence

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: vllm-production
  namespace: default
spec:
  vllm:
    enabled: true
    
    globalConfig:
      image: "vllm/vllm-openai:latest"
      service:
        type: ClusterIP
        port: 8000
      persistence:
        enabled: true
        size: "50Gi"
        storageClass: "fast-ssd"
    
    models:
      - name: "llama2-7b"
        model: "meta-llama/Llama-2-7b-chat-hf"
        replicas: 3
        service:
          type: ClusterIP
          port: 8001
        envVars:
          - name: "VLLM_USE_ASYNC_ENGINE"
            value: "true"
          - name: "VLLM_MAX_MODEL_LEN"
            value: "4096"
      
      - name: "mistral-7b"
        model: "mistralai/Mistral-7B-Instruct-v0.2"
        replicas: 2
        service:
          type: ClusterIP
          port: 8002
        envVars:
          - name: "VLLM_USE_ASYNC_ENGINE"
            value: "true"
          - name: "VLLM_MAX_MODEL_LEN"
            value: "32768"
      
      - name: "phi-2"
        model: "microsoft/phi-2"
        replicas: 1
        service:
          type: ClusterIP
          port: 8003
        envVars:
          - name: "VLLM_USE_ASYNC_ENGINE"
            value: "true"
    
    # Router for intelligent model routing
    router:
      enabled: true
      replicas: 2
      image: "lmcache/lmstack-router:latest"
      resources:
        requests:
          cpu: "500m"
          memory: "1Gi"
        limits:
          cpu: "1"
          memory: "2Gi"
      service:
        type: LoadBalancer
        port: 8000
      envVars:
        - name: "ROUTER_LOG_LEVEL"
          value: "INFO"
        - name: "ROUTER_TIMEOUT"
          value: "30"
```

## Configuration Options

### VLLMSpec Fields

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `enabled` | bool | Yes | `false` | Enable vLLM deployment |
| `models` | []VLLMModelSpec | Yes | - | List of models to deploy |
| `router` | VLLMRouterSpec | No | - | Router configuration for model routing |
| `globalConfig` | VLLMGlobalConfig | No | - | Global configuration for all models |
| `apiKey` | corev1.SecretReference | No | - | API key authentication configuration using SecretReference |

### VLLMModelSpec Fields

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `name` | string | Yes | - | Unique name for this model deployment |
| `model` | string | Yes | - | Model identifier (e.g., "meta-llama/Llama-2-7b-chat-hf") |
| `replicas` | int32 | No | `1` | Number of pods for this model (1-10) |
| `image` | string | No | Global default | Container image for this model |
| `resources` | ResourceRequirements | No | Global default | CPU/Memory/GPU requirements |
| `service` | ServiceSpec | No | Global default | Service configuration for this model |
| `affinity` | Affinity | No | - | Pod scheduling rules |
| `envVars` | []EnvVar | No | - | Custom environment variables |
| `volumeMounts` | []VolumeMount | No | - | Custom volume mounts |
| `volumes` | []Volume | No | - | Custom volumes |
| `persistence` | VLLMPersistenceSpec | No | Global default | Storage configuration |

### VLLMRouterSpec Fields

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `enabled` | bool | No | `false` | Enable the vLLM router |
| `replicas` | int32 | No | `1` | Number of router pods (1-5) |
| `image` | string | No | `lmcache/lmstack-router:latest` | Router container image |
| `resources` | ResourceRequirements | No | - | CPU/Memory requirements |
| `service` | ServiceSpec | No | ClusterIP:8000 | Service configuration |
| `affinity` | Affinity | No | - | Pod scheduling rules |
| `envVars` | []EnvVar | No | - | Custom environment variables |

### VLLMGlobalConfig Fields

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `image` | string | No | `vllm/vllm-openai:latest` | Default container image |
| `resources` | ResourceRequirements | No | - | Default resource requirements |
| `service` | ServiceSpec | No | ClusterIP:8000 | Default service configuration |
| `persistence` | VLLMPersistenceSpec | No | - | Default persistence configuration |

## Model Routing

The vLLM router provides intelligent routing between multiple models:

### **Router Features**
- **Load Balancing**: Distributes requests across available models
- **Model Selection**: Routes requests to appropriate models based on configuration
- **Health Checking**: Monitors model health and routes away from unhealthy instances
- **Failover**: Automatically fails over to healthy models if one becomes unavailable

### **Router Configuration**
```yaml
router:
  enabled: true
  replicas: 2
  image: "lmcache/lmstack-router:latest"
  service:
    type: LoadBalancer
    port: 8000
  envVars:
    - name: "ROUTER_LOG_LEVEL"
      value: "INFO"
    - name: "ROUTER_TIMEOUT"
      value: "30"
```

### **Accessing Models**
- **Router Endpoint**: `http://router-service:8000` (routes to appropriate model)
- **Direct Model Access**: `http://model-service:port` (direct access to specific model)
- **Model-Specific Endpoints**: Each model has its own service for direct access

## API Key Authentication

The operator automatically creates and manages API key secrets for vLLM deployments. If you don't specify an `apiKey`, a secure API key is generated automatically and injected as the `VLLM_API_KEY` environment variable.

### Configuration

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `name` | string | No | Auto-generated | Name of the secret containing the API key |
| `key` | string | No | `VLLM_API_KEY` | Key name in the secret |

### Examples

**Automatic (Recommended):**
```yaml
vllm:
  enabled: true
  # No apiKey specified - operator creates secret automatically
```

**Custom Secret:**
```yaml
vllm:
  enabled: true
  apiKey:
    name: "my-custom-secret"
    key: "CUSTOM_API_KEY"
```

The `VLLM_API_KEY` is automatically used by OpenWebUI and Tabby when connecting to vLLM.

## Environment Variables

vLLM supports [various environment variables](https://docs.vllm.ai/en/v0.7.3/serving/env_vars.html) for fine-tuning performance. Use the `envVars` field to set these in your deployment.

### **Common Environment Variables**
```yaml
envVars:
  - name: "VLLM_USE_ASYNC_ENGINE"
    value: "true"
  - name: "VLLM_MAX_MODEL_LEN"
    value: "4096"
  - name: "VLLM_GPU_MEMORY_UTILIZATION"
    value: "0.9"
  - name: "VLLM_TENSOR_PARALLEL_SIZE"
    value: "1"
```

## Integration with Other Components

### OpenWebUI Integration

vLLM works seamlessly with OpenWebUI for web-based chat interfaces:

```yaml
openwebui:
  enabled: true
  replicas: 1
  image: ghcr.io/open-webui/open-webui:main
  resources:
    requests:
      cpu: "500m"
      memory: "1Gi"
    limits:
      cpu: "1"
      memory: "2Gi"
  service:
    type: ClusterIP
    port: 8080
  ingress:
    host: "vllm-webui.localhost"
```

The deployed vLLM instances will be automatically connected to OpenWebUI as openapi endpoints for chat functionality.

### Tabby Integration

vLLM models can be used with Tabby for AI-powered code completion:

```yaml
tabby:
  enabled: true
  replicas: 1
  image: tabbyml/tabby:latest
  device: cuda
  chatModel: "meta-llama/Llama-2-7b-chat-hf"
  completionModel: "codellama/CodeLlama-7b-Instruct-hf"
  resources:
    requests:
      cpu: "2"
      memory: "4Gi"
      nvidia.com/gpu: "1"
    limits:
      cpu: "4"
      memory: "8Gi"
      nvidia.com/gpu: "1"
  service:
    type: ClusterIP
    port: 8080
```

## Best Practices

### **Resource Allocation**
- **GPU Memory**: Allocate sufficient GPU memory for your models
- **CPU Cores**: Reserve 2-4 CPU cores per vLLM pod
- **Memory**: Plan for 2-4x model size in RAM
- **Storage**: Use fast storage for model caching

### **Scaling Considerations**
- **Horizontal Scaling**: Use multiple replicas for high availability
- **GPU Distribution**: Spread pods across GPU nodes
- **Load Balancing**: Use LoadBalancer service type for external access
- **Resource Limits**: Set appropriate CPU/memory limits

### **Model Management**
- **Model Selection**: Choose models appropriate for your use case
- **Quantization**: Use quantized models for memory efficiency
- **Caching**: Enable persistence for faster model loading
- **Updates**: Plan for model updates and versioning

### **Production Deployment**
- **Monitoring**: Implement comprehensive monitoring and alerting
- **Logging**: Configure structured logging for debugging
- **Security**: Use network policies and RBAC
- **Backup**: Implement backup strategies for model data

## Troubleshooting

### Common Issues

#### **GPU Not Found**
```bash
# Check GPU availability
kubectl get nodes -o json | jq '.items[] | select(.status.allocatable."nvidia.com/gpu" != null)'

# Verify GPU drivers are installed
# Check node labels for GPU resources
```

#### **Out of Memory**
```yaml
# Reduce GPU memory utilization
envVars:
  - name: VLLM_GPU_MEMORY_UTILIZATION
    value: "0.8"

# Increase memory limits
resources:
  requests:
    memory: "16Gi"
  limits:
    memory: "32Gi"
```

#### **Model Loading Failures**
```bash
# Check pod logs
kubectl logs -f deployment/vllm-basic-llama2-7b

# Verify model names are correct
# Check network connectivity to Hugging Face
```

#### **Service Connection Issues**
```bash
# Verify service configuration
kubectl get svc vllm-basic-llama2-7b
kubectl describe svc vllm-basic-llama2-7b

# Check endpoint connectivity
kubectl get endpoints vllm-basic-llama2-7b
```

#### **Router Issues**
```bash
# Check router deployment
kubectl get deployment vllm-basic-vllm-router
kubectl logs -f deployment/vllm-basic-vllm-router

# Verify router service
kubectl get svc vllm-basic-vllm-router
```

## Migration from Ollama

### **When to Use vLLM vs Ollama**

| Feature | Ollama | vLLM |
|---------|--------|------|
| **Performance** | Good | Excellent |
| **GPU Support** | Basic | Advanced |
| **Memory Efficiency** | Standard | Optimized |
| **Multi-Model** | Yes | Yes (with routing) |
| **Model Routing** | No | Yes |
| **Production Ready** | Limited | Yes |
| **Resource Usage** | Higher | Lower |

### **Migration Steps**

1. **Update Configuration**
   ```yaml
   # Before (Ollama)
   spec:
     ollama:
       models: ["llama2:7b"]
   
   # After (vLLM)
   spec:
     vLLM:
       enabled: true
       models:
         - name: "llama2-7b"
           model: "meta-llama/Llama-2-7b-chat-hf"
   ```

2. **Adjust Resources** (typically 20-30% less memory needed)
3. **Update Service Ports** (11434 → 8000)
4. **Add Router** (optional, for multi-model deployments)
5. **Test Performance** and adjust as needed

## Examples

### **Complete Production Setup**

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: production-vllm
  namespace: ai-production
spec:
  vllm:
    enabled: true
    
    globalConfig:
      image: "vllm/vllm-openai:latest"
      persistence:
        enabled: true
        size: "100Gi"
        storageClass: "fast-ssd"
    
    models:
      - name: "llama2-13b"
        model: "meta-llama/Llama-2-13b-chat-hf"
        replicas: 3
        resources:
          requests:
            cpu: "4"
            memory: "16Gi"
            nvidia.com/gpu: "1"
          limits:
            cpu: "8"
            memory: "32Gi"
            nvidia.com/gpu: "1"
        affinity:
          nodeAffinity:
            requiredDuringSchedulingIgnoredDuringExecution:
              nodeSelectorTerms:
              - matchExpressions:
                - key: nvidia.com/gpu
                  operator: Exists
          podAntiAffinity:
            preferredDuringSchedulingIgnoredDuringExecution:
            - weight: 100
              podAffinityTerm:
                labelSelector:
                  matchExpressions:
                  - key: app
                    operator: In
                    values: ["vllm"]
                topologyKey: kubernetes.io/hostname
        envVars:
          - name: VLLM_USE_ASYNC_ENGINE
            value: "true"
          - name: VLLM_MAX_MODEL_LEN
            value: "8192"
          - name: VLLM_GPU_MEMORY_UTILIZATION
            value: "0.95"
      
      - name: "codellama-34b"
        model: "codellama/CodeLlama-34b-Instruct-hf"
        replicas: 2
        resources:
          requests:
            cpu: "6"
            memory: "24Gi"
            nvidia.com/gpu: "1"
          limits:
            cpu: "12"
            memory: "48Gi"
            nvidia.com/gpu: "1"
        envVars:
          - name: VLLM_USE_ASYNC_ENGINE
            value: "true"
          - name: VLLM_MAX_MODEL_LEN
            value: "8192"
    
    router:
      enabled: true
      replicas: 3
      image: "lmcache/lmstack-router:latest"
      resources:
        requests:
          cpu: "1"
          memory: "2Gi"
        limits:
          cpu: "2"
          memory: "4Gi"
      service:
        type: LoadBalancer
        port: 8000
      affinity:
        podAntiAffinity:
          preferredDuringSchedulingIgnoredDuringExecution:
          - weight: 100
            podAffinityTerm:
              labelSelector:
                matchExpressions:
                - key: app
                  operator: In
                  values: ["vllm-router"]
              topologyKey: kubernetes.io/hostname
  
  openwebui:
    enabled: true
    replicas: 2
    image: ghcr.io/open-webui/open-webui:main
    service:
      type: LoadBalancer
      port: 8080
    ingress:
      host: "ai.company.com"
  
  tabby:
    enabled: true
    replicas: 2
    image: tabbyml/tabby:latest
    device: cuda
    chatModel: "meta-llama/Llama-2-13b-chat-hf"
    completionModel: "codellama/CodeLlama-34b-Instruct-hf"
    resources:
      requests:
        cpu: "2"
        memory: "4Gi"
        nvidia.com/gpu: "1"
      limits:
        cpu: "4"
        memory: "8Gi"
        nvidia.com/gpu: "1"
    service:
      type: LoadBalancer
      port: 8080
```

## Next Steps

- **Deploy Your First vLLM**: Use the examples above to get started
- **Explore Multi-Model**: Deploy multiple models with individual scaling
- **Add Router**: Implement intelligent model routing for production
- **Monitor Performance**: Use built-in monitoring and observability
- **Scale Up**: Add more replicas and GPU nodes as needed

For more information, see the [Ollama Deployment](./ollama) guide and [API Reference](../crd-reference).


