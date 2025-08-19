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

package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	prometheusOperatorVersion = "v0.77.1"
	prometheusOperatorURL     = "https://github.com/prometheus-operator/prometheus-operator/" +
		"releases/download/%s/bundle.yaml"

	certmanagerVersion = "v1.16.3"
	certmanagerURLTmpl = "https://github.com/cert-manager/cert-manager/releases/download/%s/cert-manager.yaml"
)

func warnError(err error) {
	fmt.Printf("warning: %v\n", err)
}

// Run executes the provided command within this context
func Run(cmd *exec.Cmd) (string, error) {
	dir, _ := GetProjectDir()
	cmd.Dir = dir

	if err := os.Chdir(cmd.Dir); err != nil {
		fmt.Printf("chdir dir: %q\n", err)
	}

	cmd.Env = append(os.Environ(), "GO111MODULE=on")
	command := strings.Join(cmd.Args, " ")
	fmt.Printf("running: %q\n", command)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return string(output), fmt.Errorf("%q failed with error %q: %w", command, string(output), err)
	}

	return string(output), nil
}

// InstallPrometheusOperator installs the prometheus Operator to be used to export the enabled metrics.
func InstallPrometheusOperator() error {
	url := fmt.Sprintf(prometheusOperatorURL, prometheusOperatorVersion)
	cmd := exec.Command("kubectl", "create", "-f", url)
	_, err := Run(cmd)
	return err
}

// UninstallPrometheusOperator uninstalls the prometheus
func UninstallPrometheusOperator() {
	url := fmt.Sprintf(prometheusOperatorURL, prometheusOperatorVersion)
	cmd := exec.Command("kubectl", "delete", "-f", url)
	if _, err := Run(cmd); err != nil {
		warnError(err)
	}
}

// IsPrometheusCRDsInstalled checks if any Prometheus CRDs are installed
// by verifying the existence of key CRDs related to Prometheus.
func IsPrometheusCRDsInstalled() bool {
	// List of common Prometheus CRDs
	prometheusCRDs := []string{
		"prometheuses.monitoring.coreos.com",
		"prometheusrules.monitoring.coreos.com",
		"prometheusagents.monitoring.coreos.com",
	}

	cmd := exec.Command("kubectl", "get", "crds", "-o", "custom-columns=NAME:.metadata.name")
	output, err := Run(cmd)
	if err != nil {
		return false
	}
	crdList := GetNonEmptyLines(output)
	for _, crd := range prometheusCRDs {
		for _, line := range crdList {
			if strings.Contains(line, crd) {
				return true
			}
		}
	}

	return false
}

// InstallCertManager installs the cert-manager to be used to create the certificates.
func InstallCertManager() error {
	url := fmt.Sprintf(certmanagerURLTmpl, certmanagerVersion)
	cmd := exec.Command("kubectl", "create", "-f", url)
	_, err := Run(cmd)
	return err
}

// UninstallCertManager uninstalls the cert-manager
func UninstallCertManager() {
	url := fmt.Sprintf(certmanagerURLTmpl, certmanagerVersion)
	cmd := exec.Command("kubectl", "delete", "-f", url)
	if _, err := Run(cmd); err != nil {
		warnError(err)
	}
}

// IsCertManagerCRDsInstalled checks if any CertManager CRDs are installed
// by verifying the existence of key CRDs related to CertManager.
func IsCertManagerCRDsInstalled() bool {
	// List of common CertManager CRDs
	certManagerCRDs := []string{
		"certificates.cert-manager.io",
		"issuers.cert-manager.io",
		"clusterissuers.cert-manager.io",
	}

	cmd := exec.Command("kubectl", "get", "crds", "-o", "custom-columns=NAME:.metadata.name")
	output, err := Run(cmd)
	if err != nil {
		return false
	}
	crdList := GetNonEmptyLines(output)
	for _, crd := range certManagerCRDs {
		for _, line := range crdList {
			if strings.Contains(line, crd) {
				return true
			}
		}
	}

	return false
}

// LoadImageToKindClusterWithName loads a local docker image to the kind cluster
func LoadImageToKindClusterWithName(name string) error {
	cluster := "llm-operator-test-e2e"
	if v, ok := os.LookupEnv("KIND_CLUSTER"); ok {
		cluster = v
	}
	kindOptions := []string{"load", "docker-image", name, "--name", cluster}
	cmd := exec.Command("kind", kindOptions...)
	_, err := Run(cmd)
	return err
}

// GetNonEmptyLines converts given command output string into individual objects
// according to line breakers, and ignores the empty elements in it.
func GetNonEmptyLines(output string) []string {
	var res []string
	elements := strings.Split(output, "\n")
	for _, element := range elements {
		if element != "" {
			res = append(res, element)
		}
	}

	return res
}

// GetProjectDir will return the directory where the project is
func GetProjectDir() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return wd, fmt.Errorf("failed to get current working directory: %w", err)
	}
	wd = strings.ReplaceAll(wd, "/test/e2e", "")
	return wd, nil
}

