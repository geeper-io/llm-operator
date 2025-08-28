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

package v1alpha1

import (
	"context"
	"fmt"
	"strings"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation"
	"k8s.io/apimachinery/pkg/util/validation/field"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"sigs.k8s.io/controller-runtime/pkg/webhook/admission"

	llmgeeperiov1alpha1 "github.com/geeper-io/llm-operator/api/v1alpha1"
)

// nolint:unused
// log is for logging in this package.
var lmdeploymentlog = logf.Log.WithName("lmdeployment-resource")

// SetupLMDeploymentWebhookWithManager registers the webhook for LMDeployment in the manager.
func SetupLMDeploymentWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).For(&llmgeeperiov1alpha1.LMDeployment{}).
		WithValidator(&LMDeploymentCustomValidator{}).
		WithDefaulter(&LMDeploymentCustomDefaulter{}).
		Complete()
}

// +kubebuilder:webhook:path=/mutate-llm-geeper-io-v1alpha1-lmdeployment,mutating=true,failurePolicy=fail,sideEffects=None,groups=llm.geeper.io,resources=lmdeployments,verbs=create;update,versions=v1alpha1,name=mlmdeployment-v1alpha1.kb.io,admissionReviewVersions=v1

// LMDeploymentCustomDefaulter struct is responsible for setting default values on the custom resource of the
// Kind LMDeployment when those are created or updated.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as it is used only for temporary operations and does not need to be deeply copied.
type LMDeploymentCustomDefaulter struct{}

var _ webhook.CustomDefaulter = &LMDeploymentCustomDefaulter{}

// Default implements webhook.CustomDefaulter so a webhook will be registered for the Kind LMDeployment.
func (d *LMDeploymentCustomDefaulter) Default(_ context.Context, obj runtime.Object) error {
	lmDeployment, ok := obj.(*llmgeeperiov1alpha1.LMDeployment)
	if !ok {
		return fmt.Errorf("expected a LMDeployment but got a %T", obj)
	}

	if len(lmDeployment.Spec.Ollama.Models) > 0 {
		lmDeployment.Spec.Ollama.Enabled = true
	}

	if lmDeployment.Spec.Ollama.Enabled {
		d.defaultOllama(lmDeployment)
	}

	if lmDeployment.Spec.VLLM.Enabled {
		d.defaultVLLM(lmDeployment)
	}

	d.defaultOpenWebUI(lmDeployment)
	d.defaultTabby(lmDeployment)

	return nil
}

func (d *LMDeploymentCustomDefaulter) defaultOllama(lmDeployment *llmgeeperiov1alpha1.LMDeployment) {
	if lmDeployment.Spec.Ollama.Image == "" {
		switch lmDeployment.Spec.Ollama.Flavor {
		case "amd":
			lmDeployment.Spec.Ollama.Image = "ollama/ollama:rocm"
		case "nvidia":
			fallthrough
		default:
			lmDeployment.Spec.Ollama.Image = "ollama/ollama:latest"
		}
	}
	if lmDeployment.Spec.Ollama.Replicas == 0 {
		lmDeployment.Spec.Ollama.Replicas = 1
	}
	if lmDeployment.Spec.Ollama.Service.Type == "" {
		lmDeployment.Spec.Ollama.Service.Type = corev1.ServiceTypeClusterIP
	}
	if lmDeployment.Spec.Ollama.Service.Port == 0 {
		lmDeployment.Spec.Ollama.Service.Port = 11434
	}
}

