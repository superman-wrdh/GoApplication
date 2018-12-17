package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool)
	data := make(chan int)
	go func() {
		time.Sleep(time.Second * 10)
		timeout <- true
	}()

	go func() {
		time.Sleep(time.Second * 12)
		fmt.Print("send data")
		data <- 10
	}()

	select {
	case <-data:
		fmt.Print("get data")
	case <-timeout:
		fmt.Print("timeout")
		//default:
		//	fmt.Print("default")
	}

}
