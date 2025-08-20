# LLM Operator

The LLM Operator provides a Custom Resource Definition (CRD) called `LMDeployment` that allows you to declaratively deploy Ollama instances with specified models and optionally connect them to OpenWebUI for a web-based interface.

## Overview

The LLM Operator provides a Custom Resource Definition (CRD) called `LMDeployment` that allows you to declaratively deploy Ollama instances with specified models and optionally connect them to OpenWebUI for a web-based interface.

## Features

- **🚀 Easy Deployment**: Deploy LLM services with simple YAML configurations
- **🔧 Multi-Model Support**: Support for various LLM frameworks (Ollama, OpenWebUI, Tabby, etc.)
- **⚡ Auto-scaling**: Automatic scaling based on demand and resource usage
- **🔒 Security**: Built-in security features and RBAC integration
- **🔐 Automatic Secret Management**: Automatically generates and manages secure secrets for OpenWebUI
- **📊 Monitoring**: Comprehensive metrics and observability
- **🌐 Multi-cluster**: Support for multi-cluster LMDeployments
- **⚡ Pipelines**: OpenWebUI Pipelines for custom workflows and integrations
- **🔄 Redis Integration**: Automatic Redis deployment for multi-replica OpenWebUI

## Architecture

The operator creates and manages several Kubernetes resources:

- **Services**: For Ollama, OpenWebUI (if enabled), and Tabby (if enabled)
- **LMDeployments**: For Ollama, OpenWebUI (if enabled), and Tabby (if enabled)
- **ConfigMaps**: For configuration settings
- **Secrets**: For sensitive configuration data
- **PersistentVolumeClaims**: For persistent storage (if enabled)

## Security Features

### Automatic Secret Management

The operator automatically generates and manages secure secrets for OpenWebUI deployments:

- **WEBUI_SECRET_KEY**: A cryptographically secure random key (256-bit) is automatically generated
- **Secret Storage**: Secrets are stored as Kubernetes Secret resources with proper RBAC controls
- **Automatic Rotation**: Each deployment gets a unique secret key
- **Secure Generation**: Uses crypto/rand for cryptographically secure random key generation
- **Error Handling**: Returns errors if secret generation fails, ensuring no insecure fallbacks

The secret is automatically created with the name pattern: `{deployment-name}-openwebui-secret` and contains the `WEBUI_SECRET_KEY` environment variable that OpenWebUI requires for secure operation.

## Installation

### Prerequisites

- Kubernetes cluster (1.24+)
- kubectl configured
- make (for building)

### Build and Deploy

```bash
# Build the operator
make build

# Deploy to cluster
make deploy

# Verify installation
kubectl get crd | grep ollama
```

## Examples

### LMDeployment

The `LMDeployment` CRD supports various configurations:

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: ollama-example
spec:
  ollama:
    models:
      - "llama2:7b"
      - "mistral:7b"
    replicas: 1
    image: ollama/ollama:latest
    resources:
      requests:
        cpu: "500m"
        memory: "2Gi"
      limits:
        cpu: "2"
        memory: "4Gi"
```

### Advanced Configuration

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: production-ollama
  namespace: ai-models
spec:
  ollama:
    replicas: 3
    image: ollama/ollama
    imageTag: latest
    serviceType: LoadBalancer
    servicePort: 11434
    resources:
      requests:
        cpu: "1"
        memory: "2Gi"
      limits:
        cpu: "4"
        memory: "8Gi"
    models:
      - "llama2:13b"
      - "codellama:34b"
      - "mistral:7b"
  
  openwebui:
    enabled: true
    replicas: 2
    ingressEnabled: true
    ingressHost: "ollama-webui.local"
    resources:
      requests:
        cpu: "250m"
        memory: "512Mi"
      limits:
        cpu: "1000m"
        memory: "1Gi"
  
  tabby:
    enabled: true
    replicas: 2
    image: "tabbyml/tabby"
    imageTag: "latest"
    serviceType: ClusterIP
    servicePort: 8080
    ingressEnabled: true
    ingressHost: "tabby.local"
    # Tabby will automatically use the first model from Ollama
    # You can override with: modelName: "codellama:34b"
    resources:
      requests:
        cpu: "250m"
        memory: "512Mi"
      limits:
        cpu: "1000m"
        memory: "1Gi"
```

