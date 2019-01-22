package algorithm

/**
TODO
https://leetcode.com/explore/interview/card/top-interview-questions-easy/102/math/745/
*/
func IsPowerOfThree(n int) bool {

	maxUint64 := MaxPowerOfThree()

	if n > 0 {
		x := maxUint64 % n
		return x == 0
	} else {
		return false
	}
}

func MaxPowerOfThree() int {
	return 12378192
}
