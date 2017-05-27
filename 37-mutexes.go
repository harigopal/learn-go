package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)

	var mutex = &sync.Mutex{}

	var reads uint64 = 0
	var writes uint64 = 0

	for readOperation := 0; readOperation < 100; readOperation++ {
		go func() {
			total := 0

			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&reads, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	for writeOperation := 0; writeOperation < 10; writeOperation++ {
		go func() {
			for {
				key := rand.Intn(5)
				randomValue := rand.Intn(100)
				mutex.Lock()
				state[key] = randomValue
				mutex.Unlock()
				atomic.AddUint64(&writes, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	mutex.Lock()

	finalReadsCount := atomic.LoadUint64(&reads)
	fmt.Println("Reads:", finalReadsCount)
	finalWritesCount := atomic.LoadUint64(&writes)
	fmt.Println("Writes:", finalWritesCount)
	fmt.Println("State:", state)

	mutex.Unlock()
}
