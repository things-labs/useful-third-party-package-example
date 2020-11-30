package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, fmt.Sprintf("hello world, version V%d", 4))
	})
	http.ListenAndServe(":80", nil)
}
