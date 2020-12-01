package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
)

func main() {
	// Force the GC to work to make the plots "move".
	go func() {
		m := map[string][]byte{}

		for {
			b := make([]byte, 512+rand.Intn(16*1024))
			m[strconv.Itoa(len(m)%(10*100))] = b

			if len(m)%(10*100) == 0 {
				m = make(map[string][]byte)
			}

			time.Sleep(10 * time.Millisecond)
		}
	}()

	router := gin.New()
	router.GET("/debug/statsviz/ws", func(c *gin.Context) { statsviz.Ws(c.Writer, c.Request) })
	router.GET("/debug/statsviz/statsviz", func(c *gin.Context) { statsviz.Index.ServeHTTP(c.Writer, c.Request) })
	router.Run(":8080")
}
