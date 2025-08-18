---
id: tools
title: Tool System
sidebar_label: Tools
description: Learn how to extend Geeper.AI functionality with custom tools
---

# Tool System

Geeper.AI's tool system allows you to extend the functionality of your LLM deployments with custom features, integrations, and workflows. Tools can add new capabilities, connect to external services, or implement custom business logic.

## What are Tools?

Tools are modular components that:
- **Extend Functionality**: Add new features to your LLM deployments
- **Integrate Services**: Connect to databases, APIs, and external systems
- **Customize Behavior**: Implement domain-specific logic and workflows
- **Enhance Security**: Add authentication, authorization, and compliance features
- **Optimize Performance**: Implement caching, batching, and optimization strategies

### Tool Types

- **Input Tools**: Pre-process user input before sending to LLMs
- **Output Tools**: Post-process LLM responses
- **Integration Tools**: Connect to external services and APIs
- **Workflow Tools**: Orchestrate complex multi-step processes
- **Security Tools**: Handle authentication, encryption, and compliance

## Tool Architecture

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   User Input   │───▶│   Input Plugin   │───▶│      LLM       │
└─────────────────┘    └──────────────────┘    └─────────────────┘
                              │                        │
                              ▼                        ▼
                       ┌──────────────────┐    ┌──────────────────┐
                       │ Workflow Plugin  │    │  Output Plugin   │
                       └──────────────────┘    └──────────────────┘
                              │                        │
                              ▼                        ▼
                       ┌──────────────────┐    ┌─────────────────┐
                       │ Integration      │    │   User Output   │
                       │   Plugin        │    └─────────────────┘
                       └──────────────────┘
```

## Quick Start

### 1. Create a Plugin Directory

```bash
mkdir -p tools/my-custom-tool
cd tools/my-custom-tool
```

### 2. Define Plugin Configuration

```yaml
# plugin.yaml
name: "my-custom-plugin"
version: "1.0.0"
description: "A custom plugin for Geeper.AI"
author: "Your Name"
type: "output"
entrypoint: "main.py"
requirements:
  - "requests>=2.28.0"
  - "pandas>=1.5.0"
config:
  api_key: ""
  endpoint: "https://api.example.com"
  timeout: 30
```

### 3. Implement Plugin Logic

```python
# main.py
import requests
import json
from typing import Dict, Any

class MyCustomPlugin:
    def __init__(self, config: Dict[str, Any]):
        self.api_key = config.get("api_key")
        self.endpoint = config.get("endpoint")
        self.timeout = config.get("timeout", 30)
    
    def process(self, input_data: Dict[str, Any]) -> Dict[str, Any]:
        """Process the input data and return enhanced output"""
        try:
            # Your custom logic here
            response = requests.post(
                self.endpoint,
                headers={"Authorization": f"Bearer {self.api_key}"},
                json=input_data,
                timeout=self.timeout
            )
            response.raise_for_status()
            
            # Enhance the response
            enhanced_data = input_data.copy()
            enhanced_data["processed_by"] = "my-custom-plugin"
            enhanced_data["external_data"] = response.json()
            
            return enhanced_data
            
        except Exception as e:
            return {
                "error": str(e),
                "original_data": input_data
            }

# Plugin entry point
def create_plugin(config: Dict[str, Any]):
    return MyCustomPlugin(config)
```

### 4. Deploy Plugin

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: ollama-with-tools
  namespace: default
spec:
  ollama:
    replicas: 1
    image: ollama/ollama
    imageTag: latest
    models:
      - "llama2:7b"
      - "mistral:7b"
    resources:
      requests:
        cpu: "500m"
        memory: "2Gi"
      limits:
        cpu: "2"
        memory: "4Gi"
    serviceType: ClusterIP
    servicePort: 11434

  openwebui:
    enabled: true
    replicas: 1
    image: ghcr.io/open-webui/open-webui
    imageTag: latest
    resources:
      requests:
        cpu: "250m"
        memory: "512Mi"
      limits:
        cpu: "1"
        memory: "1Gi"
    serviceType: ClusterIP
    servicePort: 8080
    ingressEnabled: true
    ingressHost: "openwebui.example.com"
    
    # Plugin system configuration
    tools:
      # Weather API plugin
      - name: weather-api
        enabled: true
        type: openapi
        image: openapitools/openapi-generator-cli
        imageTag: latest
        replicas: 1
        port: 8081
        resources:
          requests:
            cpu: "100m"
            memory: "128Mi"
          limits:
            cpu: "500m"
            memory: "256Mi"
        serviceType: ClusterIP
        configMapName: "weather-api-config"
        secretName: "weather-api-secret"
        envVars:
          - name: "API_BASE_URL"
            value: "https://api.openweathermap.org/data/2.0"

      # Document processing plugin
      - name: document-processor
        enabled: true
        type: custom
        image: custom/document-processor
        imageTag: v1.0.0
        replicas: 2
        port: 8082
        resources:
          requests:
            cpu: "200m"
            memory: "256Mi"
          limits:
            cpu: "1"
            memory: "512Mi"
        serviceType: ClusterIP
        configMapName: "document-processor-config"
        envVars:
          - name: "PROCESSING_MODE"
            value: "async"
          - name: "MAX_FILE_SIZE"
            value: "100MB"
        volumeMounts:
          - name: shared-storage
            mountPath: /data
        volumes:
          - name: shared-storage
            persistentVolumeClaim:
              claimName: document-storage-pvc
```

