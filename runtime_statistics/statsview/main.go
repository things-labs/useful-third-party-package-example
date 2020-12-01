package main

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/go-echarts/statsview"
)

// Visit your browser at http://localhost:18066/debug/statsview
// Or debug as always via http://localhost:18066/debug/pprof, http://localhost:18066/debug/pprof/heap, ...
func main() {
	go func() {
		mgr := statsview.New()

		// Start() runs a HTTP server at `localhost:18066` by default.
		mgr.Start()

		// Stop() will shutdown the http server gracefully
		// mgr.Stop()
	}()

	// busy working....
	// Force the GC to work to make the plots "move".
	m := map[string][]byte{}

	for {
		b := make([]byte, 512+rand.Intn(16*1024))
		m[strconv.Itoa(len(m)%(10*100))] = b

		if len(m)%(10*100) == 0 {
			m = make(map[string][]byte)
		}

		time.Sleep(10 * time.Millisecond)
	}
}
