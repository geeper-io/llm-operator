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

// OpenWebUILMDeploymentTestSuite defines the test suite for OpenWebUI LMDeployment e2e tests
type OpenWebUILMDeploymentTestSuite struct {
	GlobalE2ESuite
}

// TestOpenWebUILMDeployment tests OpenWebUI with Redis deployment
func (owuiTestSuite *OpenWebUILMDeploymentTestSuite) TestOpenWebUILMDeployment() {
	deploymentName := "test-openwebui-redis"

	owuiTestSuite.T().Cleanup(func() {
		owuiTestSuite.T().Log("Cleaning up LMDeployment")
		cmd := exec.Command("kubectl", "delete", "lmdeployment", deploymentName, "-n", owuiTestSuite.testNamespace)
		_, _ = utils.Run(cmd)
	})

	owuiTestSuite.T().Log("Creating OpenWebUI LMDeployment YAML")
	openwebuiDeployment := `apiVersion: llm.geeper.io/v1alpha1
kind: LMDeployment
metadata:
  name: test-openwebui-redis
  namespace: lmdeployment-openwebui-test
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
    redis:
      enabled: true
      image: redis:7-alpine
      password: "test-redis-password"
      service:
        port: 6379
      persistence:
        enabled: true
        size: 1Gi`

	// Write YAML to temporary file
	yamlFile := filepath.Join("/tmp", "test-openwebui-redis.yaml")
	err := os.WriteFile(yamlFile, []byte(openwebuiDeployment), 0644)
	require.NoError(owuiTestSuite.T(), err)

	owuiTestSuite.T().Log("Applying OpenWebUI LMDeployment YAML")
	cmd := exec.Command("kubectl", "apply", "-f", yamlFile)
	_, err = utils.Run(cmd)
	require.NoError(owuiTestSuite.T(), err, "Failed to apply OpenWebUI LMDeployment")

	owuiTestSuite.T().Log("Waiting for LMDeployment to be ready")
	owuiTestSuite.waitForLMDeploymentReady(deploymentName, 8*time.Minute)

	owuiTestSuite.T().Log("Verifying Ollama deployment is running")
	owuiTestSuite.waitForDeploymentReady("test-openwebui-redis-ollama", 3*time.Minute)

	owuiTestSuite.T().Log("Verifying OpenWebUI deployment is running")
	owuiTestSuite.waitForDeploymentReady("test-openwebui-redis-openwebui", 5*time.Minute)

	owuiTestSuite.T().Log("Verifying Redis deployment is running")
	owuiTestSuite.waitForDeploymentReady("test-openwebui-redis-redis", 3*time.Minute)

	owuiTestSuite.T().Log("Verifying all services are created")
	services := []string{
		"test-openwebui-redis-ollama",
		"test-openwebui-redis-openwebui",
		"test-openwebui-redis-redis",
	}
	for _, serviceName := range services {
		cmd := exec.Command("kubectl", "get", "service", serviceName, "-n", owuiTestSuite.testNamespace)
		_, err := utils.Run(cmd)
		require.NoError(owuiTestSuite.T(), err, "Service "+serviceName+" not found")
	}

	owuiTestSuite.T().Log("Verifying OpenWebUI secret is created")
	owuiTestSuite.waitForSecretExists("test-openwebui-redis-openwebui-secret", 2*time.Minute)

	owuiTestSuite.T().Log("Verifying Redis PVC is created")
	cmd = exec.Command("kubectl", "get", "pvc",
		"test-openwebui-redis-redis", "-n", owuiTestSuite.testNamespace)
	_, pvcErr := utils.Run(cmd)
	require.NoError(owuiTestSuite.T(), pvcErr, "Redis PVC not found")

	// Clean up temporary file
	_ = os.Remove(yamlFile)
}

// Helper methods for the OpenWebUI test suite

func (owuiTestSuite *OpenWebUILMDeploymentTestSuite) waitForLMDeploymentReady(name string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		cmd := exec.Command("kubectl", "get", "lmdeployment", name,
			"-n", owuiTestSuite.testNamespace, "-o", "jsonpath={.status.phase}")
		output, err := utils.Run(cmd)
		if err == nil && output == "Ready" {
			return
		}
		time.Sleep(10 * time.Second)
	}
	owuiTestSuite.T().Fatalf("LMDeployment %s not ready within %v", name, timeout)
}

func (owuiTestSuite *OpenWebUILMDeploymentTestSuite) waitForDeploymentReady(name string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		cmd := exec.Command("kubectl", "get", "deployment", name,
			"-n", owuiTestSuite.testNamespace, "-o", "jsonpath={.status.readyReplicas}")
		output, err := utils.Run(cmd)
		if err == nil && output == "1" {
			return
		}
		time.Sleep(10 * time.Second)
	}
	owuiTestSuite.T().Fatalf("Deployment %s not ready within %v", name, timeout)
}

func (owuiTestSuite *OpenWebUILMDeploymentTestSuite) waitForSecretExists(name string, timeout time.Duration) {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		cmd := exec.Command("kubectl", "get", "secret", name, "-n", owuiTestSuite.testNamespace)
		_, err := utils.Run(cmd)
		if err == nil {
			return
		}
		time.Sleep(10 * time.Second)
	}
	owuiTestSuite.T().Fatalf("Secret %s not found within %v", name, timeout)
}

// TestOpenWebUILMDeploymentSuite runs the OpenWebUI test suite
func TestOpenWebUILMDeploymentSuite(t *testing.T) {
	suite.Run(t, &OpenWebUILMDeploymentTestSuite{
		GlobalE2ESuite{testNamespace: "lmdeployment-openwebui-test"},
	})
}
