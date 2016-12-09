package main

import "fmt"
import "time"

func main() {
	slowChannel := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		slowChannel <- "I'm sorry I'm so slow. :("
	}()

	select {
	case slowMessage := <-slowChannel:
		fmt.Println(slowMessage)
	case <-time.After(time.Second):
		fmt.Println("The slow channel timed out. Predictably.")
	}

	fastChannel := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		fastChannel <- "Oh yeah. I'm fast! :D"
	}()

	select {
	case fastMessage := <-fastChannel:
		fmt.Println(fastMessage)
	case <-time.After(time.Second * 3):
		fmt.Println("I'm never gonna run, am I?")
	}
}
