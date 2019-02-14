package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func producer(channel chan<- string) {
	for {
		p := rand.Float64()
		fmt.Printf("product:%v \n", p)
		channel <- strconv.FormatFloat(p, 'f', 6, 64)
		time.Sleep(time.Second * time.Duration(1))
	}
}

func consume(channel <-chan string) {
	for {
		message := <-channel // 此处会阻塞, 如果channel中没有数据的话 生产者 消费
		fmt.Printf("consume: %v \n", message)
	}
}

func main() {
	channel := make(chan string, 5)
	// 两个生产者 一个消费者
	go producer(channel)
	go producer(channel)
	consume(channel)
}
