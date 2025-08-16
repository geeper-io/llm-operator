---
id: rag
title: Retrieval-Augmented Generation (RAG)
sidebar_label: RAG Integration
description: Learn how to implement RAG to enhance LLM responses with external knowledge
---

# Retrieval-Augmented Generation (RAG)

Retrieval-Augmented Generation (RAG) is a powerful technique that combines the capabilities of Large Language Models with external knowledge sources. This approach enhances LLM responses by providing relevant, up-to-date information from your own data sources.

## What is RAG?

RAG works by:
1. **Retrieving** relevant documents or information from a knowledge base
2. **Augmenting** the LLM's context with this retrieved information
3. **Generating** more accurate and informed responses

### Benefits of RAG

- **Up-to-date Information**: Access current data without retraining models
- **Domain Expertise**: Incorporate company-specific knowledge
- **Factual Accuracy**: Reduce hallucinations with verified sources
- **Cost Efficiency**: Use smaller models with external knowledge
- **Customization**: Tailor responses to your specific use case

## RAG Architecture in Geeper.AI

```
┌─────────────────┐    ┌──────────────────┐    ┌─────────────────┐
│   User Query   │───▶│   RAG Pipeline   │───▶│   LLM Response  │
└─────────────────┘    └──────────────────┘    └─────────────────┘
                              │
                              ▼
                       ┌──────────────────┐
                       │  Vector Store    │
                       │  (Chroma, Qdrant)│
                       └──────────────────┘
```

## Quick Setup

### 1. Deploy RAG Components

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: RAGDeployment
metadata:
  name: rag-system
  namespace: default
spec:
  components:
    - name: vector-store
      type: chroma
      replicas: 1
      resources:
        requests:
          memory: "1Gi"
          cpu: "500m"
    - name: embedding-service
      type: sentence-transformers
      model: all-MiniLM-L6-v2
      replicas: 2
      resources:
        requests:
          memory: "2Gi"
          cpu: "1"
    - name: rag-api
      type: fastapi
      replicas: 2
      resources:
        requests:
          memory: "512Mi"
          cpu: "250m"
  storage:
    type: persistent
    size: "10Gi"
```

### 2. Apply the Configuration

```bash
kubectl apply -f rag-deployment.yaml
```

### 3. Verify Deployment

```bash
kubectl get pods -l app=rag-system
kubectl get svc -l app=rag-system
```

## Data Ingestion

### Document Processing

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: DocumentProcessor
metadata:
  name: docs-processor
spec:
  input:
    type: s3
    bucket: "my-documents"
    prefix: "docs/"
  processing:
    chunkSize: 1000
    chunkOverlap: 200
    embeddingModel: "all-MiniLM-L6-v2"
  output:
    vectorStore: "rag-system-vector-store"
    namespace: "company-docs"
```

### Supported Document Types

- **Text Files**: `.txt`, `.md`, `.rst`
- **Office Documents**: `.pdf`, `.docx`, `.pptx`
- **Code Files**: `.py`, `.js`, `.java`, `.go`
- **Structured Data**: `.json`, `.csv`, `.xml`
- **Web Content**: HTML, RSS feeds

### Batch Processing

```bash
# Process documents in batch
kubectl apply -f document-processor.yaml

# Monitor processing
kubectl logs -f deployment/docs-processor

# Check vector store status
kubectl exec -it deployment/rag-system-vector-store -- chroma-cli list-collections
```

## RAG Configuration

### Vector Store Options

#### Chroma (Default)
```yaml
spec:
  vectorStore:
    type: chroma
    config:
      persistDirectory: "/data"
      collectionName: "company-knowledge"
      distanceFunction: "cosine"
```

#### Qdrant
```yaml
spec:
  vectorStore:
    type: qdrant
    config:
      url: "http://qdrant:6333"
      collectionName: "company-knowledge"
      vectorSize: 384
```

#### Pinecone
```yaml
spec:
  vectorStore:
    type: pinecone
    config:
      apiKey: "your-pinecone-api-key"
      environment: "us-west1-gcp"
      indexName: "company-knowledge"
```

### Embedding Models

```yaml
spec:
  embedding:
    model: "all-MiniLM-L6-v2"  # 384 dimensions
    # Alternative models:
    # - "all-mpnet-base-v2"     # 768 dimensions
    # - "multi-qa-MiniLM-L6-v2" # 384 dimensions
    # - "text-embedding-ada-002" # OpenAI embeddings
    batchSize: 32
    maxLength: 512
```

### Retrieval Configuration

```yaml
spec:
  retrieval:
    topK: 5                    # Number of documents to retrieve
    similarityThreshold: 0.7   # Minimum similarity score
    rerank: true              # Enable reranking
    rerankModel: "ms-marco-MiniLM-L-6-v2"
    maxTokens: 4000           # Maximum context length
```

