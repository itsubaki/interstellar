package compute

import "github.com/itsubaki/interstellar/broker"

func main() {
	broker.Run(NewComputeBroker())
}
