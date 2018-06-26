package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/itsubaki/interstellar/broker"
	"github.com/itsubaki/interstellar/controller"
	"github.com/itsubaki/interstellar/util"
)

type Controller struct {
	ServiceStore []controller.Service
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Config() *controller.Config {
	return &controller.Config{
		Port: util.Getenv("PORT", ":8080"),
	}
}

func (c *Controller) Service() *controller.ServiceOutput {
	return &controller.ServiceOutput{
		Service: c.ServiceStore,
	}
}

func (c *Controller) Register(in *controller.RegisterInput) *controller.RegisterOutput {
	out, err := http.Get(fmt.Sprintf("%s/v1/catalog", in.URL))
	if err != nil {
		return &controller.RegisterOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("%v", err),
		}
	}

	b, err := ioutil.ReadAll(out.Body)
	if err != nil {
		return &controller.RegisterOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("read request body: %v", err),
		}
	}
	defer out.Body.Close()

	var res broker.Catalog
	if uerr := json.Unmarshal(b, &res); uerr != nil {
		return &controller.RegisterOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("unmarshal request body: %v", uerr),
		}
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return &controller.RegisterOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("new uuid: %v", err),
		}
	}

	c.ServiceStore = append(c.ServiceStore, controller.Service{
		Name:             res.Name,
		ServiceID:        uuid.String(),
		ServiceBrokerURL: in.URL,
	})

	return &controller.RegisterOutput{
		Status:    http.StatusOK,
		ServiceID: uuid.String(),
		Message:   fmt.Sprintf("%v", res),
	}
}
