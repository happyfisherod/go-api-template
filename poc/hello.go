package main

import (
	"fmt"
	"github.com/geometry-labs/protoserdes/postgres"
)

func main() {

	fmt.Println("starting...")
	//serdes.JsonToProto() // json to proto poc
	//mongo.MongoPoc()
	postgres.PostgresPoc()
	//kafka.ProducerPoc()
	//kafka.ConsumerPoc()
}
