package controller

import "github.com/itsubaki/interstellar/broker"

type ServiceController interface {
	Config() *Config

	Register(in *RegisterInput) *RegisterOutput
	Service() *ServiceOutput
	Catalog(id string) *CatalogOutput

	Instance() *InstanceOutput
	CreateInstance(in *CreateInstanceInput) *CreateInstanceOutput
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
	Status  int        `json:"status"`
	Service []*Service `json:"service"`
}

type CatalogOutput struct {
	Status    int             `json:"status"`
	Message   string          `json:"message,omitempty"`
	ServiceID string          `json:"service_id,omitempty"`
	Catalog   *broker.Catalog `json:"catalog,omitempty"`
}

type InstanceOutput struct {
	Status   int         `json:"status"`
	Message  string      `json:"message,omitempty"`
	Instance []*Instance `json:"instance,omitempty"`
}

type CreateInstanceInput struct {
	Name      string            `json:"name"`
	ServiceID string            `json:"service_id"`
	Parameter map[string]string `json:"parameter"`
}

type CreateInstanceOutput struct {
	Status   int       `json:"status"`
	Message  string    `json:"message,omitempty"`
	Instance *Instance `json:"instance,omitempty"`
}
