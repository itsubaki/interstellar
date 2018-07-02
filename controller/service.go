package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/itsubaki/env"
)

func Run(s ServiceController) {
	g := gin.New()

	path := env.GetValue("INDEX", "./index.html")
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("read html: %v", err)
	}

	g.GET("", func(c *gin.Context) {
		if !gin.IsDebugging() {
			c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
			c.Writer.Write(file)
			return
		}

		file, err = ioutil.ReadFile(path)
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Errorf("read html: %v", err))
			return
		}
		c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		c.Writer.Write(file)
	})

	g.POST("/v1/register", func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Errorf("read request body: %v", err))
			return
		}
		defer c.Request.Body.Close()

		var in RegisterInput
		if err := json.Unmarshal(b, &in); err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Errorf("unmarshal request body: %v", err))
			return
		}

		out := s.Register(&in)
		c.JSON(out.Status, out)
	})

	g.GET("/v1/service", func(c *gin.Context) {
		out := s.Service()
		c.JSON(out.Status, out)
	})

	g.GET("/v1/service/:service_id", func(c *gin.Context) {
		in := &CatalogInput{
			ServiceID: c.Param("service_id"),
		}
		catalog := s.Catalog(in)
		c.JSON(catalog.Status, catalog)
	})

	g.GET("/v1/instance", func(c *gin.Context) {
		out := s.Instance()
		c.JSON(out.Status, out)
	})

	g.GET("/v1/instance/:instance_id", func(c *gin.Context) {
		in := &DescribeInput{
			InstanceID: c.Param("instance_id"),
		}
		out := s.Describe(in)
		c.JSON(out.Status, out)
	})

	g.POST("/v1/instance", func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Errorf("read request body: %v", err))
			return
		}
		defer c.Request.Body.Close()

		var in CreateInput
		if err := json.Unmarshal(b, &in); err != nil {
			c.JSON(http.StatusInternalServerError, fmt.Errorf("unmarshal request body: %v", err))
			return
		}

		out := s.Create(&in)
		c.JSON(out.Status, out)
	})

	log.Printf("config=%v\n", s.Config())
	if err := g.Run(s.Config().Port); err != nil {
		log.Fatalf("run broker: %v", err)
	}
}
