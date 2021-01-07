package main

import (
	"bytes"
	"log"

	"github.com/guonaihong/gout"
)

// 用于解析 服务端 返回的http body
type RspBody struct {
	ErrMsg  string `json:"errmsg"`
	ErrCode int    `json:"errcode"`
	Data    string `json:"data"`
}

// 用于解析 服务端 返回的http header
type RspHeader struct {
	Sid  string `header:"sid"`
	Time int    `header:"time"`
}

func main() {
	buf := &bytes.Buffer{}
	code := 0
	err := gout.
		GET("www.baidu.com"). // POST请求
		Debug(true).          // 打开debug模式
		BindBody(buf).        // BindJSON解析返回的body内容,同类函数有BindBody, BindYAML, BindXML
		// http code
		Code(&code).
		// 结束函数
		Do()

	if err != nil { // 判度错误
		log.Fatalf("send fail:%s\n", err)
	}
	log.Println(buf.String())
}
