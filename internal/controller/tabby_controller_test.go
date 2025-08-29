/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	"testing"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/meta"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// testClient is a mock client for testing
type testClient struct {
	secret *corev1.Secret
}

func (c *testClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, opts ...client.GetOption) error {
	if secret, ok := obj.(*corev1.Secret); ok {
		if c.secret != nil {
			*secret = *c.secret
		}
	}
	return nil
}

func (c *testClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	return nil
}

func (c *testClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	return nil
}

func (c *testClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	return nil
}

func (c *testClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	return nil
}

func (c *testClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	return nil
}

func (c *testClient) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	return nil
}

func (c *testClient) Status() client.StatusWriter {
	return nil
}

func (c *testClient) Scheme() *runtime.Scheme {
	return nil
}

func (c *testClient) RESTMapper() meta.RESTMapper {
	return nil
}

func (c *testClient) GroupVersionKindFor(obj runtime.Object) (schema.GroupVersionKind, error) {
	return schema.GroupVersionKind{}, nil
}

func (c *testClient) IsObjectNamespaced(obj runtime.Object) (bool, error) {
	return false, nil
}

func (c *testClient) SubResource(subResource string) client.SubResourceClient {
	return nil
}

func TestTabbyController_ConfigGeneration(t *testing.T) {
	t.Run("generateTabbyConfig", func(t *testing.T) {
		reconciler := &LMDeploymentReconciler{}

		t.Run("should generate correct TOML configuration with default model", func(t *testing.T) {
			deployment := &llmgeeperiov1alpha1.LMDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-deployment",
					Namespace: "default",
				},
				Spec: llmgeeperiov1alpha1.LMDeploymentSpec{
					Ollama: llmgeeperiov1alpha1.OllamaSpec{
						Models: []string{"codellama:7b"},
						Service: llmgeeperiov1alpha1.ServiceSpec{
							Port: 11434,
						},
					},
					Tabby: llmgeeperiov1alpha1.TabbySpec{
						Enabled:         true,
						ChatModel:       "codellama:7b",
						CompletionModel: "codellama:7b",
					},
				},
			}

			config, err := reconciler.generateTabbyConfig(t.Context(), deployment)
			require.NoError(t, err)

			expectedConfig := `[model]
  [model.completion]
    [model.completion.http]
      kind = "ollama/completion"
      model_name = "codellama:7b"
      api_endpoint = "http://test-deployment-ollama.default:11434"
      prompt_template = "<PRE> {prefix} <SUF>{suffix} <MID>"
  [model.chat]
    [model.chat.http]
      kind = "openai/chat"
      model_name = "codellama:7b"
      api_endpoint = "http://test-deployment-ollama.default:11434/v1"
      supported_models = ["codellama:7b"]
  [model.embedding]
    [model.embedding.local]
      model_id = "Nomic-Embed-Text"
`
			assert.Equal(t, expectedConfig, config)
		})

		t.Run("should generate correct TOML configuration with custom model name", func(t *testing.T) {
			deployment := &llmgeeperiov1alpha1.LMDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-deployment",
					Namespace: "default",
				},
				Spec: llmgeeperiov1alpha1.LMDeploymentSpec{
					Ollama: llmgeeperiov1alpha1.OllamaSpec{
						Models: []string{
							"llama2:7b",
							"codellama:13b",
						},
						Service: llmgeeperiov1alpha1.ServiceSpec{
							Port: 11434,
						},
					},
					Tabby: llmgeeperiov1alpha1.TabbySpec{
						Enabled:         true,
						ChatModel:       "llama2:7b",
						CompletionModel: "codellama:13b",
					},
				},
			}

			config, err := reconciler.generateTabbyConfig(t.Context(), deployment)
			require.NoError(t, err)

			expectedConfig := `[model]
  [model.completion]
    [model.completion.http]
      kind = "ollama/completion"
      model_name = "codellama:13b"
      api_endpoint = "http://test-deployment-ollama.default:11434"
      prompt_template = "<PRE> {prefix} <SUF>{suffix} <MID>"
  [model.chat]
    [model.chat.http]
      kind = "openai/chat"
      model_name = "llama2:7b"
      api_endpoint = "http://test-deployment-ollama.default:11434/v1"
      supported_models = ["llama2:7b", "codellama:13b"]
  [model.embedding]
    [model.embedding.local]
      model_id = "Nomic-Embed-Text"
`
			assert.Equal(t, expectedConfig, config)
		})

		t.Run("should use correct service name and port", func(t *testing.T) {
			deployment := &llmgeeperiov1alpha1.LMDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-app",
					Namespace: "production",
				},
				Spec: llmgeeperiov1alpha1.LMDeploymentSpec{
					Ollama: llmgeeperiov1alpha1.OllamaSpec{
						Models: []string{"llama2:7b"},
						Service: llmgeeperiov1alpha1.ServiceSpec{
							Port: 8080,
						},
					},
					Tabby: llmgeeperiov1alpha1.TabbySpec{
						Enabled:         true,
						ChatModel:       "llama2:7b",
						CompletionModel: "llama2:7b",
					},
				},
			}

			config, err := reconciler.generateTabbyConfig(t.Context(), deployment)
			require.NoError(t, err)

			expectedConfig := `[model]
  [model.completion]
    [model.completion.http]
      kind = "ollama/completion"
      model_name = "llama2:7b"
      api_endpoint = "http://my-app-ollama.production:8080"
      prompt_template = "<PRE> {prefix} <SUF>{suffix} <MID>"
  [model.chat]
    [model.chat.http]
      kind = "openai/chat"
      model_name = "llama2:7b"
      api_endpoint = "http://my-app-ollama.production:8080/v1"
      supported_models = ["llama2:7b"]
  [model.embedding]
    [model.embedding.local]
      model_id = "Nomic-Embed-Text"
`
			assert.Equal(t, expectedConfig, config)
		})

		t.Run("should generate correct TOML configuration with vLLM models and API key", func(t *testing.T) {
			deployment := &llmgeeperiov1alpha1.LMDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "my-app",
					Namespace: "production",
				},
				Spec: llmgeeperiov1alpha1.LMDeploymentSpec{
					VLLM: llmgeeperiov1alpha1.VLLMSpec{
						Enabled: true,
						Models: []llmgeeperiov1alpha1.VLLMModelSpec{
							{
								Name:  "llama2-7b",
								Model: "meta-llama/Llama-2-7b-chat-hf",
							},
							{
								Name:  "codellama-7b",
								Model: "codellama/CodeLlama-7b-Instruct-hf",
							},
						},
						ApiKey: &corev1.SecretReference{
							Name: "my-app-vllm-api-key",
						},
					},
					Tabby: llmgeeperiov1alpha1.TabbySpec{
						Enabled:         true,
						ChatModel:       "meta-llama/Llama-2-7b-chat-hf",
						CompletionModel: "codellama/CodeLlama-7b-Instruct-hf",
					},
				},
			}

			// Mock the secret data
			reconciler := &LMDeploymentReconciler{
				Client: &testClient{
					secret: &corev1.Secret{
						Data: map[string][]byte{
							"VLLM_API_KEY": []byte("test-api-key-12345"),
						},
					},
				},
			}

			config, err := reconciler.generateTabbyConfig(t.Context(), deployment)
			require.NoError(t, err)

			expectedConfig := `[model]
  [model.completion]
    [model.completion.http]
      kind = "openai/completion"
      model_name = "codellama/CodeLlama-7b-Instruct-hf"
      api_endpoint = "http://my-app-vllm.production:8000/v1"
      api_key = "test-api-key-12345"
      prompt_template = "<PRE> {prefix} <SUF>{suffix} <MID>"
  [model.chat]
    [model.chat.http]
      kind = "openai/chat"
      model_name = "meta-llama/Llama-2-7b-chat-hf"
      api_endpoint = "http://my-app-vllm.production:8000/v1"
      api_key = "test-api-key-12345"
      supported_models = ["meta-llama/Llama-2-7b-chat-hf", "codellama/CodeLlama-7b-Instruct-hf"]
  [model.embedding]
    [model.embedding.local]
      model_id = "Nomic-Embed-Text"
`
			assert.Equal(t, expectedConfig, config)
		})
	})

	t.Run("buildTabbySecret", func(t *testing.T) {
		reconciler := &LMDeploymentReconciler{}

		t.Run("should create Secret with correct metadata and data", func(t *testing.T) {
			deployment := &llmgeeperiov1alpha1.LMDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-deployment",
					Namespace: "default",
				},
				Spec: llmgeeperiov1alpha1.LMDeploymentSpec{
					Ollama: llmgeeperiov1alpha1.OllamaSpec{
						Models: []string{"test-model:1.0"},
						Service: llmgeeperiov1alpha1.ServiceSpec{
							Port: 11434,
						},
					},
					Tabby: llmgeeperiov1alpha1.TabbySpec{
						Enabled:         true,
						ChatModel:       "test-model:1.0",
						CompletionModel: "test-model:1.0",
					},
				},
			}

			secret, err := reconciler.buildTabbySecret(t.Context(), deployment)
			require.NoError(t, err)
			assert.NotNil(t, secret)

			// Check metadata
			assert.Equal(t, "test-deployment-tabby-config", secret.Name)
			assert.Equal(t, "default", secret.Namespace)
			assert.Equal(t, "tabby", secret.Labels["app"])
			assert.Equal(t, "test-deployment", secret.Labels["llm-deployment"])

			// Check data
			assert.Contains(t, secret.Data, "config.toml")
			assert.NotEmpty(t, secret.Data["config.toml"])

			// Check that the TOML content matches expected format
			expectedConfig := `[model]
  [model.completion]
    [model.completion.http]
      kind = "ollama/completion"
      model_name = "test-model:1.0"
      api_endpoint = "http://test-deployment-ollama.default:11434"
      prompt_template = "<PRE> {prefix} <SUF>{suffix} <MID>"
  [model.chat]
    [model.chat.http]
      kind = "openai/chat"
      model_name = "test-model:1.0"
      api_endpoint = "http://test-deployment-ollama.default:11434/v1"
      supported_models = ["test-model:1.0"]
  [model.embedding]
    [model.embedding.local]
      model_id = "Nomic-Embed-Text"
`
			assert.Equal(t, expectedConfig, string(secret.Data["config.toml"]))
		})
	})
}
