package main

import (
	_ "embed"
	"net/http"

	"github.com/gin-gonic/gin"
)

//go:embed static/logo.jpg
var content []byte

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "image/png", content)
	})

	router.Run(":8989")
}
