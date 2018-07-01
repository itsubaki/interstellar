package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/itsubaki/env"
	"github.com/itsubaki/interstellar/broker"
)

type ProjectBroker struct {
	config   *broker.Config
	template string
}

func NewProjectBroker() (*ProjectBroker, error) {
	c := &broker.Config{
		Port:     env.GetValue("PORT", ":8080"),
		Template: env.GetValue("TEMPLATE", "./template.yml"),
	}

	f, err := ioutil.ReadFile(c.Template)
	if err != nil {
		return nil, fmt.Errorf("read file: %v", err)
	}

	return &ProjectBroker{
		config:   c,
		template: string(f),
	}, nil
}

func (b *ProjectBroker) Config() *broker.Config {
	return b.config
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
		ParameterSpec: []*broker.ParamSpec{
			{Name: "integration_role_arn", Required: false},
			{Name: "region", Required: true},

			{Name: "project_name", Required: true},
			{Name: "domain", Required: true},
			{Name: "aws_account_id", Required: true},
			{Name: "cidr", Required: true},
		},
	}
}

func (b *ProjectBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
	sess := session.Must(session.NewSession())
	cfn := cloudformation.New(
		sess,
		&aws.Config{
			Credentials: stscreds.NewCredentials(sess, in.Parameter["integration_role_arn"]),
			Region:      aws.String(in.Parameter["region"]),
		})

	param := []*cloudformation.Parameter{
		{ParameterKey: aws.String("ProjectName"), ParameterValue: aws.String(in.Parameter["project_name"])},
		{ParameterKey: aws.String("DomainName"), ParameterValue: aws.String(in.Parameter["domain"])},
		{ParameterKey: aws.String("AccountId"), ParameterValue: aws.String(in.Parameter["aws_account_id"])},
		{ParameterKey: aws.String("Region"), ParameterValue: aws.String(in.Parameter["region"])},
	}

	input := &cloudformation.CreateStackInput{
		StackName:    &in.InstanceID,
		Parameters:   param,
		TemplateBody: &b.template,
	}

	if _, err := cfn.CreateStack(input); err != nil {
		return &broker.CreateOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("create stack: %v", err),
		}
	}

	return &broker.CreateOutput{
		Status:  http.StatusAccepted,
		Message: "Accepted",
	}
}

func (b *ProjectBroker) Delete(in *broker.DeleteInput) *broker.DeleteOutput {
	return &broker.DeleteOutput{
		Status:  http.StatusAccepted,
		Message: "Accepted",
	}
}

func (b *ProjectBroker) Update(in *broker.UpdateInput) *broker.UpdateOutput {
	return &broker.UpdateOutput{
		Status:  http.StatusAccepted,
		Message: "Accepted",
	}
}

func (b *ProjectBroker) Binding(in *broker.BindingInput) *broker.BindingOutput {
	return &broker.BindingOutput{
		Status:  http.StatusAccepted,
		Message: "Accepted",
	}
}

func (b *ProjectBroker) Unbinding(in *broker.UnbindingInput) *broker.UnbindingOutput {
	return &broker.UnbindingOutput{
		Status:  http.StatusAccepted,
		Message: "Accepted",
	}
}

func (b *ProjectBroker) Describe(in *broker.DescribeInput) *broker.DescribeOutput {
	sess := session.Must(session.NewSession())
	cfn := cloudformation.New(
		sess,
		&aws.Config{
			Credentials: stscreds.NewCredentials(sess, in.Parameter["integration_role_arn"]),
			Region:      aws.String(in.Parameter["region"]),
		})

	input := &cloudformation.DescribeStacksInput{
		StackName: &in.InstanceID,
	}

	desc, err := cfn.DescribeStacks(input)
	if err != nil {
		return &broker.DescribeOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("describe stack: %v", err),
		}
	}
	s0 := desc.Stacks[0]

	out := make(map[string]string)
	for i := range s0.Outputs {
		o := s0.Outputs[i]
		out[*o.OutputKey] = *o.OutputValue
	}

	return &broker.DescribeOutput{
		Status:  http.StatusOK,
		Message: fmt.Sprintf("%s", *s0.StackStatus),
		Output:  out,
	}
}
