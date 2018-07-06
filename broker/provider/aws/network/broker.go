package main

import (
	"github.com/itsubaki/env"
	"github.com/itsubaki/interstellar/broker"
)

type NetworkBroker struct {
}

func NewNetworkBroker() (*NetworkBroker, error) {
	return &NetworkBroker{}, nil
}

func (b *NetworkBroker) Config() *broker.Config {
	return &broker.Config{
		Port: env.GetValue("PORT", ":8080"),
	}
}

func (b *NetworkBroker) Catalog() *broker.Catalog {
	return &broker.Catalog{
		Name: "aws_network",
		Tag: []string{
			"aws",
			"network",
		},
		Bindable: true,
		ParameterSpec: []broker.ParamSpec{
			{Name: "project_name", Required: true},
			{Name: "network_name", Required: true},
			{Name: "region", Required: false},
		},
	}
}

func (b *NetworkBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
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

func (b *NetworkBroker) Delete(in *broker.DeleteInput) *broker.DeleteOutput {
	return &broker.DeleteOutput{
		Status:  202,
		Message: "Accepted",
	}
}

func (b *NetworkBroker) Update(in *broker.UpdateInput) *broker.UpdateOutput {
	return &broker.UpdateOutput{}
}

func (b *NetworkBroker) Binding(in *broker.BindingInput) *broker.BindingOutput {
	return &broker.BindingOutput{}
}

func (b *NetworkBroker) Unbinding(in *broker.UnbindingInput) *broker.UnbindingOutput {
	return &broker.UnbindingOutput{}
}

func (b *NetworkBroker) Describe(in *broker.DescribeInput) *broker.DescribeOutput {
	return &broker.DescribeOutput{}
}
