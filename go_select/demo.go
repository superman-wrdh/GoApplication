package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func launch() {
	fmt.Println("发射")
}

func commencingCountDown(canLunch chan int) {
	c := time.Tick(1 * time.Second)
	for countDown := 20; countDown > 0; countDown-- {
		fmt.Println(countDown)
		<-c
	}
	canLunch <- -1
}

func isAbort(abort chan int) {
	os.Stdin.Read(make([]byte, 1))
	abort <- -1
}

func randAbort(abort chan int) {
	duration := time.Duration(rand.Intn(10))
	fmt.Println("duration:", duration)
	time.Sleep(duration * time.Second)
	abort <- 1
}

func init() {
	rand.Seed(time.Now().UnixNano())
}
func main() {
	fmt.Println("发射倒计时")

	abort := make(chan int)
	canLunch := make(chan int)
	go isAbort(abort)
	go commencingCountDown(canLunch)
	go randAbort(abort)
	select {
	case <-canLunch:

	case <-abort:
		fmt.Println("发射取消!")
		return
	}
	launch()

}
