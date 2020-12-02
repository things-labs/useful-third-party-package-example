package main

import (
	"embed"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed static
var local embed.FS

func main() {
	e := gin.New()
	e.GET("/*filepath", gin.WrapH(http.FileServer(http.FS(local))))
	log.Println(e.Run(":8989"))
}
