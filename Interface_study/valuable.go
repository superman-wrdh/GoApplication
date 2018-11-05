package main

import "fmt"

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

/* method to determine the value of a stock position */
func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valueable interface {
	getValue() float32
}

func showValue(asset valueable) {
	fmt.Printf("value of asset is %f \n", asset.getValue())
}

func main() {
	var o valueable = stockPosition{"Apple Inc", 1200, 4}
	showValue(o)
	o = car{"BMW", "x8", 120000}
	showValue(o)
}
