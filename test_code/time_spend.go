package main

import (
	"fmt"
	"math/big"
	"time"
)

func fib(n uint64) uint64 {
	/**
	 result support uint64
	if more than uint64,the result will be cut
	*/
	if n <= 2 {
		return 2
	} else {
		var content = [3]uint64{1, 1, 2}
		var i uint64 = 3
		for ; i <= n; i++ {
			r := content[0] + content[1]
			l := content[1]
			content[0] = l
			content[1] = r
		}
		return content[1]
	}
}

func fib2(n int) *big.Int {
	/*
	 support big int
	*/
	if n <= 2 {
		return big.NewInt(2)
	}
	var content = [3]*big.Int{big.NewInt(1), big.NewInt(1), big.NewInt(2)}
	var i = 3
	for ; i <= n; i++ {
		var r = big.NewInt(0)
		r.Add(content[0], content[1])
		var l = content[1]
		content[0] = l
		content[1] = r
	}
	return content[1]
}

func main() {
	t1 := time.Now()
	fmt.Print(fib2(1000000))
	t2 := time.Now()
	fmt.Print(" \n take times ")
	fmt.Println(t2.Sub(t1))
}
