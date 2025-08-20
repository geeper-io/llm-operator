---
id: observability
title: Observability with Langfuse
sidebar_label: Observability
description: Comprehensive monitoring and observability for your AI applications using Langfuse
---

# Observability with Langfuse

Langfuse provides comprehensive observability for your AI applications, giving you insights into model performance, user interactions, and system behavior. This guide shows you how to integrate Langfuse with Geeper.AI for production-ready monitoring.

## What is Langfuse?

Langfuse is an open-source LLM observability and monitoring platform that helps you:

- **Track LLM interactions** and their performance
- **Monitor costs** and usage patterns
- **Debug issues** with detailed request/response logs
- **Analyze user behavior** and model effectiveness
- **Set up alerts** for performance degradation
- **Generate reports** for stakeholders

## Quick Start

### 1. Set Up Langfuse

#### Option A: Langfuse Cloud
1. Go to [cloud.langfuse.com](https://cloud.langfuse.com)
2. Create a new account and project
3. Get your API keys from the project settings

#### Option B: Self-Hosted Langfuse 
```bash
# Install Langfuse using Helm
helm repo add langfuse https://langfuse.github.io/langfuse-k8s
helm repo update
helm install langfuse langfuse/langfuse -f values.yaml
```

See the Langfuse [documentation](https://docs.langfuse.com) and [Helm repo](https://github.com/langfuse/langfuse-k8s) for more details on self-hosting.

### 2. Enable Langfuse in Your LMDeployment

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: my-ai-app
spec:
  openwebui:
    enabled: true
    
    # Enable Langfuse monitoring with external URL
    langfuse:
      enabled: true
      url: "https://cloud.langfuse.com"  # or your self-hosted URL
      publicKey: "your-langfuse-public-key"
      secretKey: "your-langfuse-secret-key"
      projectName: "my-ai-project"
      environment: "production"
      debug: false
    
    # Pipelines will be automatically enabled for Langfuse monitoring
    # No need to explicitly set pipelines.enabled = true
```

### 3. Deploy and Monitor

```bash
kubectl apply -f your-deployment.yaml

# Check Langfuse integration
kubectl logs -l app=openwebui | grep -i langfuse

# If using self-hosted Langfuse, check the deployment
kubectl get pods -l app=langfuse
kubectl logs -l app=langfuse

# Check pipelines (automatically enabled)
kubectl get pods -l app=pipelines
kubectl logs -l app=pipelines
```

## Automatic Features

When you enable Langfuse monitoring, Geeper.AI automatically:

1. **Enables OpenWebUI Pipelines** - Required for Langfuse integration
2. **Adds Langfuse Monitoring Pipeline** - Automatically includes the official Langfuse monitoring pipeline
3. **Deploys Self-Hosted Langfuse** - If no external URL is provided, deploys a self-hosted instance
4. **Configures Environment Variables** - Automatically sets all necessary Langfuse configuration in OpenWebUI

This means you can enable comprehensive monitoring with just:

```yaml
langfuse:
  enabled: true
  projectName: "my-project"
```

## Configuration Options

### Langfuse Configuration

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `langfuse.enabled` | bool | Yes | Enable Langfuse monitoring |
| `langfuse.url` | string | Yes | Langfuse server URL |
| `langfuse.publicKey` | string | Yes | Langfuse public key |
| `langfuse.secretKey` | string | Yes | Langfuse secret key |
| `langfuse.projectName` | string | Yes | Name of the Langfuse project |
| `langfuse.environment` | string | No | Environment name (production, staging, development) |
| `langfuse.debug` | bool | No | Enable debug logging |

### Environment Variables

When Langfuse is enabled, these environment variables are automatically set in OpenWebUI:

| Variable | Description | Source |
|----------|-------------|---------|
| `LANGFUSE_PUBLIC_KEY` | Public key for authentication | `langfuse.publicKey` |
| `LANGFUSE_SECRET_KEY` | Secret key for authentication | `langfuse.secretKey` |
| `LANGFUSE_HOST` | Langfuse server URL | `langfuse.url` |
| `LANGFUSE_PROJECT` | Project name | `langfuse.projectName` |
| `LANGFUSE_ENVIRONMENT` | Environment name | `langfuse.environment` |
| `LANGFUSE_DEBUG` | Debug mode | `langfuse.debug` |

## What Gets Monitored

### 1. **LLM Interactions**
- **Request/Response pairs** with full context
- **Model performance** metrics (latency, token usage)
- **Cost tracking** per model and request
- **Prompt engineering** effectiveness

### 2. **Automatic Pipeline Integration**
- **Pipelines automatically enabled** when Langfuse is enabled
- **Langfuse monitoring pipeline** automatically added to pipeline URLs
- **Custom pipelines** can be combined with automatic monitoring

### 3. **User Behavior**
- **Session tracking** and user journeys
- **Feature usage** patterns
- **Error rates** and failure modes
- **Performance degradation** over time

### 4. **System Metrics**
- **Resource utilization** (CPU, memory, GPU)
- **Network latency** and throughput
- **Storage performance** and capacity
- **Kubernetes health** status

### 5. **Business Intelligence**
- **Usage trends** and growth patterns
- **Cost analysis** and optimization opportunities
- **User satisfaction** metrics
- **ROI calculations** for AI investments

## Advanced Configuration

### 1. **Multi-Environment Setup**

```yaml
# Development
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: dev-ai-app
spec:
  openwebui:
    langfuse:
      enabled: true
      url: "https://cloud.langfuse.com"
      publicKey: "dev-public-key"
      secretKey: "dev-secret-key"
      projectName: "my-ai-project"
      environment: "development"
      debug: true

---
# Production
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: prod-ai-app
spec:
  openwebui:
    langfuse:
      enabled: true
      url: "https://cloud.langfuse.com"
      publicKey: "prod-public-key"
      secretKey: "prod-secret-key"
      projectName: "my-ai-project"
      environment: "production"
      debug: false
```

### 2. **Using Kubernetes Secrets**

```yaml
# Create a secret for Langfuse credentials
apiVersion: v1
kind: Secret
metadata:
  name: langfuse-secrets
type: Opaque
data:
  public-key: <base64-encoded-public-key>
  secret-key: <base64-encoded-secret-key>
---
# Reference the secret in your LMDeployment
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: secure-ai-app
spec:
  openwebui:
    langfuse:
      enabled: true
      url: "https://cloud.langfuse.com"
      publicKey: "langfuse-secrets"
      secretKey: "langfuse-secrets"
      projectName: "my-ai-project"
      environment: "production"
```

### 3. **Custom Pipeline Integration**

```yaml
spec:
  openwebui:
    langfuse:
      enabled: true
      # ... Langfuse config
    
    pipelines:
      enabled: true
      pipelineUrls:
        - "https://github.com/open-webui/pipelines/blob/main/examples/monitoring/custom_metrics_pipeline.py"
        - "https://github.com/open-webui/pipelines/blob/main/examples/monitoring/performance_tracking_pipeline.py"
```

## Viewing Your Data

### 1. **Langfuse Dashboard**

Once deployed, visit your Langfuse instance to see:

- **Real-time metrics** and performance data
- **Request traces** with full context
- **Cost analysis** and usage patterns
- **User session** analytics
- **Model comparison** charts

### 2. **Key Metrics to Monitor**

#### **Performance Metrics**
- **Latency**: Response time per request
- **Throughput**: Requests per second
- **Error Rate**: Percentage of failed requests
- **Token Usage**: Input/output token consumption

#### **Cost Metrics**
- **Cost per Request**: Total cost divided by request count
- **Cost per Token**: Cost per input/output token
- **Monthly Spend**: Total cost over time periods
- **Cost by Model**: Comparison across different models

#### **Quality Metrics**
- **User Satisfaction**: Feedback scores and ratings
- **Completion Rate**: Percentage of successful conversations
- **Fallback Rate**: How often fallback responses are used
- **Context Retention**: Memory effectiveness over time

### 3. **Setting Up Alerts**

Configure alerts for:

- **High Error Rates**: >5% failure rate
- **Performance Degradation**: >2x latency increase
- **Cost Spikes**: >50% cost increase
- **Resource Exhaustion**: >90% resource usage

## Production Examples

### 1. **High-Traffic Production Setup**

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: production-ai-chat
spec:
  openwebui:
    enabled: true
    replicas: 5
    
    # Comprehensive Langfuse monitoring
    langfuse:
      enabled: true
      url: "https://cloud.langfuse.com"
      publicKey: "prod-public-key"
      secretKey: "prod-secret-key"
      projectName: "production-ai-chat"
      environment: "production"
      debug: false
    
    # Advanced pipelines for monitoring
    pipelines:
      enabled: true
      image: ghcr.io/open-webui/pipelines:main
      replicas: 3
      
      # Custom monitoring pipelines
      pipelineUrls:
        - "https://github.com/open-webui/pipelines/blob/main/examples/monitoring/performance_tracking_pipeline.py"
        - "https://github.com/open-webui/pipelines/blob/main/examples/monitoring/error_analysis_pipeline.py"
      
      # Enable persistence for monitoring data
      persistence:
        enabled: true
        size: "100Gi"
        storageClass: "fast-ssd"
    
    # Redis for session management
    redis:
      enabled: true
      image: redis:7-alpine
      replicas: 3
      persistence:
        enabled: true
        size: "50Gi"
```

### 2. **Multi-Region Monitoring**

```yaml
# US Region
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: us-ai-chat
  labels:
    region: us-east-1
spec:
  openwebui:
    langfuse:
      enabled: true
      url: "https://cloud.langfuse.com"
      projectName: "global-ai-chat"
      environment: "us-production"
      # ... other config

---
# EU Region
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: eu-ai-chat
  labels:
    region: eu-west-1
spec:
  openwebui:
    langfuse:
      enabled: true
      url: "https://cloud.langfuse.com"
      projectName: "global-ai-chat"
      environment: "eu-production"
      # ... other config
```

## Best Practices

### 1. **Security**
- **Use Secrets**: Store API keys in Kubernetes secrets
- **Network Policies**: Restrict access to Langfuse endpoints
- **RBAC**: Implement proper role-based access control
- **Audit Logging**: Track access to monitoring data

### 2. **Performance**
- **Batch Processing**: Group related metrics for efficiency
- **Sampling**: Use sampling for high-volume deployments
- **Caching**: Cache frequently accessed monitoring data
- **Compression**: Compress monitoring data in transit

### 3. **Reliability**
- **Redundancy**: Deploy multiple monitoring endpoints
- **Health Checks**: Monitor monitoring system health
- **Backup**: Regularly backup monitoring data
- **Testing**: Test monitoring setup in staging environments

### 4. **Cost Optimization**
- **Data Retention**: Set appropriate data retention policies
- **Sampling**: Use intelligent sampling for high-volume data
- **Storage Classes**: Use appropriate storage for different data types
- **Cleanup**: Implement automated data cleanup processes

## Troubleshooting

### Common Issues

1. **Langfuse Connection Failed**
   - Verify API keys are correct
   - Check network connectivity to Langfuse
   - Verify project name exists
   - Check firewall rules

2. **No Data in Dashboard**
   - Verify environment variables are set correctly
   - Check OpenWebUI logs for errors
   - Verify pipeline deployment is running
   - Check Langfuse project configuration

3. **High Latency**
   - Monitor network latency to Langfuse
   - Check pipeline performance
   - Verify resource allocation
   - Monitor queue depths

4. **Missing Metrics**
   - Verify pipeline URLs are accessible
   - Check pipeline logs for errors
   - Verify monitoring pipeline is loaded
   - Check environment variable configuration

### Debug Commands

```bash
# Check Langfuse environment variables
kubectl exec -it <openwebui-pod> -- env | grep LANGFUSE

# Test Langfuse connectivity
kubectl exec -it <openwebui-pod> -- curl -H "Authorization: Bearer $LANGFUSE_SECRET_KEY" $LANGFUSE_HOST/api/public/projects

# Check pipeline logs
kubectl logs -l app=pipelines -f

# Verify OpenWebUI configuration
kubectl get configmap -l app=openwebui -o yaml

# Check Langfuse integration in OpenWebUI
kubectl logs -l app=openwebui | grep -i langfuse
```

## Integration with Other Tools

### 1. **Grafana Dashboards**
- Import Langfuse data into Grafana
- Create custom dashboards
- Set up advanced alerting
- Historical trend analysis

### 2. **Prometheus Metrics**
- Export Langfuse metrics to Prometheus
- Use Grafana for visualization
- Set up custom alerting rules
- Long-term metric storage

### 3. **ELK Stack**
- Send Langfuse logs to Elasticsearch
- Use Logstash for processing
- Create Kibana dashboards
- Advanced log analysis

### 4. **Slack/Teams Integration**
- Real-time alerting
- Performance notifications
- Cost threshold alerts
- System health updates

## Next Steps

1. **Start Simple**: Begin with basic Langfuse monitoring
2. **Add Pipelines**: Enable pipelines for advanced monitoring
3. **Custom Metrics**: Create custom monitoring pipelines
4. **Alerting**: Set up proactive monitoring alerts
5. **Optimization**: Use data to optimize your AI applications

Langfuse provides enterprise-grade observability for your AI applications, helping you understand performance, control costs, and deliver better user experiences. Start monitoring today and gain insights into your AI operations!

