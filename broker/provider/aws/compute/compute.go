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
	return &broker.Catalog{}
}

func (b *ComputeBroker) Binding(in *broker.BindingInput) *broker.BindingOutput {
	return &broker.BindingOutput{}
}

func (b *ComputeBroker) Unbinding(in *broker.UnbindingInput) *broker.UnbindingOutput {
	return &broker.UnbindingOutput{}
}

func (b *ComputeBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	return &broker.CreateOutput{}
}

func (b *ComputeBroker) Delete(in *broker.DeleteInput) *broker.DeleteOutput {
	return &broker.DeleteOutput{}
}

func (b *ComputeBroker) Update(in *broker.UpdateInput) *broker.UpdateOutput {
	return &broker.UpdateOutput{}
}
