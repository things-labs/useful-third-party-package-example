package main

import (
	"fmt"
	"log"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

func main() {
	b, err := qrcode.NewQRCodeWriter().EncodeWithoutHint("hello world", gozxing.BarcodeFormat_QR_CODE, 30, 30)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b.ToString("██", "  "))
}
