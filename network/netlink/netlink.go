package main

import (
	"log"

	"github.com/vishvananda/netlink"
)

func main() {
	b, err := netlink.BridgeVlanList()
	if err != nil {
		panic(err)
	}

	log.Println(b)
}
