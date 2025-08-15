# OllamaDeployment CRD Reference

This document provides a comprehensive reference for the `OllamaDeployment` Custom Resource Definition (CRD) used by the LLM Operator.

## Overview

The `OllamaDeployment` CRD allows you to declaratively deploy Ollama instances with specified models and optionally connect them to OpenWebUI for a web-based interface. The operator automatically manages the underlying Kubernetes resources including Deployments, Services, and Ingresses.

## API Version

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: OllamaDeployment
```

## Schema

### Top-level Fields

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `metadata` | [ObjectMeta](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.33/#objectmeta-v1-meta) | Yes | Standard Kubernetes metadata |
| `spec` | [OllamaDeploymentSpec](#ollamadeploymentspec) | Yes | Desired state of the deployment |
| `status` | [OllamaDeploymentStatus](#ollamadeploymentstatus) | No | Observed state of the deployment (read-only) |

### OllamaDeploymentSpec

| Field | Type | Required | Description |
|-------|------|----------|-------------|
| `ollama` | [OllamaSpec](#ollamaspec) | Yes | Ollama deployment configuration |
| `openwebui` | [OpenWebUISpec](#openwebuispec) | No | OpenWebUI deployment configuration |
| `tabby` | [TabbySpec](#tabbyspec) | No | Tabby deployment configuration |

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
| `enabled` | bool | No | false | Enable OpenWebUI deployment |
| `replicas` | int32 | No | 1 | Number of OpenWebUI pods (1-5) |
| `image` | string | No | `ghcr.io/open-webui/open-webui` | OpenWebUI container image |
| `imageTag` | string | No | `main` | OpenWebUI image tag |
| `resources` | [ResourceRequirements](#resourcerequirements) | No | None | Resource limits and requests |
| `service` | [ServiceSpec](#servicespec) | No | Default service config | Service configuration |
| `ingress` | [IngressSpec](#ingressspec) | No | Default ingress config | Ingress configuration |

### TabbySpec

| Field | Type | Required | Default | Description |
|-------|------|----------|---------|-------------|
| `enabled` | bool | No | false | Enable Tabby deployment |
| `replicas` | int32 | No | 1 | Number of Tabby pods (1-5) |
| `image` | string | No | `tabbyml/tabby` | Tabby container image |
| `imageTag` | string | No | `latest` | Tabby image tag |
| `resources` | [ResourceRequirements](#resourcerequirements) | No | None | Resource limits and requests |
| `service` | [ServiceSpec](#servicespec) | No | Default service config | Service configuration |
| `ingress` | [IngressSpec](#ingressspec) | No | Default ingress config | Ingress configuration |
| `ollamaServiceName` | string | No | Auto-generated | Ollama service name to connect to |
| `ollamaServicePort` | int32 | No | Auto-detected | Ollama service port to connect to |
| `modelName` | string | No | Auto-detected | Ollama model to use for code completion |
| `envVars` | [corev1.EnvVar](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.33/#envvar-v1-core)[] | No | None | Custom environment variables |
| `volumeMounts` | [corev1.VolumeMount](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.33/#volumemount-v1-core)[] | No | None | Custom volume mounts |
| `volumes` | [corev1.Volume](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.33/#volume-v1-core)[] | No | None | Custom volumes |

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

### OllamaDeploymentStatus

| Field | Type | Description |
|-------|------|-------------|
| `phase` | string | Overall deployment phase (Pending, Progressing, Ready) |
| `conditions` | [metav1.Condition](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.33/#condition-v1-meta)[] | Latest observations of deployment state |
| `ollamaStatus` | [DeploymentComponentStatus](#deploymentcomponentstatus) | Ollama deployment status |
| `openwebuiStatus` | [DeploymentComponentStatus](#deploymentcomponentstatus) | OpenWebUI deployment status |
| `tabbyStatus` | [DeploymentComponentStatus](#deploymentcomponentstatus) | Tabby deployment status |
| `readyReplicas` | int32 | Number of ready replicas |
| `totalReplicas` | int32 | Total number of replicas |

### DeploymentComponentStatus

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
kind: OllamaDeployment
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
kind: OllamaDeployment
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

### Default Resource Values

If no resources are specified, the operator uses reasonable defaults for the main Ollama container. Model pulling happens in the postStart hook using the same resources as the main container.

## Examples

### Basic Tabby Deployment

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: Deployment
metadata:
  name: tabby-example
  namespace: default
spec:
  ollama:
    models:
      - "codellama:7b"
    replicas: 1
  
  tabby:
    enabled: true
    replicas: 1
    ingress:
      enabled: true
      host: "tabby.localhost"
    resources:
      requests:
        cpu: "250m"
        memory: "512Mi"
      limits:
        cpu: "1000m"
        memory: "1Gi"
```

### Advanced Tabby Configuration

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: Deployment
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
    imageTag: "latest"
    service:
      type: LoadBalancer
      port: 8080
    ingress:
      enabled: true
      host: "tabby.example.com"
    modelName: "codellama:13b"  # Use specific model
    ollamaServiceName: "custom-ollama"  # Custom Ollama service
    ollamaServicePort: 11434
    resources:
      requests:
        cpu: "500m"
        memory: "1Gi"
      limits:
        cpu: "2000m"
        memory: "2Gi"
    envVars:
      - name: "TABBY_LOG_LEVEL"
        value: "debug"
      - name: "TABBY_HOST"
        value: "0.0.0.0"
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

## Status Monitoring

### Phase Values

- **Pending**: No replicas are ready
- **Progressing**: Some replicas are ready but not all
- **Ready**: All replicas are ready

### Status Fields

Monitor deployment progress using:

```bash
kubectl get ollamadeployment <name> -o yaml
kubectl describe ollamadeployment <name>
```

## Best Practices

### Resource Planning

1. **Model Size**: Consider model size when setting memory limits
2. **CPU Allocation**: Allocate sufficient CPU for model inference
3. **Storage**: Use persistent volumes for production deployments

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

1. **Models Not Pulling**: Check init container logs
2. **OpenWebUI Connection**: Verify Ollama service accessibility
3. **Tabby Connection**: Verify Ollama service accessibility and model availability
4. **Resource Constraints**: Ensure sufficient cluster resources

### Debug Commands

```bash
# Check operator logs
kubectl logs -n llm-operator-system deployment/llm-operator-controller-manager

# Check CRD status
kubectl describe ollamadeployment <name>

# Check created resources
kubectl get all -l ollama-deployment=<name>

# Check postStart hook execution
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
