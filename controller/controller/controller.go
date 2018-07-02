package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/itsubaki/env"
	"github.com/itsubaki/interstellar/broker"
	"github.com/itsubaki/interstellar/controller"
	"github.com/itsubaki/interstellar/controller/repository"
)

type Controller struct {
	service  repository.ServiceRepository
	catalog  repository.CatalogRepository
	instance repository.InstanceRepository
}

func NewController() *Controller {
	return &Controller{}
}

func (c *Controller) Config() *controller.Config {
	return &controller.Config{
		Port: env.GetValue("PORT", ":8080"),
	}
}

func (c *Controller) Service() *controller.ServiceOutput {
	all := c.service.SelectAll()
	service := []controller.Service{}
	for i := range all {
		service = append(service, *all[i])
	}

	return &controller.ServiceOutput{
		Status:  http.StatusOK,
		Service: service,
	}
}

func (c *Controller) Catalog(in *controller.CatalogInput) *controller.CatalogOutput {
	s, ok := c.service.FindByID(in.ServiceID)
	if !ok {
		return &controller.CatalogOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("service=%s not found", in.ServiceID),
		}
	}

	catalog, ok := c.catalog.FindByName(s.Name)
	if !ok {
		return &controller.CatalogOutput{
			Status:    http.StatusBadRequest,
			ServiceID: s.ServiceID,
			Message:   fmt.Sprintf("catalog=%s not found", in.ServiceID),
		}
	}

	return &controller.CatalogOutput{
		Status:    http.StatusOK,
		ServiceID: s.ServiceID,
		Catalog:   catalog,
	}

}

func (c *Controller) Instance() *controller.InstanceOutput {
	all := c.instance.SelectAll()
	instance := []controller.Instance{}
	for i := range all {
		instance = append(instance, *all[i])
	}

	return &controller.InstanceOutput{
		Status:   http.StatusOK,
		Instance: instance,
	}
}

func (c *Controller) Create(in *controller.CreateInput) *controller.CreateOutput {
	s, ok := c.service.FindByID(in.ServiceID)
	if !ok {
		return &controller.CreateOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("service=%s not found", in.ServiceID),
		}
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return &controller.CreateOutput{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("new uuid: %v", err),
		}
	}
	instanceID := uuid.String()

	input := &broker.CreateInput{
		Parameter: in.Parameter,
	}

	jb, err := json.Marshal(input)
	if err != nil {
		return &controller.CreateOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("unmarshal request body: %v", err),
		}
	}

	out, err := http.Post(
		fmt.Sprintf("%s/v1/service/%s", s.ServiceBrokerURL, instanceID),
		"application/json",
		bytes.NewReader(jb),
	)
	if err != nil {
		return &controller.CreateOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("%v", err),
		}
	}

	b, err := ioutil.ReadAll(out.Body)
	if err != nil {
		return &controller.CreateOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("read request body: %v", err),
		}
	}
	defer out.Body.Close()

	var res broker.CreateOutput
	if uerr := json.Unmarshal(b, &res); uerr != nil {
		return &controller.CreateOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("unmarshal request body: %v", uerr),
		}
	}

	i := &controller.Instance{
		Name:       in.Name,
		InstanceID: instanceID,
		ServiceID:  s.ServiceID,
		Parameter:  in.Parameter,
	}

	if res.Status == http.StatusOK || res.Status == http.StatusCreated {
		i.Output = res.Instance.Output
		c.instance.Insert(i)
		return &controller.CreateOutput{
			Status:   res.Status,
			Message:  res.Message,
			Instance: i,
		}
	}

	if res.Status == http.StatusAccepted {
		c.instance.Insert(i)
		return &controller.CreateOutput{
			Status:   res.Status,
			Message:  res.Message,
			Instance: i,
		}
	}

	return &controller.CreateOutput{
		Status:  res.Status,
		Message: res.Message,
	}
}

func (c *Controller) Describe(in *controller.DescribeInput) *controller.DescribeOutput {
	i, ok := c.instance.FindByID(in.InstanceID)
	if !ok {
		return &controller.DescribeOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("instance=%s not found", in.InstanceID),
		}
	}

	s, ok := c.service.FindByID(i.ServiceID)
	if !ok {
		return &controller.DescribeOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("service=%s not found", i.ServiceID),
		}
	}

	out, err := http.Get(fmt.Sprintf("%s/v1/service/%s", s.ServiceBrokerURL, in.InstanceID))
	if err != nil {
		return &controller.DescribeOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("%v", err),
		}
	}

	b, err := ioutil.ReadAll(out.Body)
	if err != nil {
		return &controller.DescribeOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("read request body: %v", err),
		}
	}
	defer out.Body.Close()

	var res broker.DescribeOutput
	if uerr := json.Unmarshal(b, &res); uerr != nil {
		return &controller.DescribeOutput{
			Status:  http.StatusBadRequest,
			Message: fmt.Sprintf("unmarshal request body: %v", uerr),
		}
	}

	return &controller.DescribeOutput{
		Status:  res.Status,
		Message: res.Message,
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

	if s, ok := c.service.FindByName(res.Name); ok {
		return &controller.RegisterOutput{
			Status:    http.StatusConflict,
			ServiceID: s.ServiceID,
			Message:   fmt.Sprintf("%s already exists", res.Name),
		}
	}

	uuid, err := uuid.NewUUID()
	if err != nil {
		return &controller.RegisterOutput{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("new uuid: %v", err),
		}
	}

	c.catalog.Insert(&res)
	c.service.Insert(&controller.Service{
		Name:             res.Name,
		ServiceID:        uuid.String(),
		ServiceBrokerURL: in.URL,
	})

	return &controller.RegisterOutput{
		Status:    http.StatusOK,
		ServiceID: uuid.String(),
	}
}
