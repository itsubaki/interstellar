package launcher

import (
	"log"

	"github.com/gin-gonic/gin"
)

func Run(l Launcher) {
	g := gin.New()

	if err := g.Run(l.Config().Port); err != nil {
		log.Fatalf("run broker: %v", err)
	}
}
