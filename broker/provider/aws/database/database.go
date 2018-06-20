package database

import "github.com/itsubaki/interstellar/broker"

type DatabaseBroker struct {
}

func NewDatabaseBroker() *DatabaseBroker {
	return &DatabaseBroker{}
}

func (b *DatabaseBroker) Config() *broker.Config {
	return &broker.Config{}
}

func (b *DatabaseBroker) Catalog() *broker.Catalog {
	return &broker.Catalog{}
}

func (b *DatabaseBroker) Binding(in *broker.BindingInput) *broker.BindingOutput {
	return &broker.BindingOutput{}
}

func (b *DatabaseBroker) Unbinding(in *broker.UnbindingInput) *broker.UnbindingOutput {
	return &broker.UnbindingOutput{}
}

func (b *DatabaseBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	return &broker.CreateOutput{}
}

func (b *DatabaseBroker) Delete(in *broker.DeleteInput) *broker.DeleteOutput {
	return &broker.DeleteOutput{}
}

func (b *DatabaseBroker) Update(in *broker.UpdateInput) *broker.UpdateOutput {
	return &broker.UpdateOutput{}
}
