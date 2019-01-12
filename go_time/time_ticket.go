package main

import (
	"fmt"
	"time"
)

func statusUpdate() string { return time.Now().String() }

func main() {
	c := time.Tick(2 * time.Second)
	for now := range c {
		fmt.Printf("%v %s\n", now, statusUpdate())
	}
}
