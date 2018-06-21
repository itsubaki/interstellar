package main

import "github.com/itsubaki/interstellar/launcher"

type Interstellar struct {
}

func NewInterstellar() *Interstellar {
	return &Interstellar{}
}

func (i *Interstellar) Config() *launcher.Config {
	return &launcher.Config{
		Port: ":8080",
	}
}

func (i *Interstellar) Register(in *launcher.RegisterInput) *launcher.RegisterOutput {
	return &launcher.RegisterOutput{}
}

func (i *Interstellar) List(in *launcher.ListInput) *launcher.ListOutput {
	return &launcher.ListOutput{}
}
