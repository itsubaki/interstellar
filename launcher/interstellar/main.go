package interstellar

import (
	"log"

	"github.com/itsubaki/interstellar/launcher"
)

func main() {
	c := NewConfig()
	log.Printf("config=%v\n", c)

	i := NewInterstellar(c)
	if err := launcher.Run(i); err != nil {
		log.Println(err)
	}
}
