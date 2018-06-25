package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/itsubaki/interstellar/broker"
	"github.com/itsubaki/interstellar/launcher"
	"github.com/itsubaki/interstellar/util"
)

type Interstellar struct {
}

func NewInterstellar() *Interstellar {
	return &Interstellar{}
}

func (i *Interstellar) Config() *launcher.Config {
	return &launcher.Config{
		Port: util.Getenv("PORT", ":8080"),
	}
}

func (i *Interstellar) Register(in *launcher.RegisterInput) *launcher.RegisterOutput {
	out, err := http.Get(in.CatalogURL)
	if err != nil {
		return &launcher.RegisterOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("%v", err),
		}
	}

	b, err := ioutil.ReadAll(out.Body)
	if err != nil {
		return &launcher.RegisterOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("read request body: %v", err),
		}
	}
	defer out.Body.Close()

	var res broker.Catalog
	if uerr := json.Unmarshal(b, &res); uerr != nil {
		return &launcher.RegisterOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("unmarshal request body: %v", uerr),
		}
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return &launcher.RegisterOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("new uuid: %v", err),
		}
	}

	return &launcher.RegisterOutput{
		Status:    http.StatusOK,
		ServiceID: uuid.String(),
		Message:   fmt.Sprintf("%v", res),
	}
}
