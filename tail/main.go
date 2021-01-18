package main

import (
	"fmt"

	"github.com/nxadm/tail"
)

func main() {
	t, err := tail.TailFile("./tail.log", tail.Config{
		Follow: true,
		ReOpen: true,
	})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		fmt.Printf(line.Text)
	}
}
