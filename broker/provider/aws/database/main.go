package database

import "github.com/itsubaki/interstellar/broker"

func main() {
	broker.Run(NewDatabaseBroker())
}
