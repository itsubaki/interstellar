package controller

type ServiceController interface {
	Config() *Config
	Register(in *RegisterInput) *RegisterOutput
	Service() *ServiceOutput
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
	Service []Service `json:"service"`
}

type Service struct {
	Name             string `json:"name"`
	ServiceID        string `json:"service_id"`
	ServiceBrokerURL string `json:"service_broker_url"`
}
