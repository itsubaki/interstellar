package main

import "github.com/itsubaki/interstellar/broker"

type ProjectBroker struct {
}

func NewProjectBroker() *ProjectBroker {
	return &ProjectBroker{}
}

func (b *ProjectBroker) Config() *broker.Config {
	return &broker.Config{
		Port: ":9083",
	}
}

func (b *ProjectBroker) Catalog() *broker.Catalog {
	return &broker.Catalog{
		Name: "aws_project",
		Tag: []string{
			"aws",
			"project",
		},
		Require:  []string{},
		Bindable: false,
	}
}

func (b *ProjectBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	// in.Parameter["integration_role_arn"]
	// in.Parameter["project_name"]
	// in.Parameter["vpc_cidr"]
	// in.Parameter["domain"]
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
			{Key: "project_name", Value: "${project_name}"},
			{Key: "domain", Value: "${domain}"},
			{Key: "hostedzone", Value: "${project_name}.${domain}."},
			{Key: "nameserver", Value: "ns-1,ns-2,ns-3,ns-4"},
			{Key: "certificate", Value: "acm-12345678"},
			{Key: "bucket_log", Value: "s3://log.${project_name}.${domain}"},
			{Key: "bucket_deploy", Value: "s3://deploy.${project_name}.${domain}"},
			{Key: "bucket_config", Value: "s3://config.${project_name}.${domain}"},
		},
	}
}

func (b *ProjectBroker) Delete(in *broker.DeleteInput) *broker.DeleteOutput {
	return &broker.DeleteOutput{
		Status:  202,
		Message: "Accepted",
		Input:   in,
	}
}

func (b *ProjectBroker) Update(in *broker.UpdateInput) *broker.UpdateOutput {
	return &broker.UpdateOutput{}
}

func (b *ProjectBroker) Binding(in *broker.BindingInput) *broker.BindingOutput {
	return &broker.BindingOutput{}
}

func (b *ProjectBroker) Unbinding(in *broker.UnbindingInput) *broker.UnbindingOutput {
	return &broker.UnbindingOutput{}
}
