package main

import "fmt"
import "time"

func main() {
	quarterSecondRepeater := time.NewTicker(time.Millisecond * 250)

	go func() {
		for t := range quarterSecondRepeater.C {
			fmt.Println("BAM! at", t)
		}
	}()

	time.Sleep(time.Millisecond * 1100)
	quarterSecondRepeater.Stop()
	fmt.Println("All right, that's enough.")
}
