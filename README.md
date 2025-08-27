# LLM Operator 🚀

**Deploy and manage AI/LLM applications on Kubernetes with zero configuration complexity.**

The LLM Operator transforms Kubernetes into a powerful AI platform, letting you deploy production-ready AI applications with simple YAML configurations. No more complex Helm charts, manual setup, or infrastructure headaches.

## ✨ What You Get

### 🎯 **One-Click AI Deployments**
Deploy complete AI stacks with a single YAML file:
- **Ollama** - Run any open-source LLM locally
- **vLLM** - High-performance model serving with GPU acceleration
- **OpenWebUI** - Beautiful web interface for your models
- **Tabby** - AI-powered code completion for your IDE
- **Langfuse** - Monitor and analyze AI performance
- **Redis** - High-performance caching and session storage

### 🚀 **Production-Ready Out of the Box**
- **Auto-scaling** - Scale based on demand automatically
- **Load balancing** - Built-in traffic distribution
- **Monitoring** - Integrated metrics and observability
- **Security** - RBAC, network policies, and secret management
- **Persistence** - Data survives pod restarts and updates

### 🔧 **Zero Configuration Complexity**
- **Smart defaults** - Works immediately with sensible settings
- **Auto-discovery** - Components find each other automatically
- **Built-in networking** - Services, ingresses, and DNS just work
- **Resource optimization** - Right-sized containers and storage

## 🎨 **Use Cases**

### **AI Development & Testing**
Perfect for developers who want to:
- Test different LLM models quickly
- Experiment with AI applications
- Build and iterate on AI features
- Share AI environments with teams

### **Production AI Services**
Enterprise-ready for:
- Customer-facing AI applications
- Internal AI tools and assistants
- AI-powered workflows and automation
- Multi-tenant AI platforms

### **AI Research & Education**
Ideal for:
- Research teams testing new models
- Educational institutions teaching AI
- Hackathons and AI competitions
- Proof-of-concept development

## 🚀 **Quick Start**

### **Option 1: Using Published Helm Chart (Recommended)**

```bash
# Install from GitHub Container Registry
helm registry login ghcr.io
helm install llm-operator oci://ghcr.io/geeper-io/llm-operator \
  --version latest \
  --namespace llm-operator \
  --create-namespace
```

### **Option 2: Deploy Your First AI Stack**

#### **Using Ollama (Default)**
```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: my-ai-stack
spec:
  ollama:
    models:
      - "llama2:7b"
      - "codellama:7b"
  
  openwebui:
    enabled: true
  
  tabby:
    enabled: true
    chatModel: "llama2:7b"
    completionModel: "codellama:7b"
```

#### **Using vLLM (High Performance)**
```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: my-vllm-stack
spec:
  vllm:
    enabled: true
    models:
      - "meta-llama/Llama-2-7b-chat-hf"
      - "codellama/CodeLlama-7b-Instruct-hf"
  
  openwebui:
    enabled: true
  
  tabby:
    enabled: true
    chatModel: "meta-llama/Llama-2-7b-chat-hf"
    completionModel: "codellama/CodeLlama-7b-Instruct-hf"
```

Apply with:
```bash
kubectl apply -f my-ai-stack.yaml
```

**📚 [Full Deployment Guide](docs/QUICKSTART_DEPLOYMENT.md)**

**That's it!** Your AI stack is now running with:
- Ollama or vLLM serving your models
- OpenWebUI providing a web chat interface
- Tabby offering AI code completion
- Automatic networking and monitoring

## 🌟 **Key Features**

### **🧠 Multi-Model Support**
- **Ollama**: Run any open-source LLM locally with easy model management
- **vLLM**: High-performance model serving with GPU acceleration and optimized inference
- Run multiple LLM models simultaneously
- Switch between models instantly
- Mix different model types (chat, code, embedding)
- Automatic model management and updates

### **🔄 CI/CD Pipeline**
- **Automated Builds**: Docker images on every commit and tag
- **Helm Charts**: OCI artifacts published to GHCR
- **Multi-Arch**: AMD64 and ARM64 support
- **Security**: Automated vulnerability scanning
- **Testing**: Comprehensive test suite with E2E validation

### **🚀 GPU Acceleration**
- **NVIDIA GPU Support** - Accelerate inference with CUDA
- **AMD GPU Support** - Use ROCm for AMD graphics cards
- **Intel GPU Support** - Leverage Intel graphics acceleration
- **Resource Management** - Standard Kubernetes resource specification

### **💻 Developer Experience**
- **Tabby Integration** - Get AI code completion in VS Code, Vim, and more
- **WebSocket Support** - Real-time streaming for IDE extensions
- **Custom Prompts** - Tailor AI responses to your needs
- **API Access** - Integrate with your existing applications

### **📊 Observability & Monitoring**
- **Langfuse Integration** - Track AI request performance
- **Request Analytics** - Understand usage patterns
- **Pipeline Monitoring** - Monitor multi-step AI workflows
- **Cost Tracking** - Monitor AI usage and costs

### **🔒 Enterprise Security**
- **RBAC Integration** - Control who accesses what
- **Network Policies** - Secure communication between components
- **Secret Management** - Secure API keys and credentials
- **Audit Logging** - Track all AI interactions

### **📈 Scalability & Reliability**
- **Horizontal Scaling** - Add more replicas as needed
- **Auto-recovery** - Automatic restart on failures
- **Load Distribution** - Spread traffic across multiple pods
- **Data Persistence** - Your data survives restarts

## 🎯 **Why Choose LLM Operator?**

### **vs. Manual Kubernetes Setup**
- **10x faster** deployment
- **Zero configuration** errors
- **Built-in best practices**
- **Automatic updates** and maintenance

### **vs. Traditional VM/Container Deployments**
- **Kubernetes-native** scaling and reliability
- **Integrated monitoring** and logging
- **Easy backup** and disaster recovery
- **Multi-environment** consistency

### **vs. Cloud AI Services**
- **Full control** over your models and data
- **Cost-effective** for high-volume usage
- **Privacy** - keep data in your infrastructure
- **Customization** - modify and extend as needed

## 🚀 **Getting Started**

### **Prerequisites**
- Kubernetes cluster (1.20+)
- `kubectl` configured
- Storage class for persistence (optional)

### **Installation**
```bash
# Install the operator
kubectl apply -f https://raw.githubusercontent.com/your-repo/main/config/crd/bases/llm.geeper.io_lmdeployments.yaml

# Deploy your first AI stack
kubectl apply -f examples/minimal-ollama.yaml
```

### **Next Steps**
1. **Explore Examples** - Check out the `examples/` directory
2. **Customize** - Modify configurations for your needs
3. **Scale** - Add more models and components
4. **Monitor** - Set up Langfuse for observability

## 📚 **Examples & Templates**

### **Basic Deployments**
- **`minimal-ollama.yaml`** - Just Ollama with basic models
- **`ollama-with-openwebui.yaml`** - Ollama + web interface
- **`ollama-with-tabby.yaml`** - Ollama + code completion

### **Advanced Deployments**
- **`openwebui-with-langfuse.yaml`** - Full monitoring stack
- **`production-ollama.yaml`** - Production-ready configuration
- **`openwebui-complete.yaml`** - All features enabled

## 🤝 **Community & Support**

- **Documentation** - Comprehensive guides and examples
- **Examples** - Ready-to-use configurations
- **Issues** - Report bugs and request features
- **Discussions** - Ask questions and share solutions

## 📄 **License**

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Ready to deploy AI at scale?** Start with the [Quick Start Guide](docs/quickstart.md) or jump straight into [examples](examples/) to see what's possible! 🚀