func (d *LMDeploymentCustomDefaulter) defaultVLLM(lmDeployment *llmgeeperiov1alpha1.LMDeployment) {
	// Set global defaults if global config is not specified
	if lmDeployment.Spec.VLLM.GlobalConfig == nil {
		lmDeployment.Spec.VLLM.GlobalConfig = &llmgeeperiov1alpha1.VLLMGlobalConfig{}
	}

	// Set global image default
	if lmDeployment.Spec.VLLM.GlobalConfig.Image == "" {
		lmDeployment.Spec.VLLM.GlobalConfig.Image = "vllm/vllm-openai:latest"
	}

	// Set global service defaults
	if lmDeployment.Spec.VLLM.GlobalConfig.Service.Type == "" {
		lmDeployment.Spec.VLLM.GlobalConfig.Service.Type = corev1.ServiceTypeClusterIP
	}
	if lmDeployment.Spec.VLLM.GlobalConfig.Service.Port == 0 {
		lmDeployment.Spec.VLLM.GlobalConfig.Service.Port = 8000
	}

	// Set global persistence defaults
	if lmDeployment.Spec.VLLM.GlobalConfig.Persistence != nil && lmDeployment.Spec.VLLM.GlobalConfig.Persistence.Enabled {
		if lmDeployment.Spec.VLLM.GlobalConfig.Persistence.Size == "" {
			lmDeployment.Spec.VLLM.GlobalConfig.Persistence.Size = "10Gi"
		}
	}

	// Set defaults for each model
	for i := range lmDeployment.Spec.VLLM.Models {
		modelSpec := &lmDeployment.Spec.VLLM.Models[i]

		// Set model image default (fall back to global)
		if modelSpec.Image == "" {
			modelSpec.Image = lmDeployment.Spec.VLLM.GlobalConfig.Image
		}

		// Set model replicas default
		if modelSpec.Replicas == 0 {
			modelSpec.Replicas = 1
		}

		// Set model service defaults (fall back to global)
		if modelSpec.Service.Type == "" {
			modelSpec.Service.Type = lmDeployment.Spec.VLLM.GlobalConfig.Service.Type
		}
		if modelSpec.Service.Port == 0 {
			modelSpec.Service.Port = lmDeployment.Spec.VLLM.GlobalConfig.Service.Port
		}

		// Set model persistence defaults (fall back to global)
		if modelSpec.Persistence != nil && modelSpec.Persistence.Enabled {
			if modelSpec.Persistence.Size == "" {
				modelSpec.Persistence.Size = lmDeployment.Spec.VLLM.GlobalConfig.Persistence.Size
			}
		}
	}

	// Set router defaults if enabled
	if lmDeployment.Spec.VLLM.Router.Image == "" {
		lmDeployment.Spec.VLLM.Router.Image = "lmcache/lmstack-router:latest"
	}
	if lmDeployment.Spec.VLLM.Router.Replicas == 0 {
		lmDeployment.Spec.VLLM.Router.Replicas = 1
	}
	if lmDeployment.Spec.VLLM.Router.Service.Type == "" {
		lmDeployment.Spec.VLLM.Router.Service.Type = corev1.ServiceTypeClusterIP
	}
	if lmDeployment.Spec.VLLM.Router.Service.Port == 0 {
		lmDeployment.Spec.VLLM.Router.Service.Port = 8000
	}
}

func (d *LMDeploymentCustomDefaulter) defaultOpenWebUI(lmDeployment *llmgeeperiov1alpha1.LMDeployment) {
	if lmDeployment.Spec.OpenWebUI.Image == "" {
		lmDeployment.Spec.OpenWebUI.Image = "ghcr.io/open-webui/open-webui:main"
	}
	if lmDeployment.Spec.OpenWebUI.Replicas == 0 {
		lmDeployment.Spec.OpenWebUI.Replicas = 1
	}
	if lmDeployment.Spec.OpenWebUI.Service.Type == "" {
		lmDeployment.Spec.OpenWebUI.Service.Type = corev1.ServiceTypeClusterIP
	}
	if lmDeployment.Spec.OpenWebUI.Service.Port == 0 {
		lmDeployment.Spec.OpenWebUI.Service.Port = 8080
	}

	// Set OpenWebUI Redis defaults
	if lmDeployment.Spec.OpenWebUI.Redis.Image == "" {
		lmDeployment.Spec.OpenWebUI.Redis.Image = "redis:7-alpine"
	}
	if lmDeployment.Spec.OpenWebUI.Redis.Service.Port == 0 {
		lmDeployment.Spec.OpenWebUI.Redis.Service.Port = 6379
	}
	if lmDeployment.Spec.OpenWebUI.Redis.Service.Type == "" {
		lmDeployment.Spec.OpenWebUI.Redis.Service.Type = corev1.ServiceTypeClusterIP
	}
	if lmDeployment.Spec.OpenWebUI.Redis.Persistence.Size == "" {
		lmDeployment.Spec.OpenWebUI.Redis.Persistence.Size = "1Gi"
	}

	// Set default Langfuse configuration if enabled
	if lmDeployment.Spec.OpenWebUI.Langfuse != nil && lmDeployment.Spec.OpenWebUI.Langfuse.Enabled {
		// Set default project name if not provided
		if lmDeployment.Spec.OpenWebUI.Langfuse.ProjectName == "" {
			lmDeployment.Spec.OpenWebUI.Langfuse.ProjectName = lmDeployment.Name
		}

		// Set default environment if not provided
		if lmDeployment.Spec.OpenWebUI.Langfuse.Environment == "" {
			lmDeployment.Spec.OpenWebUI.Langfuse.Environment = "development"
		}
	}

	// Set OpenWebUI Pipelines defaults (for both manual and auto-enabled)
	if lmDeployment.Spec.OpenWebUI.Pipelines != nil && lmDeployment.Spec.OpenWebUI.Pipelines.Enabled {
		if lmDeployment.Spec.OpenWebUI.Pipelines.Image == "" {
			lmDeployment.Spec.OpenWebUI.Pipelines.Image = "ghcr.io/open-webui/pipelines:main"
		}
		if lmDeployment.Spec.OpenWebUI.Pipelines.Replicas == 0 {
			lmDeployment.Spec.OpenWebUI.Pipelines.Replicas = 1
		}
		if lmDeployment.Spec.OpenWebUI.Pipelines.Port == 0 {
			lmDeployment.Spec.OpenWebUI.Pipelines.Port = 9099
		}
		if lmDeployment.Spec.OpenWebUI.Pipelines.Service.Type == "" {
			lmDeployment.Spec.OpenWebUI.Pipelines.Service.Type = corev1.ServiceTypeClusterIP
		}
		if lmDeployment.Spec.OpenWebUI.Pipelines.PipelinesDir == "" {
			lmDeployment.Spec.OpenWebUI.Pipelines.PipelinesDir = "/app/pipelines"
		}
	}
}

