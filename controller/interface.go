package controller

import "github.com/itsubaki/interstellar/broker"

type ServiceController interface {
	Config() *Config

	Register(in *RegisterInput) *RegisterOutput
	Service() *ServiceOutput
	Catalog(in *CatalogInput) *CatalogOutput

	Instance() *InstanceOutput
	Describe(in *DescribeInput) *DescribeOutput
	Create(in *CreateInput) *CreateOutput
}

type Config struct {
	Port string
}

type Service struct {
	Name             string `json:"name"`
	ServiceID        string `json:"service_id"`
	ServiceBrokerURL string `json:"service_broker_url"`
}

type Instance struct {
	Name       string            `json:"name"`
	ServiceID  string            `json:"service_id"`
	InstanceID string            `json:"instance_id"`
	Parameter  map[string]string `json:"parameter"`
	Output     map[string]string `json:"output,omitempty"`
}

type RegisterInput struct {
	URL string `json:"url"`
}

type RegisterOutput struct {
	Status    int    `json:"status"`
	Message   string `json:"message,omitempty"`
	ServiceID string `json:"service_id,omitempty"`
}

type ServiceOutput struct {
	Status  int       `json:"status"`
	Service []Service `json:"service"`
}

type CatalogInput struct {
	ServiceID string `json:"service_id"`
}

type CatalogOutput struct {
	Status    int             `json:"status"`
	Message   string          `json:"message,omitempty"`
	ServiceID string          `json:"service_id,omitempty"`
	Catalog   *broker.Catalog `json:"catalog,omitempty"`
}

type InstanceOutput struct {
	Status   int        `json:"status"`
	Message  string     `json:"message,omitempty"`
	Instance []Instance `json:"instance"`
}

type CreateInput struct {
	Name      string            `json:"name"`
	ServiceID string            `json:"service_id"`
	Parameter map[string]string `json:"parameter"`
}

type CreateOutput struct {
	Status   int       `json:"status"`
	Message  string    `json:"message,omitempty"`
	Instance *Instance `json:"instance,omitempty"`
}

type DescribeInput struct {
	InstanceID string `json:"instance_id"`
}

type DescribeOutput struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}
