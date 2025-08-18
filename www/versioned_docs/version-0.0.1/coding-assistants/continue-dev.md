---
id: continue-dev
title: Continue.dev Integration
sidebar_label: Continue.dev
description: Integrate Continue.dev VSCode extension with your Geeper.AI LMDeployment
---

# Continue.dev Integration

Continue.dev is a powerful VSCode extension that provides AI-powered code completion, chat, and editing capabilities. This guide shows you how to deploy a simple LMDeployment with Geeper.AI and integrate it with Continue.dev for a seamless development experience.

## What is Continue.dev?

Continue.dev is an open-source AI coding assistant that integrates directly into VSCode, providing:

- **AI Chat**: Interactive conversations about your code
- **Code Editing**: AI-powered code modifications and improvements
- **Context Awareness**: Understands your codebase, terminal, and development environment
- **Multi-Model Support**: Works with various LLM providers including OpenAI-compatible APIs

## Prerequisites

- VSCode installed
- Kubernetes cluster with Geeper.AI operator deployed
- kubectl configured
- Basic knowledge of Kubernetes and YAML

## Step 1: Deploy an LMDeployment

First, let's create a basic LMDeployment that will serve as the backend for Continue.dev:

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: continue-dev-example
  namespace: default
spec:
  ollama:
    models:
      - "codellama:7b"
  
  openwebui:
    enabled: true
    ingress:
      enabled: true
      host: "my-chat-ui.com"
```

Apply the configuration:

```bash
kubectl apply -f continue-dev-example.yaml
```

## Step 2: Create OpenWebUI API Key

1. Access your OpenWebUI instance at the ingress URL (e.g., `http://my-chat-ui.com`)
2. Go to **Settings** → **Account** → **API keys**
3. Click on **Create new secret key**
4. Copy the generated API key - you'll need it for the next step

## Step 3: Install Continue.dev Extension

1. Open VSCode
2. Go to Extensions (Ctrl+Shift+X)
3. Search for "Continue"
4. Install the "Continue" extension by Continue AI

## Step 4: Configure Continue.dev

1. In VSCode, click on the Continue tab in the sidebar
2. Click on the assistant selector next to the chat input
3. Hover over `Local Assistant` and click the settings icon (⚙️)
4. This opens the `config.yaml` file in your editor

Here's an example complete configuration for your Geeper.AI deployment:

```yaml
name: Local Assistant
version: 1.0.0
schema: v1
models:
  - name: CodeLlama
    provider: openai
    model: codellama:7b
    env:
      useLegacyCompletionsEndpoint: false
    apiBase: http://my-chat-ui.com/api
    apiKey: YOUR_OPEN_WEBUI_API_KEY
    roles:
      - chat
      - edit
context:
  - provider: code
  - provider: docs
  - provider: diff
  - provider: terminal
  - provider: problems
  - provider: folder
  - provider: codebase
```

## Step 5: Test the Integration

1. Open any code file in VSCode
2. Go to the Continue tab
3. Start a conversation with your AI assistant
4. Try asking it to:
   - Explain the current code
   - Suggest improvements
   - Help with debugging
   - Generate new functions
