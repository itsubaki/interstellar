package main

import "github.com/itsubaki/interstellar/broker"

type DatabaseBroker struct {
}

func NewDatabaseBroker() *DatabaseBroker {
	return &DatabaseBroker{}
}

func (b *DatabaseBroker) Config() *broker.Config {
	return &broker.Config{
		Port: ":9082",
	}
}

func (b *DatabaseBroker) Catalog() *broker.Catalog {
	return &broker.Catalog{
		Name: "aws_database",
		Tag: []string{
			"aws",
			"database",
		},
		Require:  []string{"aws_project"},
		Bindable: true,
	}
}

func (b *DatabaseBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	// in.Parameter["project_name"]
	// in.Parameter["environment"]
	// in.Parameter["instance_name"]
	// in.Parameter["master_username"]
	// in.Parameter["master_password"]
	return &broker.CreateOutput{
		Status:  201,
		Message: "Created",
		Input:   in,
		Output: []*broker.Output{
			{Key: "username", Value: "foobar"},
			{Key: "password", Value: "hogehoge"},
			{Key: "endpoint_write", Value: "db://${environ}-${instance_name}.write.${project_name}.${domain}"},
			{Key: "endpoint_read", Value: "db://${environ}-${instance_name}.read.${project_name}.${domain}"},
		},
	}
}

func (b *DatabaseBroker) Delete(in *broker.DeleteInput) *broker.DeleteOutput {
	return &broker.DeleteOutput{}
}

func (b *DatabaseBroker) Update(in *broker.UpdateInput) *broker.UpdateOutput {
	return &broker.UpdateOutput{}
}

func (b *DatabaseBroker) Binding(in *broker.BindingInput) *broker.BindingOutput {
	// sg := in.Parameter["securitygroup_id"]
	// Add sg with 3306 to database_sg
	return &broker.BindingOutput{}
}

func (b *DatabaseBroker) Unbinding(in *broker.UnbindingInput) *broker.UnbindingOutput {
	// sg := in.Parameter["securitygroup_id"]
	// Delete sg from database_sg
	return &broker.UnbindingOutput{}
}
