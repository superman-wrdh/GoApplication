package main

import (
	"fmt"
	"time"
)

var c chan int

func handle(int) {
	fmt.Println("do some thing")
	c <- 1
}

func main() {
	select {
	case m := <-c:
		handle(m)
	case <-time.After(5 * time.Second):
		fmt.Println("timed out")
	}
}
