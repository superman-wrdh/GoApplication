package main

import (
	"time"
	"fmt"
)

/**
FIFO
 */

func sample(message chan string) {
	message <- "hello goroutine1"
	message <- "hello goroutine2"
	message <- "hello goroutine3"
	message <- "hello goroutine4"
}

func sample2(message chan string) {
	time.Sleep(time.Second)
	str := <-message
	str = str + " I am go goroutine"
	message <- str
	close(message) //最后一次往里面写
}
func main() {
	var message = make(chan string, 3)
	go sample(message)
	go sample2(message)
	time.Sleep(time.Second)
	for str := range message {
		fmt.Println(str)
	}
	fmt.Println("hello world")
}
