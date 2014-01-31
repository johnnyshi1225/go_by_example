package main

import (
	. "fmt"
	"time"
	"sync/atomic"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)

	var ops uint64 = 0
	for i := 0; i < 20; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}

	time.Sleep(time.Second)

	opsFinal := atomic.LoadUint64(&ops)
	Println("ops:", opsFinal)
}
