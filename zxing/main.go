package main

import (
	"fmt"
	"log"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/makiuchi-d/gozxing"
)

func main() {
	br, err := qr.Encode("hello world", qr.M, qr.Auto)
	if err != nil {
		log.Fatal(err)
	}

	br, err = barcode.Scale(br, 25, 25)
	if err != nil {
		log.Fatal(err)
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(br)
	if err != nil {
		log.Fatal(err)
	}
	bmxt, _ := bmp.GetBlackMatrix()
	fmt.Println(bmxt.ToString("██", "  "))
}
