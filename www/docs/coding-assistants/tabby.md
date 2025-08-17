---
id: tabby
title: Tabby Integration
sidebar_label: Tabby
description: Learn how to integrate Tabby with Geeper.AI for intelligent code completion
---

# Tabby Integration

Tabby is an open-source, self-hosted AI coding assistant that provides intelligent code completion and generation. It's designed to work seamlessly with your Geeper.AI deployments, offering privacy-focused AI coding assistance.

## What is Tabby?

Tabby is a self-hosted AI coding assistant that provides:

- **🔒 Privacy-First**: Your code never leaves your infrastructure
- **🚀 Fast Completion**: Real-time code suggestions as you type
- **🎯 Context-Aware**: Understands your codebase for better suggestions
- **🔧 Customizable**: Configurable models and parameters
- **📱 Multi-Platform**: Works with VS Code, Vim, and other editors
- **🌐 Self-Hosted**: Full control over your AI coding environment

## Key Features

- **💻 Intelligent Code Completion**: Context-aware suggestions
- **📝 Code Generation**: Generate functions, classes, and entire files
- **🔍 Code Understanding**: Analyze and explain existing code
- **🔄 Refactoring**: Suggest code improvements and optimizations
- **📚 Multi-Language Support**: Works with Python, JavaScript, Go, Rust, and more
- **⚡ Low Latency**: Optimized for real-time development workflows

## Installation

### VS Code Extension

1. **Open VS Code Extensions**:
   - Press `Ctrl+Shift+X` (Windows/Linux) or `Cmd+Shift+X` (Mac)
   - Or go to View → Extensions

2. **Search for Tabby**:
   - Type "Tabby" in the search box
   - Look for "Tabby - AI Code Completion"

3. **Install Extension**:
   - Click "Install" on the Tabby extension
   - Restart VS Code when prompted

4. **Verify Installation**:
   - Check the Extensions panel for Tabby
   - Look for Tabby icon in the status bar

### Vim/Neovim Plugin

#### Using vim-plug
```vim
" Add to your .vimrc or init.vim
Plug 'TabbyML/vim-tabby'

" Or for Neovim
Plug 'TabbyML/vim-tabby', { 'do': ':UpdateRemotePlugins' }
```

#### Using lazy.nvim
```lua
-- Add to your Neovim configuration
{
  'TabbyML/vim-tabby',
  config = function()
    require('tabby').setup()
  end
}
```

### Command Line Installation

```bash
# Install Tabby CLI globally
npm install -g @tabby/cli

# Verify installation
tabby --version

# Initialize Tabby in your project
cd your-project
tabby init
```

## Configuration

### Basic Setup

1. **Open Tabby Settings**:
   - VS Code: `Ctrl+,` (Windows/Linux) or `Cmd+,` (Mac)
   - Search for "Tabby" in settings

2. **Configure Server Endpoint**:
   - Set your Geeper.AI Tabby endpoint
   - Configure authentication if required
   - Set completion preferences

3. **Customize Behavior**:
   - Adjust completion triggers
   - Set language-specific settings
   - Configure file exclusions

### Geeper.AI Integration

```json
{
  "tabby": {
    "server": {
      "endpoint": "http://your-tabby-endpoint:8080",
      "token": "your-auth-token"
    },
    "completion": {
      "triggerMode": "automatic",
      "maxLines": 50,
      "temperature": 0.1
    },
    "languages": {
      "python": {
        "enabled": true,
        "maxTokens": 100
      },
      "javascript": {
        "enabled": true,
        "maxTokens": 80
      },
      "go": {
        "enabled": true,
        "maxTokens": 120
      }
    }
  }
}
```

### Vim Configuration

```vim
" Basic Tabby configuration
let g:tabby_server = 'http://localhost:8080'
let g:tabby_key = 'your-auth-token'

" Completion settings
let g:tabby_trigger_mode = 'automatic'
let g:tabby_max_lines = 50

" Language-specific settings
let g:tabby_languages = {
  \ 'python': {'enabled': 1, 'max_tokens': 100},
  \ 'javascript': {'enabled': 1, 'max_tokens': 80},
  \ 'go': {'enabled': 1, 'max_tokens': 120}
  \ }
```

## Usage

### VS Code

#### Basic Completion
1. **Start Typing**: Begin typing code in any supported file
2. **Accept Suggestions**: Press `Tab` to accept suggestions
3. **Cycle Options**: Use `Ctrl+Shift+]` to cycle through alternatives
4. **Trigger Manually**: Press `Ctrl+Shift+Space` for manual completion

#### Commands
- **`Tabby: Complete`**: Manually trigger completion
- **`Tabby: Explain Code`**: Get explanation of selected code
- **`Tabby: Generate Tests`**: Generate tests for selected code
- **`Tabby: Refactor`**: Suggest refactoring improvements

#### Keyboard Shortcuts
- **`Tab`**: Accept current suggestion
- **`Ctrl+Shift+]`**: Next suggestion
- **`Ctrl+Shift+[`**: Previous suggestion
- **`Ctrl+Shift+Space`**: Manual completion trigger

### Vim/Neovim

#### Basic Commands
```vim
" Trigger completion
<Tab>

" Accept suggestion
<CR>

" Next suggestion
<C-n>

" Previous suggestion
<C-p>

" Manual completion
<Tab>s
```

#### Key Mappings
```vim
" Custom key mappings
inoremap <silent> <Tab> <Tab>
inoremap <silent> <S-Tab> <C-p>
inoremap <silent> <C-n> <C-n>
inoremap <silent> <C-p> <C-p>
```

## Deployment with Geeper.AI

### Basic Tabby Deployment

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: TabbyDeployment
metadata:
  name: tabby-server
  namespace: default
spec:
  tabby:
    replicas: 2
    image: tabbyml/tabby
    imageTag: latest
    resources:
      requests:
        cpu: "1"
        memory: "2Gi"
      limits:
        cpu: "4"
        memory: "8Gi"
    service:
      type: ClusterIP
      port: 8080
    ingress:
      enabled: true
      host: "tabby.example.com"
```

### Tabby with Ollama Backend

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
    image: tabbyml/tabby
    imageTag: latest
    resources:
      requests:
        cpu: "500m"
        memory: "1Gi"
      limits:
        cpu: "2"
        memory: "4Gi"
    service:
      type: ClusterIP
      port: 8080
    ingress:
      enabled: true
      host: "tabby.localhost"
```

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
