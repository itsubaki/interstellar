package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

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
			{Name: "aws_account_id", Required: true},
			{Name: "integration_role_arn", Required: false},
			{Name: "project_name", Required: true},
			{Name: "cidr", Required: true},
			{Name: "domain", Required: true},
		},
	}
}

func (b *ProjectBroker) Create(in *broker.CreateInput) *broker.CreateOutput {
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

func (b *ProjectBroker) Status(in *broker.StatusInput) *broker.StatusOutput {
	out := make(map[string]string)
	out["nameserver"] = "ns-1,ns-2,ns-3,ns-4"

	return &broker.StatusOutput{
		Status: http.StatusOK,
		Output: out,
	}
}
