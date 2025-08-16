---
id: openwebui
title: Chat with LLMs using OpenWebUI
sidebar_label: OpenWebUI Chat
description: Learn how to deploy and use OpenWebUI for chatting with LLMs
---

# Chat with LLMs using OpenWebUI

OpenWebUI is a powerful web-based chat interface that provides an intuitive way to interact with Large Language Models. With Geeper.AI, you can easily deploy OpenWebUI instances and start chatting with your deployed LLMs.

## What is OpenWebUI?

OpenWebUI is an open-source web interface for LLMs that offers:
- **Modern Chat Interface**: Clean, responsive design similar to ChatGPT
- **Multi-Model Support**: Connect to various LLM backends
- **Conversation Management**: Save, export, and manage chat histories
- **Customizable UI**: Themes, layouts, and personalization options
- **API Integration**: RESTful API for programmatic access
- **User Management**: Multi-user support with authentication

## Quick Deployment

### 1. Deploy OpenWebUI with Geeper.AI

#### Minimal Setup
```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: Deployment
metadata:
  name: minimal-ollama
  namespace: default
spec:
  ollama:
    models:
      - "gemma3:270m"
  openwebui:
    enabled: true
    ingress:
      enabled: true
      host: "openwebui.localhost"
```

#### Production Setup
```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: OllamaDeployment
metadata:
  name: production-ollama
  namespace: ai-production
spec:
  ollama:
    replicas: 3
    image: ollama/ollama
    imageTag: latest
    service:
      type: LoadBalancer
      port: 11434
    resources:
      requests:
        cpu: "2"
        memory: "4Gi"
      limits:
        cpu: "8"
        memory: "16Gi"
    models:
      - "llama2:13b"
      - "codellama:34b"
      - "mistral:7b"
      - "phi:2.7b"
  
  openwebui:
    enabled: true
    replicas: 2
    image: ghcr.io/open-webui/open-webui
    imageTag: main
    service:
      type: ClusterIP
      port: 8080
    ingress:
      enabled: true
      host: "ai.company.com"
    resources:
      requests:
        cpu: "500m"
        memory: "1Gi"
      limits:
        cpu: "2"
        memory: "2Gi"
```

### 2. Apply the Configuration

```bash
kubectl apply -f openwebui-deployment.yaml
```

### 3. Access OpenWebUI

```bash
# Get the external IP
kubectl get svc openwebui-chat

# Open in browser
open http://<EXTERNAL-IP>:8080
```

## Configuration Options

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `WEBUI_SECRET_KEY` | Secret key for session management | Required |
| `DEFAULT_MODELS` | Comma-separated list of default models | `ollama/llama2` |
| `OLLAMA_BASE_URL` | Ollama API endpoint | `http://ollama:11434` |
| `WEBUI_AUTH` | Enable authentication | `false` |
| `WEBUI_AUTH_SECRET_KEY` | Authentication secret | Required if auth enabled |

### Resource Configuration

```yaml
spec:
  resources:
    requests:
      memory: "1Gi"
      cpu: "500m"
      nvidia.com/gpu: "1"  # If GPU acceleration needed
    limits:
      memory: "2Gi"
      cpu: "1000m"
      nvidia.com/gpu: "1"
```

### Service Configuration

```yaml
spec:
  service:
    type: ClusterIP  # or NodePort, LoadBalancer
    port: 8080
    targetPort: 8080
    annotations:
      # Custom annotations for your cloud provider
      service.beta.kubernetes.io/aws-load-balancer-type: nlb
```

## Connecting to LLM Backends

### Ollama Backend

1. **Deploy Ollama Service**:
```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: Deployment
metadata:
  name: production-ollama
  namespace: ai-production
spec:
  ollama:
    replicas: 3
    image: ollama/ollama
    imageTag: latest
    service:
      type: LoadBalancer
      port: 11434
    resources:
      requests:
        cpu: "2"
        memory: "4Gi"
      limits:
        cpu: "8"
        memory: "16Gi"
    models:
      - "llama2:13b"
      - "codellama:34b"
      - "mistral:7b"
      - "phi:2.7b"
```

2. **Configure OpenWebUI**:
```yaml
spec:
  openwebui:
    enabled: true
    replicas: 2
    image: ghcr.io/open-webui/open-webui
    imageTag: main
    service:
      type: ClusterIP
      port: 8080
    ingress:
      enabled: true
      host: "ai.company.com"
    resources:
      requests:
        cpu: "500m"
        memory: "1Gi"
      limits:
        cpu: "2"
        memory: "2Gi"
```

### External API Backend

```yaml
spec:
  env:
    - name: OPENAI_API_KEY
      value: "your-openai-api-key"
    - name: OPENAI_API_BASE_URL
      value: "https://api.openai.com/v1"
```

## Advanced Features

### Authentication

Enable user authentication:

```yaml
spec:
  env:
    - name: WEBUI_AUTH
      value: "true"
    - name: WEBUI_AUTH_SECRET_KEY
      value: "your-auth-secret"
    - name: WEBUI_AUTH_DISABLE_SIGNUP
      value: "false"
```

### Custom Themes

```yaml
spec:
  env:
    - name: WEBUI_THEME
      value: "dark"  # or light, auto
    - name: WEBUI_THEME_COLOR
      value: "#6366f1"
```

### Persistent Storage

```yaml
spec:
  volumeMounts:
    - name: chat-data
      mountPath: /app/backend/data
  volumes:
    - name: chat-data
      persistentVolumeClaim:
        claimName: openwebui-pvc
```

## Usage Examples

### Basic Chat

1. Open OpenWebUI in your browser
2. Select a model from the dropdown
3. Type your message and press Enter
4. The LLM will respond with generated text

### Conversation Management

- **Save Conversations**: Click the save icon to store chat history
- **Export Chats**: Export conversations as JSON or Markdown
- **Load Previous**: Access saved conversations from the sidebar

### Model Switching

- **Multiple Models**: Switch between different LLMs during a conversation
- **Model Comparison**: Compare responses from different models
- **Custom Models**: Add your own fine-tuned models

## Troubleshooting

### Common Issues

1. **Connection Refused**:
   - Check if Ollama service is running
   - Verify network policies allow communication
   - Check service endpoints

2. **Authentication Errors**:
   - Ensure `WEBUI_SECRET_KEY` is set
   - Check secret key format and length
   - Verify environment variables are loaded

3. **Resource Issues**:
   - Monitor resource usage with `kubectl top pods`
   - Adjust resource limits if needed
   - Check for OOMKilled events

### Debug Commands

```bash
# Check pod logs
kubectl logs -f deployment/openwebui-chat

# Check service endpoints
kubectl get endpoints openwebui-chat

# Test connectivity
kubectl exec -it deployment/openwebui-chat -- curl ollama-backend:11434/api/tags
```

## Best Practices

1. **Security**:
   - Use strong secret keys
   - Enable authentication for production
   - Restrict network access with NetworkPolicies

2. **Performance**:
   - Use appropriate resource limits
   - Enable GPU acceleration when available
   - Monitor and scale based on usage

3. **Monitoring**:
   - Set up Prometheus metrics
   - Configure alerting for failures
   - Monitor chat usage patterns

## Next Steps

- [RAG Integration](/docs/chat/rag) - Learn how to add Retrieval-Augmented Generation
- [Plugin System](/docs/chat/plugins) - Extend functionality with plugins
- [Advanced Configuration](/docs/chat/advanced) - Deep dive into configuration options
- [API Reference](/docs/api/openwebui) - Complete API documentation

---

*OpenWebUI provides a powerful and intuitive way to interact with your LLMs through Geeper.AI*
