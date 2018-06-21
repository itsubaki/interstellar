package launcher

type Launcher interface {
	Config() *Config
	Register(in *RegisterInput) *RegisterOutput
	List(in *ListInput) *ListOutput
}

type Config struct {
	Port string
}

type RegisterInput struct {
	CatalogURL string
}

type RegisterOutput struct {
	Status    int
	Message   string
	ServiceID string
	Input     *RegisterInput
}

type ListInput struct {
}

type ListOutput struct {
}
