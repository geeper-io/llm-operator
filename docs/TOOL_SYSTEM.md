# OpenWebUI Tool System

The OpenWebUI tool system allows you to deploy additional services alongside OpenWebUI and automatically configure them as tools in OpenWebUI's configuration. This enables you to extend OpenWebUI's functionality with external APIs, custom services, and other containerized applications.

## Overview

The plugin system provides:
- **Automatic deployment** of tool services
- **Automatic configuration** of OpenWebUI via config.json
- **Service discovery** and configuration
- **Configuration management** via ConfigMaps
- **Credential management** via Secrets
- **Resource management** with proper ownership

## How It Works

### 1. Tool Deployment
- Each plugin is deployed as a **separate Kubernetes deployment**
- Tools run independently with their own resources and scaling
- Each tool gets its own service for internal communication

### 2. OpenWebUI Configuration
- The controller automatically generates a `config.json` file
- This file is mounted as a volume in OpenWebUI's `/app/backend/data` directory
- OpenWebUI reads this configuration and enables the tools automatically
- No environment variables needed - OpenWebUI uses its native config.json format

### 3. Configuration Structure
The generated `config.json` follows OpenWebUI's native format:

```json
{
  "version": 0,
  "ui": {
    "enable_signup": false
  },
  "openai": {
    "enable": false,
    "api_base_urls": ["https://api.openai.com/v1"],
    "api_keys": [""],
    "api_configs": {"0": {}}
  },
  "tool_server": {
    "connections": [
      {
        "url": "http://tool-service:port",
        "path": "openapi.json",
        "auth_type": "bearer",
        "key": "",
        "config": {
          "enable": true,
          "access_control": {
            "read": {"group_ids": [], "user_ids": []},
            "write": {"group_ids": [], "user_ids": []}
          }
        },
        "info": {
          "name": "tool-name",
          "description": "tool description"
        }
      }
    ]
  }
}
```

## Plugin Types

### 1. OpenAPI Tools
OpenAPI tools integrate external APIs by providing:
- OpenAPI specification endpoints
- Configuration via ConfigMaps
- Credentials via Secrets
- Service endpoints

### 2. Custom Tools
Custom tools for any service that:
- Exposes HTTP endpoints
- Can be containerized
- Integrates with OpenWebUI

## Configuration

### Basic Plugin Structure

```yaml
tools:
  - name: "tool-name"
    enabled: true
    type: "openapi"  # or "custom"
    image: "tool-image"
    replicas: 1
    port: 8080
    serviceType: "ClusterIP"
    configMapName: "tool-config"  # Optional
    secretName: "tool-secret"     # Optional
    resources:
      requests:
        cpu: "100m"
        memory: "128Mi"
      limits:
        cpu: "500m"
        memory: "256Mi"
```

### Tool Configuration via ConfigMaps

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: weather-tool-config
data:
  config: |
    {
      "api_spec_url": "https://api.openweathermap.org/data/2.0/openapi.json",
      "base_url": "https://api.openweathermap.org/data/2.0",
      "default_units": "metric",
      "cache_ttl": 300
    }
```

### Tool Credentials via Secrets

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: weather-tool-secret
type: Opaque
data:
  credentials: <base64-encoded-api-key>
  # echo -n "your-actual-api-key" | base64
```

## Example Use Cases

### 1. Weather API Tool
Deploy a weather API tool that provides weather data:

```yaml
tools:
  - name: "weather-tool"
    type: "openapi"
    image: "openapitools/openapi-generator-cli"
    port: 8081
    configMapName: "weather-tool-config"
    secretName: "weather-tool-secret"
```

### 2. Document Processing Tool
Deploy a custom document processing tool:

