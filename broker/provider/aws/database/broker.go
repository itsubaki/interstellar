package main

import (
	"github.com/itsubaki/interstellar/broker"
	"github.com/itsubaki/interstellar/util"
)

type DatabaseBroker struct {
}

func NewDatabaseBroker() *DatabaseBroker {
	return &DatabaseBroker{}
}

func (b *DatabaseBroker) Config() *broker.Config {
	return &broker.Config{
		Port: util.Getenv("PORT", ":8080"),
	}
}

func (b *DatabaseBroker) Catalog() *broker.Catalog {
	return &broker.Catalog{
		Name: "aws_database",
		Tag: []string{
			"aws",
			"database",
		},
		Require:  []string{"aws_project", "aws_environ"},
		Bindable: true,
		ParameterSpec: []broker.ParamSpec{
			{Name: "project_name", Required: true},
			{Name: "environ_name", Required: true},
			{Name: "instance_name", Required: true},
			{Name: "master_username", Required: true},
			{Name: "master_password", Required: true},
		},
	}
}

// ExportName is related with project_name, environ_name, instance_name
// ExportValue
//  - master_username
//  - master_password
//  - endpoint_write
//  - endpoint_read
func (b *DatabaseBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	out := make(map[string]string)
	out["endpoint_write"] = "${environ}-${instance_name}.write.${project_name}.${domain}"
	out["endpoint_read"] = "${environ}-${instance_name}.read.${project_name}.${domain}"

	return &broker.CreateOutput{
		Status:  202,
		Message: "Accepted",
		Output:  out,
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

func (b *DatabaseBroker) Status(in *broker.StatusInput) *broker.StatusOutput {
	return &broker.StatusOutput{}
}
