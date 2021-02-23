package main

import (
	"fmt"
	"io"

	"github.com/nxadm/tail"
)

func main() {
	t, err := tail.TailFile("tail.log", tail.Config{
		Location: &tail.SeekInfo{0, io.SeekEnd},
		Follow:   true,
		ReOpen:   true,
	})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		fmt.Printf(line.Text)
	}
}
