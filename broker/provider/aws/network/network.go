package network

import "github.com/itsubaki/interstellar/broker"

type NetworkBroker struct {
}

func NewNetworkBroker() *NetworkBroker {
	return &NetworkBroker{}
}

func (b *NetworkBroker) Catalog() *broker.Catalog {
	return &broker.Catalog{}
}

func (b *NetworkBroker) Binding(in *broker.BindingInput) *broker.BindingOutput {
	return &broker.BindingOutput{}
}

func (b *NetworkBroker) Unbinding(in *broker.UnbindingInput) *broker.UnbindingOutput {
	return &broker.UnbindingOutput{}
}

func (b *NetworkBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	return &broker.CreateOutput{}
}

func (b *NetworkBroker) Delete(in *broker.DeleteInput) *broker.DeleteOutput {
	return &broker.DeleteOutput{}
}

func (b *NetworkBroker) Update(in *broker.UpdateInput) *broker.UpdateOutput {
	return &broker.UpdateOutput{}
}
