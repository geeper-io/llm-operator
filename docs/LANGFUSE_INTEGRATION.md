# Langfuse Integration Guide

Langfuse provides comprehensive monitoring and observability for your LLM applications. This guide covers how to integrate Langfuse with the LLM Operator.

## Overview

Langfuse tracks and analyzes:
- **LLM API calls** and responses
- **Performance metrics** (latency, token usage, costs)
- **Error monitoring** and debugging
- **Usage patterns** and analytics
- **Pipeline execution** tracking

## Deployment Options

### 1. Langfuse Cloud (Recommended)

**Best for**: Production environments, teams, enterprise use
**Benefits**: Fully managed, automatic updates, enterprise features

```bash
# 1. Create account at https://langfuse.com
# 2. Get your API keys from the dashboard
# 3. Create Kubernetes secret
kubectl create secret generic langfuse-credentials \
  --from-literal=LANGFUSE_PUBLIC_KEY="your-public-key" \
  --from-literal=LANGFUSE_SECRET_KEY="your-secret-key" \
  -n default

# 4. Configure in your LMDeployment
```

**Configuration:**
```yaml
langfuse:
  enabled: true
  url: "https://cloud.langfuse.com"
  projectName: "my-llm-project"
  environment: "production"
  secretRef:
    name: "langfuse-credentials"
    namespace: "default"
```

### 2. Self-Hosted Langfuse

**Best for**: Data sovereignty, custom configurations, air-gapped environments

#### Quick Start with Helm

```bash
# Add Helm repository
helm repo add langfuse https://langfuse.github.io/langfuse-k8s
helm repo update

# Install with default configuration
helm install langfuse langfuse/langfuse -n default

# Get the service URL
kubectl get svc langfuse-web -n default
```

#### Production Configuration

Create `langfuse-values.yaml`:

```yaml
langfuse:
  resources:
    limits:
      cpu: "2"
      memory: "4Gi"
    requests:
      cpu: "2"
      memory: "4Gi"
  
  ingress:
    enabled: true
    hosts:
    - host: langfuse.your-domain.com
      paths:
      - path: /
        pathType: Prefix

postgresql:
  primary:
    persistence:
      enabled: true
      size: "20Gi"
      storageClass: "fast-ssd"

clickhouse:
  persistence:
    enabled: true
    size: "50Gi"
    storageClass: "fast-ssd"

redis:
  primary:
    persistence:
      enabled: true
      size: "10Gi"
      storageClass: "fast-ssd"
```

Install with custom values:
```bash
helm install langfuse langfuse/langfuse -f langfuse-values.yaml -n default
```

#### Self-Hosting Resources

- [Official Helm Chart](https://github.com/langfuse/langfuse-k8s)
- [Self-Hosting Documentation](https://langfuse.com/self-hosting)
- [Production Sizing Guide](https://github.com/langfuse/langfuse-k8s#sizing)
- [Troubleshooting Guide](https://github.com/langfuse/langfuse-k8s/blob/main/TROUBLESHOOTING.md)

## Integration with LLM Operator

### Automatic Pipeline Configuration

When Langfuse is enabled, the operator automatically:
1. **Configures OpenWebUI pipelines** with Langfuse monitoring
2. **Sets environment variables** for Langfuse connection
3. **Mounts credentials** from Kubernetes secrets
4. **Enables request tracking** for all LLM interactions

### Configuration Example

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: monitored-llm-app
  namespace: default
spec:
  ollama:
    models:
      - "llama3.2:1b"
      - "gemma3:270m"
  
  openwebui:
    enabled: true
    image: ghcr.io/open-webui/open-webui:main
    
    # Langfuse monitoring
    langfuse:
      enabled: true
      url: "https://langfuse.your-domain.com"  # Self-hosted
      # url: "https://cloud.langfuse.com"      # Cloud
      projectName: "llm-operator-demo"
      environment: "staging"
      secretRef:
        name: "langfuse-credentials"
        namespace: "default"
      debug: true
    
    # Pipelines (auto-enabled with Langfuse)
    pipelines:
      enabled: true
      image: ghcr.io/open-webui/pipelines:main
      persistence:
        enabled: true
        size: "5Gi"
```

### Secret Management

The operator expects a Kubernetes secret with Langfuse credentials:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: langfuse-credentials
  namespace: default
type: Opaque
stringData:
  LANGFUSE_PUBLIC_KEY: "your-public-key"
  LANGFUSE_SECRET_KEY: "your-secret-key"
```

## Monitoring Features

### Request Tracking

- **Input/Output**: Track all prompts and responses
- **Metadata**: Model used, parameters, timestamps
- **Performance**: Latency, token usage, costs
- **Errors**: Failed requests and error details

### Analytics Dashboard

Access your Langfuse dashboard to view:
- **Request Volume**: Number of API calls over time
- **Performance Metrics**: Average latency, token usage
- **Cost Analysis**: Token costs and usage patterns
- **Error Rates**: Failed request percentages
- **Model Usage**: Which models are used most

### Pipeline Monitoring

Track OpenWebUI pipeline execution:
- **Pipeline Steps**: Monitor each step in your RAG workflows
- **Execution Time**: Identify bottlenecks in your pipelines
- **Success Rates**: Track pipeline completion rates
- **Resource Usage**: Monitor pipeline resource consumption

## Best Practices

### 1. Environment Separation

```yaml
# Development
langfuse:
  environment: "development"
  projectName: "llm-app-dev"

# Staging  
langfuse:
  environment: "staging"
  projectName: "llm-app-staging"

# Production
langfuse:
  environment: "production"
  projectName: "llm-app-prod"
```

### 2. Resource Planning

**Self-Hosting Requirements:**
- **Langfuse**: 2 CPU, 4Gi RAM minimum
- **PostgreSQL**: 2 CPU, 8Gi RAM minimum
- **ClickHouse**: 2 CPU, 8Gi RAM minimum
- **Redis**: 1 CPU, 1.5Gi RAM minimum

### 3. Security

- **Use Kubernetes secrets** for credential storage
- **Enable RBAC** for secret access control
- **Network policies** to restrict access
- **TLS encryption** for external access

### 4. Monitoring

- **Health checks** for all components
- **Resource monitoring** (CPU, memory, disk)
- **Log aggregation** and analysis
- **Alerting** for critical issues

## Troubleshooting

### Common Issues

1. **Connection Failed**
   - Verify Langfuse URL is accessible
   - Check secret credentials are correct
   - Ensure network policies allow connection

2. **No Data in Dashboard**
   - Verify Langfuse is enabled in configuration
   - Check pipeline logs for errors
   - Verify secret is properly mounted

3. **High Resource Usage**
   - Adjust resource limits in Helm values
   - Consider scaling up database resources
   - Monitor ClickHouse performance

### Debug Mode

Enable debug logging to troubleshoot issues:

```yaml
langfuse:
  enabled: true
  debug: true  # Enables verbose logging
```

### Getting Help

- [Langfuse Documentation](https://langfuse.com/docs)
- [Langfuse Community](https://github.com/langfuse/langfuse/discussions)
- [Helm Chart Issues](https://github.com/langfuse/langfuse-k8s/issues)

## Migration from Self-Hosted to Cloud

If you want to migrate from self-hosted to Langfuse Cloud:

1. **Export Data**: Use Langfuse export features
2. **Update Configuration**: Change URL to cloud instance
3. **Migrate Credentials**: Update secret with cloud API keys
4. **Verify Integration**: Test monitoring functionality

The operator handles the transition seamlessly - just update the URL and credentials!
