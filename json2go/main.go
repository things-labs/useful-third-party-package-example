package main

import (
	"fmt"

	"github.com/m-zajac/json2go"
)

func main() {
	v := `{"id": 123, "uid": 112, "value": false}`
	parse := json2go.NewJSONParser("Document")
	err := parse.FeedBytes([]byte(v))
	if err != nil {
		panic(err)
	}
	fmt.Println(parse.String())
}
