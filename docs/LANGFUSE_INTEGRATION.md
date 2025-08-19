# Langfuse Integration with Geeper.AI

Langfuse is a powerful LLM observability and monitoring platform that provides comprehensive insights into your AI applications. This guide shows you how to integrate Langfuse with your Geeper.AI deployments using OpenWebUI Pipelines.

## What is Langfuse?

Langfuse is an open-source LLM observability platform that helps you:

- **Track LLM Requests**: Monitor all interactions with your language models
- **Analyze Performance**: Measure latency, token usage, and costs
- **Evaluate Quality**: Assess response relevance and accuracy
- **Optimize Costs**: Monitor and control AI spending
- **Debug Issues**: Identify and resolve problems quickly

## Prerequisites

- Geeper.AI operator deployed in your cluster
- Langfuse account (cloud or self-hosted)
- OpenWebUI Pipelines enabled in your deployment

## Quick Start

### 1. Get Langfuse Credentials

1. **Sign up** at [Langfuse Cloud](https://cloud.langfuse.com) or deploy self-hosted
2. **Create a project** for your AI application
3. **Generate API keys** (Public Key and Secret Key)
4. **Note your project name** and environment

### 2. Enable Langfuse in Your LMDeployment

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: my-ai-app
spec:
  openwebui:
    enabled: true
    
    # Enable Langfuse monitoring
    langfuse:
      enabled: true
      url: "https://cloud.langfuse.com"  # or your self-hosted URL
      publicKey: "your-public-key"
      secretKey: "your-secret-key"
      projectName: "my-ai-project"
      environment: "production"
      debug: false
    
    # Enable Pipelines (optional, for additional functionality)
    pipelines:
      enabled: true
      image: ghcr.io/open-webui/pipelines:main
      
      # Add Langfuse monitoring pipeline
      pipelineUrls:
        - "https://github.com/open-webui/pipelines/blob/main/examples/monitoring/langfuse_monitor_pipeline.py"
```

**💡 Automatic Pipeline Inclusion**: When you enable Langfuse monitoring, the Langfuse monitoring pipeline is automatically added to your pipelines configuration. You don't need to manually specify it unless you want to customize the pipeline URL.

### 3. Apply the Configuration

```bash
kubectl apply -f deployment-with-langfuse.yaml
```

### 4. Verify Integration

```bash
# Check pipeline deployment
kubectl get deployment -l app=pipelines

# Check pipeline logs for Langfuse connection
kubectl logs -l app=pipelines -f

# Look for Langfuse connection messages
```

## Configuration Options

### Basic Langfuse Configuration

| Field | Required | Description |
|-------|----------|-------------|
| `enabled` | Yes | Set to `true` to enable Langfuse monitoring |
| `url` | Yes | Langfuse server URL |
| `publicKey` | Yes | Your Langfuse public key |
| `secretKey` | Yes | Your Langfuse secret key |
| `projectName` | Yes | Name of your Langfuse project |
| `environment` | No | Environment name (default: "production") |
| `debug` | No | Enable debug logging (default: false) |

### Langfuse URL Options

- **Langfuse Cloud**: `https://cloud.langfuse.com`
- **Self-hosted**: `http://your-langfuse-instance:3000`
- **Custom domain**: `https://langfuse.yourdomain.com`

## Advanced Configuration

### Production Setup with Secrets

For production deployments, use Kubernetes secrets for sensitive data:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: langfuse-credentials
type: Opaque
data:
  public-key: <base64-encoded-public-key>
  secret-key: <base64-encoded-secret-key>
---
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: production-ai-app
spec:
  openwebui:
    enabled: true
    
    # Langfuse monitoring configuration
    langfuse:
      enabled: true
      url: "https://cloud.langfuse.com"
      publicKey: "your-public-key"  # Use secret in production
      secretKey: "your-secret-key"  # Use secret in production
      projectName: "production-ai"
      environment: "production"
    
    # Enable pipelines for additional functionality
    pipelines:
      enabled: true
      image: ghcr.io/open-webui/pipelines:main
      pipelineUrls:
        - "https://github.com/open-webui/pipelines/blob/main/examples/filters/langfuse_filter_pipeline.py"
```

### Multi-Environment Setup

Configure different environments for development, staging, and production:

```yaml
# Development
langfuse:
  enabled: true
  url: "https://cloud.langfuse.com"
  publicKey: "dev-public-key"
  secretKey: "dev-secret-key"
  projectName: "ai-app-dev"
  environment: "development"

# Staging
langfuse:
  enabled: true
  url: "https://cloud.langfuse.com"
  publicKey: "staging-public-key"
  secretKey: "staging-secret-key"
  projectName: "ai-app-staging"
  environment: "staging"

# Production
langfuse:
  enabled: true
  url: "https://cloud.langfuse.com"
  publicKey: "prod-public-key"
  secretKey: "prod-secret-key"
  projectName: "ai-app-prod"
  environment: "production"
```

## What Gets Monitored

When Langfuse is enabled, the following data is automatically tracked:

### Request Information
- **User Input**: The original user message or prompt
- **Model Used**: Which LLM model processed the request
- **Timestamp**: When the request was made
- **User ID**: Identifier for the user (if available)

### Response Data
- **Model Output**: The generated response
- **Token Usage**: Number of input and output tokens
- **Latency**: Time taken to generate the response
- **Cost**: Estimated cost of the request

### Metadata
- **Pipeline Information**: Which pipelines processed the request
- **Filter Results**: Content filtering and rate limiting outcomes
- **Error Information**: Any errors that occurred during processing

## Viewing Your Data

### 1. Langfuse Dashboard

1. **Log into** your Langfuse account
2. **Navigate to** your project
3. **View the dashboard** showing:
   - Request volume and trends
   - Performance metrics
   - Cost analysis
   - Quality scores

### 2. Key Metrics

- **Request Volume**: Number of requests over time
- **Latency**: Response time distribution
- **Token Usage**: Input/output token consumption
- **Costs**: Spending trends and breakdowns
- **Quality**: Response relevance scores

### 3. Request Details

Click on individual requests to see:
- **Full conversation context**
- **Model parameters used**
- **Pipeline processing steps**
- **Performance metrics**
- **Cost breakdown**

## Best Practices

### 1. Security
- **Use secrets** for API keys in production
- **Rotate keys** regularly
- **Limit access** to Langfuse credentials
- **Monitor usage** for unusual patterns

### 2. Data Management
- **Set retention policies** for old data
- **Anonymize sensitive information** if needed
- **Comply with data regulations** (GDPR, CCPA, etc.)
- **Regular backups** of important data

### 3. Performance
- **Monitor pipeline latency** impact
- **Optimize batch sizes** for high-volume deployments
- **Use appropriate log levels** (avoid debug in production)
- **Scale pipelines** based on monitoring needs

### 4. Cost Optimization
- **Track token usage** patterns
- **Identify expensive models** and requests
- **Optimize prompts** to reduce token consumption
- **Set up alerts** for cost thresholds

## Troubleshooting

### Common Issues

1. **Connection Failed**
   - Verify Langfuse URL is accessible
   - Check API keys are correct
   - Ensure network policies allow outbound connections

2. **No Data Appearing**
   - Check pipeline logs for errors
   - Verify project name matches exactly
   - Ensure environment is set correctly

3. **High Latency**
   - Monitor pipeline resource usage
   - Check Langfuse server performance
   - Consider using self-hosted instance for better performance

### Debug Commands

```bash
# Check pipeline status
kubectl get pods -l app=pipelines

# View pipeline logs
kubectl logs -l app=pipelines -f

# Check environment variables
kubectl exec -it <pipeline-pod> -- env | grep LANGFUSE

# Test Langfuse connectivity
kubectl run test --rm -i --tty --image=curlimages/curl -- \
  curl -H "Authorization: Bearer <your-secret-key>" \
  https://cloud.langfuse.com/api/public/projects
```

### Log Analysis

Look for these messages in pipeline logs:

```
✅ Langfuse connection established
✅ Project "my-project" found
✅ Monitoring enabled for environment "production"
❌ Langfuse connection failed: Invalid API key
❌ Project "invalid-project" not found
```

## Next Steps

- [OpenWebUI Pipelines](PIPELINES.md) - Learn more about pipeline capabilities
- [Example Configurations](../examples/openwebui-with-langfuse.yaml) - Complete setup examples
- [Langfuse Documentation](https://langfuse.com/docs) - Official Langfuse guides
- [Monitoring Best Practices](MONITORING.md) - Production monitoring strategies

---

*Langfuse integration provides comprehensive observability into your AI applications, helping you optimize performance, control costs, and deliver better user experiences.*
