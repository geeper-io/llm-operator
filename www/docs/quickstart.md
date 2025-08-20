---
id: quickstart
title: Quick Start Guide
sidebar_label: Quick Start
sidebar_position: 2
description: Get up and running with LLM Operator in 5 minutes
---

# Quick Start Guide

Get started with LLM Operator in just 5 minutes! This guide will walk you through deploying your first LLM service on Kubernetes.

## Prerequisites

- Kubernetes cluster (1.24+)
- `kubectl` configured to communicate with your cluster
- `helm` (optional, for easy installation)

## 1. Install the Operator

### Option A: Using Helm (Recommended)

```bash
# Add the Helm repository
helm repo add llm-operator https://geeper-ai.github.io/llm-operator

# Install the operator
helm install llm-operator llm-operator/llm-operator \
  --namespace llm-operator \
  --create-namespace
```

### Option B: Using kubectl

```bash
# Apply the operator manifests
kubectl apply -f https://raw.githubusercontent.com/geeper-io/llm-operator/main/install.yaml
```

## 2. Deploy Your First LLM

Create a simple Ollama deployment:

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: my-first-llm
  namespace: default
spec:
  ollama:
    models:
      - "llama2:7b"
    service:
      type: ClusterIP
      port: 11434
    ingress:
      enabled: true
      host: "llm.localhost"
  
  openwebui:
    enabled: true
    replicas: 1
    image: ghcr.io/open-webui/open-webui:main
    service:
      type: ClusterIP
      port: 8080
    ingress:
      enabled: true
      host: "chat.localhost"
```

Save this as `my-first-llm.yaml` and apply it:

```bash
kubectl apply -f my-first-llm.yaml
```

## 3. Verify the Deployment

Check the status of your deployment:

```bash
kubectl get lmdeployments
kubectl get pods -l app=my-first-llm
kubectl get services -l app=my-first-llm
```

## 4. Access Your LLM

Once all pods are running, you can access:

- **Ollama API**: `http://llm.localhost:11434`
- **OpenWebUI Chat**: `http://chat.localhost:8080`

## 5. Test the Setup

Pull a model and start chatting:

```bash
# Pull the model
curl -X POST http://llm.localhost:11434/api/pull \
  -H "Content-Type: application/json" \
  -d '{"name": "llama2:7b"}'

# Test a simple completion
curl -X POST http://llm.localhost:11434/api/generate \
  -H "Content-Type: application/json" \
  -d '{
    "model": "llama2:7b",
    "prompt": "Hello, how are you?",
    "stream": false
  }'
```

## What's Next?

- [Advanced Configuration](/docs/advanced/) - Learn about pipelines, observability, and more
- [Examples](https://github.com/geeper-io/llm-operator/tree/main/examples) - Explore production-ready configurations
- [API Reference](/docs/crd-reference) - Complete CRD documentation

## Troubleshooting

If you encounter issues:

1. Check pod logs: `kubectl logs -l app=my-first-llm`
2. Verify operator status: `kubectl get pods -n llm-operator`
3. Check events: `kubectl describe lmdeployment my-first-llm`

For more help, visit our [GitHub Discussions](https://github.com/geeper-io/llm-operator/discussions) or [Issues](https://github.com/geeper-io/llm-operator/issues).


