---
id: pipelines
title: OpenWebUI Pipelines
sidebar_label: Pipelines
description: Advanced pipeline functionality for OpenWebUI with custom workflows and integrations
---

# OpenWebUI Pipelines

OpenWebUI Pipelines is a powerful feature that extends your AI chat application with custom workflows, filters, and integrations. It provides a UI-agnostic OpenAI API plugin framework that can be deployed alongside your OpenWebUI instance.

## What are Pipelines?

OpenWebUI Pipelines is a powerful feature that extends your AI chat application with custom workflows, filters, and integrations. It provides a UI-agnostic OpenAI API plugin framework that can be deployed alongside your OpenWebUI instance.

**Note**: Pipelines are automatically enabled when you enable Langfuse monitoring, as they are required for the integration.

Pipelines are Python-based workflows that can:

- **Filter and modify** requests/responses
- **Integrate with external services** (APIs, databases, etc.)
- **Add custom logic** to your AI conversations
- **Implement monitoring and observability**
- **Create custom authentication flows**
- **Add rate limiting and content filtering**

## Quick Start

### 1. Enable Pipelines in Your LMDeployment

#### Option A: Manual Pipeline Configuration

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: my-ai-app
spec:
  openwebui:
    enabled: true
    
    # Enable Pipelines manually
    pipelines:
      enabled: true
      image: ghcr.io/open-webui/pipelines:main
      replicas: 1
      port: 9099
      serviceType: ClusterIP
      
      # Configure pipeline directory
      pipelinesDir: "/app/pipelines"
      
      # Add custom pipeline URLs
      pipelineUrls:
        - "https://github.com/open-webui/pipelines/blob/main/examples/filters/detoxify_filter_pipeline.py"
        - "https://github.com/open-webui/pipelines/blob/main/examples/filters/rate_limit_filter_pipeline.py"
      
      # Enable persistence
      persistence:
        enabled: true
        size: "20Gi"
        storageClass: "fast-ssd"
```

#### Option B: Automatic Pipeline Enablement (Recommended)

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: my-ai-app
spec:
  openwebui:
    enabled: true
    
    # Enable Langfuse monitoring - Pipelines will be automatically enabled
    langfuse:
      enabled: true
      projectName: "my-project"
      # Pipelines will be automatically enabled and configured
```

### 2. Deploy and Verify

```bash
kubectl apply -f your-deployment.yaml
kubectl get pods -l app=pipelines
kubectl logs -l app=pipelines
```

## Configuration Options

### Basic Configuration

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `pipelines.enabled` | bool | false | Enable OpenWebUI Pipelines |
| `pipelines.image` | string | `ghcr.io/open-webui/pipelines:main` | Container image to use |
| `pipelines.replicas` | int32 | 1 | Number of pipeline pods (1-3) |
| `pipelines.port` | int32 | 9099 | Port for the pipeline service |
| `pipelines.serviceType` | string | ClusterIP | Kubernetes service type |

### Pipeline Sources

| Field | Type | Description |
|-------|------|-------------|
| `pipelines.pipelinesDir` | string | Directory containing pipeline definitions |
| `pipelines.pipelineUrls` | []string | URLs to fetch pipeline definitions from |

### Persistence Configuration

| Field | Type | Description |
|-------|------|-------------|
| `pipelines.persistence.enabled` | bool | Enable persistent storage |
| `pipelines.persistence.storageClass` | string | Storage class for PVCs |
| `pipelines.persistence.size` | string | Size of persistent volume |

### Resource Configuration

| Field | Type | Description |
|-------|------|-------------|
| `pipelines.resources.requests.cpu` | string | CPU requests |
| `pipelines.resources.requests.memory` | string | Memory requests |
| `pipelines.resources.limits.cpu` | string | CPU limits |
| `pipelines.resources.limits.memory` | string | Memory limits |

## Example Use Cases

### 1. Content Filtering

Filter inappropriate content from AI responses:

```yaml
pipelineUrls:
  - "https://github.com/open-webui/pipelines/blob/main/examples/filters/detoxify_filter_pipeline.py"
```

### 2. Rate Limiting

Implement rate limiting for API requests:

```yaml
pipelineUrls:
  - "https://github.com/open-webui/pipelines/blob/main/examples/filters/rate_limit_filter_pipeline.py"
```

### 3. Custom Monitoring

Add custom logging and monitoring:

```yaml
pipelineUrls:
  - "https://github.com/open-webui/pipelines/blob/main/examples/monitoring/custom_logging_pipeline.py"
```

### 4. External API Integration

Connect to external services:

