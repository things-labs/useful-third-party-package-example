package main

import (
	"net/http"

	"github.com/gin-contrib/httpsign"
	"github.com/gin-contrib/httpsign/crypto"
	"github.com/gin-gonic/gin"
)

func main() {
	// Define algorithm
	hmacsha256 := &crypto.HmacSha256{}
	hmacsha512 := &crypto.HmacSha512{}
	// Init define secret params
	readKeyID := httpsign.KeyID("read")
	writeKeyID := httpsign.KeyID("write")
	secrets := httpsign.Secrets{
		readKeyID: &httpsign.Secret{
			Key:       "HMACSHA256-SecretKey",
			Algorithm: hmacsha256, // You could using other algo with interface Crypto
		},
		writeKeyID: &httpsign.Secret{
			Key:       "HMACSHA512-SecretKey",
			Algorithm: hmacsha512,
		},
	}

	// Init server
	r := gin.Default()

	// Create middleware with default rule. Could modify by parse Option func
	auth := httpsign.NewAuthenticator(secrets)

	r.Use(auth.Authenticated())
	r.GET("/a", func(c *gin.Context) {
		c.String(http.StatusOK, "a")
	})
	r.POST("/b", func(c *gin.Context) {
		c.String(http.StatusOK, "b")
	})
	r.POST("/c", func(c *gin.Context) {
		c.String(http.StatusOK, "c")
	})

	r.Run(":8080")
}