func (d *LMDeploymentCustomDefaulter) defaultTabby(lmDeployment *llmgeeperiov1alpha1.LMDeployment) {
	if lmDeployment.Spec.Tabby.Image == "" {
		lmDeployment.Spec.Tabby.Image = "tabbyml/tabby:latest"
	}
	if lmDeployment.Spec.Tabby.Replicas == 0 {
		lmDeployment.Spec.Tabby.Replicas = 1
	}
	if lmDeployment.Spec.Tabby.Service.Type == "" {
		lmDeployment.Spec.Tabby.Service.Type = corev1.ServiceTypeClusterIP
	}
	if lmDeployment.Spec.Tabby.Service.Port == 0 {
		lmDeployment.Spec.Tabby.Service.Port = 8080
	}
}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// NOTE: The 'path' attribute must follow a specific pattern and should not be modified directly here.
// Modifying the path for an invalid path can cause API server errors; failing to locate the webhook.
// +kubebuilder:webhook:path=/validate-llm-geeper-io-v1alpha1-lmdeployment,mutating=false,failurePolicy=fail,sideEffects=None,groups=llm.geeper.io,resources=lmdeployments,verbs=create;update,versions=v1alpha1,name=vlmdeployment-v1alpha1.kb.io,admissionReviewVersions=v1

// LMDeploymentCustomValidator struct is responsible for validating the LMDeployment resource
// when it is created, updated, or deleted.
//
// NOTE: The +kubebuilder:object:generate=false marker prevents controller-gen from generating DeepCopy methods,
// as this struct is used only for temporary operations and does not need to be deeply copied.
type LMDeploymentCustomValidator struct{}

var _ webhook.CustomValidator = &LMDeploymentCustomValidator{}

// ValidateCreate implements webhook.CustomValidator so a webhook will be registered for the type LMDeployment.
func (v *LMDeploymentCustomValidator) ValidateCreate(_ context.Context, obj runtime.Object) (admission.Warnings, error) {
	lmDeployment, ok := obj.(*llmgeeperiov1alpha1.LMDeployment)
	if !ok {
		return nil, fmt.Errorf("expected a LMDeployment object but got %T", obj)
	}
	lmdeploymentlog.Info("Validation for LMDeployment upon creation", "name", lmDeployment.GetName())

	return v.validate(lmDeployment)
}

