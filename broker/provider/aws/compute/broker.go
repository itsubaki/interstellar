package main

import "github.com/itsubaki/interstellar/broker"

type ComputeBroker struct {
}

func NewComputeBroker() *ComputeBroker {
	return &ComputeBroker{}
}

func (b *ComputeBroker) Config() *broker.Config {
	return &broker.Config{
		Port: ":8080",
	}
}

func (b *ComputeBroker) Catalog() *broker.Catalog {
	return &broker.Catalog{
		Name: "aws_compute",
		Tag: []string{
			"aws",
			"compute",
		},
		Require:  []string{"aws_project", "aws_environ"},
		Bindable: false,
		ParameterSpec: []*broker.ParamSpec{
			{Name: "project_name", Required: true},
			{Name: "environ_name", Required: true},
			{Name: "instance_name", Required: true},
			{Name: "image_id", Required: true},
			{Name: "package", Required: true},
		},
	}
}

// ExportName is related with project_name, environ_name, instance_name
// ExportValue
//  - endpoint
func (b *ComputeBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	out := make(map[string]string)
	out["endpoint"] = "https://${environ}-${instance_name}.${project_name}.{domain}"

	return &broker.CreateOutput{
		Status:  202,
		Message: "Accepted",
		Input:   in,
		Output:  out,
	}
}

func (b *ComputeBroker) Delete(in *broker.DeleteInput) *broker.DeleteOutput {
	return &broker.DeleteOutput{}
}

func (b *ComputeBroker) Update(in *broker.UpdateInput) *broker.UpdateOutput {
	return &broker.UpdateOutput{}
}

func (b *ComputeBroker) Binding(in *broker.BindingInput) *broker.BindingOutput {
	return &broker.BindingOutput{}
}

func (b *ComputeBroker) Unbinding(in *broker.UnbindingInput) *broker.UnbindingOutput {
	return &broker.UnbindingOutput{}
}

func (b *ComputeBroker) Status(in *broker.StatusInput) *broker.StatusOutput {
	return &broker.StatusOutput{}
}
