package main

import "github.com/itsubaki/interstellar/broker"

type ComputeBroker struct {
}

func NewComputeBroker() *ComputeBroker {
	return &ComputeBroker{}
}

func (b *ComputeBroker) Config() *broker.Config {
	return &broker.Config{
		Port: ":9081",
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

func (b *ComputeBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	return &broker.CreateOutput{
		Status:  202,
		Message: "Accepted",
		Input:   in,
		Output: []*broker.Output{
			{Key: "endpoint", Value: "https://${environ}-${instance_name}.${project_name}.{domain}"},
		},
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
