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
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/geeper-io/llm-operator/test/utils"
)

var (
	// Optional Environment Variables:
	// - CERT_MANAGER_INSTALL_SKIP=true: Skips CertManager installation during test setup.
	// These variables are useful if CertManager is already installed, avoiding
	// re-installation and conflicts.
	skipCertManagerInstall = os.Getenv("CERT_MANAGER_INSTALL_SKIP") == "true"
	// isCertManagerAlreadyInstalled will be set true when CertManager CRDs be found on the cluster
	isCertManagerAlreadyInstalled = false

	// projectImage is the name of the image which will be build and loaded
	// with the code source changes to be tested.
	projectImage = "example.com/llm-operator:v0.0.1"
)

// GlobalE2ESuite handles global setup and teardown for all e2e tests
type GlobalE2ESuite struct {
	suite.Suite
	testNamespace string
}

// SetupSuite runs once before all test suites
func (suite *GlobalE2ESuite) SetupSuite() {
	suite.T().Log("Starting llm-operator integration test suite")

	suite.T().Log("Building the manager(Operator) image")
	cmd := exec.Command("make", "docker-build", fmt.Sprintf("IMG=%s", projectImage))
	//cmd := exec.Command("make", "docker-build")
	_, err := utils.Run(cmd)
	suite.Require().NoError(err, "Failed to build the manager(Operator) image")

	// TODO(user): If you want to change the e2e test vendor from Kind, ensure the image is
	// built and available before running the tests. Also, remove the following block.
	suite.T().Log("Loading the manager(Operator) image on Kind")
	err = utils.LoadImageToKindClusterWithName(projectImage)
	suite.Require().NoError(err, "Failed to load the manager(Operator) image into Kind")

	cmd = exec.Command("make", "install.yaml", fmt.Sprintf("IMG=%s", projectImage))
	_, err = utils.Run(cmd)
	suite.Require().NoError(err, "Failed to build the install.yaml file")

	cmd = exec.Command("kubectl", "apply", "-f", "install.yaml")
	_, err = utils.Run(cmd)
	suite.Require().NoError(err, "Failed to apply the install.yaml file")

	// The tests-e2e are intended to run on a temporary cluster that is created and destroyed for testing.
	// To prevent errors when tests run in environments with CertManager already installed,
	// we check for its presence before execution.
	// Setup CertManager before the suite if not skipped and if not already installed
	if !skipCertManagerInstall {
		suite.T().Log("Checking if cert manager is installed already")
		isCertManagerAlreadyInstalled = utils.IsCertManagerCRDsInstalled()
		if !isCertManagerAlreadyInstalled {
			suite.T().Log("Installing CertManager...")
			suite.Require().NoError(utils.InstallCertManager(), "Failed to install CertManager")
		} else {
			suite.T().Log("WARNING: CertManager is already installed. Skipping installation...")
		}
	}

	suite.T().Log("Creating test namespace")
	cmd = exec.Command("kubectl", "create", "ns", suite.testNamespace)
	_, err = utils.Run(cmd)
	require.NoError(suite.T(), err, "Failed to create test namespace")

	//suite.T().Log("Labeling the namespace to enforce the restricted security policy")
	//cmd = exec.Command("kubectl", "label", "--overwrite", "ns", suite.testNamespace,
	//	"pod-security.kubernetes.io/enforce=restricted")
	//_, err = utils.Run(cmd)
	//require.NoError(suite.T(), err, "Failed to label namespace with restricted policy")
}

// TearDownSuite runs once after all test suites
func (suite *GlobalE2ESuite) TearDownSuite() {
	// Teardown CertManager after the suite if not skipped and if it was not already installed
	if !skipCertManagerInstall && !isCertManagerAlreadyInstalled {
		suite.T().Log("Uninstalling CertManager...")
		utils.UninstallCertManager()
	}
	suite.T().Log("Cleaning up test namespace")
	cmd := exec.Command("kubectl", "delete", "ns", suite.testNamespace)
	_, _ = utils.Run(cmd)
}

// TestE2E runs the end-to-end (e2e) test suite for the project. These tests execute in an isolated,
// temporary environment to validate project changes with the purposed to be used in CI jobs.
// The default setup requires Kind, builds/loads the Manager Docker image locally, and installs
// CertManager.
func TestE2E(t *testing.T) {
	suite.Run(t, new(GlobalE2ESuite))
}
