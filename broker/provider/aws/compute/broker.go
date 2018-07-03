package main

import (
	"github.com/itsubaki/env"
	"github.com/itsubaki/interstellar/broker"
)

type ComputeBroker struct {
}

func NewComputeBroker() (*ComputeBroker, error) {
	return &ComputeBroker{}, nil
}

func (b *ComputeBroker) Config() *broker.Config {
	return &broker.Config{
		Port: env.GetValue("PORT", ":8080"),
	}
}

func (b *ComputeBroker) Catalog() *broker.Catalog {
	return &broker.Catalog{
		Name: "aws_compute",
		Tag: []string{
			"aws",
			"compute",
		},
		Bindable: false,
		ParameterSpec: []broker.ParamSpec{
			{Name: "project_name", Required: true},
			{Name: "environ_name", Required: true},
			{Name: "instance_name", Required: true},
			{Name: "region", Required: false},
		},
	}
}

// ExportName is related with project_name, environ_name, instance_name
// ExportValue
//  - endpoint
func (b *ComputeBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	i := &broker.Instance{
		InstanceID: in.InstanceID,
		Parameter:  in.Parameter,
	}

	return &broker.CreateOutput{
		Status:   202,
		Message:  "Accepted",
		Instance: i,
	}
}

func (b *ComputeBroker) Delete(in *broker.DeleteInput) *broker.DeleteOutput {
	return &broker.DeleteOutput{}
}

func (b *ComputeBroker) Update(in *broker.UpdateInput) *broker.UpdateOutput {
	return &broker.UpdateOutput{}
}

func (b *ComputeBroker) Binding(in *broker.BindingInput) *broker.BindingOutput {
	// project_name
	// environ_name
	// instance_name
	// run_command -> docker environment
	return &broker.BindingOutput{}
}

func (b *ComputeBroker) Unbinding(in *broker.UnbindingInput) *broker.UnbindingOutput {
	return &broker.UnbindingOutput{}
}

func (b *ComputeBroker) Describe(in *broker.DescribeInput) *broker.DescribeOutput {
	return &broker.DescribeOutput{}
}
