package main

import (
	"github.com/itsubaki/env"
	"github.com/itsubaki/interstellar/broker"
)

type EnvironBroker struct {
}

func NewEnvironBroker() *EnvironBroker {
	return &EnvironBroker{}
}

func (b *EnvironBroker) Config() *broker.Config {
	return &broker.Config{
		Port: env.GetValue("PORT", ":8080"),
	}
}

func (b *EnvironBroker) Catalog() *broker.Catalog {
	return &broker.Catalog{
		Name: "aws_environ",
		Tag: []string{
			"aws",
			"environ",
		},
		Bindable: true,
		ParameterSpec: []broker.ParamSpec{
			{Name: "integration_role_arn", Required: true},
			{Name: "region", Required: false},

			{Name: "project_name", Required: true},
			{Name: "environ_name", Required: true},
		},
	}
}

func (b *EnvironBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
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

func (b *EnvironBroker) Delete(in *broker.DeleteInput) *broker.DeleteOutput {
	return &broker.DeleteOutput{
		Status:  202,
		Message: "Accepted",
	}
}

func (b *EnvironBroker) Update(in *broker.UpdateInput) *broker.UpdateOutput {
	return &broker.UpdateOutput{}
}

func (b *EnvironBroker) Binding(in *broker.BindingInput) *broker.BindingOutput {
	return &broker.BindingOutput{}
}

func (b *EnvironBroker) Unbinding(in *broker.UnbindingInput) *broker.UnbindingOutput {
	return &broker.UnbindingOutput{}
}

func (b *EnvironBroker) Describe(in *broker.DescribeInput) *broker.DescribeOutput {
	return &broker.DescribeOutput{}
}
