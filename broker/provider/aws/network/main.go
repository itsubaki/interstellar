package network

import "github.com/itsubaki/interstellar/broker"

func main() {
	broker.Run(NewNetworkBroker())
}