// ValidateUpdate implements webhook.CustomValidator so a webhook will be registered for the type LMDeployment.
func (v *LMDeploymentCustomValidator) ValidateUpdate(_ context.Context, oldObj, newObj runtime.Object) (admission.Warnings, error) {
	lmDeployment, ok := newObj.(*llmgeeperiov1alpha1.LMDeployment)
	if !ok {
		return nil, fmt.Errorf("expected a LMDeployment object for the newObj but got %T", newObj)
	}
	lmdeploymentlog.Info("Validation for LMDeployment upon update", "name", lmDeployment.GetName())

	return v.validate(lmDeployment)
}

// ValidateDelete implements webhook.CustomValidator so a webhook will be registered for the type LMDeployment.
func (v *LMDeploymentCustomValidator) ValidateDelete(ctx context.Context, obj runtime.Object) (admission.Warnings, error) {
	return nil, nil
}

func (l *LMDeploymentCustomValidator) validate(lmDeployment *llmgeeperiov1alpha1.LMDeployment) (admission.Warnings, error) {
	var allErrs field.ErrorList

	if !lmDeployment.Spec.Ollama.Enabled && !lmDeployment.Spec.VLLM.Enabled {
		allErrs = append(allErrs, field.Required(field.NewPath("spec"), "at least one of Ollama or vLLM must be enabled"))
	}

	if lmDeployment.Spec.VLLM.Enabled {
		if err := l.validateVLLM(lmDeployment); err != nil {
			allErrs = append(allErrs, err...)
		}
	}
	if lmDeployment.Spec.Ollama.Enabled {
		// Default to Ollama if neither is explicitly enabled
		if err := l.validateOllama(lmDeployment); err != nil {
			allErrs = append(allErrs, err...)
		}
	}

	// Validate OpenWebUI configuration if enabled
	if lmDeployment.Spec.OpenWebUI.Enabled {
		if err := l.validateOpenWebUI(lmDeployment); err != nil {
			allErrs = append(allErrs, err...)
		}
	}

	// Validate Tabby configuration if enabled
	if lmDeployment.Spec.Tabby.Enabled {
		if err := l.validateTabby(lmDeployment); err != nil {
			allErrs = append(allErrs, err...)
		}
	}

	if len(allErrs) == 0 {
		return nil, nil
	}

	return nil, &field.Error{Type: field.ErrorTypeInvalid, Field: "spec", Detail: allErrs.ToAggregate().Error()}
}

// validateOllama validates Ollama configuration
func (l *LMDeploymentCustomValidator) validateOllama(lmDeployment *llmgeeperiov1alpha1.LMDeployment) field.ErrorList {
	var allErrs field.ErrorList
	ollamaPath := field.NewPath("spec", "ollama")

	// Validate models
	if len(lmDeployment.Spec.Ollama.Models) == 0 {
		allErrs = append(allErrs, field.Required(ollamaPath.Child("models"), "at least one model must be specified"))
	}

	return allErrs
}

