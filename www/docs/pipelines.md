---
id: pipelines
title: OpenWebUI Pipelines
sidebar_label: Pipelines
description: Extend OpenWebUI functionality with custom workflows, filters, and integrations
---

# OpenWebUI Pipelines

OpenWebUI Pipelines is a powerful plugin framework that extends OpenWebUI functionality with custom workflows, filters, and integrations. Based on the [official OpenWebUI Pipelines documentation](https://docs.openwebui.com/pipelines/), this feature allows you to easily add custom logic and integrate Python libraries.

## What are Pipelines?

Pipelines bring modular, customizable workflows to any UI client supporting OpenAI API specs. They allow you to:

- **Extend Functionality**: Easily add custom logic and integrate Python libraries
- **Create Custom Workflows**: Build sophisticated Retrieval-Augmented Generation (RAG) pipelines
- **Add Filters**: Implement rate limiting, toxic message detection, and more
- **Integrate External Services**: Connect with home automation APIs, monitoring tools, and more

## Key Features

- **Limitless Possibilities**: Easily add custom logic and integrate Python libraries
- **Seamless Integration**: Compatible with any UI/client supporting OpenAI API specs
- **Custom Hooks**: Build and integrate custom pipelines
- **Performance**: Offload computationally heavy tasks from your main OpenWebUI instance

## Quick Start

### 1. Enable Pipelines in Your LMDeployment

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: example-with-pipelines
spec:
  openwebui:
    enabled: true
    pipelines:
      enabled: true
      image: ghcr.io/open-webui/pipelines:main
      replicas: 1
      port: 9099
      resources:
        requests:
          cpu: "500m"
          memory: "1Gi"
        limits:
          cpu: "1"
          memory: "2Gi"
```

### 2. Apply the Configuration

```bash
kubectl apply -f deployment-with-pipelines.yaml
```

### 3. Verify Deployment

```bash
# Check pipeline deployment
kubectl get deployment -l app=pipelines

# Check pipeline service
kubectl get svc -l app=pipelines

# View pipeline logs
kubectl logs -l app=pipelines -f
```

## Configuration Options

### Basic Settings

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `enabled` | bool | false | Enable OpenWebUI Pipelines |
| `image` | string | `ghcr.io/open-webui/pipelines:main` | Pipelines container image |
| `replicas` | int32 | 1 | Number of Pipelines pods (1-3) |
| `port` | int32 | 9099 | Port the Pipelines service exposes |
| `serviceType` | string | ClusterIP | Type of service to expose |

### Advanced Configuration

| Field | Type | Description |
|-------|------|-------------|
| `pipelinesDir` | string | Directory containing pipeline definitions (default: `/app/pipelines`) |
| `pipelineUrls` | []string | List of URLs to fetch pipeline definitions from |
| `resources` | ResourceRequirements | CPU and memory requirements for Pipelines pods |
| `envVars` | []EnvVar | Custom environment variables |
| `volumeMounts` | []VolumeMount | Custom volume mounts |
| `volumes` | []Volume | Custom volumes |

### Persistence

| Field | Type | Description |
|-------|------|-------------|
| `persistence.enabled` | bool | Enable persistent storage for pipeline data |
| `persistence.size` | string | Size of persistent volume (default: "10Gi") |
| `persistence.storageClass` | string | Storage class for persistent volumes |

## Example Use Cases

### 1. Message Filtering

Add content filtering to prevent toxic messages:

```yaml
pipelineUrls:
  - "https://github.com/open-webui/pipelines/blob/main/examples/filters/detoxify_filter_pipeline.py"
```

### 2. Rate Limiting

Implement rate limiting to prevent abuse:

```yaml
pipelineUrls:
  - "https://github.com/open-webui/pipelines/blob/main/examples/filters/rate_limit_filter_pipeline.py"
```

### 3. Custom RAG Pipeline

Build sophisticated retrieval-augmented generation:

```yaml
pipelineUrls:
  - "https://github.com/open-webui/pipelines/blob/main/examples/rag/custom_rag_pipeline.py"
```

### 4. Function Calling

Handle function calls with custom logic:

```yaml
pipelineUrls:
  - "https://github.com/open-webui/pipelines/blob/main/examples/function_calling/function_pipeline.py"
```

## Production Example

Here's a complete production-ready configuration:

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: production-with-pipelines
  namespace: ai-production
spec:
  openwebui:
    enabled: true
    replicas: 3
    pipelines:
      enabled: true
      image: ghcr.io/open-webui/pipelines:main
      replicas: 2
      port: 9099
      serviceType: ClusterIP
      resources:
        requests:
          cpu: "1"
          memory: "2Gi"
        limits:
          cpu: "2"
          memory: "4Gi"
      
      # Production pipelines
      pipelineUrls:
        - "https://github.com/open-webui/pipelines/blob/main/examples/filters/detoxify_filter_pipeline.py"
        - "https://github.com/open-webui/pipelines/blob/main/examples/filters/rate_limit_filter_pipeline.py"
        - "https://github.com/open-webui/pipelines/blob/main/examples/monitoring/langfuse_monitor_pipeline.py"
      
      # Enable persistence
      persistence:
        enabled: true
        size: "50Gi"
        storageClass: "fast-ssd"
      
      # Custom environment variables
      envVars:
        - name: PIPELINES_DEBUG
          value: "false"
        - name: PIPELINES_LOG_LEVEL
          value: "warn"
        - name: LANGFUSE_PUBLIC_KEY
          valueFrom:
            secretKeyRef:
              name: langfuse-secrets
              key: public-key
        - name: LANGFUSE_SECRET_KEY
          valueFrom:
            secretKeyRef:
              name: langfuse-secrets
              key: secret-key
```

## How It Works

1. **Deployment**: The operator creates a separate Pipelines deployment alongside OpenWebUI
2. **Configuration**: OpenWebUI is automatically configured to connect to the Pipelines service
3. **Integration**: All requests flow through the pipeline before reaching the LLM
4. **Customization**: You can add custom Python code and dependencies

## Architecture

```
User Request → OpenWebUI → Pipelines → LLM (Ollama)
                ↓
            Pipeline Processing
            - Filters
            - Custom Logic
            - External Integrations
```

## Best Practices

1. **Start Simple**: Begin with basic filters before building complex pipelines
2. **Resource Management**: Allocate sufficient CPU and memory for pipeline processing
3. **Persistence**: Enable persistence for production deployments to preserve pipeline data
4. **Monitoring**: Monitor pipeline performance and resource usage
5. **Security**: Only use pipelines from trusted sources

## Troubleshooting

### Common Issues

1. **Pipeline Not Loading**
   - Check if the pipeline URL is accessible
   - Verify the pipeline has no additional dependencies
   - Check pipeline logs: `kubectl logs -l app=pipelines`

2. **Performance Issues**
   - Monitor resource usage: `kubectl top pods -l app=pipelines`
   - Consider increasing CPU/memory limits
   - Check if pipelines are processing requests correctly

3. **Connection Issues**
   - Verify the Pipelines service is running
   - Check OpenWebUI configuration includes the pipeline connection
   - Test connectivity: `kubectl run test --rm -i --tty --image=curlimages/curl -- curl http://<pipelines-service>:9099/health`

### Debug Commands

```bash
# Check pipeline deployment status
kubectl get deployment -l app=pipelines

# View pipeline logs
kubectl logs -l app=pipelines -f

# Check pipeline service
kubectl get svc -l app=pipelines

# Test pipeline endpoint
kubectl run test --rm -i --tty --image=curlimages/curl -- \
  curl http://<pipelines-service>:9099/health
```

## Next Steps

- [Example Configuration](../examples/openwebui-with-pipelines.yaml) - Complete pipeline setup
- [OpenWebUI Documentation](https://docs.openwebui.com/pipelines/) - Official pipeline documentation
- [Pipeline Examples](https://github.com/open-webui/pipelines/tree/main/examples) - Community pipeline examples
- [Plugin System](../plugin-system) - Learn about OpenWebUI plugins
- [Continue.dev Integration](../coding-assistants/continue-dev) - Use pipelines with Continue.dev

---

*OpenWebUI Pipelines provide limitless possibilities for extending your AI chat experience with custom workflows and integrations.*
