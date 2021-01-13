package main

import (
	"fmt"
	"log"

	"github.com/skip2/go-qrcode"
)

func main() {
	qr, err := qrcode.New("hello world", qrcode.Medium)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(qr.ToSmallString(false))
}
