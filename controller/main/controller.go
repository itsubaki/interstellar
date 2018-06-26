package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/itsubaki/interstellar/broker"
	"github.com/itsubaki/interstellar/controller"
	"github.com/itsubaki/interstellar/controller/repo"
	"github.com/itsubaki/interstellar/util"
)

type Controller struct {
	ServiceRepository repo.ServiceRepository
	CatalogRepository repo.CatalogRepository
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
		Status:  http.StatusOK,
		Service: c.ServiceRepository.SelectAll(),
	}
}

func (c *Controller) Catalog(id string) *controller.CatalogOutput {
	s, ok := c.ServiceRepository.FindByInstanceID(id)
	if !ok {
		return &controller.CatalogOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("service=%s not found", id),
		}
	}

	catalog, ok := c.CatalogRepository.FindByName(s.Name)
	if !ok {
		return &controller.CatalogOutput{
			Status:    http.StatusBadRequest,
			ServiceID: s.ServiceID,
			Message:   fmt.Sprintf("catalog=%s not found", id),
		}
	}

	return &controller.CatalogOutput{
		Status:    http.StatusOK,
		ServiceID: s.ServiceID,
		Catalog:   *catalog,
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

	var catalog broker.Catalog
	if uerr := json.Unmarshal(b, &catalog); uerr != nil {
		return &controller.RegisterOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("unmarshal request body: %v", uerr),
		}
	}

	if s, ok := c.ServiceRepository.FindByName(catalog.Name); ok {
		return &controller.RegisterOutput{
			Status:    http.StatusConflict,
			ServiceID: s.ServiceID,
			Message:   fmt.Sprintf("%s already exists", catalog.Name),
		}
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return &controller.RegisterOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("new uuid: %v", err),
		}
	}

	c.CatalogRepository.Insert(&catalog)
	c.ServiceRepository.Insert(&controller.Service{
		Name:             catalog.Name,
		ServiceID:        uuid.String(),
		ServiceBrokerURL: in.URL,
	})

	return &controller.RegisterOutput{
		Status:    http.StatusOK,
		ServiceID: uuid.String(),
		Message:   fmt.Sprintf("%v", catalog),
	}
}
