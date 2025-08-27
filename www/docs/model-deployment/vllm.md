---
id: vllm-deployment
title: vLLM Deployment
sidebar_label: vLLM
description: Deploy high-performance vLLM model serving with GPU acceleration
---

# vLLM Deployment

vLLM (Very Large Language Model) is a high-performance inference engine that provides optimized model serving with GPU acceleration. The LLM Operator supports vLLM as an alternative to Ollama for production environments that require maximum performance and efficiency.

## What is vLLM?

vLLM is a high-performance inference engine for large language models that provides:

- **🚀 Optimized Inference**: PagedAttention algorithm for efficient memory management
- **⚡ High Throughput**: Significantly faster inference compared to standard serving
- **🎯 GPU Optimization**: CUDA, ROCm, and Metal support for hardware acceleration
- **🔄 Continuous Batching**: Efficient handling of multiple concurrent requests
- **📊 Production Ready**: Built for high-scale production deployments

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

### **Model Compatibility**
- **Hugging Face Models**: Direct support for HF model formats
- **Custom Models**: Support for custom model architectures
- **Multi-Model**: Serve multiple models simultaneously
- **Model Switching**: Dynamic model loading and switching

## Configuration

### Basic vLLM Deployment

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: vllm-basic
  namespace: default
spec:
  vllm:
    enabled: true
    replicas: 1
    image: vllm/vllm-openai:latest
    models:
      - "meta-llama/Llama-2-7b-chat-hf"
      - "microsoft/DialoGPT-medium"
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
```

### GPU-Accelerated Deployment

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: vllm-gpu
  namespace: default
spec:
  vllm:
    enabled: true
    replicas: 2
    image: vllm/vllm-openai:latest
    models:
      - "meta-llama/Llama-2-13b-chat-hf"
      - "codellama/CodeLlama-34b-Instruct-hf"
      - "microsoft/DialoGPT-large"
    service:
      type: ClusterIP
      port: 8000
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
```

### Persistent Storage Configuration

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: vllm-persistent
  namespace: default
spec:
  vllm:
    enabled: true
    replicas: 1
    image: vllm/vllm-openai:latest
    models:
      - "meta-llama/Llama-2-7b-chat-hf"
      - "microsoft/DialoGPT-medium"
      - "gpt2"
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
    persistence:
      enabled: true
      size: "20Gi"
      storageClass: "standard"
    envVars:
      - name: VLLM_USE_ASYNC_ENGINE
        value: "true"
      - name: VLLM_MAX_MODEL_LEN
        value: "4096"
      - name: VLLM_GPU_MEMORY_UTILIZATION
        value: "0.9"
```

## Configuration Options

### VLLMSpec Fields

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `enabled` | bool | Yes | `false` | Enable vLLM deployment |
| `replicas` | int32 | No | `1` | Number of vLLM pods (1-10) |
| `image` | string | No | `vllm/vllm-openai:latest` | vLLM container image |
| `models` | []string | Yes | - | List of models to deploy |
| `service` | ServiceSpec | No | ClusterIP:8000 | Service configuration |
| `resources` | ResourceRequirements | No | - | CPU/Memory/GPU requirements |
| `affinity` | Affinity | No | - | Pod scheduling rules |
| `envVars` | []EnvVar | No | - | Custom environment variables |
| `volumeMounts` | []VolumeMount | No | - | Custom volume mounts |
| `volumes` | []Volume | No | - | Custom volumes |
| `persistence` | VLLMPersistenceSpec | No | - | Storage configuration |

### Service Configuration

```yaml
service:
  type: ClusterIP        # ClusterIP, NodePort, or LoadBalancer
  port: 8000            # Service port (1-65535)
```

### Resource Requirements

```yaml
resources:
  requests:
    cpu: "4"
    memory: "16Gi"
    nvidia.com/gpu: "1"    # GPU resource request
  limits:
    cpu: "8"
    memory: "32Gi"
    nvidia.com/gpu: "1"    # GPU resource limit
```

### Persistence Configuration

```yaml
persistence:
  enabled: true
  size: "20Gi"              # Storage size
  storageClass: "fast-ssd"  # Storage class (optional)
