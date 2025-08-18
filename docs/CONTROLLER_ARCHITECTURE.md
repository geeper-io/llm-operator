# Controller Architecture

The LLM Operator controller has been refactored into a modular architecture with separate, focused controllers for better maintainability and separation of concerns.

## Architecture Overview

```
internal/controller/
├── deployment_controller.go    # Main orchestrator controller for LMDeployments
├── ollama_controller.go        # Ollama-specific operations
├── openwebui_controller.go     # OpenWebUI-specific operations
└── tool_controller.go        # Tool management operations
```

## Controller Responsibilities

### 1. Main LMDeployment Controller (`deployment_controller.go`)

**Purpose**: Main orchestrator that coordinates all operations

**Responsibilities**:
- Main reconciliation loop
- Finalizer management
- Default value setting
- Status updates
- Resource creation/update utilities
- Controller initialization and setup

**Key Functions**:
- `Reconcile()` - Main reconciliation loop
- `setDefaults()` - Set default values for all components
- `updateStatus()` - Update overall LMDeployment status
- `createOrUpdate*()` - Generic resource management utilities
- `SetupWithManager()` - Initialize specialized controllers

### 2. Ollama Controller (`ollama_controller.go`)

**Purpose**: Handle all Ollama-specific operations

**Responsibilities**:
- Ollama LMDeployment creation and management
- Ollama service creation and management
- Model pulling configuration
- Resource requirements for Ollama

**Key Functions**:
- `ReconcileOllama()` - Main Ollama reconciliation
- `buildOllamaDeployment()` - Build Ollama LMDeployment spec
- `buildOllamaService()` - Build Ollama service spec

### 3. OpenWebUI Controller (`openwebui_controller.go`)

**Purpose**: Handle all OpenWebUI-specific operations

**Responsibilities**:
- OpenWebUI LMDeployment creation and management
- OpenWebUI service creation and management
- Ingress configuration
- Configuration ConfigMap generation
- Volume mounting for configuration

**Key Functions**:
- `ReconcileOpenWebUI()` - Main OpenWebUI reconciliation
- `buildOpenWebUIDeployment()` - Build OpenWebUI LMDeployment spec
- `buildOpenWebUIService()` - Build OpenWebUI service spec
- `buildOpenWebUIIngress()` - Build OpenWebUI ingress spec
- `buildOpenWebUIConfigMap()` - Generate OpenWebUI configuration

### 4. Tool Controller (`tool_controller.go`)

**Purpose**: Handle all tool management operations

**Responsibilities**:
- Tool LMDeployment creation and management
- Tool service creation and management
- Tool configuration and credential management
- Environment variable setup for tools

**Key Functions**:
- `ReconcileTools()` - Main tool reconciliation
- `buildToolDeployment()` - Build tool LMDeployment spec
- `buildToolService()` - Build tool service spec

## Benefits of This Architecture

### 1. **Separation of Concerns**
- Each controller has a single, well-defined responsibility
- Clear boundaries between different types of operations
- Easier to understand and maintain

### 2. **Modularity**
- Controllers can be developed and tested independently
- Easy to add new functionality to specific areas
- Reduced coupling between different components

### 3. **Maintainability**
- Smaller, focused files are easier to navigate
- Changes to one component don't affect others
- Clear ownership of functionality

### 4. **Testability**
- Each controller can be unit tested independently
- Easier to mock dependencies
- Better test coverage and isolation

### 5. **Extensibility**
- New controllers can be added easily
- Existing controllers can be enhanced without affecting others
- Clear patterns for adding new functionality

## How It Works

### 1. **Initialization**
```go
func (r *OllamaDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
    // Initialize specialized controllers
    r.ollamaController = NewOllamaController()
    r.openwebuiController = NewOpenWebUIController()
    		r.toolController = NewToolController()
    
    // ... rest of setup
}
```

### 2. **Coordination**
```go
func (r *OllamaDeploymentReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
    // ... setup and validation ...
    
    // Delegate to specialized controllers
    if err := r.ollamaController.ReconcileOllama(ctx, deployment, r); err != nil {
        return ctrl.Result{RequeueAfter: time.Minute}, err
    }
    
    if deployment.Spec.OpenWebUI.Enabled {
        if err := r.openwebuiController.ReconcileOpenWebUI(ctx, deployment, r); err != nil {
            return ctrl.Result{RequeueAfter: time.Minute}, err
        }
        
        		if len(deployment.Spec.OpenWebUI.Tools) > 0 {
            		if err := r.toolController.ReconcileTools(ctx, deployment, r); err != nil {
                return ctrl.Result{RequeueAfter: time.Minute}, err
            }
        }
    }
    
    // ... status update ...
}
```

### 3. **Resource Management**
- Main controller provides utility functions for resource creation/updates
- Specialized controllers use these utilities to manage their resources
- All resources maintain proper owner references for cleanup

## Adding New Controllers

To add a new specialized controller:

1. **Create the controller file**:
   ```go
   // new_component_controller.go
   package controller
   
   type NewComponentController struct{}
   
   func NewNewComponentController() *NewComponentController {
       return &NewComponentController{}
   }
   
   func (c *NewComponentController) ReconcileNewComponent(ctx context.Context, deployment *llmgeeperiov1alpha1.Deployment, r *OllamaDeploymentReconciler) error {
       // Implementation
   }
   ```

2. **Add to main controller**:
   ```go
   type OllamaDeploymentReconciler struct {
       // ... existing fields ...
       newComponentController *NewComponentController
   }
   ```

3. **Initialize in SetupWithManager**:
   ```go
   func (r *OllamaDeploymentReconciler) SetupWithManager(mgr ctrl.Manager) error {
       // ... existing initialization ...
       r.newComponentController = NewNewComponentController()
       
       // ... rest of setup
   }
   ```

4. **Use in reconciliation**:
   ```go
   if err := r.newComponentController.ReconcileNewComponent(ctx, deployment, r); err != nil {
       return ctrl.Result{RequeueAfter: time.Minute}, err
   }
   ```

## Best Practices

### 1. **Controller Design**
- Keep controllers focused on a single responsibility
- Use clear, descriptive names for controllers and methods
- Follow consistent patterns across all controllers

### 2. **Resource Management**
- Always use the main controller's utility functions for resource operations
- Maintain proper owner references for all resources
- Handle errors appropriately and return meaningful error messages

### 3. **Configuration**
- Keep configuration logic in the appropriate specialized controller
- Use consistent patterns for building Kubernetes resources
- Validate configuration before creating resources

### 4. **Testing**
- Test each controller independently
- Mock dependencies appropriately
- Test both success and failure scenarios

## Future Enhancements

This architecture provides a solid foundation for future enhancements:

- **Metrics and Monitoring**: Each controller can expose its own metrics
- **Health Checks**: Individual controller health can be monitored
- **Configuration Validation**: Each controller can validate its own configuration
- **Tool System**: Easy to add new types of components
- **Multi-version Support**: Different API versions can be handled by different controllers


