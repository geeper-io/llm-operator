---
id: tabby
title: Tabby Integration
sidebar_label: Tabby
description: Learn how to integrate Tabby with Geeper.AI for intelligent code completion
---

# Tabby Integration

Tabby is an open-source, self-hosted AI coding assistant that provides intelligent code completion and generation. It's designed to work seamlessly with your Geeper.AI LMDeployments, offering privacy-focused AI coding assistance.

## What is Tabby?

Tabby is a self-hosted AI coding assistant that provides:

- **Intelligent Code Completion**: Context-aware suggestions
- **Multi-language Support**: Python, JavaScript, Go, Rust, and more
- **Privacy-First**: All code stays within your infrastructure
- **Customizable**: Configurable completion behavior and models
- **Enterprise Ready**: Built for production use

## Example deployment with Tabby

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: tabby-with-ollama
  namespace: default
spec:
  ollama:
    models:
      - "codellama:7b"
      - "llama2:7b"
    resources:
      requests:
        cpu: "2"
        memory: "4Gi"
      limits:
        cpu: "8"
        memory: "16Gi"
  
  tabby:
    enabled: true
    replicas: 2
    image: tabbyml/tabby:latest
    resources:
      requests:
        cpu: "500m"
        memory: "1Gi"
      limits:
        cpu: "2"
        memory: "4Gi"
    ingress:
      enabled: true
      host: "tabby.localhost"
```

## IDE Extension Support with Ingress

For IDE extensions to work properly through ingress, you need to enable WebSocket support. Tabby uses WebSockets for streaming responses and real-time communication with IDE extensions.

### Service Naming Convention

The operator automatically creates services with the pattern: `<deployment-name>-tabby`

**Example**: For a deployment named `tabby-with-ollama`, the Tabby service will be `tabby-with-ollama-tabby`

### WebSocket Annotations

#### HAProxy
```yaml
tabby:
  enabled: true
  ingress:
    host: "tabby.example.com"
    annotations:
      haproxy.org/ssl-passthrough: "true" # enables TCP mode
```

#### NGINX
```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
   name: tabby-with-ollama
   namespace: default
spec:
   …
   tabby:
     enabled: true
     ingress:
       host: "tabby.example.com"
       annotations:
         nginx.org/websocket-services: "tabby-with-ollama-tabby"  # Use actual service name
```

#### Custom NGINX Configuration
```nginx
location / {
    proxy_pass       http://tabby-with-ollama-tabby:8080;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
}
```

### Why WebSocket Support is Important

- **Real-time Communication**: IDE extensions need WebSockets for live code completion
- **Streaming Responses**: Tabby streams completion suggestions as you type
- **Low Latency**: WebSockets provide faster response times than HTTP polling
- **IDE Integration**: Most modern IDEs expect WebSocket connections for AI assistants

For detailed reverse proxy configuration, see the [official Tabby documentation](https://tabby.tabbyml.com/docs/administration/reverse-proxy/).

## Best Practices

### 1. Model Selection
- **Code-Specific Models**: Use models trained on code (e.g., CodeLlama)
- **Size vs. Speed**: Balance model size with completion speed
- **Domain Specialization**: Choose models for your programming languages

### 2. Performance Optimization
- **Resource Allocation**: Allocate sufficient CPU and memory
- **Scaling**: Scale based on user demand
- **Caching**: Enable completion caching for better performance

### 3. Security
- **Network Policies**: Restrict access to Tabby services
- **Authentication**: Implement proper authentication mechanisms
- **Code Privacy**: Ensure code never leaves your infrastructure

### 4. User Experience
- **Trigger Modes**: Choose between automatic and manual completion
- **Suggestion Quality**: Configure temperature and max tokens
- **Language Support**: Enable languages relevant to your team

## Troubleshooting

### Common Issues

1. **No Completions**:
   - Check Tabby server connectivity
   - Verify authentication credentials
   - Check file type support

2. **Slow Completions**:
   - Monitor resource usage
   - Check network latency
   - Optimize model configuration

3. **Poor Quality Suggestions**:
   - Adjust temperature settings
   - Check model quality
   - Verify context inclusion

### Debug Commands

```bash
# Check Tabby server status
curl http://your-tabby-endpoint:8080/health

# Test completion API
curl -X POST http://your-tabby-endpoint:8080/v1/completions \
  -H "Content-Type: application/json" \
  -d '{"prompt": "def hello", "max_tokens": 10}'

# View Tabby logs
kubectl logs -f deployment/tabby-server

# Check resource usage
kubectl top pods -l app=tabby
```

## Next Steps

- [Continue.dev Integration](/docs/coding-assistants/continue-dev) - Learn about Continue.dev
- [Advanced Configuration](/docs/coding-assistants/advanced-config) - Deep dive into settings
- [Custom Models](/docs/coding-assistants/custom-models) - Train custom completion models
- [API Reference](/docs/api/tabby) - Complete Tabby API documentation

---

*Tabby provides fast, private, and intelligent code completion powered by your Geeper.AI infrastructure*