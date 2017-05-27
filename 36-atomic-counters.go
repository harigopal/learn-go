package main

import "fmt"
import "time"
import "sync/atomic"

func main() {
	var counter uint64 = 0

	for i := 0; i < 50; i++ {
		go func() {
			for {
				atomic.AddUint64(&counter, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	counterFinalValue := atomic.LoadUint64(&counter)
	fmt.Println("Counter:", counterFinalValue)
}