```yaml
tools:
  - name: "document-tool"
    type: "custom"
    image: "custom/document-processor"
    port: 8082
    replicas: 2
    configMapName: "document-tool-config"
    envVars:
      - name: "PROCESSING_MODE"
        value: "async"
    volumeMounts:
      - name: shared-storage
        mountPath: /data
    volumes:
      - name: shared-storage
        persistentVolumeClaim:
          claimName: document-storage-pvc
```

## What Gets Created

### 1. Tool Deployments
- Each tool becomes a separate Kubernetes deployment
- Tools run independently and can scale separately
- Each tool has its own service for internal communication

### 2. OpenWebUI Configuration
- **ConfigMap**: Contains the `config.json` with tool server connections
- **Volume Mount**: The config.json is mounted in OpenWebUI's data directory
- **Automatic Discovery**: OpenWebUI automatically discovers and enables the tools

### 3. Service Discovery
Tools are accessible via internal service names:
- `{deployment-name}-plugin-{tool-name}:{port}`
- Example: `ollama-with-tools-plugin-weather-tool:8081`

## Resource Management

### Automatic Cleanup
All resources are automatically managed with proper owner references:
- When you delete the main deployment, all tools are automatically deleted
- ConfigMaps, services, and other resources are properly cleaned up
- No orphaned resources left behind

### Resource Limits
Configure resource requirements for each tool:
```yaml
resources:
  requests:
    cpu: "100m"
    memory: "128Mi"
  limits:
    cpu: "500m"
    memory: "256Mi"
```

## Monitoring and Status

### Tool Status Tracking
Tool status is tracked and included in the overall deployment status:
- Available replicas
- Ready replicas
- Updated replicas
- Overall health status

### Status Annotations
Tool statuses are stored in annotations for easy access:
```yaml
annotations:
  llm.geeper.io/plugin-statuses: '{"weather-tool":{"availableReplicas":1,"readyReplicas":1}}'
```

## Best Practices

### 1. Configuration Management
- Store tool-specific configuration in ConfigMaps
- Store credentials in Secrets
- Use environment variables for runtime configuration
- Version control your configurations

### 2. Security
- Never store sensitive data in ConfigMaps
- Use Kubernetes Secrets for all credentials
- Limit tool access to necessary resources
- Rotate credentials regularly

### 3. Resource Planning
- Set appropriate resource limits for each tool
- Monitor resource usage and adjust as needed
- Use resource requests to ensure proper scheduling

### 4. Monitoring
- Monitor tool health and performance
- Set up alerts for tool failures
- Track API usage and rate limits

## Troubleshooting

### Common Issues

#### Tool Not Starting
- Check resource limits and requests
- Verify image and tag availability
- Check logs for startup errors
- Verify ConfigMap and Secret references

#### Configuration Issues
- Verify ConfigMap names and keys
- Check Secret names and keys
- Ensure proper permissions
- Validate JSON configuration format

#### Service Discovery Issues
- Verify service names and ports
- Check network policies
- Ensure services are in the same namespace

### Debugging Commands

```bash
# Check tool deployment status
kubectl get deployments -l app=openwebui-plugin

# Check tool service status
kubectl get services -l app=openwebui-plugin

# Check tool logs
kubectl logs -l app=openwebui-plugin

# Check OpenWebUI configuration
kubectl get configmap <deployment-name>-openwebui-config -o yaml

# Check OpenWebUI logs
kubectl logs -l app=openwebui

# Check OpenWebUI config volume mount
kubectl describe pod -l app=openwebui
```

## How OpenWebUI Uses the Configuration

1. **Startup**: OpenWebUI reads the mounted `config.json` file
2. **Tool Discovery**: Automatically discovers tool server connections
3. **Connection**: Establishes connections to the configured tools
4. **Integration**: Makes tools available in the UI for users
5. **Authentication**: Uses configured auth types and keys for secure access

## Future Enhancements

The plugin system is designed to be extensible. Future versions may include:
- Advanced tool configuration options
- Tool dependency management
- Automatic scaling based on usage
- Tool health checks and auto-recovery
- Integration with external monitoring systems
