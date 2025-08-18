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
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

func TestTabbyController_UnitTests(t *testing.T) {
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
						Models: []llmgeeperiov1alpha1.OllamaModel{
							"codellama:7b",
						},
						Service: llmgeeperiov1alpha1.ServiceSpec{
							Port: 11434,
						},
					},
					Tabby: llmgeeperiov1alpha1.TabbySpec{
						Enabled: true,
					},
				},
			}

			config, err := reconciler.generateTabbyConfig(deployment)
			require.NoError(t, err)

			expectedConfig := `[model]
  [model.completion]
    [model.completion.http]
      kind = "ollama/completion"
      model_name = "codellama:7b"
      api_endpoint = "http://test-deployment-ollama.default.svc.cluster.local:11434"
      prompt_template = "<PRE> {prefix} <SUF>{suffix} <MID>"
  [model.chat]
    [model.chat.http]
      kind = "ollama/chat"
      model_name = "codellama:7b"
      api_endpoint = "http://test-deployment-ollama.default.svc.cluster.local:11434"
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
						Models: []llmgeeperiov1alpha1.OllamaModel{
							"llama2:7b",
							"codellama:13b",
						},
						Service: llmgeeperiov1alpha1.ServiceSpec{
							Port: 11434,
						},
					},
					Tabby: llmgeeperiov1alpha1.TabbySpec{
						Enabled:   true,
						ModelName: "custom-model:latest",
					},
				},
			}

			config, err := reconciler.generateTabbyConfig(deployment)
			require.NoError(t, err)

			expectedConfig := `[model]
  [model.completion]
    [model.completion.http]
      kind = "ollama/completion"
      model_name = "custom-model:latest"
      api_endpoint = "http://test-deployment-ollama.default.svc.cluster.local:11434"
      prompt_template = "<PRE> {prefix} <SUF>{suffix} <MID>"
  [model.chat]
    [model.chat.http]
      kind = "ollama/chat"
      model_name = "custom-model:latest"
      api_endpoint = "http://test-deployment-ollama.default.svc.cluster.local:11434"
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
						Models: []llmgeeperiov1alpha1.OllamaModel{
							"llama2:7b",
						},
						Service: llmgeeperiov1alpha1.ServiceSpec{
							Port: 8080,
						},
					},
					Tabby: llmgeeperiov1alpha1.TabbySpec{
						Enabled: true,
					},
				},
			}

			config, err := reconciler.generateTabbyConfig(deployment)
			require.NoError(t, err)

			expectedConfig := `[model]
  [model.completion]
    [model.completion.http]
      kind = "ollama/completion"
      model_name = "llama2:7b"
      api_endpoint = "http://my-app-ollama.production.svc.cluster.local:8080"
      prompt_template = "<PRE> {prefix} <SUF>{suffix} <MID>"
  [model.chat]
    [model.chat.http]
      kind = "ollama/chat"
      model_name = "llama2:7b"
      api_endpoint = "http://my-app-ollama.production.svc.cluster.local:8080"
  [model.embedding]
    [model.embedding.local]
      model_id = "Nomic-Embed-Text"
`
			assert.Equal(t, expectedConfig, config)
		})
	})

	t.Run("buildTabbyConfigMap", func(t *testing.T) {
		reconciler := &LMDeploymentReconciler{}

		t.Run("should create ConfigMap with correct metadata and data", func(t *testing.T) {
			deployment := &llmgeeperiov1alpha1.LMDeployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-deployment",
					Namespace: "default",
				},
				Spec: llmgeeperiov1alpha1.LMDeploymentSpec{
					Ollama: llmgeeperiov1alpha1.OllamaSpec{
						Models: []llmgeeperiov1alpha1.OllamaModel{
							"test-model:1.0",
						},
						Service: llmgeeperiov1alpha1.ServiceSpec{
							Port: 11434,
						},
					},
					Tabby: llmgeeperiov1alpha1.TabbySpec{
						Enabled: true,
					},
				},
			}

			configMap := reconciler.buildTabbyConfigMap(deployment)
			assert.NotNil(t, configMap)

			// Check metadata
			assert.Equal(t, "test-deployment-tabby-config", configMap.Name)
			assert.Equal(t, "default", configMap.Namespace)
			assert.Equal(t, "tabby", configMap.Labels["app"])
			assert.Equal(t, "test-deployment", configMap.Labels["llm-deployment"])

			// Check data
			assert.Contains(t, configMap.Data, "config.toml")
			assert.NotEmpty(t, configMap.Data["config.toml"])

			// Check that the TOML content matches expected format
			expectedConfig := `[model]
  [model.completion]
    [model.completion.http]
      kind = "ollama/completion"
      model_name = "test-model:1.0"
      api_endpoint = "http://test-deployment-ollama.default.svc.cluster.local:11434"
      prompt_template = "<PRE> {prefix} <SUF>{suffix} <MID>"
  [model.chat]
    [model.chat.http]
      kind = "ollama/chat"
      model_name = "test-model:1.0"
      api_endpoint = "http://test-deployment-ollama.default.svc.cluster.local:11434"
  [model.embedding]
    [model.embedding.local]
      model_id = "Nomic-Embed-Text"
`
			assert.Equal(t, expectedConfig, configMap.Data["config.toml"])
		})
	})
}
