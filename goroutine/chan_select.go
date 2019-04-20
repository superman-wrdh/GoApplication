package main

import (
	"fmt"
	"time"
)

func main() {
	print("run \n")
	intChan := make(chan int, 1)
	// 一秒后关闭通道。
	time.AfterFunc(time.Second*5, func() {
		close(intChan)
	})
	select {
	case _, ok := <-intChan:
		if !ok {
			fmt.Println("The candidate case is closed.")
			break
		}
		fmt.Println("The candidate case is selected.")
	}
	print("finished")
}
