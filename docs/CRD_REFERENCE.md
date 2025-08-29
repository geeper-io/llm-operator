---
id: crd-reference
title: CRD Reference
sidebar_label: CRD Reference
description: Complete reference for Geeper.AI Custom Resource Definitions
---

# CRD Reference

## Overview

The LLM Operator provides a Custom Resource Definition (CRD) called `LMDeployment` that allows you to declaratively deploy Ollama instances with specified models and optionally connect them to OpenWebUI for a web-based interface, Tabby for code completion, and custom components. The operator automatically manages the underlying Kubernetes resources including Deployments, Services, and Ingresses.

## LMDeployment

The `LMDeployment` resource defines the complete configuration for deploying LLM services.

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: example-deployment
spec:
  # ... configuration details
```

## API Version

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
```

## Schema

### Top-level Fields

| Field | Type | Description | Required |
|-------|------|-------------|----------|
| `apiVersion` | string | `llm.geeper.io/v1alpha1` | Yes |
| `kind` | string | `LMDeployment` | Yes |
| `metadata` | [ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.27/#objectmeta-v1-meta) | Standard object metadata | Yes |
| `spec` | [LMDeploymentSpec](#lmdeploymentspec) | Yes | Desired state of the LMDeployment |
| `status` | [LMDeploymentStatus](#lmdeploymentstatus) | No | Observed state of the LMDeployment (read-only) |

### LMDeploymentSpec

The `LMDeploymentSpec` defines the desired state of the LMDeployment.

| Field | Type | Description | Required |
|-------|------|-------------|----------|
| `ollama` | [OllamaSpec](#ollamaspec) | Yes | Ollama LMDeployment configuration |
| `openwebui` | [OpenWebUISpec](#openwebuispec) | No | OpenWebUI LMDeployment configuration |
| `tabby` | [TabbySpec](#tabbyspec) | No | Tabby LMDeployment configuration |

### OllamaSpec

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `replicas` | int32 | No | 1 | Number of Ollama pods to run (1-10) |
| `image` | string | No | `ollama/ollama` | Ollama container image |
| `imageTag` | string | No | `latest` | Ollama image tag |
| `resources` | [ResourceRequirements](#resourcerequirements) | No | None | Resource limits and requests |
| `models` | [OllamaModel](#ollamamodel)[] | Yes | - | List of models to deploy |
| `service` | [ServiceSpec](#servicespec) | No | Default service config | Service configuration |

### OllamaModel

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `OllamaModel` | string | Yes | Model specification in "modelname:tag" format (e.g., "llama2:7b", "mistral:7b") |

### ServiceSpec

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `type` | string | No | `ClusterIP` | Service type (ClusterIP, NodePort, LoadBalancer) |
| `port` | int32 | No | Component-specific | Service port (1-65535) |

### IngressSpec

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `enabled` | bool | No | false | Enable ingress for the component |
| `host` | string | No | None | Hostname for the ingress |
| `annotations` | map[string]string | No | None | Custom annotations for the ingress |
| `tls` | [IngressTLS](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.33/#ingresstls-v1-networking) | No | None | TLS configuration for the ingress |

### OpenWebUISpec

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `enabled` | bool | No | false | Enable OpenWebUI LMDeployment |
| `replicas` | int32 | No | 1 | Number of OpenWebUI pods (1-5) |
| `image` | string | No | `ghcr.io/open-webui/open-webui` | OpenWebUI container image |
| `imageTag` | string | No | `main` | OpenWebUI image tag |
| `resources` | [ResourceRequirements](#resourcerequirements) | No | None | Resource limits and requests |
| `service` | [ServiceSpec](#servicespec) | No | Default service config | Service configuration |
| `ingress` | [IngressSpec](#ingressspec) | No | Default ingress config | Ingress configuration |

### TabbySpec

| Field | Required | Type | Default | Description |
|-------|----------|------|---------|-------------|
| `enabled` | No | bool | false | Enable Tabby LMDeployment |
| `replicas` | No | int32 | 1 | Number of Tabby pods (1-5) |
| `image` | No | string | `tabbyml/tabby` | Tabby container image |
| `imageTag` | No | string | `latest` | Tabby image tag |
| `chatModel` | **Yes** | string | - | Ollama model for chat functionality (must be in spec.ollama.models) |
| `completionModel` | **Yes** | string | - | Ollama model for code completion (must be in spec.ollama.models) |
| `resources` | No | ResourceRequirements | - | Resource limits and requests |
| `service` | No | ServiceSpec | - | Service configuration |
| `ingress` | No | IngressSpec | - | Ingress configuration |
| `envVars` | No | []EnvVar | - | Environment variables |
| `volumeMounts` | No | []VolumeMount | - | Volume mounts |
| `volumes` | No | []Volume | - | Volumes |
| `configMapName` | No | string | - | Custom ConfigMap for configuration |

### ResourceRequirements

| Field | Type | Description |
|-------|------|-------------|
| `limits` | [ResourceList](#resourcelist) | Maximum amount of compute resources allowed |
| `requests` | [ResourceList](#resourcelist) | Minimum amount of compute resources required |

### ResourceList

| Field | Type | Description |
|-------|------|-------------|
| `cpu` | string | CPU resource (e.g., "100m", "2") |
| `memory` | string | Memory resource (e.g., "128Mi", "2Gi") |
| `storage` | string | Storage resource (e.g., "1Gi", "100Gi") |

### LMDeploymentStatus

| Field | Type | Description |
|-------|------|-------------|
| `phase` | string | Overall LMDeployment phase (Pending, Progressing, Ready) |
| `conditions` | [metav1.Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.33/#condition-v1-meta)[] | Latest observations of LMDeployment state |
| `ollamaStatus` | [LMDeploymentComponentStatus](#lmdeploymentcomponentstatus) | Ollama LMDeployment status |
| `openwebuiStatus` | [LMDeploymentComponentStatus](#lmdeploymentcomponentstatus) | OpenWebUI LMDeployment status |
| `tabbyStatus` | [LMDeploymentComponentStatus](#lmdeploymentcomponentstatus) | Tabby LMDeployment status |
| `readyReplicas` | int32 | Number of ready replicas |
| `totalReplicas` | int32 | Total number of replicas |

### LMDeploymentComponentStatus

The `LMDeploymentComponentStatus` represents the status of a component within the LMDeployment.

| Field | Type | Description |
|-------|------|-------------|
| `availableReplicas` | int32 | Number of available replicas |
| `readyReplicas` | int32 | Number of ready replicas |
| `updatedReplicas` | int32 | Number of updated replicas |
| `conditions` | [metav1.Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.33/#condition-v1-meta)[] | Component state conditions |

## Examples

### Minimal Configuration

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: minimal-ollama
spec:
  ollama:
    models:
      - "llama2:7b"
```

### Full Configuration

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: full-ollama
  namespace: ai-models
spec:
  ollama:
    replicas: 2
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
      - "mistral:7b"
  
  openwebui:
    enabled: true
    replicas: 1
    image: ghcr.io/open-webui/open-webui
    imageTag: main
    serviceType: ClusterIP
    servicePort: 8080
    ingressEnabled: true
    ingressHost: "ollama.example.com"
    resources:
      requests:
        cpu: "200m"
        memory: "512Mi"
      limits:
        cpu: "1"
        memory: "1Gi"
```

## Resource Management

### Model Pulling

The operator automatically creates postStart hooks to pull specified models after the Ollama container starts. This ensures models are downloaded and available for use.

### Resource Allocation

- **CPU**: Specified in cores (e.g., "1" = 1 core, "500m" = 0.5 cores)
- **Memory**: Specified in bytes with SI suffixes (e.g., "1Gi", "512Mi")
- **Storage**: Specified in bytes with SI suffixes (e.g., "10Gi", "100Mi")

### Resource Configuration

Resources can be specified for CPU and memory. If no resources are specified, Kubernetes will use its default resource allocation.

## Examples

### Basic Tabby LMDeployment

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: tabby-example
spec:
  tabby:
    enabled: true
    replicas: 1
    image: tabbyml/tabby:latest
    resources:
      requests:
        cpu: "500m"
        memory: "1Gi"
      limits:
        cpu: "2"
        memory: "2Gi"
```

### Advanced Tabby Configuration

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: production-tabby
  namespace: ai-models
spec:
  ollama:
    models:
      - "codellama:13b"
      - "llama2:13b"
    replicas: 2
    resources:
      requests:
        cpu: "1000m"
        memory: "4Gi"
      limits:
        cpu: "4000m"
        memory: "8Gi"
  
  tabby:
    enabled: true
    replicas: 2
    image: "tabbyml/tabby"
    chatModel: "codellama:7b"      # Must be in spec.ollama.models list
    completionModel: "codellama:7b" # Must be in spec.ollama.models list
    resources:
      limits:
        cpu: "2"
        memory: "4Gi"
      requests:
        cpu: "500m"
        memory: "1Gi"
    service:
      type: LoadBalancer
      port: 8080
    ingress:
      host: "tabby.example.com"
      annotations:
        nginx.ingress.kubernetes.io/ssl-redirect: "true"
```

## Service Configuration

### Service Types

- **ClusterIP**: Internal cluster access only
- **NodePort**: Accessible from outside the cluster via node IP
- **LoadBalancer**: Cloud provider load balancer (if available)

### Port Configuration

- **Ollama**: Default port 11434 (configurable)
- **OpenWebUI**: Default port 8080 (configurable)
- **Tabby**: Default port 8080 (configurable)

## Ingress Configuration

When `ingressEnabled: true` and `ingressHost` is specified, the operator creates Ingress resources for external access.

### OpenWebUI Ingress

- Path: `/` (root)
- PathType: `Prefix`
- Backend: OpenWebUI service

### Tabby Ingress

- Path: `/` (root)
- PathType: `Prefix`
- Backend: Tabby service

## Monitoring

Monitor LMDeployment progress using:

```bash
# Check the status
kubectl get lmdeployment <name> -o yaml
kubectl describe lmdeployment <name>

# Watch the progress
kubectl get lmdeployment <name> -w
```

## Best Practices

### Resource Planning

1. **Model Size**: Consider model size when setting memory limits
2. **CPU Allocation**: Allocate sufficient CPU for model inference
3. **Storage**: Use persistent volumes for production LMDeployments

### High Availability

1. **Replicas**: Use multiple replicas for production workloads
2. **Resource Limits**: Set appropriate limits to prevent resource exhaustion
3. **Service Types**: Use LoadBalancer for external access

### Security

1. **Image Tags**: Use specific image tags instead of `latest`
2. **Resource Limits**: Always set resource limits
3. **Network Policies**: Consider implementing network policies

## Troubleshooting

### Common Issues

1. **Models not pulling**: Check Ollama container logs for postStart hook execution
2. **OpenWebUI not connecting**: Verify Ollama service is accessible
3. **Tabby not connecting**: Verify Ollama service is accessible and model is available
4. **Resource constraints**: Ensure sufficient CPU/memory for model loading

### Debug Commands

```bash
# Check operator logs
kubectl logs -n llm-operator-system deployment/llm-operator-controller-manager

# Check CRD status
kubectl describe lmdeployment <name>

# Check created resources
kubectl get all -l ollama-deployment=<name>

# Check model pulling in Ollama container logs
kubectl logs <pod-name> -c ollama
```

## Migration from Previous Versions

If you're upgrading from a previous version:

1. Backup your existing configurations
2. Update the CRD schema
3. Verify resource compatibility
4. Test in a non-production environment first

## Support

For issues and questions:
1. Check the troubleshooting section
2. Review operator logs
3. Open an issue in the project repository
