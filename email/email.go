package main

import (
	"net/smtp"
	"net/textproto"
	"time"

	"github.com/jordan-wright/email"
)

func main() {
	e := &email.Email{
		To:      []string{"jgb40@qq.com"},
		From:    "jiang <thinkgo@aliyun.com>",
		Subject: "Awesome Subject",
		Text:    []byte("Text Body is, of course, supported!"),
		// HTML:    []byte("<h1>Fancy HTML is supported, too!</h1>"),
		Headers: textproto.MIMEHeader{},
	}

	// err := e.Send("smtp.aliyun.com:25",
	// 	smtp.PlainAuth("", "youremail@aliyun.com", "yourpassword", "smtp.aliyun.com"))
	// log.Println(err)

	pool, err := email.NewPool("smtp.aliyun.com:25", 10, smtp.PlainAuth("", "thinkgo@aliyun.com", "yourpassword!", "smtp.aliyun.com"))
	if err != nil {
		panic(err)
	}
	err = pool.Send(e, time.Second*5)
	if err != nil {
		panic(err)
	}
}
