package broker

type ServiceBroker interface {
	Config() *Config
	Catalog() *Catalog

	// 200 OK: binding already exists
	// 201 Created: binding was created
	// 409 Conflict: binding already exists and there is some difference between the input
	Binding(in *BindingInput) *BindingOutput

	// 200 OK: binding was deleted
	// 410 Gone: binding does not exist
	Unbinding(in *UnbindingInput) *UnbindingOutput

	// 200 OK: instance already exists
	// 201 Created: instance was created
	// 202 Accepted: instance creation is in progress
	// 409 Conflict: instance already exists and there is some difference between the input
	Create(in *CreateInput) *CreateOutput

	// 200 OK: instance was deleted
	// 202 Accepted: instance deletion is in progress
	// 410 Gone: instance does not exist
	Delete(in *DeleteInput) *DeleteOutput

	// 200 OK: changes have been applied
	// 202 Accepted: instance update is in progress
	Update(in *UpdateInput) *UpdateOutput

	Describe(in *DescribeInput) *DescribeOutput
}

type Config struct {
	Port     string
	Template string
}

type Catalog struct {
	Name          string       `json:"name"`
	DisplayName   string       `json:"display_name"`
	Description   string       `json:"description"`
	Tag           []string     `json:"tag"`
	Require       []string     `json:"require"`
	Bindable      bool         `json:"bindable"`
	ParameterSpec []*ParamSpec `json:"parameter"`
}

type ParamSpec struct {
	Name         string   `json:"name"`
	Required     bool     `json:"required"`
	DefaultValue string   `json:"default_value"`
	AllowedValue []string `json:"allowed_value"`
	Description  string   `json:"description"`
}

type CreateInput struct {
	InstanceID string            `json:"instance_id,omitempty"`
	Parameter  map[string]string `json:"parameter"`
}

type CreateOutput struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Output  map[string]string `json:"output,omitempty"`
}

type DeleteInput struct {
	InstanceID string `json:"instance_id"`
}

type DeleteOutput struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type UpdateInput struct {
	InstanceID string            `json:"instance_id"`
	Parameter  map[string]string `json:"parameter"`
}

type UpdateOutput struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Output  map[string]string `json:"output,omitempty"`
}

type BindingInput struct {
	InstanceID string            `json:"instance_id"`
	Parameter  map[string]string `json:"parameter"`
}

type BindingOutput struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Output  map[string]string `json:"output,omitempty"`
}

type UnbindingInput struct {
}

type UnbindingOutput struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Output  map[string]string `json:"output,omitempty"`
}

type DescribeInput struct {
	InstanceID string            `json:"instance_id"`
	Parameter  map[string]string `json:"parameter"`
}

type DescribeOutput struct {
	Status  int               `json:"status"`
	Message string            `json:"message"`
	Output  map[string]string `json:"output,omitempty"`
}
