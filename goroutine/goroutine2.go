package main

import (
	"fmt"
	"time"
)

func sendDate(ch chan string) {
	ch <- " I "
	ch <- " am "
	ch <- " a "
	ch <- " chan "
}

func getDate(ch chan string) {
	var input string
	for {
		input = <-ch
		fmt.Printf("%s", input)
	}
}

func main() {
	fmt.Print("start \n")
	ch := make(chan string)
	go sendDate(ch)
	go getDate(ch)
	time.Sleep(1e9)
	fmt.Print("\n end")
}
