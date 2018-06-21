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
}

type RegisterOutput struct {
}

type ListInput struct {
}

type ListOutput struct {
}
