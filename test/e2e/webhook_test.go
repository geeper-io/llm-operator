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

package e2e

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/geeper-io/llm-operator/test/utils"
)

// WebhookTestSuite defines the test suite for webhook e2e tests
type WebhookTestSuite struct {
	GlobalE2ESuite
}

// TestWebhookDefaulting tests that the defaulting webhook correctly sets default values
func (webhookTestSuite *WebhookTestSuite) TestWebhookDefaulting() {
	deploymentName := "test-webhook-defaulting"

	webhookTestSuite.T().Cleanup(func() {
		webhookTestSuite.T().Log("Cleaning up LMDeployment")
		cmd := exec.Command("kubectl", "delete", "lmdeployment", deploymentName, "-n", webhookTestSuite.testNamespace)
		_, _ = utils.Run(cmd)
	})

	webhookTestSuite.T().Log("Creating LMDeployment with minimal configuration to test defaults")
	minimalDeployment := `apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: test-webhook-defaulting
  namespace: lmdeployment-webhook-test
spec:
  ollama:
    models:
    - "hf.co/amakhov/tiny-random-llama:F16"`

	// Write YAML to temporary file
	yamlFile := filepath.Join("/tmp", "test-webhook-defaulting.yaml")
	err := os.WriteFile(yamlFile, []byte(minimalDeployment), 0644)
	require.NoError(webhookTestSuite.T(), err)

	webhookTestSuite.T().Log("Applying LMDeployment YAML")
	cmd := exec.Command("kubectl", "apply", "-f", yamlFile)
	_, err = utils.Run(cmd)
	require.NoError(webhookTestSuite.T(), err, "Failed to apply LMDeployment")

	webhookTestSuite.T().Log("Waiting for LMDeployment to be ready")
	webhookTestSuite.waitForLMDeploymentReady(deploymentName, 5*time.Minute)

	webhookTestSuite.T().Log("Verifying that defaults were applied by the webhook")
	webhookTestSuite.verifyWebhookDefaults(deploymentName)

	// Clean up temporary file
	_ = os.Remove(yamlFile)
}

// TestWebhookValidation tests that the validation webhook correctly rejects invalid configurations
func (webhookTestSuite *WebhookTestSuite) TestWebhookValidation() {
	webhookTestSuite.T().Log("Testing webhook validation with invalid configurations")

	webhookTestSuite.testInvalidConfiguration(
		"test-webhook-validation-invalid-tabby-model",
		`apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: test-webhook-validation-invalid-tabby-model
  namespace: lmdeployment-webhook-test
spec:
  ollama:
    models:
    - "hf.co/amakhov/tiny-random-llama:F16"
  tabby:
    enabled: true
    chatModel: "non-existent-model"`,
		"chat model must be one of the models specified in spec.ollama.models",
	)
}

// testInvalidConfiguration tests that an invalid configuration is rejected by the webhook
func (webhookTestSuite *WebhookTestSuite) testInvalidConfiguration(name, yamlContent, expectedError string) {
	webhookTestSuite.T().Cleanup(func() {
		webhookTestSuite.T().Logf("Cleaning up invalid LMDeployment: %s", name)
		cmd := exec.Command("kubectl", "delete", "lmdeployment", name, "-n", webhookTestSuite.testNamespace)
		_, _ = utils.Run(cmd)
	})

	webhookTestSuite.T().Logf("Testing invalid configuration: %s", name)

	// Write YAML to temporary file
	yamlFile := filepath.Join("/tmp", name+".yaml")
	err := os.WriteFile(yamlFile, []byte(yamlContent), 0644)
	require.NoError(webhookTestSuite.T(), err)

	webhookTestSuite.T().Logf("Attempting to apply invalid LMDeployment (should be rejected)")
	cmd := exec.Command("kubectl", "apply", "-f", yamlFile)
	output, err := utils.Run(cmd)

	// The creation should be rejected by the webhook
	require.Error(webhookTestSuite.T(), err, "Invalid configuration should have been rejected")
	require.Contains(webhookTestSuite.T(), output, expectedError, "Error message should contain expected validation details")

	webhookTestSuite.T().Logf("Webhook correctly rejected invalid configuration: %s", name)

	// Clean up temporary file
	_ = os.Remove(yamlFile)
}

// verifyWebhookDefaults verifies that the webhook correctly applied default values
func (webhookTestSuite *WebhookTestSuite) verifyWebhookDefaults(deploymentName string) {
	webhookTestSuite.T().Log("Verifying webhook defaults")

	// Check that Ollama defaults were applied
	webhookTestSuite.T().Log("Checking Ollama defaults")

	// Check replicas default
	cmd := exec.Command("kubectl", "get", "lmdeployment", deploymentName,
		"-n", webhookTestSuite.testNamespace, "-o", "jsonpath={.spec.ollama.replicas}")
	output, err := utils.Run(cmd)
	require.NoError(webhookTestSuite.T(), err, "Failed to get Ollama replicas")
	require.Equal(webhookTestSuite.T(), "1", output, "Ollama replicas should default to 1")

	// Check image default
	cmd = exec.Command("kubectl", "get", "lmdeployment", deploymentName,
		"-n", webhookTestSuite.testNamespace, "-o", "jsonpath={.spec.ollama.image}")
	output, err = utils.Run(cmd)
	require.NoError(webhookTestSuite.T(), err, "Failed to get Ollama image")
	require.Equal(webhookTestSuite.T(), "ollama/ollama:latest", output, "Ollama image should default to ollama/ollama:latest")

	// Check service type default
	cmd = exec.Command("kubectl", "get", "lmdeployment", deploymentName,
		"-n", webhookTestSuite.testNamespace, "-o", "jsonpath={.spec.ollama.service.type}")
	output, err = utils.Run(cmd)
	require.NoError(webhookTestSuite.T(), err, "Failed to get Ollama service type")
	require.Equal(webhookTestSuite.T(), "ClusterIP", output, "Ollama service type should default to ClusterIP")

	// Check service port default
	cmd = exec.Command("kubectl", "get", "lmdeployment", deploymentName,
		"-n", webhookTestSuite.testNamespace, "-o", "jsonpath={.spec.ollama.service.port}")
	output, err = utils.Run(cmd)
	require.NoError(webhookTestSuite.T(), err, "Failed to get Ollama service port")
	require.Equal(webhookTestSuite.T(), "11434", output, "Ollama service port should default to 11434")

	webhookTestSuite.T().Log("All webhook defaults verified successfully")
}

// waitForLMDeploymentReady waits for an LMDeployment to be ready
func (webhookTestSuite *WebhookTestSuite) waitForLMDeploymentReady(name string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		cmd := exec.Command("kubectl", "get", "lmdeployment", name,
			"-n", webhookTestSuite.testNamespace, "-o", "jsonpath={.status.phase}")
		output, err := utils.Run(cmd)
		if err == nil && output == "Ready" {
			webhookTestSuite.T().Logf("LMDeployment %s is ready", name)
			return
		}
		time.Sleep(5 * time.Second)
	}
	webhookTestSuite.T().Fatalf("LMDeployment %s did not become ready within %v", name, timeout)
}

// TestWebhookLMDeploymentSuite runs the webhook test suite
func TestWebhookLMDeploymentSuite(t *testing.T) {
	suite.Run(t, &WebhookTestSuite{
		GlobalE2ESuite{testNamespace: "lmdeployment-webhook-test"},
	})
}
