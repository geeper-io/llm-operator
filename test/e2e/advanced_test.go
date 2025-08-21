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

// AdvancedLMDeploymentTestSuite defines the test suite for advanced LMDeployment e2e tests
type AdvancedLMDeploymentTestSuite struct {
	GlobalE2ESuite
}

// TestAdvancedLMDeployment tests advanced deployment with all components
func (suite *AdvancedLMDeploymentTestSuite) TestAdvancedLMDeployment() {
	deploymentName := "test-advanced-all"

	suite.T().Cleanup(func() {
		suite.T().Log("Cleaning up LMDeployment")
		cmd := exec.Command("kubectl", "delete", "lmdeployment", deploymentName, "-n", suite.testNamespace)
		_, _ = utils.Run(cmd)
	})

	suite.T().Log("Creating advanced LMDeployment YAML")
	advancedDeployment := `apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: test-advanced-all
  namespace: lmdeployment-advanced-test
spec:
  ollama:
    replicas: 1
    image: ollama/ollama:latest
    models:
    - "gemma3:270m"
    service:
      type: ClusterIP
      port: 11434
  
  openwebui:
    enabled: true
    replicas: 1
    image: ghcr.io/open-webui/open-webui:main
    service:
      type: ClusterIP
      port: 8080
    redis:
      enabled: true
      image: redis:7-alpine
      password: "test-redis-password"
      persistence:
        enabled: true
        size: 1Gi
    
    pipelines:
      enabled: true
      image: ghcr.io/open-webui/pipelines:main
      replicas: 1
      persistence:
        enabled: true
        size: 1Gi`

	// Write YAML to temporary file
	yamlFile := filepath.Join("/tmp", "test-advanced-all.yaml")
	err := os.WriteFile(yamlFile, []byte(advancedDeployment), 0644)
	require.NoError(suite.T(), err)

	suite.T().Log("Applying advanced LMDeployment YAML")
	cmd := exec.Command("kubectl", "apply", "-f", yamlFile)
	_, err = utils.Run(cmd)
	require.NoError(suite.T(), err, "Failed to apply advanced LMDeployment")

	suite.T().Log("Waiting for LMDeployment to be ready")
	suite.waitForLMDeploymentReady(deploymentName, 10*time.Minute)

	suite.T().Log("Verifying all deployments are running")
	deployments := []string{
		"test-advanced-all-ollama",
		"test-advanced-all-openwebui",
		"test-advanced-all-redis",
		"test-advanced-all-pipelines",
	}
	for _, deploymentName := range deployments {
		suite.waitForDeploymentReady(deploymentName, 5*time.Minute)
	}

	suite.T().Log("Verifying all services are created")
	services := []string{
		"test-advanced-all-ollama",
		"test-advanced-all-openwebui",
		"test-advanced-all-redis",
		"test-advanced-all-pipelines",
	}
	for _, serviceName := range services {
		cmd := exec.Command("kubectl", "get", "service", serviceName, "-n", suite.testNamespace)
		_, err := utils.Run(cmd)
		require.NoError(suite.T(), err, "Service "+serviceName+" not found")
	}

	suite.T().Log("Verifying all PVCs are created")
	pvcs := []string{
		"test-advanced-all-redis",
		"test-advanced-all-pipelines-data",
	}
	for _, pvcName := range pvcs {
		cmd := exec.Command("kubectl", "get", "pvc", pvcName, "-n", suite.testNamespace)
		_, err := utils.Run(cmd)
		require.NoError(suite.T(), err, "PVC "+pvcName+" not found")
	}

	// Clean up temporary file
	os.Remove(yamlFile)
}

// Helper methods for the advanced test suite

func (suite *AdvancedLMDeploymentTestSuite) waitForLMDeploymentReady(name string, timeout time.Duration) {
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

func (suite *AdvancedLMDeploymentTestSuite) waitForDeploymentReady(name string, timeout time.Duration) {
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

// TestAdvancedLMDeploymentSuite runs the advanced test suite
func TestAdvancedLMDeploymentSuite(t *testing.T) {
	suite.Run(t, &AdvancedLMDeploymentTestSuite{
		GlobalE2ESuite{testNamespace: "lmdeployment-advanced-test"},
	})
}
