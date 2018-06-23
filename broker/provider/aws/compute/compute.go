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
		Require:  []string{"aws_initialize"},
		Bindable: false,
	}
}

func (b *ComputeBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	// in.Parameter["project_name"]
	// in.Parameter["environment"]
	// in.Parameter["instance_name"]
	// in.Parameter["image_id"]
	// s3://deploy.${project_name}.${domain}/application/release/v1.0.tar.gz
	// in.Parameter["application_key"]
	return &broker.CreateOutput{
		Status:  201,
		Message: "Created",
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
