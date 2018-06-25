package launcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
)

func Run(l Launcher) {
	g := gin.New()

	g.POST("/v1/register", func(c *gin.Context) {
		b, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(500, fmt.Errorf("read request body: %v", err))
			return
		}
		defer c.Request.Body.Close()

		var req RegisterInput
		if err := json.Unmarshal(b, &req); err != nil {
			c.JSON(500, fmt.Errorf("unmarshal request body: %v", err))
			return
		}

		out := l.Register(&req)
		c.JSON(out.Status, out)
	})

	if err := g.Run(l.Config().Port); err != nil {
		log.Fatalf("run broker: %v", err)
	}
}
