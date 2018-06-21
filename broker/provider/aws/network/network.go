package main

import "github.com/itsubaki/interstellar/broker"

type NetworkBroker struct {
}

func NewNetworkBroker() *NetworkBroker {
	return &NetworkBroker{}
}

func (b *NetworkBroker) Config() *broker.Config {
	return &broker.Config{
		Port: ":9083",
	}
}

func (b *NetworkBroker) Catalog() *broker.Catalog {
	return &broker.Catalog{
		Name: "aws_network",
		Tag: []string{
			"aws",
			"network",
		},
		Require:  []string{},
		Bindable: true,
	}
}

func (b *NetworkBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	return &broker.CreateOutput{
		Status:  201,
		Message: "Created",
		Input:   in,
		Output: []*broker.Output{
			{Key: "vpc_id", Value: "vpc-12345678"},
		},
	}
}

func (b *NetworkBroker) Delete(in *broker.DeleteInput) *broker.DeleteOutput {
	return &broker.DeleteOutput{
		Status:  202,
		Message: "Accepted",
		Input:   in,
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
