package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/things-go/cdpui"
)

func main() {
	ui := cdpui.New("http://localhost:8080")
	defer ui.Close()

	go func() {
		http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprint(w, indexHTML)
		})
		http.ListenAndServe(":8080", nil)
	}()
	time.Sleep(time.Millisecond * 5)
	ui.Run()
	<-ui.Wait()
}

const indexHTML = `<!doctype html>
<html>
<head>
  <title>example</title>
</head>
<body>
  <div id="box3">
    <h2>box3</h3>
    <p id="box4">
      box4 text
      <input id="input1" value="some value"><br><br>
      <textarea id="textarea1" style="width:500px;height:400px">textarea</textarea><br><br>
      <input id="input2" type="submit" value="Next">
      <select id="select1">
        <option value="one">1</option>
        <option value="two">2</option>
        <option value="three">3</option>
        <option value="four">4</option>
      </select>
    </p>
  </div>
</body>
</html>`
