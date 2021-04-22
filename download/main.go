package main

import (
	"fmt"
	"log"
	"os/exec"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	log.Println("----> download starting...")

	cnt := 35
	wg.Add(cnt)
	for i := 1; i <= cnt; i++ {
		num := fmt.Sprintf("%02d", i)
		url := "https://www.bilibili.com/video/BV1iV411t7Vi?p=" + num
		cmd := exec.Command("youtube-dl", url, "-o", "electron-p"+num)
		go func() {
			cmd.Run()
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("----> download done")
}
