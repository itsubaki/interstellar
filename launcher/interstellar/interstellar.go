package interstellar

import "github.com/itsubaki/interstellar/launcher"

type Interstellar struct {
}

func NewInterstellar(conf *Config) *Interstellar {
	return &Interstellar{}
}

func (i *Interstellar) Register(in *launcher.RegisterInput) *launcher.RegisterOutput {
	return &launcher.RegisterOutput{}
}

func (i *Interstellar) List(in *launcher.ListInput) *launcher.ListOutput {
	return &launcher.ListOutput{}
}
