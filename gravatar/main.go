package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/things-go/x/extos"
)

// 获取头像
func Gravatar(email string, size uint16, isHttps bool) string {
	gravatarDomain := "http://gravatar.com/avatar"
	if isHttps {
		gravatarDomain = "https://secure.gravatar.com/avatar"
	}
	v := md5.Sum([]byte(email))
	return fmt.Sprintf("%s/%s?s=%d", gravatarDomain, string(v[:]), size)
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
