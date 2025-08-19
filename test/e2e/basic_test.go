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

// BasicLMDeploymentTestSuite defines the test suite for basic LMDeployment e2e tests
type BasicLMDeploymentTestSuite struct {
	GlobalE2ESuite
}

// TestBasicLMDeployment tests basic Ollama deployment
func (suite *BasicLMDeploymentTestSuite) TestBasicLMDeployment() {
	deploymentName := "test-basic-ollama"

	suite.T().Cleanup(func() {
		suite.T().Log("Cleaning up LMDeployment")
		cmd := exec.Command("kubectl", "delete", "lmdeployment", deploymentName, "-n", suite.testNamespace)
		_, _ = utils.Run(cmd)
	})

	suite.T().Log("Creating basic LMDeployment YAML")
	basicDeployment := `apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: test-basic-ollama
  namespace: lmdeployment-basic-test
spec:
  ollama:
    replicas: 1
    image: ollama/ollama:latest
    models:
    - "gemma3:270m"
    service:
      type: ClusterIP
      port: 11434`

	// Write YAML to temporary file
	yamlFile := filepath.Join("/tmp", "test-basic-ollama.yaml")
	err := os.WriteFile(yamlFile, []byte(basicDeployment), 0644)
	require.NoError(suite.T(), err)

	suite.T().Log("Applying LMDeployment YAML")
	cmd := exec.Command("kubectl", "apply", "-f", yamlFile)
	_, err = utils.Run(cmd)
	require.NoError(suite.T(), err, "Failed to apply LMDeployment")

	suite.T().Log("Waiting for LMDeployment to be ready")
	suite.waitForLMDeploymentReady(deploymentName, 5*time.Minute)

	suite.T().Log("Verifying Ollama deployment is created and running")
	suite.waitForDeploymentReady("test-basic-ollama-ollama", 3*time.Minute)

	suite.T().Log("Verifying Ollama service is created")
	cmd = exec.Command("kubectl", "get", "service",
		"test-basic-ollama-ollama", "-n", suite.testNamespace)
	_, err = utils.Run(cmd)
	require.NoError(suite.T(), err, "Ollama service not found")

	suite.T().Log("Verifying Ollama pod is running")
	suite.waitForPodRunning("app=ollama,llm-deployment=test-basic-ollama", 3*time.Minute)

	// Clean up temporary file
	os.Remove(yamlFile)
}

// Helper methods for the basic test suite

func (suite *BasicLMDeploymentTestSuite) waitForLMDeploymentReady(name string, timeout time.Duration) {
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

func (suite *BasicLMDeploymentTestSuite) waitForDeploymentReady(name string, timeout time.Duration) {
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

func (suite *BasicLMDeploymentTestSuite) waitForPodRunning(labels string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		cmd := exec.Command("kubectl", "get", "pods", "-l", labels,
			"-n", suite.testNamespace, "-o", "jsonpath={.items[0].status.phase}")
		output, err := utils.Run(cmd)
		if err == nil && output == "Running" {
			return
		}
		time.Sleep(10 * time.Second)
	}
	suite.T().Fatalf("Pod with labels %s not running within %v", labels, timeout)
}

// TestBasicLMDeploymentSuite runs the basic test suite
func TestBasicLMDeploymentSuite(t *testing.T) {
	suite.Run(t, &BasicLMDeploymentTestSuite{
		GlobalE2ESuite{testNamespace: "lmdeployment-basic-test"},
	})
}
