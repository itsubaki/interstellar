package broker

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func Run(b ServiceBroker) {
	g := gin.New()

	g.GET("/v1/catalog", func(c *gin.Context) {
		c.JSON(200, b.Catalog())
	})

	g.GET("/v1/service/:instance_id", func(c *gin.Context) {
		in := &DescribeInput{
			InstanceID: c.Param("instance_id"),
		}
		out := b.Describe(in)
		c.JSON(out.Status, out)
	})

	g.POST("/v1/service/:instance_id", func(c *gin.Context) {
		in := &CreateInput{
			InstanceID: c.Param("instance_id"),
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
