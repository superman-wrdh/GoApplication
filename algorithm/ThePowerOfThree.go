package algorithm

import (
	"math"
)

func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	} else {
		return 1162261467%n == 0
	}
}

func TypeMaxOfPowerOfTree() int {
	// 1162261467 int 32
	// 4052555153018976256 int 64
	var a int
	a = math.MaxInt32
	var num = 0
	for i := 0; i < 1000000; i++ {
		tmp := math.Pow(3, float64(i))
		if float64(a)-tmp > 0 {
			num = int(tmp)
		} else {
			break
		}
	}
	return num
}
