/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the Apache License is distributed on an "AS IS" BASIS,
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

// TabbyLMDeploymentTestSuite defines the test suite for Tabby LMDeployment e2e tests
type TabbyLMDeploymentTestSuite struct {
	GlobalE2ESuite
}

// TestTabbyLMDeployment tests Tabby deployment
func (suite *TabbyLMDeploymentTestSuite) TestTabbyLMDeployment() {
	deploymentName := "test-tabby"

	suite.T().Cleanup(func() {
		suite.T().Log("Cleaning up LMDeployment")
		cmd := exec.Command("kubectl", "delete", "lmdeployment", deploymentName, "-n", suite.testNamespace)
		_, _ = utils.Run(cmd)
	})

	suite.T().Log("Creating Tabby LMDeployment YAML")
	tabbyDeployment := `apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: test-tabby
  namespace: lmdeployment-tabby-test
spec:
  ollama:
    replicas: 1
    image: ollama/ollama:latest
    models:
    - "gemma3:270m"
    service:
      type: ClusterIP
      port: 11434
  
  tabby:
    enabled: true
    replicas: 1
    image: tabbyml/tabby:latest
    service:
      type: ClusterIP
      port: 8080
    modelName: gemma3:270m`

	// Write YAML to temporary file
	yamlFile := filepath.Join("/tmp", "test-tabby.yaml")
	err := os.WriteFile(yamlFile, []byte(tabbyDeployment), 0644)
	require.NoError(suite.T(), err)

	suite.T().Log("Applying Tabby LMDeployment YAML")
	cmd := exec.Command("kubectl", "apply", "-f", yamlFile)
	_, err = utils.Run(cmd)
	require.NoError(suite.T(), err, "Failed to apply Tabby LMDeployment")

	suite.T().Log("Waiting for LMDeployment to be ready")
	suite.waitForLMDeploymentReady(deploymentName, 8*time.Minute)

	suite.T().Log("Verifying Ollama deployment is running")
	suite.waitForDeploymentReady("test-tabby-ollama", 3*time.Minute)

	suite.T().Log("Verifying Tabby deployment is running")
	suite.waitForDeploymentReady("test-tabby-tabby", 5*time.Minute)

	suite.T().Log("Verifying all services are created")
	services := []string{
		"test-tabby-ollama",
		"test-tabby-tabby",
	}
	for _, serviceName := range services {
		cmd := exec.Command("kubectl", "get", "service", serviceName, "-n", suite.testNamespace)
		_, err := utils.Run(cmd)
		require.NoError(suite.T(), err, "Service "+serviceName+" not found")
	}

	// Clean up temporary file
	os.Remove(yamlFile)
}

// Helper methods for the Tabby test suite

func (suite *TabbyLMDeploymentTestSuite) waitForLMDeploymentReady(name string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		cmd := exec.Command("kubectl", "get", "lmdeployment", name,
			"-n", suite.testNamespace, "-o", "jsonpath={.status.phase}")
		output, err := utils.Run(cmd)
		if err == nil && output == "Ready" {
			return
		}
		time.Sleep(10 * time.Second)
	}
	suite.T().Fatalf("LMDeployment %s not ready within %v", name, timeout)
}

func (suite *TabbyLMDeploymentTestSuite) waitForDeploymentReady(name string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		cmd := exec.Command("kubectl", "get", "deployment", name,
			"-n", suite.testNamespace, "-o", "jsonpath={.status.readyReplicas}")
		output, err := utils.Run(cmd)
		if err == nil && output == "1" {
			return
		}
		time.Sleep(10 * time.Second)
	}
	suite.T().Fatalf("Deployment %s not ready within %v", name, timeout)
}

// TestTabbyLMDeploymentSuite runs the Tabby test suite
func TestTabbyLMDeploymentSuite(t *testing.T) {
	suite.Run(t, &TabbyLMDeploymentTestSuite{
		GlobalE2ESuite{testNamespace: "lmdeployment-tabby-test"},
	})
}
