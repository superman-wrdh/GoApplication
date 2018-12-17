package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ch := make(chan int)

	for {
		go product(ch)
		go consumer(ch, "A")
		go consumer(ch, "B")
		time.Sleep(3 * time.Second)
		fmt.Println("ticket....")
	}
}
func product(c chan<- int) {
	i := rand.Int()
	fmt.Printf("product %d \n", i)
	c <- i
	time.Sleep(time.Second)

}
func consumer(c <-chan int, label string) {
	fmt.Printf("consumer %s %d \n", label, <-c)
}