// validateVLLM validates vLLM configuration
func (l *LMDeploymentCustomValidator) validateVLLM(lmDeployment *llmgeeperiov1alpha1.LMDeployment) field.ErrorList {
	var allErrs field.ErrorList
	vllmPath := field.NewPath("spec", "vllm")

	// Validate models
	if len(lmDeployment.Spec.VLLM.Models) == 0 {
		allErrs = append(allErrs, field.Required(vllmPath.Child("models"), "at least one model must be specified"))
	} else {
		// Validate each model
		for i, modelSpec := range lmDeployment.Spec.VLLM.Models {
			modelPath := vllmPath.Child("models").Index(i)

			// Validate model name
			if modelSpec.Name == "" {
				allErrs = append(allErrs, field.Required(modelPath.Child("name"), "model name must be specified"))
			}

			// Validate model identifier
			if modelSpec.Model == "" {
				allErrs = append(allErrs, field.Required(modelPath.Child("model"), "model identifier must be specified"))
			}

			// Validate replicas
			if modelSpec.Replicas < 0 {
				allErrs = append(allErrs, field.Invalid(modelPath.Child("replicas"), modelSpec.Replicas, "replicas must be non-negative"))
			}

			// Validate service port
			if modelSpec.Service.Port > 0 && (modelSpec.Service.Port < 1 || modelSpec.Service.Port > 65535) {
				allErrs = append(allErrs, field.Invalid(modelPath.Child("service", "port"), modelSpec.Service.Port, "service port must be between 1 and 65535"))
			}

			// Validate persistence configuration if enabled
			if modelSpec.Persistence != nil && modelSpec.Persistence.Enabled {
				if modelSpec.Persistence.Size == "" {
					allErrs = append(allErrs, field.Required(modelPath.Child("persistence", "size"), "persistence size must be specified when persistence is enabled"))
				}
			}
		}
	}

	// Validate router configuration if enabled
	if lmDeployment.Spec.VLLM.Router != nil && lmDeployment.Spec.VLLM.Router.Enabled {
		routerPath := vllmPath.Child("router")

		// Validate router replicas
		if lmDeployment.Spec.VLLM.Router.Replicas < 0 {
			allErrs = append(allErrs, field.Invalid(routerPath.Child("replicas"), lmDeployment.Spec.VLLM.Router.Replicas, "router replicas must be non-negative"))
		}

		// Validate router service port
		if lmDeployment.Spec.VLLM.Router.Service.Port > 0 && (lmDeployment.Spec.VLLM.Router.Service.Port < 1 || lmDeployment.Spec.VLLM.Router.Service.Port > 65535) {
			allErrs = append(allErrs, field.Invalid(routerPath.Child("service", "port"), lmDeployment.Spec.VLLM.Router.Service.Port, "router service port must be between 1 and 65535"))
		}
	}

	// Validate global configuration if specified
	if lmDeployment.Spec.VLLM.GlobalConfig != nil {
		globalPath := vllmPath.Child("globalConfig")

		// Validate global service port
		if lmDeployment.Spec.VLLM.GlobalConfig.Service.Port > 0 && (lmDeployment.Spec.VLLM.GlobalConfig.Service.Port < 1 || lmDeployment.Spec.VLLM.GlobalConfig.Service.Port > 65535) {
			allErrs = append(allErrs, field.Invalid(globalPath.Child("service", "port"), lmDeployment.Spec.VLLM.GlobalConfig.Service.Port, "global service port must be between 1 and 65535"))
		}

		// Validate global persistence configuration if enabled
		if lmDeployment.Spec.VLLM.GlobalConfig.Persistence != nil && lmDeployment.Spec.VLLM.GlobalConfig.Persistence.Enabled {
			if lmDeployment.Spec.VLLM.GlobalConfig.Persistence.Size == "" {
				allErrs = append(allErrs, field.Required(globalPath.Child("persistence", "size"), "global persistence size must be specified when persistence is enabled"))
			}
		}
	}

	return allErrs
}

// validateOpenWebUI validates OpenWebUI configuration
func (l *LMDeploymentCustomValidator) validateOpenWebUI(lmDeployment *llmgeeperiov1alpha1.LMDeployment) field.ErrorList {
	var allErrs field.ErrorList
	openwebuiPath := field.NewPath("spec", "openwebui")

	// Validate Redis configuration if enabled
	if lmDeployment.Spec.OpenWebUI.Redis.Enabled {
		redisPath := openwebuiPath.Child("redis")

		// Validate Redis persistence size if enabled
		if lmDeployment.Spec.OpenWebUI.Redis.Persistence.Enabled {
			if lmDeployment.Spec.OpenWebUI.Redis.Persistence.Size == "" {
				allErrs = append(allErrs, field.Required(redisPath.Child("persistence", "size"), "Redis persistence size must be specified when persistence is enabled"))
			}
		}
	}

	// Validate Pipelines configuration if enabled
	if lmDeployment.Spec.OpenWebUI.Pipelines != nil && lmDeployment.Spec.OpenWebUI.Pipelines.Enabled {
		pipelinesPath := openwebuiPath.Child("pipelines")

		// Validate Pipelines port
		if lmDeployment.Spec.OpenWebUI.Pipelines.Port < 1 || lmDeployment.Spec.OpenWebUI.Pipelines.Port > 65535 {
			allErrs = append(allErrs, field.Invalid(pipelinesPath.Child("port"), lmDeployment.Spec.OpenWebUI.Pipelines.Port, "Pipelines port must be between 1 and 65535"))
		}

		// Validate Pipelines service port
		if lmDeployment.Spec.OpenWebUI.Pipelines.Service.Port > 65535 {
			allErrs = append(allErrs, field.Invalid(pipelinesPath.Child("service", "port"), lmDeployment.Spec.OpenWebUI.Pipelines.Service.Port, "Pipelines service port must be below 65535"))
		}
	}

	// Validate Langfuse configuration if enabled
	if lmDeployment.Spec.OpenWebUI.Langfuse != nil && lmDeployment.Spec.OpenWebUI.Langfuse.Enabled {
		langfusePath := openwebuiPath.Child("langfuse")

		// Validate Langfuse URL
		if lmDeployment.Spec.OpenWebUI.Langfuse.URL == "" {
			allErrs = append(allErrs, field.Required(langfusePath.Child("url"), "Langfuse URL must be specified when Langfuse is enabled"))
		}

		// Validate Langfuse project name
		if lmDeployment.Spec.OpenWebUI.Langfuse.ProjectName != "" {
			if errs := validation.IsDNS1123Subdomain(lmDeployment.Spec.OpenWebUI.Langfuse.ProjectName); len(errs) > 0 {
				allErrs = append(allErrs, field.Invalid(langfusePath.Child("projectName"), lmDeployment.Spec.OpenWebUI.Langfuse.ProjectName, fmt.Sprintf("project name is invalid: %s", strings.Join(errs, ", "))))
			}
		}
	}

	return allErrs
}

