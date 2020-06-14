package main

import (
	"fmt"
	"math"
	"time"
)

var count = 0

type QueenN struct {
	count int
	N     int
}

func (qn *QueenN) TotalCount() int {
	columns := make([]int, qn.N)
	qn.backTracing(0, columns)
	return qn.count
}

func (qn *QueenN) check(row, col int, columns []int) bool {
	for r := 0; r < row; r++ {
		rx := int64(row - r)
		rxi := int64(math.Abs(float64(columns[r]) - float64(col)))
		if columns[r] == col || rx == rxi {
			return false
		}
	}
	return true
}

func (qn *QueenN) backTracing(row int, columns []int) {
	if row == qn.N {
		qn.count++
		return
	}
	for col := 0; col < qn.N; col++ {
		columns[row] = col
		if qn.check(row, col, columns) {
			qn.backTracing(row+1, columns)
		}
		columns[row] = -1
	}
}

func main() {
	for i := 8; i < 20; i++ {
		n := i
		start := time.Now()
		qn := QueenN{0, n}
		qn.TotalCount()
		fmt.Print("queen ", i, " ")
		fmt.Print("cnt ", qn.count)
		cost := time.Since(start)
		fmt.Print(" cost:", cost)
		fmt.Println()
		/**
		benchmark
		MacBook Pro (16-inch, 2019) 2.6Ghz/6 core

		queen 8 cnt 92 cost:400.8µs
		queen 9 cnt 352 cost:1.684013ms
		queen 10 cnt 724 cost:8.080724ms
		queen 11 cnt 2680 cost:39.477366ms
		queen 12 cnt 14200 cost:168.708918ms
		queen 13 cnt 73712 cost:936.029856ms
		queen 14 cnt 365596 cost:5.935042083s
		queen 15 cnt 2279184 cost:40.397361785s
		*/

	}

}
