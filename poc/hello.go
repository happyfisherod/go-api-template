package main

import (
	"fmt"
	"github.com/geometry-labs/protoserdes/mongo"
)

func main() {
	fmt.Println("starting...")
	//serdes.JsonToProto() // json to proto poc
	mongo.MongoPoc()

}
