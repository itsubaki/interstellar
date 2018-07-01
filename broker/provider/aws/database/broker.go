package main

import (
	"github.com/itsubaki/env"
	"github.com/itsubaki/interstellar/broker"
)

type DatabaseBroker struct {
}

func NewDatabaseBroker() *DatabaseBroker {
	return &DatabaseBroker{}
}

func (b *DatabaseBroker) Config() *broker.Config {
	return &broker.Config{
		Port: env.GetValue("PORT", ":8080"),
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
		ParameterSpec: []*broker.ParamSpec{
			{Name: "integration_role_arn", Required: false},
			{Name: "region", Required: true},

			{Name: "project_name", Required: true},
			{Name: "environ_name", Required: true},
			{Name: "instance_name", Required: true},
			{Name: "master_username", Required: true},
			{Name: "master_password", Required: true},
		},
	}
}

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
	return &broker.BindingOutput{}
}

func (b *DatabaseBroker) Unbinding(in *broker.UnbindingInput) *broker.UnbindingOutput {
	return &broker.UnbindingOutput{}
}

func (b *DatabaseBroker) Describe(in *broker.DescribeInput) *broker.DescribeOutput {
	return &broker.DescribeOutput{}
}
