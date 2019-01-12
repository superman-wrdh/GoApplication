package main

import (
	"fmt"
	"time"
)

func expensiveCall() {
	fmt.Println("task run")
}

func main() {
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("The call took %v to run \n", t1.Sub(t0))
}
