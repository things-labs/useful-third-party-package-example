package main

import (
	"log"
	"net/smtp"
	"net/textproto"

	"github.com/jordan-wright/email"
)

func main() {
	e := &email.Email{
		To:      []string{"destexmail@aliyun.com"},
		From:    "jiang <youremail@aliyun.com>",
		Subject: "Awesome Subject",
		Text:    []byte("Text Body is, of course, supported!"),
		// HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
		Headers: textproto.MIMEHeader{},
	}
	err := e.Send("smtp.aliyun.com:25",
		smtp.PlainAuth("", "youremail@aliyun.com", "yourpassword", "smtp.aliyun.com"))
	log.Println(err)
}
