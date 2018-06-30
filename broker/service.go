package broker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Run(b ServiceBroker) {
	g := gin.New()

	g.GET("/v1/catalog", func(c *gin.Context) {
		c.JSON(200, b.Catalog())
	})

	g.POST("/v1/service/:instance_id/describe", func(c *gin.Context) {
		in := &DescribeInput{
			InstanceID: c.Param("instance_id"),
			Parameter:  make(map[string]string),
		}
		out := b.Describe(in)
		c.JSON(out.Status, out)
	})

	g.POST("/v1/service/:instance_id", func(c *gin.Context) {
		bytea, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, fmt.Errorf("read request body: %v", err))
			return
		}
		defer c.Request.Body.Close()

		in := &CreateInput{
			InstanceID: c.Param("instance_id"),
		}
		if err := json.Unmarshal(bytea, &in); err != nil {
			c.JSON(500, fmt.Errorf("unmarshal request body: %v", err))
			return
		}

		out := b.Create(in)
		c.JSON(out.Status, out)
	})

	g.DELETE("/v1/service/:instance_id", func(c *gin.Context) {
		in := &DeleteInput{
			InstanceID: c.Param("instance_id"),
		}
		out := b.Delete(in)
		c.JSON(out.Status, out)
	})

	log.Printf("config=%v\n", b.Config())
	if err := g.Run(b.Config().Port); err != nil {
		log.Printf("run broker: %v", err)
		os.Exit(1)
	}
}