// validateTabby validates Tabby configuration
func (l *LMDeploymentCustomValidator) validateTabby(lmDeployment *llmgeeperiov1alpha1.LMDeployment) field.ErrorList {
	var allErrs field.ErrorList
	tabbyPath := field.NewPath("spec", "tabby")

	// Validate device
	if lmDeployment.Spec.Tabby.Device != "" {
		validDevices := []string{"cpu", "cuda", "rocm", "metal", "vulkan"}
		valid := false
		for _, device := range validDevices {
			if lmDeployment.Spec.Tabby.Device == device {
				valid = true
				break
			}
		}
		if !valid {
			allErrs = append(allErrs, field.NotSupported(tabbyPath.Child("device"), lmDeployment.Spec.Tabby.Device, validDevices))
		}
	}

	// Validate chat model if specified
	if lmDeployment.Spec.Tabby.ChatModel != "" {
		// Check if the chat model exists in Ollama or vLLM models
		found := false
		if lmDeployment.Spec.VLLM.Enabled {
			if lmDeployment.Spec.VLLM.Models != nil {
				for _, modelSpec := range lmDeployment.Spec.VLLM.Models {
					if modelSpec.Model == lmDeployment.Spec.Tabby.ChatModel {
						found = true
						break
					}
				}
			}
		}
		if lmDeployment.Spec.Ollama.Enabled {
			for _, model := range lmDeployment.Spec.Ollama.Models {
				if model == lmDeployment.Spec.Tabby.ChatModel {
					found = true
					break
				}
			}
		}
		if !found {
			modelSource := "spec.vllm.models"
			if !lmDeployment.Spec.VLLM.Enabled {
				modelSource = "spec.ollama.models"
			}
			allErrs = append(allErrs, field.Invalid(tabbyPath.Child("chatModel"), lmDeployment.Spec.Tabby.ChatModel, fmt.Sprintf("chat model must be one of the models specified in %s", modelSource)))
		}
	}

	// Validate completion model if specified
	if lmDeployment.Spec.Tabby.CompletionModel != "" {
		// Check if the completion model exists in Ollama or vLLM models
		found := false
		if lmDeployment.Spec.VLLM.Enabled {
			if lmDeployment.Spec.VLLM.Models != nil {
				for _, modelSpec := range lmDeployment.Spec.VLLM.Models {
					if modelSpec.Model == lmDeployment.Spec.Tabby.CompletionModel {
						found = true
						break
					}
				}
			}
		}
		if lmDeployment.Spec.Ollama.Enabled {
			for _, model := range lmDeployment.Spec.Ollama.Models {
				if model == lmDeployment.Spec.Tabby.CompletionModel {
					found = true
					break
				}
			}
		}
		if !found {
			modelSource := "spec.vllm.models"
			if !lmDeployment.Spec.VLLM.Enabled {
				modelSource = "spec.ollama.models"
			}
			allErrs = append(allErrs, field.Invalid(tabbyPath.Child("completionModel"), lmDeployment.Spec.Tabby.CompletionModel, fmt.Sprintf("completion model must be one of the models specified in %s", modelSource)))
		}
	}

	return allErrs
}
