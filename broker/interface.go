package broker

type ServiceBroker interface {
	Config() *Config
	Catalog() *Catalog
	Binding(in *BindingInput) *BindingOutput
	Unbinding(in *UnbindingInput) *UnbindingOutput
	Create(in *CreateInput) *CreateOutput
	Delete(in *DeleteInput) *DeleteOutput
	Update(in *UpdateInput) *UpdateOutput
}

type Config struct {
	Port string
}

type Catalog struct {
}

type BindingInput struct {
}

type BindingOutput struct {
}

type UnbindingInput struct {
}

type UnbindingOutput struct {
}

type CreateInput struct {
}

type CreateOutput struct {
}

type DeleteInput struct {
}

type DeleteOutput struct {
}

type UpdateInput struct {
}

type UpdateOutput struct {
}