### Tabby Code Completion

Tabby provides AI-powered code completion by connecting to your Ollama models. It automatically:

- Connects to the deployed Ollama service via configuration file
- Uses the first specified model (or a custom one you specify)
- Provides a REST API for code completion
- Supports multiple programming languages
- Can be accessed via ingress for external integration
- Mounts configuration to `/data/config.toml` for easy customization
- Includes local embedding model (Nomic-Embed-Text) by default

#### IDE Extension Support with Ingress

For IDE extensions to work properly through ingress, you need to enable WebSocket support. Tabby uses WebSockets for streaming responses and real-time communication.

**Service Naming Convention:**
The operator automatically creates services with the pattern: `<deployment-name>-tabby`

**Example**: For a deployment named `minimal-ollama`, the Tabby service will be `minimal-ollama-tabby`

**WebSocket Annotations:**

**HAProxy:**
```yaml
tabby:
  enabled: true
  ingress:
    host: "tabby.example.com"
    annotations:
      haproxy.org/ssl-passthrough: "true"
      haproxy.org/websocket-services: "ws-svc"
```

**NGINX:**
```yaml
tabby:
  enabled: true
  ingress:
    host: "tabby.example.com"
    annotations:
      nginx.org/websocket-services: "ws-svc"
      nginx.org/websocket-services: "minimal-ollama-tabby"  # Use actual service name
```

**Custom NGINX Configuration:**
```nginx
location / {
    proxy_pass       http://minimal-ollama-tabby:8080;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection "upgrade";
}
```

