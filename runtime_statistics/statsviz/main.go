package main

import (
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/arl/statsviz"
	"github.com/gin-gonic/gin"
	"github.com/pkg/browser"
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

	go func() {
		time.Sleep(time.Second)
		err := browser.OpenURL("http://localhost:8080/debug/statsviz")
		if err != nil {
			log.Fatal(err)
		}
	}()

	router := gin.New()
	router.GET("/debug/statsviz/*filepath", func(c *gin.Context) {
		if c.Param("filepath") == "/ws" {
			statsviz.Ws(c.Writer, c.Request)
			return
		}
		statsviz.IndexAtRoot("/debug/statsviz").ServeHTTP(c.Writer, c.Request)
	})
	router.Run(":8080")
}
