---
id: continue-dev
title: Continue.dev Integration
sidebar_label: Continue.dev
description: Learn how to integrate Continue.dev with Geeper.AI for enhanced coding assistance
---

# Continue.dev Integration

Continue.dev is a powerful AI-powered coding assistant that integrates seamlessly with your development environment. It provides intelligent code completion, refactoring suggestions, and coding guidance directly in your IDE.

## What is Continue.dev?

Continue.dev is an open-source AI coding assistant that:

- **Understands Context**: Analyzes your entire codebase for intelligent suggestions
- **Multi-Model Support**: Works with various LLM backends including your Geeper.AI deployments
- **IDE Integration**: Available as extensions for VS Code, JetBrains IDEs, and more
- **Real-time Assistance**: Provides instant feedback and suggestions as you code
- **Customizable**: Configurable prompts and workflows for your specific needs
- **Privacy-First**: Runs locally or connects to your private LLM instances

## Key Features

- **🎯 Smart Code Completion**: Context-aware suggestions based on your codebase
- **🔧 Refactoring Assistance**: AI-powered code improvements and optimizations
- **📝 Documentation Generation**: Automatic documentation and comment generation
- **🐛 Bug Detection**: Identify potential issues and suggest fixes
- **🚀 Code Generation**: Generate boilerplate code and implementations
- **📚 Learning**: Learn from your coding patterns and preferences

## Installation

### VS Code Extension

1. **Open VS Code Extensions**:
   - Press `Ctrl+Shift+X` (Windows/Linux) or `Cmd+Shift+X` (Mac)
   - Or go to View → Extensions

2. **Search for Continue**:
   - Type "Continue" in the search box
   - Look for "Continue - AI-powered coding assistant"

3. **Install Extension**:
   - Click "Install" on the Continue extension
   - Restart VS Code when prompted

4. **Verify Installation**:
   - Check the Extensions panel for Continue
   - Look for Continue icon in the sidebar

### JetBrains IDEs (IntelliJ, PyCharm, etc.)

1. **Open Plugin Manager**:
   - Go to File → Settings → Plugins (Windows/Linux)
   - Or File → Preferences → Plugins (Mac)

2. **Search for Continue**:
   - Click "Marketplace" tab
   - Search for "Continue"

3. **Install Plugin**:
   - Click "Install" on the Continue plugin
   - Restart your IDE

4. **Verify Installation**:
   - Check the Plugins panel for Continue
   - Look for Continue in the Tools menu

### Command Line Installation

```bash
# Install Continue CLI globally
npm install -g @continue/cli

# Verify installation
continue --version

# Initialize Continue in your project
cd your-project
continue init
```

## Configuration

### Basic Setup

1. **Open Continue Settings**:
   - VS Code: `Ctrl+,` (Windows/Linux) or `Cmd+,` (Mac)
   - Search for "Continue" in settings

2. **Configure LLM Backend**:
   - Set your Geeper.AI endpoint
   - Configure authentication if required
   - Choose your preferred model

3. **Customize Behavior**:
   - Adjust response length
   - Set coding style preferences
   - Configure file exclusions

### Geeper.AI Integration

```json
{
  "continue": {
    "llm": {
      "provider": "custom",
      "endpoint": "http://your-geeper-ai-endpoint:8080",
      "model": "codellama:34b",
      "apiKey": "your-api-key"
    },
    "context": {
      "maxTokens": 4000,
      "includePatterns": ["**/*.{js,ts,py,go,rs,java}"],
      "excludePatterns": ["**/node_modules/**", "**/dist/**"]
    }
  }
}
```

## Usage

### Basic Commands

- **`/edit`**: Edit the current selection or file
- **`/explain`**: Explain the selected code
- **`/fix`**: Fix issues in the selected code
- **`/test`**: Generate tests for the selected code
- **`/doc`**: Generate documentation

### Keyboard Shortcuts

- **`Ctrl+Shift+L`** (Windows/Linux) or **`Cmd+Shift+L`** (Mac): Open Continue chat
- **`Ctrl+Shift+E`** (Windows/Linux) or **`Cmd+Shift+E`** (Mac): Edit with Continue
- **`Ctrl+Shift+X`** (Windows/Linux) or **`Cmd+Shift+X`** (Mac): Explain code

### Chat Interface

1. **Open Continue Chat**:
   - Click Continue icon in sidebar
   - Or use keyboard shortcut

2. **Ask Questions**:
   - "How do I implement authentication?"
   - "What's wrong with this function?"
   - "Generate a test for this class"

3. **Review Suggestions**:
   - Accept or reject changes
   - Modify suggestions as needed
   - Apply changes to your code

## Best Practices

### 1. Context Management
- **Include Relevant Files**: Add related files to provide better context
- **Exclude Unnecessary Files**: Avoid including build artifacts and dependencies
- **Use Clear Prompts**: Be specific about what you want to accomplish

### 2. Code Quality
- **Review Suggestions**: Always review AI-generated code before applying
- **Test Changes**: Run tests after applying AI suggestions
- **Iterate**: Use feedback to improve future suggestions

### 3. Security
- **Private Endpoints**: Use your Geeper.AI instances for sensitive code
- **Code Review**: Don't blindly accept AI suggestions
- **Audit Regularly**: Review AI-generated code for security issues

## Troubleshooting

### Common Issues

1. **Extension Not Loading**:
   - Restart VS Code/IDE
   - Check extension compatibility
   - Verify installation

2. **Connection Errors**:
   - Check Geeper.AI endpoint
   - Verify network connectivity
   - Check authentication credentials

3. **Poor Suggestions**:
   - Improve context by including more files
   - Use clearer prompts
   - Check LLM model quality

### Debug Commands

```bash
# Check Continue status
continue status

# Test connection to Geeper.AI
continue test-connection

# View logs
continue logs

# Reset configuration
continue reset
```

## Next Steps

- [Tabby Integration](/docs/coding-assistants/tabby) - Learn about Tabby code completion
- [Advanced Configuration](/docs/coding-assistants/advanced-config) - Deep dive into settings
- [Custom Workflows](/docs/coding-assistants/workflows) - Create custom coding workflows
- [API Reference](/docs/api/continue-dev) - Complete Continue.dev API documentation

---

*Continue.dev transforms your coding experience with AI-powered assistance powered by Geeper.AI*
