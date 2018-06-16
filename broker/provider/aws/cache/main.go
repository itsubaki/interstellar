package cache

import "github.com/itsubaki/interstellar/broker"

func main() {
	broker.Run(NewCacheBroker())
}
