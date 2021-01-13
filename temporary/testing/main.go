package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	B int64 `json:"b,string"`
}

func main() {
	ss := `{"b":"1342368628290957313"}`
	a := A{}
	json.Unmarshal([]byte(ss), &a)

	fmt.Printf("value: %+v", a)
}
