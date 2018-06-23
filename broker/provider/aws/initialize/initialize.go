package main

import "github.com/itsubaki/interstellar/broker"

type InitializeBroker struct {
}

func NewInitializeBroker() *InitializeBroker {
	return &InitializeBroker{}
}

func (b *InitializeBroker) Config() *broker.Config {
	return &broker.Config{
		Port: ":9083",
	}
}

func (b *InitializeBroker) Catalog() *broker.Catalog {
	return &broker.Catalog{
		Name: "aws_initialize",
		Tag: []string{
			"aws",
			"initialize",
		},
		Require:  []string{},
		Bindable: false,
	}
}

func (b *InitializeBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	// in.Parameter["integration_role_arn"]
	// in.Parameter["vpc_cidr"]
	return &broker.CreateOutput{
		Status:  201,
		Message: "Created",
		Input:   in,
		Output: []*broker.Output{
			{Key: "vpc_id", Value: "vpc-12345678"},
			{Key: "subnet_a_public", Value: "subnet-12345678"},
			{Key: "subnet_b_public", Value: "subnet-12345678"},
			{Key: "subnet_c_public", Value: "subnet-12345678"},
			{Key: "subnet_a_private", Value: "subnet-12345678"},
			{Key: "subnet_b_private", Value: "subnet-12345678"},
			{Key: "subnet_c_private", Value: "subnet-12345678"},
		},
	}
}

func (b *InitializeBroker) Delete(in *broker.DeleteInput) *broker.DeleteOutput {
	return &broker.DeleteOutput{
		Status:  202,
		Message: "Accepted",
		Input:   in,
	}
}

func (b *InitializeBroker) Update(in *broker.UpdateInput) *broker.UpdateOutput {
	return &broker.UpdateOutput{}
}

func (b *InitializeBroker) Binding(in *broker.BindingInput) *broker.BindingOutput {
	return &broker.BindingOutput{}
}

func (b *InitializeBroker) Unbinding(in *broker.UnbindingInput) *broker.UnbindingOutput {
	return &broker.UnbindingOutput{}
}
