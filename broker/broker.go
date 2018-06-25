package broker

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Run(b ServiceBroker) {
	g := gin.New()

	g.GET("/v1/catalog", func(c *gin.Context) {
		c.JSON(200, b.Catalog())
	})

	g.PUT("/v1/service/:instance_id", func(c *gin.Context) {
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

	log.Printf("%v\n", b.Config())
	if err := g.Run(b.Config().Port); err != nil {
		log.Fatalf("run broker: %v", err)
	}
}
