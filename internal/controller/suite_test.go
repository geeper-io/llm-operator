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
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

// These tests use standard Go testing with Testify for assertions.

var (
	ctx       context.Context
	cancel    context.CancelFunc
	testEnv   *envtest.Environment
	cfg       *rest.Config
	k8sClient client.Client
)

func TestMain(m *testing.M) {
	// Set up test environment
	setupTestEnv()

	// Run tests
	code := m.Run()

	// Clean up test environment
	teardownTestEnv()

	os.Exit(code)
}

func setupTestEnv() {
	logf.SetLogger(zap.New(zap.WriteTo(os.Stdout), zap.UseDevMode(true)))

	ctx, cancel = context.WithCancel(context.TODO())

	var err error
	err = llmgeeperiov1alpha1.AddToScheme(scheme.Scheme)
	require.NoError(nil, err)

	// Bootstrap test environment
	testEnv = &envtest.Environment{
		CRDDirectoryPaths:     []string{filepath.Join("..", "..", "config", "crd", "bases")},
		ErrorIfCRDPathMissing: true,
	}

	// Retrieve the first found binary directory to allow running tests from IDEs
	if getFirstFoundEnvTestBinaryDir() != "" {
		testEnv.BinaryAssetsDirectory = getFirstFoundEnvTestBinaryDir()
	}

	// cfg is defined in this file globally.
	cfg, err = testEnv.Start()
	require.NoError(nil, err)
	require.NotNil(nil, cfg)

	k8sClient, err = client.New(cfg, client.Options{Scheme: scheme.Scheme})
	require.NoError(nil, err)
	require.NotNil(nil, k8sClient)
}

func teardownTestEnv() {
	if cancel != nil {
		cancel()
	}
	if testEnv != nil {
		err := testEnv.Stop()
		require.NoError(nil, err)
	}
}

// getFirstFoundEnvTestBinaryDir locates the first binary in the specified path.
// ENVTEST-based tests depend on specific binaries, usually located in paths set by
// controller-runtime. When running tests directly (e.g., via an IDE) without using
// Makefile targets, the 'BinaryAssetsDirectory' must be explicitly configured.
//
// This function streamlines the process by finding the required binaries, similar to
// how the Makefile target works.
func getFirstFoundEnvTestBinaryDir() string {
	// Define the paths where the binaries might be located
	possiblePaths := []string{
		filepath.Join("..", "..", "hack", "tools"),
		filepath.Join("..", "..", "bin"),
		filepath.Join("..", "..", "..", "bin"),
	}

	// Check each path for the required binaries
	for _, path := range possiblePaths {
		if _, err := os.Stat(filepath.Join(path, "etcd")); err == nil {
			if _, err := os.Stat(filepath.Join(path, "kube-apiserver")); err == nil {
				return path
			}
		}
	}

	return ""
}
