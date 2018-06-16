package launcher

type Launcher struct {
}

func NewLauncher(conf *Config) *Launcher {
	return &Launcher{}
}

func (l *Launcher) Run() error {
	return nil
}

func (l *Launcher) Register(in *RegisterInput) *RegisterOutput {
	return &RegisterOutput{}
}

func (l *Launcher) List(in *ListInput) *ListOutput {
	return &ListOutput{}
}
