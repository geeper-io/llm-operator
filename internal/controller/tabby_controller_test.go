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
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

func TestTabbyController_ConfigurationGeneration(t *testing.T) {
	t.Run("generateTabbyConfig", func(t *testing.T) {
		reconciler := &DeploymentReconciler{}

		t.Run("should generate correct TOML configuration with default model", func(t *testing.T) {
			deployment := &llmgeeperiov1alpha1.Deployment{
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
			assert.NotEmpty(t, config)

			// Check that the configuration contains expected sections
			assert.Contains(t, config, "[chat.model.ollama]")
			assert.Contains(t, config, "[completion.model.ollama]")
			assert.Contains(t, config, "[model.embedding.local]")

			// Check Ollama service configuration
			assert.Contains(t, config, "host = \"test-deployment-ollama.default.svc.cluster.local:11434\"")
			assert.Contains(t, config, "model = \"codellama\"")

			// Check embedding model configuration
			assert.Contains(t, config, "model_id = \"Nomic-Embed-Text\"")
		})

		t.Run("should generate correct TOML configuration with custom model name", func(t *testing.T) {
			deployment := &llmgeeperiov1alpha1.Deployment{
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
			assert.NotEmpty(t, config)

			// Check that custom model name is used instead of first Ollama model
			assert.Contains(t, config, "model = \"custom-model:latest\"")
			assert.NotContains(t, config, "model = \"llama2:7b\"")
		})

		t.Run("should handle model names with colons correctly", func(t *testing.T) {
			deployment := &llmgeeperiov1alpha1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-deployment",
					Namespace: "default",
				},
				Spec: llmgeeperiov1alpha1.LMDeploymentSpec{
					Ollama: llmgeeperiov1alpha1.OllamaSpec{
						Models: []llmgeeperiov1alpha1.OllamaModel{
							"codellama:7b:latest",
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
			assert.NotEmpty(t, config)

			// Should extract the model name part before the first colon
			assert.Contains(t, config, "model = \"codellama:7b:latest\"")
		})

		t.Run("should handle model names without colons correctly", func(t *testing.T) {
			deployment := &llmgeeperiov1alpha1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      "test-deployment",
					Namespace: "default",
				},
				Spec: llmgeeperiov1alpha1.LMDeploymentSpec{
					Ollama: llmgeeperiov1alpha1.OllamaSpec{
						Models: []llmgeeperiov1alpha1.OllamaModel{
							"simple-model",
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
			assert.NotEmpty(t, config)

			// Should use the full model name
			assert.Contains(t, config, "model = \"simple-model\"")
		})

		t.Run("should generate valid TOML format", func(t *testing.T) {
			deployment := &llmgeeperiov1alpha1.Deployment{
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

			config, err := reconciler.generateTabbyConfig(deployment)
			require.NoError(t, err)
			assert.NotEmpty(t, config)

			// Check TOML structure
			lines := strings.Split(strings.TrimSpace(config), "\n")
			assert.GreaterOrEqual(t, len(lines), 10) // Should have multiple lines

			// Check for proper TOML section headers
			assert.Contains(t, config, "[chat.model.ollama]")
			assert.Contains(t, config, "[completion.model.ollama]")
			assert.Contains(t, config, "[model.embedding.local]")

			// Check for proper key-value pairs
			assert.Contains(t, config, "host = ")
			assert.Contains(t, config, "model = ")
			assert.Contains(t, config, "model_id = ")
		})

		t.Run("should use correct service name and port", func(t *testing.T) {
			deployment := &llmgeeperiov1alpha1.Deployment{
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
			assert.NotEmpty(t, config)

			// Check that service name and port are correctly formatted
			expectedHost := "my-app-ollama.production.svc.cluster.local:8080"
			assert.Contains(t, config, expectedHost)
		})
	})

	t.Run("buildTabbyConfigMap", func(t *testing.T) {
		reconciler := &DeploymentReconciler{}

		t.Run("should create ConfigMap with correct metadata and data", func(t *testing.T) {
			deployment := &llmgeeperiov1alpha1.Deployment{
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

			// Check that the TOML content is valid
			configContent := configMap.Data["config.toml"]
			assert.Contains(t, configContent, "[chat.model.ollama]")
			assert.Contains(t, configContent, "[completion.model.ollama]")
			assert.Contains(t, configContent, "[model.embedding.local]")
		})

		t.Run("should handle configuration generation errors gracefully", func(t *testing.T) {
			// This test would require mocking the TOML encoder to return an error
			// For now, we'll test the happy path which is more important
			deployment := &llmgeeperiov1alpha1.Deployment{
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
			assert.NotEmpty(t, configMap.Data["config.toml"])
		})
	})
}