For detailed reverse proxy configuration, see the [official Tabby documentation](https://tabby.tabbyml.com/docs/administration/reverse-proxy/).

### OpenWebUI with Automatic Secret Management

OpenWebUI deployments automatically include secure secret management:

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: openwebui-secure
  namespace: default
spec:
  ollama:
    enabled: true
    replicas: 1
    image: ollama/ollama:latest
  
  openwebui:
    enabled: true
    replicas: 1
    image: ghcr.io/open-webui/open-webui:main
    service:
      type: ClusterIP
      port: 8080
    
    # The WEBUI_SECRET_KEY is automatically generated and managed
    # No need to specify it manually - the operator handles everything
    
    ingress:
      host: openwebui.example.com
    
    # Optional: Enable Redis for multi-replica deployments
    redis:
      enabled: true
      image: redis:7-alpine
      password: "your-redis-password"
```

**Automatic Features:**
- 🔐 **WEBUI_SECRET_KEY**: Automatically generated as a 256-bit secure random key
- 🔒 **Secret Storage**: Stored as Kubernetes Secret with proper RBAC controls
- 🔄 **Unique Keys**: Each deployment gets its own unique secret key
- 🛡️ **Secure Generation**: Uses cryptographically secure random number generation
- ❌ **No Fallbacks**: Returns errors if generation fails, ensuring security

## API Reference

### Deployment

| Field | Type | Description | Default |
|-------|------|-------------|---------|
| `spec.ollama.replicas` | int32 | Number of Ollama pods | 1 |
| `spec.ollama.image` | string | Ollama container image | `ollama/ollama` |
| `spec.ollama.imageTag` | string | Ollama image tag | `latest` |
| `spec.ollama.service.type` | string | Service type (ClusterIP, NodePort, LoadBalancer) | `ClusterIP` |
| `spec.ollama.service.port` | int32 | Service port | `11434` |
| `spec.ollama.models` | []OllamaModel | List of models to deploy | Required |
| `spec.ollama.resources` | ResourceRequirements | Resource limits and requests | None |

### OllamaModel

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | Model name (e.g., "llama2", "mistral") |
| `tag` | string | Model tag/version (e.g., "7b", "13b") |
| `pullPolicy` | string | Pull policy (Always, IfNotPresent, Never) |

### OpenWebUI Configuration

| Field | Type | Description | Default |
|-------|------|-------------|---------|
| `spec.openwebui.enabled` | bool | Enable OpenWebUI deployment | false |
| `spec.openwebui.replicas` | int32 | Number of OpenWebUI pods | 1 |
| `spec.openwebui.image` | string | OpenWebUI container image | `ghcr.io/open-webui/open-webui` |
| `spec.openwebui.imageTag` | string | OpenWebUI image tag | `main` |
| `spec.openwebui.ingress.enabled` | bool | Enable ingress | false |
| `spec.openwebui.ingress.host` | string | Ingress hostname | None |

### Tabby Configuration

| Field | Type | Description | Default |
|-------|------|-------------|---------|
| `spec.tabby.enabled` | bool | Enable Tabby deployment | false |
| `spec.tabby.replicas` | int32 | Number of Tabby pods | 1 |
| `spec.tabby.image` | string | Tabby container image | `tabbyml/tabby` |
| `spec.tabby.imageTag` | string | Tabby image tag | `latest` |
| `spec.tabby.chatModel` | string | Ollama model for chat functionality | Required |
| `spec.tabby.completionModel` | string | Ollama model for code completion | Required |
| `spec.tabby.persistence.enabled` | bool | Enable data persistence | false |
| `spec.tabby.persistence.size` | string | Storage size for persistence | "10Gi" |

**Note**: Both `chatModel` and `completionModel` must be specified and must exist in `spec.ollama.models`.

### Langfuse Monitoring

Langfuse provides comprehensive monitoring and observability for your LLM applications:

| Field | Type | Description | Default |
|-------|------|-------------|---------|
| `spec.openwebui.langfuse.enabled` | bool | Enable Langfuse monitoring | false |
| `spec.openwebui.langfuse.url` | string | External Langfuse instance URL | Required |
| `spec.openwebui.langfuse.projectName` | string | Langfuse project name | Deployment name |
| `spec.openwebui.langfuse.environment` | string | Environment name | "development" |
| `spec.openwebui.langfuse.secretRef.name` | string | Kubernetes secret name | Required |
| `spec.openwebui.langfuse.secretRef.namespace` | string | Secret namespace | "default" |
| `spec.openwebui.langfuse.debug` | bool | Enable debug logging | false |

**Langfuse Options:**
- **Langfuse Cloud**: Use managed service at `https://cloud.langfuse.com`
- **Self-Hosted**: Deploy using [official Helm chart](https://github.com/langfuse/langfuse-k8s)

**Features:**
- 📊 **Request Tracking**: Monitor all LLM API calls and responses
- 🔍 **Performance Analytics**: Track latency, token usage, and costs
- 🚨 **Error Monitoring**: Identify and debug failed requests
- 📈 **Usage Metrics**: Understand application usage patterns
- 🔗 **Pipeline Integration**: Automatic OpenWebUI pipeline configuration

## Monitoring

Check the status of your deployment:

```bash
# Get the deployment status
kubectl get lmdeployment ollama-example -o yaml

# Check the status fields:
- `status.phase`: Overall deployment phase (Pending, Progressing, Ready)
- `status.ollamaStatus`: Ollama deployment status
- `status.openwebuiStatus`: OpenWebUI deployment status
- `status.tabbyStatus`: Tabby deployment status
```

## Development

### Prerequisites

- Go 1.24+
- kubebuilder
- controller-gen

### Local Development

```bash
# Install dependencies
go mod tidy

# Run tests
make test

# Run locally
make run

# Generate code
make generate
```

### Building

```bash
# Build binary
make build

# Build Docker image
make docker-build

# Push Docker image
make docker-push
```

## Troubleshooting

Check operator logs and deployment status:

```bash
# Check operator logs
kubectl logs -n llm-operator-system deployment/llm-operator-controller-manager

# Describe the deployment
kubectl describe lmdeployment ollama-example

# Get all resources with the deployment label
kubectl get all -l ollama-deployment=ollama-example
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the Apache License 2.0.

## Documentation

- [Quick Start Guide](docs/QUICKSTART.md) - Get started with the operator
- [CRD Reference](docs/CRD_REFERENCE.md) - Complete API documentation
- [Controller Architecture](docs/CONTROLLER_ARCHITECTURE.md) - How the operator works
- [Pipelines Guide](docs/PIPELINES.md) - OpenWebUI Pipelines integration
- [Langfuse Integration](docs/LANGFUSE_INTEGRATION.md) - Monitoring and observability
- [Security Guide](docs/SECURITY_IMPROVEMENTS.md) - Security best practices