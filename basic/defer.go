package main

import (
	"time"
	"runtime"
	"fmt"
)

func main() {
	var a [10]int
	for i := 0; i < 100; i++ {
		go func(i int) {
			for {
				a[i]++
				runtime.Gosched()
			}
		}(i)
	}
	fmt.Println(a)
	time.Sleep(time.Millisecond)
}
