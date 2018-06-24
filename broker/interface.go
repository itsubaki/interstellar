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

	Status(in *StatusInput) *StatusOutput
}

type Config struct {
	Port string
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

type Parameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Output struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type CreateInput struct {
	InstanceID string       `json:"instance_id"`
	Parameter  []*Parameter `json:"parameter"`
}

type CreateOutput struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Input   *CreateInput `json:"input"`
	Output  []*Output    `json:"output"`
}

type DeleteInput struct {
	InstanceID string `json:"instance_id"`
}

type DeleteOutput struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Input   *DeleteInput `json:"input"`
}

type UpdateInput struct {
	InstanceID string       `json:"instance_id"`
	Parameter  []*Parameter `json:"parameter"`
}

type UpdateOutput struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Input   *UpdateInput `json:"input"`
	Output  []*Output    `json:"output"`
}

type BindingInput struct {
	InstanceID string       `json:"instance_id"`
	Parameter  []*Parameter `json:"parameter"`
}

type BindingOutput struct {
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Input   *UpdateInput `json:"input"`
	Output  []*Output    `json:"output"`
}

type UnbindingInput struct {
}

type UnbindingOutput struct {
}

type StatusInput struct {
	InstanceID string `json:"instance_id"`
}

type StatusOutput struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
