package main

import (
	"fmt"
	"log"

	"github.com/bwmarrin/snowflake"
	"github.com/nats-io/nuid"
	"github.com/rs/xid"
)

func main() {
	fmt.Println(nuid.Next())
	fmt.Println(nuid.Next())
	fmt.Println(xid.New().String())
	fmt.Println(xid.New().String())

	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(node.Generate())
	fmt.Println(node.Generate())

}
