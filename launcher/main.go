package launcher

import "log"

func main() {
	c := NewConfig()
	log.Printf("config=%v\n", c)

	l := NewLauncher(c)
	if err := l.Run(); err != nil {
		log.Println(err)
	}
}
