package main

import "GoApplication/algorithm"

func main() {
	//algorithm.T1()
	//algorithm.T2()
	//input := []int{7, 6, 4, 3, 1}
	//fmt.Println(maxProfit(input))
	algorithm.TestCaseBf()
}

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	var n1 = 1
	var n2 = 2
	var result = 2
	for i := 3; i <= n; i++ {
		result = n2 + n1
		n1 = n2
		n2 = result
	}
	return result
}

func maxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrices := prices[0]
	profit := 0
	for i := range prices {
		minPrices = min(minPrices, prices[i])
		profit = max(prices[i]-minPrices, profit)
	}
	return profit
}

func isDeseaseArr(arr []int) {

}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a >= b {
		return b
	} else {
		return a
	}
}
