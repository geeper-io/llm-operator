---
id: ollama-deployment
title: Ollama Deployment
sidebar_label: Ollama
sidebar_position: 2
description: Deploy and configure Ollama for local model serving with the LLM Operator
---

# Ollama Deployment

Ollama is the default model serving backend in the LLM Operator, providing a simple and efficient way to run open-source language models locally on Kubernetes. Ollama is perfect for development, testing, and production workloads that require easy model management and deployment.

## What is Ollama?

Ollama is an open-source framework that makes it easy to run large language models locally. It provides:

- **🚀 Easy Model Management**: Simple commands to pull, run, and manage models
- **⚡ Local Inference**: Run models directly on your infrastructure
- **🎯 Wide Model Support**: Access to thousands of open-source models
- **🔄 Simple API**: RESTful API for easy integration
- **📦 Container Ready**: Optimized for containerized deployments

## Key Features

### **Model Management**
- **Easy Pulling**: Simple commands to download models from Ollama library
- **Version Control**: Support for model tags and versions
- **Local Storage**: Models stored locally for fast access
- **Model Switching**: Quickly switch between different models

### **Performance**
- **Optimized Inference**: Efficient model serving with minimal overhead
- **Memory Management**: Smart memory allocation for optimal performance
- **Multi-Model Support**: Run multiple models simultaneously
- **Resource Efficient**: Minimal resource footprint

### **Integration**
- **REST API**: Standard HTTP API for easy integration
- **WebSocket Support**: Real-time streaming capabilities
- **Multiple Clients**: Support for various client libraries
- **Kubernetes Native**: Designed for container orchestration

## Configuration

### Basic Ollama Deployment

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: ollama-basic
  namespace: default
spec:
  ollama:
    replicas: 1
    image: ollama/ollama:latest
    models:
      - "llama2:7b"
      - "codellama:7b"
      - "mistral:7b"
    service:
      type: ClusterIP
      port: 11434
    resources:
      requests:
        cpu: "2"
        memory: "8Gi"
      limits:
        cpu: "4"
        memory: "16Gi"
```

### Production Ollama Deployment

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: production-ollama
  namespace: ai-production
spec:
  ollama:
    replicas: 3
    image: ollama/ollama:latest
    models:
      - "llama2:13b"
      - "codellama:34b"
      - "mistral:7b"
      - "phi:2.7b"
    service:
      type: LoadBalancer
      port: 11434
    resources:
      requests:
        cpu: "4"
        memory: "16Gi"
      limits:
        cpu: "8"
        memory: "32Gi"
    affinity:
      podAntiAffinity:
        preferredDuringSchedulingIgnoredDuringExecution:
        - weight: 100
          podAffinityTerm:
            labelSelector:
              matchExpressions:
              - key: app
                operator: In
                values: ["ollama"]
            topologyKey: kubernetes.io/hostname
```

### GPU-Accelerated Deployment

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: ollama-gpu
  namespace: default
spec:
  ollama:
    replicas: 2
    image: ollama/ollama:latest
    models:
      - "llama2:13b"
      - "codellama:34b"
      - "mistral:7b"
    resources:
      requests:
        cpu: "2"
        memory: "8Gi"
        nvidia.com/gpu: "1"
      limits:
        cpu: "4"
        memory: "16Gi"
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
                values: ["ollama"]
            topologyKey: kubernetes.io/hostname
    service:
      type: ClusterIP
      port: 11434
```

## Configuration Options

### OllamaSpec Fields

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `replicas` | int32 | No | `1` | Number of Ollama pods (1-10) |
| `image` | string | No | `ollama/ollama:latest` | Ollama container image |
| `models` | []string | Yes | - | List of models to deploy |
| `service` | ServiceSpec | No | ClusterIP:11434 | Service configuration |
| `resources` | ResourceRequirements | No | - | CPU/Memory/GPU requirements |
| `affinity` | Affinity | No | - | Pod scheduling rules |

### Service Configuration

```yaml
service:
  type: ClusterIP        # ClusterIP, NodePort, or LoadBalancer
  port: 11434           # Service port (1-65535)
```

### Resource Requirements

```yaml
resources:
  requests:
    cpu: "2"
    memory: "8Gi"
    nvidia.com/gpu: "1"    # GPU resource request (optional)
  limits:
    cpu: "4"
    memory: "16Gi"
    nvidia.com/gpu: "1"    # GPU resource limit (optional)
```

## Model Management

### **Supported Model Formats**

Ollama supports various model formats and sources:

- **Ollama Models**: `llama2:7b`, `codellama:34b`, `mistral:7b`
- **Custom Models**: `my-model:latest`, `company/model:v1.0`
- **Hugging Face**: `hf.co/username/model:tag`
- **Local Models**: `./path/to/model:tag`

### **Model Pulling Strategy**

The LLM Operator automatically pulls models during pod startup:

1. **PostStart Hook**: Models are pulled automatically when pods start
2. **Parallel Downloading**: Multiple models can be pulled simultaneously
3. **Caching**: Models are cached locally for faster subsequent starts
4. **Error Handling**: Failed model pulls are retried automatically

### **Model Configuration Examples**

```yaml
# Basic models
models:
  - "llama2:7b"
  - "mistral:7b"

