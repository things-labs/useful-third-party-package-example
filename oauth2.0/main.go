package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

func main() {
	conf := &oauth2.Config{
		ClientID:     "ClientID _id",
		ClientSecret: "ClientSecret",
		Scopes:       []string{"repo", "user"},
		RedirectURL:  "https://thinkgos.cn",
		Endpoint:     github.Endpoint,
	}

	log.Println(conf.AuthCodeURL("abcdefg")) // 1: 通过获取跳到第三方的url,附带重定向url,state等,
	//2: 当用户授权后会重定向到指定的url并附带上code,state等信息
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Println("得到code")
		return
	}
	tok, err := conf.Exchange(context.Background(), "code") // 3: 服务端将得到state,code,验证state并使用code换取token
	if err != nil {
		log.Println(err)
	}

	// 4: 使用 token换取授权的信息
	// 5: 颁发服务器自己的认证token
	client := conf.Client(context.Background(), tok)
	if err := GetUsers(client); err != nil {
		panic(err)
	}

}

// GetUsers 使用oauth2获取用户信息
func GetUsers(client *http.Client) error {
	url := fmt.Sprintf("https://api.github.com/user")

	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("Status Code from", url, ":", resp.StatusCode)
	io.Copy(os.Stdout, resp.Body)
	return nil
}
