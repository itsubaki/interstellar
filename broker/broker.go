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

	if err := g.Run(b.Config().Port); err != nil {
		log.Fatalf("run broker: %v", err)
	}
}
