---
id: overview
title: Overview
sidebar_label: Overview
sidebar_position: 1
description: Learn about Geeper.AI - the LLM Operator for Kubernetes
---

# Overview

Geeper.AI is a powerful Kubernetes operator that simplifies the LMDeployment and management of Large Language Models (LLMs) in Kubernetes clusters. It provides a declarative way to deploy, configure, and manage various LLM services with enterprise-grade reliability and scalability.

## What is Geeper.AI?

Geeper.AI is an open-source Kubernetes operator that extends Kubernetes with custom resources for managing LLM LMDeployments. It automates the complex process of deploying and managing LLM services, making it easy for developers and DevOps teams to run AI workloads in production environments.

## Key Features

- **🚀 Easy LMDeployment**: Deploy LLM services with simple YAML configurations
- **🔧 Multi-Model Support**: Support for various LLM frameworks (Ollama, OpenWebUI, Tabby, etc.)
- **⚡ Auto-scaling**: Automatic scaling based on demand and resource usage
- **🔄 Lifecycle Management**: Automated updates, rollbacks, and health monitoring
- **🔒 Security**: Built-in security features and RBAC integration
- **📊 Monitoring**: Comprehensive metrics and observability
- **🌐 Multi-cluster**: Support for multi-cluster LMDeployments
- **🔧 Tool System**: Extensible architecture with tool support
- **⚡ Pipelines**: OpenWebUI Pipelines for custom workflows and integrations
- **🔄 Redis Integration**: Automatic Redis deployment for multi-replica OpenWebUI

## What You Can Do

- **Deploy LLMs**: Run Ollama instances with your preferred models
- **Web Interface**: Add OpenWebUI for a chat-based interface
- **Code Completion**: Use Tabby for intelligent code suggestions
- **Plugin System**: Extend functionality with custom plugins
- **RAG Integration**: Connect to external knowledge sources
- **Pipelines**: Advanced workflow and integration capabilities
- **Monitoring**: Comprehensive observability with Langfuse
- **Multi-Instance**: Scale across multiple replicas with Redis

## Quick Installation

### Prerequisites
- Kubernetes cluster (v1.20+)
- kubectl configured
- Helm (optional, for advanced LMDeployments)

### Install with kubectl
```bash
# Clone the repository
git clone https://github.com/your-org/llm-operator.git
cd llm-operator

# Apply the CRDs
kubectl apply -f config/crd/bases/

# Deploy the operator
kubectl apply -f config/default/

# Verify installation
kubectl get pods -n llm-operator-system
```

### Install with Helm
```bash
# Add the Helm repository
helm repo add geeper-ai https://your-org.github.io/llm-operator

# Install the operator
helm install llm-operator geeper-ai/llm-operator \
  --namespace llm-operator-system \
  --create-namespace
```

## Quick Start Example

Deploy your first LLM service:

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: llama2-lmdeployment
spec:
  ollama:
    models:
      - "llama2:7b"
    replicas: 2
    resources:
      requests:
        memory: "4Gi"
        cpu: "2"
      limits:
        memory: "8Gi"
        cpu: "4"
    gpu:
      enabled: true
      count: 1
```

Apply the configuration:
```bash
kubectl apply -f lmdeployment.yaml
```

## Architecture

Geeper.AI follows the Kubernetes operator pattern:

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   Custom        │    │   Geeper.AI      │    │   Kubernetes    │
│   Resources     │───▶│   Operator       │───▶│   Resources     │
│   (CRDs)       │    │   (Controller)   │    │   (Pods, SVCs)  │
└─────────────────┘    └──────────────────┘    └─────────────────┘
```

The operator watches for changes to custom resources and automatically:
1. Creates the necessary Kubernetes resources
2. Manages the lifecycle of LLM services
3. Handles scaling, updates, and health monitoring
4. Provides status and metrics information

## Next Steps

- [Chat & Interaction](/docs/chat/openwebui) - Chat with LLMs using OpenWebUI
- [RAG Integration](/docs/chat/rag) - Add Retrieval-Augmented Generation
- [Tool System](/docs/chat/tools) - Extend functionality with tools
- [Coding Assistants](/docs/coding-assistants/continue-dev) - AI-powered coding with Continue.dev and Tabby
- [Installation Guide](/docs/installation) - Detailed installation instructions
- [User Guide](/docs/usage) - Learn how to use Geeper.AI
- [CRD Reference](/docs/crd-reference) - Complete API documentation
- [Examples](/docs/examples) - Real-world LMDeployment examples
- [Troubleshooting](/docs/troubleshooting) - Common issues and solutions

## Community and Support

- **GitHub**: [github.com/your-org/llm-operator](https://github.com/your-org/llm-operator)
- **Issues**: [GitHub Issues](https://github.com/your-org/llm-operator/issues)
- **Discussions**: [GitHub Discussions](https://github.com/your-org/llm-operator/discussions)
- **Documentation**: [docs.geeper.ai](https://docs.geeper.ai)

## Contributing

We welcome contributions! Please see our [Contributing Guide](/docs/contributing) for details on how to:
- Report bugs
- Suggest new features
- Submit pull requests
- Join our community

---

*Geeper.AI - Making LLM LMDeployment simple and reliable in Kubernetes*
