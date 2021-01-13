package main

import (
	"fmt"

	"github.com/nats-io/nuid"
	"github.com/rs/xid"
)

func main() {
	fmt.Println(nuid.Next())
	fmt.Println(nuid.Next())
	fmt.Println(xid.New().String())
	fmt.Println(xid.New().String())
}