# Custom models with tags
models:
  - "my-llama:latest"
  - "company/model:v2.1"

# Hugging Face models
models:
  - "hf.co/meta-llama/Llama-2-7b-chat-hf"
  - "hf.co/microsoft/DialoGPT-medium"
```

## Integration with Other Components

### OpenWebUI Integration

Ollama works seamlessly with OpenWebUI for web-based chat interfaces:

```yaml
openwebui:
  enabled: true
  replicas: 1
  image: ghcr.io/open-webui/open-webui:main
  service:
    type: ClusterIP
    port: 8080
  ingress:
    host: "ollama-webui.localhost"
```

### Tabby Integration

Ollama models can be used with Tabby for AI-powered code completion:

```yaml
tabby:
  enabled: true
  replicas: 1
  image: tabbyml/tabby:latest
  device: cuda
  chatModel: "llama2:7b"
  completionModel: "codellama:7b"
  resources:
    requests:
      cpu: "2"
      memory: "4Gi"
      nvidia.com/gpu: "1"
```

## Best Practices

### **Resource Allocation**

- **CPU**: Allocate 2-4 CPU cores per Ollama pod
- **Memory**: Plan for 2-4x model size in RAM
- **GPU**: Use GPU resources for larger models (13B+ parameters)
- **Storage**: Use local storage for model caching

### **Scaling Considerations**

- **Horizontal Scaling**: Use multiple replicas for high availability
- **Load Distribution**: Spread pods across different nodes
- **Resource Limits**: Set appropriate CPU/memory limits
- **Monitoring**: Implement comprehensive monitoring and alerting

### **Model Selection**

- **Use Case**: Choose models appropriate for your specific needs
- **Resource Constraints**: Consider available CPU/memory/GPU resources
- **Performance**: Balance model size with inference speed
- **Licensing**: Ensure compliance with model licenses

### **Production Deployment**

- **High Availability**: Use multiple replicas and pod anti-affinity
- **Monitoring**: Implement logging, metrics, and alerting
- **Security**: Use network policies and RBAC
- **Backup**: Plan for model data backup and recovery

## Troubleshooting

### Common Issues

#### **Model Pull Failures**
```bash
# Check pod logs for pull errors
kubectl logs -f deployment/ollama-basic-ollama

# Verify model names are correct
# Check network connectivity to Ollama registry
```

#### **Out of Memory**
```yaml
# Increase memory limits
resources:
  requests:
    memory: "16Gi"
  limits:
    memory: "32Gi"
```

#### **GPU Not Found**
```bash
# Check GPU availability
kubectl get nodes -o json | jq '.items[] | select(.status.allocatable."nvidia.com/gpu" != null)'

# Verify GPU drivers are installed
# Check node labels for GPU resources
```

#### **Service Connection Issues**
```bash
# Verify service configuration
kubectl get svc ollama-basic-ollama
kubectl describe svc ollama-basic-ollama

# Check endpoint connectivity
kubectl get endpoints ollama-basic-ollama
```

### Performance Optimization

#### **Memory Optimization**
```yaml
# Use quantized models when possible
models:
  - "llama2:7b-q4_0"    # Quantized version
  - "codellama:7b-q4_0"

# Adjust resource allocation
resources:
  requests:
    memory: "8Gi"        # Minimum required
  limits:
    memory: "16Gi"       # Maximum allowed
```

#### **Scaling for Performance**
```yaml
# Increase replicas for higher throughput
replicas: 3

# Use pod anti-affinity for better distribution
affinity:
  podAntiAffinity:
    preferredDuringSchedulingIgnoredDuringExecution:
    - weight: 100
      podAffinityTerm:
        labelSelector:
          matchExpressions:
          - key: app
            operator: In
            values: ["ollama"]
        topologyKey: kubernetes.io/hostname
```

## Examples

### **Development Environment**
```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: dev-ollama
  namespace: development
spec:
  ollama:
    replicas: 1
    image: ollama/ollama:latest
    models:
      - "llama2:7b"
    service:
      type: NodePort
      port: 11434
    resources:
      requests:
        cpu: "1"
        memory: "4Gi"
      limits:
        cpu: "2"
        memory: "8Gi"
```

### **Production Environment**
```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: production-ollama
  namespace: ai-production
spec:
  ollama:
    replicas: 3
    image: ollama/ollama:latest
    models:
      - "llama2:13b"
      - "codellama:34b"
      - "mistral:7b"
      - "phi:2.7b"
    service:
      type: LoadBalancer
      port: 11434
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
                values: ["ollama"]
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
```

## Next Steps

- **Deploy Your First Ollama**: Use the examples above to get started
- **Explore Model Options**: Try different models for various use cases
- **Integrate Components**: Connect with OpenWebUI and Tabby
- **Scale Up**: Add more replicas and resources as needed
- **Monitor Performance**: Use built-in monitoring and observability

For more information, see the [vLLM Deployment](./vllm) guide and [API Reference](../crd-reference).