```

## Environment Variables

vLLM supports various environment variables for fine-tuning performance:

### Performance Tuning

| Variable | Description | Default | Example |
|----------|-------------|---------|---------|
| `VLLM_USE_ASYNC_ENGINE` | Enable async engine | `false` | `true` |
| `VLLM_MAX_MODEL_LEN` | Maximum sequence length | `8192` | `4096` |
| `VLLM_GPU_MEMORY_UTILIZATION` | GPU memory usage | `0.9` | `0.95` |
| `VLLM_TENSOR_PARALLEL_SIZE` | Tensor parallelism | `1` | `2` |
| `VLLM_BLOCK_SIZE` | Attention block size | `16` | `32` |

### Model Configuration

| Variable | Description | Default | Example |
|----------|-------------|---------|---------|
| `MODEL_NAME` | Primary model name | First model | `llama2-7b` |
| `MODEL_LIST` | Comma-separated models | All models | `llama2-7b,code-7b` |
| `HOST` | Bind address | `0.0.0.0` | `0.0.0.0` |
| `PORT` | Service port | `8000` | `8000` |

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
  persistence:
    enabled: true
    size: "5Gi"
    storageClass: "standard"
  ingress:
    host: "vllm-webui.localhost"
```

### Tabby Integration

vLLM models can be used with Tabby for AI-powered code completion:

```yaml
tabby:
  enabled: true
  replicas: 1
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
kubectl logs -f deployment/vllm-basic-vllm

# Verify model names are correct
# Check network connectivity to Hugging Face
```

#### **Service Connection Issues**
```bash
# Verify service configuration
kubectl get svc vllm-basic-vllm
kubectl describe svc vllm-basic-vllm

# Check endpoint connectivity
kubectl get endpoints vllm-basic-vllm
```

### Performance Tuning

#### **Optimize Throughput**
```yaml
envVars:
  - name: VLLM_USE_ASYNC_ENGINE
    value: "true"
  - name: VLLM_MAX_MODEL_LEN
    value: "4096"
  - name: VLLM_BLOCK_SIZE
    value: "32"
```

#### **Memory Optimization**
```yaml
envVars:
  - name: VLLM_GPU_MEMORY_UTILIZATION
    value: "0.9"
  - name: VLLM_TENSOR_PARALLEL_SIZE
    value: "2"
```

## Migration from Ollama

### **When to Use vLLM vs Ollama**

| Feature | Ollama | vLLM |
|---------|--------|------|
| **Performance** | Good | Excellent |
| **GPU Support** | Basic | Advanced |
| **Memory Efficiency** | Standard | Optimized |
| **Ease of Use** | Simple | Moderate |
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
     vllm:
       enabled: true
       models: ["meta-llama/Llama-2-7b-chat-hf"]
   ```

2. **Adjust Resources**
   ```yaml
   # vLLM typically needs less memory
   resources:
     requests:
       memory: "8Gi"    # vs 16Gi for Ollama
     limits:
       memory: "16Gi"   # vs 32Gi for Ollama
   ```

3. **Update Service Ports**
   ```yaml
   # Ollama uses port 11434
   # vLLM uses port 8000
   service:
     port: 8000
   ```

4. **Update Model References**
   ```yaml
   # Update Tabby model references
   tabby:
     chatModel: "meta-llama/Llama-2-7b-chat-hf"
     completionModel: "codellama/CodeLlama-7b-Instruct-hf"
   ```

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
    replicas: 3
    image: vllm/vllm-openai:latest
    models:
      - "meta-llama/Llama-2-13b-chat-hf"
      - "codellama/CodeLlama-34b-Instruct-hf"
      - "microsoft/DialoGPT-large"
    service:
      type: LoadBalancer
      port: 8000
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
    persistence:
      enabled: true
      size: "100Gi"
      storageClass: "fast-ssd"
    envVars:
      - name: VLLM_USE_ASYNC_ENGINE
        value: "true"
      - name: VLLM_MAX_MODEL_LEN
        value: "8192"
      - name: VLLM_GPU_MEMORY_UTILIZATION
        value: "0.95"
  
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

### **Development Environment**

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: dev-vllm
  namespace: development
spec:
  vllm:
    enabled: true
    replicas: 1
    image: vllm/vllm-openai:latest
    models:
      - "meta-llama/Llama-2-7b-chat-hf"
    service:
      type: NodePort
      port: 8000
    resources:
      requests:
        cpu: "2"
        memory: "8Gi"
      limits:
        cpu: "4"
        memory: "16Gi"
  
  openwebui:
    enabled: true
    replicas: 1
    image: ghcr.io/open-webui/open-webui:main
    service:
      type: NodePort
      port: 8080
```

## Next Steps

- **Deploy Your First vLLM**: Use the examples above to get started
- **Explore Advanced Features**: Experiment with different configurations
- **Monitor Performance**: Use built-in monitoring and observability
- **Scale Up**: Add more replicas and GPU nodes as needed
- **Join Community**: Get help and share experiences with the community

For more information, see the [Ollama Deployment](./ollama) guide and [API Reference](../crd-reference).
