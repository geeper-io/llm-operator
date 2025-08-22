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

// TestLMDeploymentStatusAndMetrics tests status and metrics reporting
func (statusTestSuite *StatusLMDeploymentTestSuite) TestLMDeploymentStatusAndMetrics() {
	deploymentName := "test-status-metrics"

	statusTestSuite.T().Cleanup(func() {
		statusTestSuite.T().Log("Cleaning up LMDeployment")
		cmd := exec.Command("kubectl", "delete", "lmdeployment", deploymentName, "-n", statusTestSuite.testNamespace)
		_, _ = utils.Run(cmd)
	})

	statusTestSuite.T().Log("Creating LMDeployment for status testing")
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
    - "hf.co/amakhov/tiny-random-llama:F16"
  
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
	require.NoError(statusTestSuite.T(), err)

	statusTestSuite.T().Log("Applying LMDeployment YAML")
	cmd := exec.Command("kubectl", "apply", "-f", yamlFile)
	_, err = utils.Run(cmd)
	require.NoError(statusTestSuite.T(), err, "Failed to apply LMDeployment")

	statusTestSuite.T().Log("Waiting for LMDeployment to be ready")
	statusTestSuite.waitForLMDeploymentReady(deploymentName, 8*time.Minute)

	statusTestSuite.T().Log("Verifying total replicas count")
	statusTestSuite.waitForStatusField(deploymentName, "totalReplicas", "4", 2*time.Minute)

	statusTestSuite.T().Log("Verifying ready replicas count")
	statusTestSuite.waitForStatusField(deploymentName, "readyReplicas", "4", 2*time.Minute)

	statusTestSuite.T().Log("Verifying component statuses")
	statusTestSuite.waitForStatusField(deploymentName, "ollamaStatus.readyReplicas", "2", 2*time.Minute)
	statusTestSuite.waitForStatusField(deploymentName, "openwebuiStatus.readyReplicas", "2", 2*time.Minute)

	// Clean up temporary file
	_ = os.Remove(yamlFile)
}

// Helper methods for the status test suite

func (statusTestSuite *StatusLMDeploymentTestSuite) waitForLMDeploymentReady(name string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		cmd := exec.Command("kubectl", "get", "lmdeployment", name,
			"-n", statusTestSuite.testNamespace, "-o", "jsonpath={.status.phase}")
		output, err := utils.Run(cmd)
		if err == nil && output == "Ready" {
			return
		}
		time.Sleep(10 * time.Second)
	}
	statusTestSuite.T().Fatalf("LMDeployment %s not ready within %v", name, timeout)
}

func (statusTestSuite *StatusLMDeploymentTestSuite) waitForStatusField(name, field, expectedValue string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		cmd := exec.Command("kubectl", "get", "lmdeployment", name,
			"-n", statusTestSuite.testNamespace, "-o", "jsonpath={.status."+field+"}")
		output, err := utils.Run(cmd)
		if err == nil && output == expectedValue {
			return
		}
		time.Sleep(10 * time.Second)
	}
	statusTestSuite.T().Fatalf("Status field %s for LMDeployment %s not equal to %s within %v", field, name, expectedValue, timeout)
}

// TestStatusLMDeploymentSuite runs the status test suite
func TestStatusLMDeploymentSuite(t *testing.T) {
	suite.Run(t, &StatusLMDeploymentTestSuite{
		GlobalE2ESuite{testNamespace: "lmdeployment-status-test"},
	})
}
