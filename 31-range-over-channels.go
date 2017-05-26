package main

import "fmt"

func main() {
  queue := make(chan string, 2)
  queue <- "first string"
  queue <- "second string"
  close(queue)

  for elem := range queue {
    fmt.Println("Element: " + elem)
  }
}