## Integration with Chat

### OpenWebUI RAG Integration

```yaml
apiVersion: llm.geeper.io/v1alpha1
kind: OpenWebUIDeployment
metadata:
  name: openwebui-with-rag
spec:
  # ... other config ...
  env:
    - name: RAG_ENABLED
      value: "true"
    - name: RAG_API_URL
      value: "http://rag-system-rag-api:8000"
    - name: RAG_COLLECTION
      value: "company-knowledge"
    - name: RAG_AUTO_RETRIEVE
      value: "true"
```

### Custom RAG API

```python
# Example RAG API endpoint
@app.post("/rag/query")
async def rag_query(query: str, top_k: int = 5):
    # 1. Generate embeddings for query
    query_embedding = embedding_model.encode(query)
    
    # 2. Search vector store
    results = vector_store.similarity_search(
        query_embedding, 
        k=top_k
    )
    
    # 3. Format context
    context = format_documents(results)
    
    # 4. Return augmented prompt
    return {
        "context": context,
        "augmented_prompt": f"Context: {context}\n\nQuery: {query}"
    }
```

## Advanced RAG Features

### Hybrid Search

```yaml
spec:
  search:
    type: hybrid
    components:
      - type: vector
        weight: 0.7
      - type: keyword
        weight: 0.3
        algorithm: bm25
      - type: semantic
        weight: 0.5
        model: "sentence-transformers/all-mpnet-base-v2"
```

### Multi-Modal RAG

```yaml
spec:
  multimodal:
    enabled: true
    models:
      - type: text
        model: "all-MiniLM-L6-v2"
      - type: image
        model: "clip-ViT-B-32"
      - type: audio
        model: "wav2vec2-base"
```

### RAG Chaining

```yaml
spec:
  chaining:
    enabled: true
    steps:
      - name: "document-retrieval"
        type: "vector-search"
      - name: "fact-checking"
        type: "external-api"
        endpoint: "https://fact-check-api.com"
      - name: "response-generation"
        type: "llm-completion"
```

## Performance Optimization

### Caching

```yaml
spec:
  caching:
    enabled: true
    type: redis
    ttl: 3600  # 1 hour
    maxSize: "1Gi"
```

### Scaling

```yaml
spec:
  scaling:
    autoScaling: true
    minReplicas: 2
    maxReplicas: 10
    targetCPUUtilizationPercentage: 70
    targetMemoryUtilizationPercentage: 80
```

### Monitoring

```yaml
spec:
  monitoring:
    metrics:
      - type: retrieval_latency
      - type: embedding_generation_time
      - type: vector_search_accuracy
      - type: cache_hit_rate
    alerting:
      - metric: retrieval_latency
        threshold: 1000ms
        severity: warning
```

## Best Practices

### 1. Data Quality
- **Clean Documents**: Remove duplicates and irrelevant content
- **Structured Format**: Use consistent document structure
- **Metadata**: Include source, date, and version information

### 2. Chunking Strategy
- **Semantic Chunking**: Split by meaning, not just length
- **Overlap**: Use 10-20% overlap between chunks
- **Context Preservation**: Keep related information together

### 3. Embedding Selection
- **Domain-Specific**: Choose models trained on similar content
- **Dimension Balance**: Balance accuracy vs. performance
- **Fine-tuning**: Consider fine-tuning for your specific domain

### 4. Retrieval Optimization
- **Reranking**: Use reranking models for better relevance
- **Hybrid Search**: Combine vector and keyword search
- **Feedback Loop**: Collect user feedback to improve retrieval

## Troubleshooting

### Common Issues

1. **Low Retrieval Quality**:
   - Check embedding model suitability
   - Verify chunking strategy
   - Review similarity thresholds

2. **Slow Performance**:
   - Enable caching
   - Optimize vector store configuration
   - Scale embedding services

3. **Memory Issues**:
   - Reduce batch sizes
   - Use smaller embedding models
   - Implement streaming for large documents

### Debug Commands

```bash
# Check vector store health
kubectl exec -it deployment/rag-system-vector-store -- chroma-cli health

# Test embedding service
kubectl exec -it deployment/rag-system-embedding-service -- curl -X POST http://localhost:8000/embed -d '{"text": "test"}'

# Monitor RAG API
kubectl logs -f deployment/rag-system-rag-api
```

## Next Steps

- [Plugin System](/docs/chat/plugins) - Extend RAG with custom plugins
- [Advanced RAG Patterns](/docs/chat/advanced-rag) - Advanced RAG architectures
- [Performance Tuning](/docs/chat/performance) - Optimize RAG performance
- [API Reference](/docs/api/rag) - Complete RAG API documentation

---

*RAG transforms your LLMs from general-purpose tools to domain experts with access to your specific knowledge*
