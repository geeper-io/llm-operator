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
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

func TestDeploymentController_Reconciliation(t *testing.T) {
	const resourceName = "test-resource"

	ctx := context.Background()

	typeNamespacedName := types.NamespacedName{
		Name:      resourceName,
		Namespace: "default",
	}

	t.Run("should successfully reconcile the resource", func(t *testing.T) {
		deployment := &llmgeeperiov1alpha1.Deployment{}

		// Create the custom resource for the Kind Deployment
		err := k8sClient.Get(ctx, typeNamespacedName, deployment)
		if err != nil && errors.IsNotFound(err) {
			resource := &llmgeeperiov1alpha1.Deployment{
				ObjectMeta: metav1.ObjectMeta{
					Name:      resourceName,
					Namespace: "default",
				},
				Spec: llmgeeperiov1alpha1.DeploymentSpec{
					Ollama: llmgeeperiov1alpha1.OllamaSpec{
						Models: []llmgeeperiov1alpha1.OllamaModel{
							"llama2:7b",
						},
					},
				},
			}
			require.NoError(t, k8sClient.Create(ctx, resource))
		}

		// Cleanup after test
		t.Cleanup(func() {
			resource := &llmgeeperiov1alpha1.Deployment{}
			err2 := k8sClient.Get(ctx, typeNamespacedName, resource)
			require.NoError(t, err2)

			// Cleanup the specific resource instance Deployment
			require.NoError(t, k8sClient.Delete(ctx, resource))
		})

		// Reconcile the created resource
		controllerReconciler := &OllamaDeploymentReconciler{
			Client: k8sClient,
			Scheme: k8sClient.Scheme(),
		}

		_, reconcileErr := controllerReconciler.Reconcile(ctx, reconcile.Request{
			NamespacedName: typeNamespacedName,
		})
		require.NoError(t, reconcileErr)
		// TODO(user): Add more specific assertions depending on your controller's reconciliation logic.
		// Example: If you expect a certain status condition after reconciliation, verify it here.
	})
}
