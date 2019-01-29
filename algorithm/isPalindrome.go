package algorithm

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	new_x, old_x := 0, x
	var rem int
	for x > 0 {
		rem = x % 10
		new_x = new_x*10 + rem
		x /= 10
	}
	return new_x == old_x
}
