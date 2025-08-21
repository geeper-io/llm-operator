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

// StatusLMDeploymentTestSuite defines the test suite for status and metrics LMDeployment e2e tests
type StatusLMDeploymentTestSuite struct {
	GlobalE2ESuite
}

// SetupSuite runs once before all tests
func (suite *StatusLMDeploymentTestSuite) SetupSuite() {
	suite.testNamespace = "lmdeployment-status-test"

	suite.T().Log("Creating test namespace")
	cmd := exec.Command("kubectl", "create", "ns", suite.testNamespace)
	_, err := utils.Run(cmd)
	require.NoError(suite.T(), err, "Failed to create test namespace")

	suite.T().Log("Labeling the namespace to enforce the restricted security policy")
	cmd = exec.Command("kubectl", "label", "--overwrite", "ns", suite.testNamespace,
		"pod-security.kubernetes.io/enforce=restricted")
	_, err = utils.Run(cmd)
	require.NoError(suite.T(), err, "Failed to label namespace with restricted policy")
}

// TearDownSuite runs once after all tests
func (suite *StatusLMDeploymentTestSuite) TearDownSuite() {
	suite.T().Log("Cleaning up test namespace")
	cmd := exec.Command("kubectl", "delete", "ns", suite.testNamespace)
	_, _ = utils.Run(cmd)
}

// TestLMDeploymentStatusAndMetrics tests status and metrics reporting
func (suite *StatusLMDeploymentTestSuite) TestLMDeploymentStatusAndMetrics() {
	deploymentName := "test-status-metrics"

	suite.T().Cleanup(func() {
		suite.T().Log("Cleaning up LMDeployment")
		cmd := exec.Command("kubectl", "delete", "lmdeployment", deploymentName, "-n", suite.testNamespace)
		_, _ = utils.Run(cmd)
	})

	suite.T().Log("Creating LMDeployment for status testing")
	statusDeployment := `apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: test-status-metrics
  namespace: lmdeployment-status-test
spec:
  ollama:
    replicas: 2
    image: ollama/ollama:latest
    models:
    - "gemma3:270m"
  
  openwebui:
    enabled: true
    replicas: 2
    image: ghcr.io/open-webui/open-webui:main
    redis:
      enabled: true
      image: redis:7-alpine
      password: "test-redis-password"
      service:
        port: 6379`

	// Write YAML to temporary file
	yamlFile := filepath.Join("/tmp", "test-status-metrics.yaml")
	err := os.WriteFile(yamlFile, []byte(statusDeployment), 0644)
	require.NoError(suite.T(), err)

	suite.T().Log("Applying LMDeployment YAML")
	cmd := exec.Command("kubectl", "apply", "-f", yamlFile)
	_, err = utils.Run(cmd)
	require.NoError(suite.T(), err, "Failed to apply LMDeployment")

	suite.T().Log("Waiting for LMDeployment to be ready")
	suite.waitForLMDeploymentReady(deploymentName, 8*time.Minute)

	suite.T().Log("Verifying total replicas count")
	suite.waitForStatusField(deploymentName, "totalReplicas", "6", 2*time.Minute)

	suite.T().Log("Verifying ready replicas count")
	suite.waitForStatusField(deploymentName, "readyReplicas", "6", 2*time.Minute)

	suite.T().Log("Verifying component statuses")
	suite.waitForStatusField(deploymentName, "ollamaStatus.readyReplicas", "2", 2*time.Minute)
	suite.waitForStatusField(deploymentName, "openwebuiStatus.readyReplicas", "2", 2*time.Minute)

	// Clean up temporary file
	os.Remove(yamlFile)
}

// Helper methods for the status test suite

func (suite *StatusLMDeploymentTestSuite) waitForLMDeploymentReady(name string, timeout time.Duration) {
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

func (suite *StatusLMDeploymentTestSuite) waitForStatusField(name, field, expectedValue string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		cmd := exec.Command("kubectl", "get", "lmdeployment", name,
			"-n", suite.testNamespace, "-o", "jsonpath={.status."+field+"}")
		output, err := utils.Run(cmd)
		if err == nil && output == expectedValue {
			return
		}
		time.Sleep(10 * time.Second)
	}
	suite.T().Fatalf("Status field %s for LMDeployment %s not equal to %s within %v", field, name, expectedValue, timeout)
}

// TestStatusLMDeploymentSuite runs the status test suite
func TestStatusLMDeploymentSuite(t *testing.T) {
	suite.Run(t, &StatusLMDeploymentTestSuite{
		GlobalE2ESuite{testNamespace: "lmdeployment-status-test"},
	})
}
