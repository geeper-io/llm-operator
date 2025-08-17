# Quick Start Guide

This guide will help you get up and running with the LLM Operator in under 10 minutes.

## Prerequisites

- Kubernetes cluster (1.24+)
- kubectl configured and pointing to your cluster
- make installed
- At least 2GB of available memory in your cluster

## Step 1: Clone and Build

```bash
# Clone the repository
git clone <repository-url>
cd llm-operator

# Build the operator
make build
```

## Step 2: Deploy the Operator

```bash
# Install CRDs
make install

# Deploy the operator
make deploy

# Wait for the operator to be ready
kubectl wait --for=condition=available --timeout=300s deployment/llm-operator-controller-manager -n llm-operator-system
```

## Step 3: Deploy Your First Ollama Instance

Create a simple LMDeployment:

```bash
cat <<EOF | kubectl apply -f -
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: my-first-ollama
  namespace: default
spec:
  ollama:
    models:
      - "llama2:7b"
      - "mistral:7b"
EOF
```

## Step 4: Monitor the LMDeployment

Check the status of your LMDeployment:

```bash
# Check the status
kubectl get lmdeployment my-first-ollama

# Watch the LMDeployment progress
kubectl get lmdeployment my-first-ollama -w

# Check created resources
kubectl get all -l ollama-deployment=my-first-ollama
```

## Step 5: Access Ollama

Once the LMDeployment is ready, you can access Ollama:

```bash
# Port forward to the Ollama service
kubectl port-forward svc/my-first-ollama-ollama 11434:11434

# In another terminal, test with curl
curl -X POST http://localhost:11434/api/generate \
  -H "Content-Type: application/json" \
  -d '{
    "model": "llama2:7b",
    "prompt": "Hello, how are you?",
    "stream": false
  }'
```

## Step 6: Add OpenWebUI (Optional)

To add a web interface, update your LMDeployment:

```bash
kubectl patch lmdeployment my-first-ollama --type='merge' -p='{
  "spec": {
    "openwebui": {
      "enabled": true,
      "ingressEnabled": true,
      "ingressHost": "ollama-webui.local"
    }
  }
}'
```

## Step 7: Access OpenWebUI

```bash
# Port forward to OpenWebUI
kubectl port-forward svc/my-first-ollama-openwebui 8080:8080

# Open http://localhost:8080 in your browser
```

## What Happens Under the Hood

The operator automatically creates:

1. **Ollama LMDeployment**: Runs the Ollama server with postStart hooks to pull models
2. **Ollama Service**: Exposes Ollama on port 11434
3. **OpenWebUI LMDeployment**: Web interface (if enabled)
4. **OpenWebUI Service**: Exposes the web UI on port 8080
5. **Ingress**: External access (if configured)

## Advanced: Enable Pipelines

OpenWebUI Pipelines provide powerful extensibility for custom workflows and integrations. To enable pipelines:

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: my-first-ollama
spec:
  ollama:
    models:
      - "llama2:7b"
  
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
      
      # Add example pipelines
      pipelineUrls:
        - "https://github.com/open-webui/pipelines/blob/main/examples/filters/detoxify_filter_pipeline.py"
      
      # Enable persistence
      persistence:
        enabled: true
        size: "10Gi"
```

Apply the updated configuration:

```bash
kubectl apply -f deployment-with-pipelines.yaml
```

The operator will automatically:
1. **Deploy Pipelines**: Create a separate Pipelines deployment
2. **Configure OpenWebUI**: Automatically connect OpenWebUI to Pipelines
3. **Set up Persistence**: Create persistent storage for pipeline data

For complete pipeline documentation, see [OpenWebUI Pipelines](PIPELINES.md).

## Next Steps

### Scale Your LMDeployment

```bash
# Scale to 3 replicas
kubectl patch lmdeployment my-first-ollama --type='merge' -p='{
  "spec": {
    "ollama": {
      "replicas": 3
    }
  }
}'
```

### Add More Models

```bash
# Add a new model
kubectl patch lmdeployment my-first-ollama --type='merge' -p='{
  "spec": {
    "ollama": {
              "models": [
          "llama2:7b",
          "mistral:7b",
          "codellama:7b"
        ]
    }
  }
}'
```

### Configure Resources

```bash
# Set resource limits
kubectl patch lmdeployment my-first-ollama --type='merge' -p='{
  "spec": {
    "ollama": {
      "resources": {
        "requests": {"cpu": "1", "memory": "2Gi"},
        "limits": {"cpu": "4", "memory": "8Gi"}
      }
    }
  }
}'
```

## Troubleshooting

### Check Operator Logs

```bash
kubectl logs -n llm-operator-system deployment/llm-operator-controller-manager
```

### Check Pod Status

```bash
kubectl get pods -l ollama-deployment=my-first-ollama
kubectl describe pod <pod-name>
```

### Check Model Pulling

```bash
# Check postStart hook execution in Ollama container logs
kubectl logs <pod-name> -c ollama
```

## Cleanup

```bash
# Delete the LMDeployment
kubectl delete lmdeployment my-first-ollama

# Undeploy the operator
make undeploy

# Uninstall CRDs
make uninstall
```

## Common Issues

1. **Out of Memory**: Ensure your cluster has sufficient memory for models
2. **Model Pull Failures**: Check network connectivity and model availability
3. **Service Not Accessible**: Verify service types and port configurations

## Getting Help

- Check the [CRD Reference](CRD_REFERENCE.md) for detailed API documentation
- Review the [README](../README.md) for comprehensive information
- Open an issue in the project repository for bugs or questions

## Production Considerations

For production LMDeployments:

1. **Use specific image tags** instead of `latest`
2. **Set appropriate resource limits** based on model requirements
3. **Configure persistent storage** for model persistence
4. **Implement network policies** for security
5. **Use multiple replicas** for high availability
6. **Monitor resource usage** and scale accordingly