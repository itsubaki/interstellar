package controller

import "github.com/itsubaki/interstellar/broker"

type ServiceController interface {
	Config() *Config
	Register(in *RegisterInput) *RegisterOutput
	Service() *ServiceOutput
	Catalog(id string) *CatalogOutput
}

type Config struct {
	Port string
}

type RegisterInput struct {
	URL string `json:"url"`
}

type RegisterOutput struct {
	Status    int    `json:"status"`
	Message   string `json:"message"`
	ServiceID string `json:"service_id,omitempty"`
}

type ServiceOutput struct {
	Status  int       `json:"status"`
	Service []Service `json:"service"`
}

type Service struct {
	Name             string `json:"name"`
	ServiceID        string `json:"service_id"`
	ServiceBrokerURL string `json:"service_broker_url"`
}

type CatalogOutput struct {
	Status    int            `json:"status"`
	Message   string         `json:"message,omitempty"`
	ServiceID string         `json:"service_id,omitempty"`
	Catalog   broker.Catalog `json:"catalog,omitempty"`
}