## Plugin Categories

### Input Processing Tools

#### Text Preprocessing
```python
class TextPreprocessorPlugin:
    def process(self, text: str) -> str:
        # Clean and normalize text
        cleaned = text.strip().lower()
        # Remove special characters
        cleaned = re.sub(r'[^\w\s]', '', cleaned)
        return cleaned
```

#### Content Filtering
```python
class ContentFilterPlugin:
    def __init__(self, blocked_words: List[str]):
        self.blocked_words = set(blocked_words)
    
    def process(self, text: str) -> Dict[str, Any]:
        found_words = [word for word in self.blocked_words if word in text.lower()]
        if found_words:
            return {
                "blocked": True,
                "reason": f"Contains blocked words: {found_words}",
                "original_text": text
            }
        return {"blocked": False, "text": text}
```

### Output Enhancement Tools

#### Response Formatting
```python
class ResponseFormatterPlugin:
    def process(self, response: str) -> str:
        # Format response as markdown
        formatted = f"# Response\n\n{response}\n\n---\n*Generated by Geeper.AI*"
        return formatted
```

#### Sentiment Analysis
```python
class SentimentAnalysisPlugin:
    def process(self, text: str) -> Dict[str, Any]:
        # Analyze sentiment using a simple approach
        positive_words = ["good", "great", "excellent", "amazing"]
        negative_words = ["bad", "terrible", "awful", "horrible"]
        
        text_lower = text.lower()
        positive_count = sum(1 for word in positive_words if word in text_lower)
        negative_count = sum(1 for word in negative_words if word in text_lower)
        
        if positive_count > negative_count:
            sentiment = "positive"
        elif negative_count > positive_count:
            sentiment = "negative"
        else:
            sentiment = "neutral"
        
        return {
            "sentiment": sentiment,
            "positive_score": positive_count,
            "negative_score": negative_count,
            "text": text
        }
```

### Integration Tools

#### Database Integration
```python
class DatabasePlugin:
    def __init__(self, connection_string: str):
        self.connection_string = connection_string
        self.engine = create_engine(connection_string)
    
    def query(self, sql: str) -> List[Dict[str, Any]]:
        with self.engine.connect() as conn:
            result = conn.execute(text(sql))
            return [dict(row) for row in result]
    
    def insert(self, table: str, data: Dict[str, Any]) -> bool:
        with self.engine.connect() as conn:
            conn.execute(text(f"INSERT INTO {table} VALUES (:data)"), data)
            conn.commit()
            return True
```

#### API Integration
```python
class WeatherAPIPlugin:
    def __init__(self, api_key: str):
        self.api_key = api_key
        self.base_url = "https://api.weatherapi.com/v1"
    
    def get_weather(self, city: str) -> Dict[str, Any]:
        url = f"{self.base_url}/current.json"
        params = {
            "key": self.api_key,
            "q": city,
            "aqi": "no"
        }
        
        response = requests.get(url, params=params)
        response.raise_for_status()
        
        data = response.json()
        return {
            "city": data["location"]["name"],
            "temperature": data["current"]["temp_c"],
            "condition": data["current"]["condition"]["text"],
            "humidity": data["current"]["humidity"]
        }
```

## Plugin Configuration

### Environment Variables

```yaml
spec:
  env:
    - name: PLUGIN_API_KEY
      valueFrom:
        secretKeyRef:
          name: plugin-secrets
          key: api-key
    - name: PLUGIN_ENDPOINT
      value: "https://api.example.com"
    - name: PLUGIN_TIMEOUT
      value: "30"
```

### ConfigMaps

```yaml
# Weather API Plugin Configuration
apiVersion: v1
kind: ConfigMap
metadata:
  name: weather-api-config
  namespace: default
data:
  config: |
    {
      "api_spec_url": "https://api.openweathermap.org/data/2.0/openapi.json",
      "base_url": "https://api.openweathermap.org/data/2.0",
      "default_units": "metric",
      "cache_ttl": 300
    }

---
# Document Processor Configuration
apiVersion: v1
kind: ConfigMap
metadata:
  name: document-processor-config
  namespace: default
data:
  config: |
    {
      "supported_formats": ["pdf", "docx", "txt", "md"],
      "max_file_size": "100MB",
      "processing_timeout": 300,
      "output_format": "json"
    }

---
# Stripe Configuration
apiVersion: v1
kind: ConfigMap
metadata:
  name: stripe-config
  namespace: default
data:
  config: |
    {
      "api_version": "2023-10-16",
      "webhook_endpoint": "/webhooks/stripe",
      "currency": "usd",
      "timeout": 30
    }
```

### Secrets

