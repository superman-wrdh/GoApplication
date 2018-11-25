package main

import (
	"time"
	"fmt"
)

var message = make(chan string)

func sample() {
	message <- "hello world"
}

func sample2() {
	time.Sleep(time.Second)
	str := <-message
	str = str + " I am go goroutine"
	message <- str
}
func main() {
	go sample()
	go sample2()
	time.Sleep(3 * time.Second)
	fmt.Println(<-message)
}
