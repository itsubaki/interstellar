package launcher

import "log"

func main() {
	conf := NewConfig()
	log.Printf("config=%v\n", conf)

	launcher := NewLauncher(conf)
	if err := launcher.Run(); err != nil {
		log.Println(err)
	}
}
