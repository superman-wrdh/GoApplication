package algorithm

/**
https://leetcode.com/problems/house-robber/
You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security system connected and it will automatically contact the police if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house, determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.
*/

// 方法一 暴力搜索 可以解决 时间超时(重复计算)
//*****************************************
//func solve(idx int, nums []int) int {
//	if idx < 0 {
//		return 0
//	}
//	maxNum := Max(nums[idx]+solve(idx-2, nums), solve(idx-1, nums))
//	return maxNum
//}
//
//func rob(nums []int) int {
//	return solve(len(nums)-1, nums)
//}
//
//func Max(x, y int) int {
//	if x < y {
//		return y
//	}
//	return x
//}
//******************************************

// 方法二 在方法一基础上加入缓存 耗时 0 ms
func solve(idx int, nums []int, resultCache map[int]int) int {
	if _, ok := resultCache[idx]; ok {
		return resultCache[idx]
	}
	if idx < 0 {
		return 0
	}
	maxNum := Max(nums[idx]+solve(idx-2, nums, resultCache), solve(idx-1, nums, resultCache))
	resultCache[idx] = maxNum
	return maxNum
}

func rob(nums []int) int {
	var resultCache = make(map[int]int)
	return solve(len(nums)-1, nums, resultCache)
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