```yaml
pipelineUrls:
  - "https://github.com/open-webui/pipelines/blob/main/examples/integrations/weather_api_pipeline.py"
```

## Production Example

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: production-ai-app
spec:
  openwebui:
    enabled: true
    replicas: 3
    
    # Production pipelines
    pipelines:
      enabled: true
      image: ghcr.io/open-webui/pipelines:main
      replicas: 2
      port: 9099
      serviceType: ClusterIP
      
      # Production pipeline URLs
      pipelineUrls:
        - "https://github.com/open-webui/pipelines/blob/main/examples/filters/detoxify_filter_pipeline.py"
        - "https://github.com/open-webui/pipelines/blob/main/examples/filters/rate_limit_filter_pipeline.py"
        # Note: Langfuse monitoring pipeline is automatically added when langfuse.enabled=true
      
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
    
    # Langfuse monitoring configuration
    langfuse:
      enabled: true
      url: "https://cloud.langfuse.com"
      publicKey: "your-langfuse-public-key"
      secretKey: "your-langfuse-secret-key"
      projectName: "production-ai-chat"
      environment: "production"
      debug: false
```

## How It Works

### Architecture

1. **Pipeline Deployment**: Separate Kubernetes deployment for pipeline processing
2. **Service Discovery**: OpenWebUI automatically connects to pipeline service
3. **Request Flow**: User requests → OpenWebUI → Pipeline → LLM → Pipeline → OpenWebUI → User
4. **Configuration**: Pipeline settings stored in ConfigMap and mounted to containers

### Automatic Integration

When you enable pipelines:

1. **Pipeline Deployment**: Creates a dedicated deployment for pipeline processing
2. **Pipeline Service**: Exposes pipelines on the configured port
3. **OpenWebUI Configuration**: Automatically configures OpenWebUI to use pipelines
4. **Persistent Storage**: Optionally creates PVCs for pipeline data

## Best Practices

### 1. Resource Management

- **Start Small**: Begin with 1 replica and scale based on usage
- **Monitor Resources**: Use resource limits to prevent pipeline pods from consuming too many resources
- **Storage Planning**: Plan persistent storage based on your pipeline data requirements

### 2. Pipeline Development

- **Use Official Examples**: Start with official pipeline examples from the repository
- **Test Locally**: Test pipelines locally before deploying to production
- **Version Control**: Store custom pipelines in version control
- **Documentation**: Document custom pipeline behavior and configuration

### 3. Security

- **Network Policies**: Restrict pipeline pod network access
- **RBAC**: Use appropriate service accounts and roles
- **Secrets**: Store sensitive configuration in Kubernetes secrets
- **Image Security**: Use trusted base images and scan for vulnerabilities

### 4. Monitoring

- **Logs**: Monitor pipeline pod logs for errors and performance
- **Metrics**: Use Kubernetes metrics to track resource usage
- **Health Checks**: Implement health checks for pipeline endpoints
- **Alerting**: Set up alerts for pipeline failures

## Troubleshooting

### Common Issues

1. **Pipeline Pod Not Starting**
   - Check resource limits and requests
   - Verify image pull permissions
   - Check pod events: `kubectl describe pod <pipeline-pod>`

2. **OpenWebUI Not Connecting to Pipelines**
   - Verify pipeline service is running
   - Check OpenWebUI configuration
   - Verify network policies allow communication

3. **Pipeline Errors**
   - Check pipeline logs: `kubectl logs <pipeline-pod>`
   - Verify pipeline URLs are accessible
   - Check Python dependencies in pipeline code

4. **Performance Issues**
   - Scale pipeline replicas
   - Optimize resource allocation
   - Review pipeline code for bottlenecks

### Debug Commands

```bash
# Check pipeline deployment status
kubectl get deployment -l app=pipelines

# View pipeline logs
kubectl logs -l app=pipelines -f

# Check pipeline service
kubectl get service -l app=pipelines

# Verify OpenWebUI configuration
kubectl get configmap -l app=openwebui -o yaml

# Test pipeline connectivity
kubectl exec -it <openwebui-pod> -- curl http://<pipeline-service>:9099/health
```

## Next Steps

- **Explore Examples**: Check out the [official pipeline examples](https://github.com/open-webui/pipelines/tree/main/examples)
- **Custom Development**: Create custom pipelines for your specific use cases
- **Integration**: Connect pipelines with external monitoring and observability tools
- **Community**: Join the OpenWebUI community for support and ideas

Pipelines provide a powerful way to extend your AI chat application with custom functionality, monitoring, and integrations. Start simple and gradually add complexity as you become more familiar with the system.

