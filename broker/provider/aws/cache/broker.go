package main

import (
	"github.com/itsubaki/env"
	"github.com/itsubaki/interstellar/broker"
)

type CacheBroker struct {
}

func NewCacheBroker() *CacheBroker {
	return &CacheBroker{}
}

func (b *CacheBroker) Config() *broker.Config {
	return &broker.Config{
		Port: env.GetValue("PORT", ":8080"),
	}
}

func (b *CacheBroker) Catalog() *broker.Catalog {
	return &broker.Catalog{
		Name: "aws_cache",
		Tag: []string{
			"aws",
			"cache",
		},
		Bindable: false,
		ParameterSpec: []broker.ParamSpec{
			{Name: "integration_role_arn", Required: true},
			{Name: "region", Required: false},

			{Name: "project_name", Required: true},
			{Name: "environ_name", Required: true},
			{Name: "instance_name", Required: true},
		},
	}
}

func (b *CacheBroker) Binding(in *broker.BindingInput) *broker.BindingOutput {
	return &broker.BindingOutput{}
}

func (b *CacheBroker) Unbinding(in *broker.UnbindingInput) *broker.UnbindingOutput {
	return &broker.UnbindingOutput{}
}

func (b *CacheBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	return &broker.CreateOutput{}
}

func (b *CacheBroker) Delete(in *broker.DeleteInput) *broker.DeleteOutput {
	return &broker.DeleteOutput{}
}

func (b *CacheBroker) Update(in *broker.UpdateInput) *broker.UpdateOutput {
	return &broker.UpdateOutput{}
}

func (b *CacheBroker) Describe(in *broker.DescribeInput) *broker.DescribeOutput {
	return &broker.DescribeOutput{}
}
