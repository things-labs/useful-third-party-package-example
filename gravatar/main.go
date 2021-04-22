package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/thinkgos/x/extos"
	"github.com/thinkgos/x/lib/algo"
)

// 获取头像
func Gravatar(email string, size uint16, isHttps bool) string {
	gravatarDomain := "http://gravatar.com/avatar"
	if isHttps {
		gravatarDomain = "https://secure.gravatar.com/avatar"
	}
	return fmt.Sprintf("%s/%s?s=%d", gravatarDomain, algo.MD5(email), size)
}

func main() {
	ul := Gravatar("jgb40@aliyun.com", 40, true)
	log.Println(ul)
	rsp, err := http.Get(ul)
	if err != nil {
		log.Println(err)
		return
	}
	defer rsp.Body.Close()

	b, err := io.ReadAll(rsp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	extos.WriteFile("./test.jpg", b)
}
