#!/bin/bash

# Deploy Example Script for LLM Operator
# This script demonstrates how to deploy and use the OllamaDeployment CRD

set -e

echo "🚀 Deploying LLM Operator with OllamaDeployment CRD..."

# Build and deploy the operator
echo "📦 Building operator..."
make build

echo "🔧 Installing CRDs..."
make install

echo "🚀 Deploying operator..."
make deploy

echo "⏳ Waiting for operator to be ready..."
kubectl wait --for=condition=available --timeout=300s deployment/llm-operator-controller-manager -n llm-operator-system

echo "✅ Operator deployed successfully!"

# Create a namespace for our example
echo "🏗️ Creating example namespace..."
kubectl create namespace ollama-example --dry-run=client -o yaml | kubectl apply -f -

# Deploy the example OllamaDeployment
echo "📋 Deploying example OllamaDeployment..."
kubectl apply -f config/samples/v1alpha1_ollama_deployment.yaml

echo "⏳ Waiting for OllamaDeployment to be ready..."
kubectl wait --for=condition=ready --timeout=600s ollamadeployment/ollama-example -n default

echo "🎉 Deployment complete!"
echo ""
echo "📊 Check the status:"
echo "kubectl get ollamadeployment ollama-example -o yaml"
echo ""
echo "🔍 View created resources:"
echo "kubectl get all -l ollama-deployment=ollama-example"
echo ""
echo "🌐 Access OpenWebUI (if ingress is configured):"
echo "Add 'ollama-webui.local' to your /etc/hosts file pointing to your cluster IP"
echo ""
echo "🧹 To clean up:"
echo "kubectl delete -f config/samples/v1alpha1_ollama_deployment.yaml"
echo "make undeploy"
echo "make uninstall"
