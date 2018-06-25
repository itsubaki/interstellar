package launcher

type Launcher interface {
	Config() *Config
	Register(in *RegisterInput) *RegisterOutput
}

type Config struct {
	Port string
}

type RegisterInput struct {
	CatalogURL string `json:"catalog_url"`
}

type RegisterOutput struct {
	Status    int            `json:"status"`
	Message   string         `json:"message"`
	ServiceID string         `json:"service_id"`
	Input     *RegisterInput `json:"input"`
}
