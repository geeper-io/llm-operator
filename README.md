# LLM Operator

A Kubernetes operator for deploying and managing Ollama with OpenWebUI.

## Overview

The LLM Operator provides a Custom Resource Definition (CRD) called `OllamaDeployment` that allows you to declaratively deploy Ollama instances with specified models and optionally connect them to OpenWebUI for a web-based interface.

## Features

- **Ollama Deployment**: Deploy Ollama instances with custom configurations
- **Model Management**: Automatically pull and manage specified models
- **OpenWebUI Integration**: Optional web-based UI for interacting with Ollama
- **Tabby Integration**: Optional code completion server using Ollama models
- **Resource Management**: Configurable resource limits and requests
- **Service Configuration**: Flexible service types (ClusterIP, NodePort, LoadBalancer)
- **Ingress Support**: Optional ingress configuration for OpenWebUI and Tabby

## Architecture

The operator creates the following Kubernetes resources:

- **Deployments**: For Ollama, OpenWebUI (if enabled), and Tabby (if enabled)
- **Services**: To expose Ollama, OpenWebUI, and Tabby
- **Ingress**: For external access to OpenWebUI and Tabby (optional)
- **PostStart Hooks**: To pull specified models after Ollama starts
- **Automatic Integration**: Tabby automatically connects to Ollama for code completion

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

## Usage

### Basic Example

Create an `OllamaDeployment` resource:

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: OllamaDeployment
metadata:
  name: ollama-example
  namespace: default
spec:
  ollama:
    replicas: 1
    models:
      - "llama2:7b"
      - "mistral:7b"
  
  openwebui:
    enabled: true
    ingressEnabled: true
    ingressHost: "ollama-webui.local"
```

### Advanced Configuration

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: OllamaDeployment
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

- Connects to the deployed Ollama service in the same deployment
- Uses the first specified model (or a custom one you specify)
- Provides a REST API for code completion
- Supports multiple programming languages
- Can be accessed via ingress for external integration
    image: ghcr.io/open-webui/open-webui
    imageTag: main
    serviceType: ClusterIP
    servicePort: 8080
    ingressEnabled: true
    ingressHost: "ai.example.com"
    resources:
      requests:
        cpu: "200m"
        memory: "512Mi"
      limits:
        cpu: "1"
        memory: "1Gi"
```

## API Reference

### OllamaDeployment

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
| `spec.tabby.service.type` | string | Service type (ClusterIP, NodePort, LoadBalancer) | `ClusterIP` |
| `spec.tabby.service.port` | int32 | Service port | `8080` |
| `spec.tabby.ingress.enabled` | bool | Enable ingress | false |
| `spec.tabby.ingress.host` | string | Ingress hostname | None |
| `spec.tabby.modelName` | string | Ollama model to use | Auto-detected |

## Monitoring

The operator provides status information through the CRD status field:

```bash
kubectl get ollamadeployment ollama-example -o yaml
```

Status fields include:
- `status.phase`: Overall deployment phase (Pending, Progressing, Ready)
- `status.ollamaStatus`: Ollama deployment status
- `status.openwebuiStatus`: OpenWebUI deployment status
- `status.tabbyStatus`: Tabby deployment status
- `status.readyReplicas`: Number of ready replicas
- `status.totalReplicas`: Total number of replicas

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

### Common Issues

1. **Models not pulling**: Check Ollama container logs for postStart hook execution
2. **OpenWebUI not connecting**: Verify Ollama service is accessible
3. **Tabby not connecting**: Verify Ollama service is accessible and model is available
3. **Resource constraints**: Ensure sufficient CPU/memory for model loading

### Debug Commands

```bash
# Check operator logs
kubectl logs -n llm-operator-system deployment/llm-operator-controller-manager

# Check CRD status
kubectl describe ollamadeployment ollama-example

# Check created resources
kubectl get all -l ollama-deployment=ollama-example

# Check model pulling in Ollama container logs
kubectl logs <pod-name> -c ollama
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests
5. Submit a pull request

## License

This project is licensed under the Apache License 2.0.