// UncommentCode searches for target in the file and remove the comment prefix
// of the target content. The target content may span multiple lines.
func UncommentCode(filename, target, prefix string) error {
	// false positive
	// nolint:gosec
	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("failed to read file %q: %w", filename, err)
	}
	strContent := string(content)

	idx := strings.Index(strContent, target)
	if idx < 0 {
		return fmt.Errorf("unable to find the code %q to be uncomment", target)
	}

	out := new(bytes.Buffer)
	_, err = out.Write(content[:idx])
	if err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	scanner := bufio.NewScanner(bytes.NewBufferString(target))
	if !scanner.Scan() {
		return nil
	}
	for {
		if _, err = out.WriteString(strings.TrimPrefix(scanner.Text(), prefix)); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}
		// Avoid writing a newline in case the previous line was the last in target.
		if !scanner.Scan() {
			break
		}
		if _, err = out.WriteString("\n"); err != nil {
			return fmt.Errorf("failed to write to output: %w", err)
		}
	}

	if _, err = out.Write(content[idx+len(target):]); err != nil {
		return fmt.Errorf("failed to write to output: %w", err)
	}

	// false positive
	// nolint:gosec
	if err = os.WriteFile(filename, out.Bytes(), 0644); err != nil {
		return fmt.Errorf("failed to write file %q: %w", filename, err)
	}

	return nil
}

// WaitForLMDeploymentReady waits for an LMDeployment to be ready
func WaitForLMDeploymentReady(name, namespace string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		cmd := exec.Command("kubectl", "get", "lmdeployment", name,
			"-n", namespace, "-o", "jsonpath={.status.phase}")
		output, err := Run(cmd)
		if err == nil && strings.TrimSpace(output) == "Ready" {
			return nil
		}
		time.Sleep(10 * time.Second)
	}
	return fmt.Errorf("LMDeployment %s in namespace %s not ready within %v", name, namespace, timeout)
}

// GetLMDeploymentStatus returns the status of an LMDeployment
func GetLMDeploymentStatus(name, namespace string) (map[string]string, error) {
	cmd := exec.Command("kubectl", "get", "lmdeployment", name,
		"-n", namespace, "-o", "jsonpath={.status}")
	output, err := Run(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to get LMDeployment status: %w", err)
	}

	// Parse the status output into a map
	status := make(map[string]string)
	// This is a simplified parser - in production you might want to use proper JSON parsing
	lines := strings.Split(output, " ")
	for _, line := range lines {
		if strings.Contains(line, ":") {
			parts := strings.SplitN(line, ":", 2)
			if len(parts) == 2 {
				status[parts[0]] = parts[1]
			}
		}
	}

	return status, nil
}

// VerifyDeploymentReady checks if a Kubernetes deployment is ready
func VerifyDeploymentReady(name, namespace string) error {
	cmd := exec.Command("kubectl", "get", "deployment", name,
		"-n", namespace, "-o", "jsonpath={.status.readyReplicas}")
	output, err := Run(cmd)
	if err != nil {
		return fmt.Errorf("failed to get deployment %s status: %w", name, err)
	}

	if strings.TrimSpace(output) == "0" {
		return fmt.Errorf("deployment %s has 0 ready replicas", name)
	}

	return nil
}

// VerifyServiceExists checks if a Kubernetes service exists
func VerifyServiceExists(name, namespace string) error {
	cmd := exec.Command("kubectl", "get", "service", name, "-n", namespace)
	_, err := Run(cmd)
	if err != nil {
		return fmt.Errorf("service %s not found in namespace %s: %w", name, namespace, err)
	}
	return nil
}

// VerifyPVCExists checks if a Kubernetes PVC exists
func VerifyPVCExists(name, namespace string) error {
	cmd := exec.Command("kubectl", "get", "pvc", name, "-n", namespace)
	_, err := Run(cmd)
	if err != nil {
		return fmt.Errorf("PVC %s not found in namespace %s: %w", name, namespace, err)
	}
	return nil
}

// VerifySecretExists checks if a Kubernetes secret exists
func VerifySecretExists(name, namespace string) error {
	cmd := exec.Command("kubectl", "get", "secret", name, "-n", namespace)
	_, err := Run(cmd)
	if err != nil {
		return fmt.Errorf("secret %s not found in namespace %s: %w", name, namespace, err)
	}
	return nil
}

// GetPodStatus returns the status of pods with specific labels
func GetPodStatus(labels, namespace string) ([]string, error) {
	cmd := exec.Command("kubectl", "get", "pods", "-l", labels,
		"-n", namespace, "-o", "jsonpath={.items[*].status.phase}")
	output, err := Run(cmd)
	if err != nil {
		return nil, fmt.Errorf("failed to get pod status: %w", err)
	}

	if strings.TrimSpace(output) == "" {
		return []string{}, nil
	}

	return strings.Split(output, " "), nil
}

// CleanupLMDeployment removes an LMDeployment and waits for cleanup
func CleanupLMDeployment(name, namespace string) error {
	fmt.Printf("cleaning up LMDeployment %s\n", name)
	cmd := exec.Command("kubectl", "delete", "lmdeployment", name, "-n", namespace)
	_, err := Run(cmd)
	if err != nil {
		return fmt.Errorf("failed to delete LMDeployment: %w", err)
	}

	// Wait for the deployment to be fully deleted
	deadline := time.Now().Add(2 * time.Minute)
	for time.Now().Before(deadline) {
		cmd := exec.Command("kubectl", "get", "lmdeployment", name, "-n", namespace)
		_, err := Run(cmd)
		if err != nil {
			// Deployment is deleted
			return nil
		}
		time.Sleep(5 * time.Second)
	}

	return fmt.Errorf("LMDeployment %s not deleted within timeout", name)
}