```yaml
# Weather API Secret
apiVersion: v1
kind: Secret
metadata:
  name: weather-api-secret
  namespace: default
type: Opaque
data:
  api-key: <base64-encoded-api-key>
  # echo -n "your-actual-api-key" | base64

---
# Stripe Secret
apiVersion: v1
kind: Secret
metadata:
  name: stripe-secret
  namespace: default
type: Opaque
data:
  api-key: <base64-encoded-stripe-secret-key>
  webhook-secret: <base64-encoded-webhook-secret>

---
# GitHub Token Secret
apiVersion: v1
kind: Secret
metadata:
  name: github-token-secret
  namespace: default
type: Opaque
data:
  token: <base64-encoded-github-token>
  # echo -n "ghp_your_github_token" | base64
```

## Plugin Lifecycle

### Installation

```bash
# Install plugin from local directory
kubectl apply -f plugin-deployment.yaml

# Install plugin from registry
kubectl apply -f https://raw.githubusercontent.com/your-org/geeper-tools/main/weather-tool.yaml
```

### Updates

```bash
# Update plugin configuration
kubectl patch plugindeployment my-plugin --type='merge' -p='{"spec":{"config":{"timeout":"60"}}}'

# Update plugin version
kubectl set image plugindeployment/my-plugin my-plugin=your-org/plugin:v2.0.0
```

### Removal

```bash
# Remove plugin
kubectl delete plugindeployment my-plugin

# Remove plugin resources
kubectl delete -f plugin-deployment.yaml
```

## Plugin Development

### Testing

```python
# test_plugin.py
import pytest
from main import MyCustomPlugin

def test_plugin_initialization():
    config = {"api_key": "test", "endpoint": "https://test.com"}
    plugin = MyCustomPlugin(config)
    assert plugin.api_key == "test"
    assert plugin.endpoint == "https://test.com"

def test_plugin_processing():
    config = {"api_key": "test", "endpoint": "https://test.com"}
    plugin = MyCustomPlugin(config)
    
    input_data = {"text": "Hello, world!"}
    result = plugin.process(input_data)
    
    assert "processed_by" in result
    assert result["processed_by"] == "my-custom-plugin"
```

### Debugging

```python
import logging

logging.basicConfig(level=logging.DEBUG)
logger = logging.getLogger(__name__)

class MyCustomPlugin:
    def process(self, input_data: Dict[str, Any]) -> Dict[str, Any]:
        logger.debug(f"Processing input: {input_data}")
        
        try:
            # Your logic here
            result = self._process_data(input_data)
            logger.debug(f"Processing result: {result}")
            return result
            
        except Exception as e:
            logger.error(f"Error processing input: {e}")
            raise
```

## Best Practices

### 1. Error Handling
- Always handle exceptions gracefully
- Provide meaningful error messages
- Log errors for debugging
- Return fallback responses when possible

### 2. Configuration
- Use environment variables for sensitive data
- Provide sensible defaults
- Validate configuration on startup
- Document all configuration options

### 3. Performance
- Implement caching where appropriate
- Use async operations for I/O
- Batch operations when possible
- Monitor resource usage

### 4. Security
- Validate all inputs
- Sanitize outputs
- Use secure communication protocols
- Implement proper authentication

### 5. Testing
- Write unit tests for all functions
- Test error conditions
- Mock external dependencies
- Test with realistic data

## Plugin Registry

### Public Tools

Geeper.AI maintains a registry of community-contributed tools:

- **Weather Plugin**: Get real-time weather information
- **Translation Plugin**: Multi-language translation support
- **Math Plugin**: Mathematical computation and visualization
- **Calendar Plugin**: Calendar integration and scheduling
- **Email Plugin**: Email composition and management

### Contributing Tools

1. **Fork the Repository**: Create your own fork of the plugin registry
2. **Develop Your Plugin**: Implement your plugin following the guidelines
3. **Test Thoroughly**: Ensure your plugin works correctly
4. **Submit Pull Request**: Submit your plugin for review
5. **Documentation**: Provide clear documentation and examples

## Troubleshooting

### Common Issues

1. **Plugin Not Loading**:
   - Check plugin configuration
   - Verify entry point file exists
   - Check for syntax errors
   - Review plugin logs

2. **Configuration Errors**:
   - Validate configuration format
   - Check required fields
   - Verify secret/configmap references
   - Test configuration values

3. **Performance Issues**:
   - Monitor resource usage
   - Check for memory leaks
   - Optimize database queries
   - Implement caching

### Debug Commands

```bash
# Check plugin status
kubectl get plugindeployments

# View plugin logs
kubectl logs -f plugindeployment/my-plugin

# Check plugin configuration
kubectl describe plugindeployment my-plugin

# Test plugin endpoint
kubectl port-forward plugindeployment/my-plugin 8080:8080
curl http://localhost:8080/health
```

## Next Steps

- [Tool API Reference](/docs/api/tools) - Complete tool API documentation
- [Advanced Tool Patterns](/docs/chat/advanced-tools) - Advanced tool development
- [Tool Examples](/docs/chat/tool-examples) - Real-world tool examples
- [Tool Marketplace](/docs/chat/tool-marketplace) - Browse available tools

---

*Tools transform Geeper.AI from a simple LLM operator into a powerful, extensible AI platform*
