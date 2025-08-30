---
id: intro
title: Model Deployment Overview
sidebar_label: Overview
description: Choose between Ollama and vLLM for your model serving needs
---

# Model Deployment Overview

The LLM Operator makes deploying both **Ollama** and **vLLM** equally simple and straightforward. With just a few lines
of YAML, you can deploy either backend (or both) on Kubernetes. The choice between them comes down to your specific
performance needs, resource constraints, and use case requirements.

:::warning

Everything below is a general guideline. Actual performance and resource usage will vary based on model size, hardware, 
and your use-case. Always benchmark with your specific models and scenarios.

:::

## Quick Comparison

| Feature                   | Ollama     | vLLM       |
|---------------------------|------------|------------|
| **Performance**           | Good       | Excellent  |
| **GPU Support**           | Basic      | Advanced   |
| **Memory Efficiency**     | Standard   | Optimized  |
| **Resource Usage**        | Higher     | Lower      |
| **Model Loading**         | Automatic  | Manual     |
| **Multi-Model**           | Yes        | Yes        |
| **API Compatibility**     | Ollama API | OpenAI API |
| **Production Readiness**  | Limited    | Yes        |

## When to Use Ollama

**Choose Ollama if you need:**

- **🧪 Development & Testing**: Perfect for prototyping and development environments
- **💻 Simple Deployments**: Basic model serving without complex requirements
- **🔄 Rapid Iteration**: Easy to switch between models and configurations
- **📖 Learning**: Great for understanding model serving concepts
- **🚀 Quick Model Switching**: Models can be changed on-the-fly

**Best Use Cases:**

- Development environments
- Testing and prototyping
- Small-scale deployments
- Learning and experimentation
- Simple chat applications
- Personal projects
- Scenarios where model flexibility is key

## When to Use vLLM

**Choose vLLM if you need:**

- **⚡ Maximum Performance**: Highest throughput and lowest latency
- **🎯 GPU Optimization**: Advanced GPU utilization and memory management
- **🏭 Production Scale**: Built for high-scale production deployments
- **📊 Resource Efficiency**: Better memory and GPU utilization
- **🔄 Continuous Batching**: Efficient handling of multiple concurrent requests
- **🔧 Advanced Features**: Tensor parallelism, quantization, and more
- **📈 High Throughput**: Handle many requests simultaneously

**Best Use Cases:**

- Production environments
- High-traffic applications
- GPU-intensive workloads
- Large model deployments
- Enterprise applications
- Performance-critical systems
- Scenarios where maximum efficiency is required

## Performance Comparison

### **Throughput (Requests/Second)**

- **Ollama**: 10-50 requests/second (depending on model size)
- **vLLM**: 50-200+ requests/second (depending on hardware)

### **Memory Usage**

- **Ollama**: 2-4x model size in RAM
- **vLLM**: 1.5-2.5x model size in RAM

### **GPU Utilization**

- **Ollama**: Basic GPU support, limited optimization
- **vLLM**: Advanced GPU optimization, PagedAttention algorithm

### **Startup Time**

- **Ollama**: Fast startup, models loaded on-demand
- **vLLM**: Slower startup, models pre-loaded for performance

## Resource Requirements

### **Ollama Requirements**

```yaml
resources:
  requests:
    cpu: "2-4"
    memory: "8-16Gi"
  limits:
    cpu: "4-8"
    memory: "16-32Gi"
```

### **vLLM Requirements**

```yaml
resources:
  requests:
    cpu: "2-4"
    memory: "6-12Gi"
    nvidia.com/gpu: "1"  # Recommended
  limits:
    cpu: "4-8"
    memory: "12-24Gi"
    nvidia.com/gpu: "1"
```

## Key Differences Beyond Deployment

### **API Compatibility**

- **Ollama**: Uses Ollama's native API format, great if you're already familiar with Ollama.
- **vLLM**: Compatible with OpenAI's API format, making it easier to integrate with existing OpenAI-based applications.

### **Resource Efficiency**

- **Ollama**: More memory usage but simpler resource management.
- **vLLM**: Better memory efficiency and GPU utilization, but requires more careful resource planning.

### **Scaling Characteristics**

- **Ollama**: Good for horizontal scaling with multiple replicas.
- **vLLM**: Excellent for both horizontal scaling and single-instance performance optimization.

## Migration Path

### **From Ollama to vLLM**

If you're currently using Ollama and need better performance:

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
       model: "meta-llama/Llama-2-7b-chat-hf"
   ```

2. **Adjust Resources** (typically 20-30% less memory needed)
3. **Update Service Ports** (11434 → 8000)
4. **Test Performance** and adjust as needed

### **From vLLM to Ollama**

If you need more flexibility in model management:

1. **Simplify Configuration**
   ```yaml
   # Before (vLLM)
   spec:
     vLLM:
       enabled: true
       model: "meta-llama/Llama-2-7b-chat-hf"
   
   # After (Ollama)
   spec:
     ollama:
       models: ["llama2:7b"]
   ```

2. **Increase Memory Allocation** (typically 20-30% more needed)
3. **Update Service Ports** (8000 → 11434)

## Decision Matrix

### **Choose Ollama When:**

- ✅ You need flexibility in model management
- ✅ You're in development or testing phase
- ✅ You want to experiment with different models
- ✅ You have limited GPU resources
- ✅ You prefer Ollama's API format
- ✅ You want on-the-fly model switching

### **Choose vLLM When:**

- ✅ You need maximum performance
- ✅ You're deploying to production
- ✅ You have GPU resources available
- ✅ You need to handle high traffic
- ✅ You want advanced optimization features
- ✅ You prefer OpenAI-compatible API

### **Consider Both When:**

- 🔄 You want to A/B test performance
- 🔄 You have mixed workload requirements
- 🔄 You're migrating between environments
- 🔄 You want to compare user experience

## Getting Started

Both Ollama and vLLM are equally easy to deploy with the LLM Operator. Choose your path based on your performance and
flexibility needs:

- **[Ollama Deployment](./ollama)** - Flexible model management for development and testing
- **[vLLM Deployment](./vllm)** - High-performance setup for production workloads

Both guides include complete examples, configuration options, and best practices to get you up and running quickly.

