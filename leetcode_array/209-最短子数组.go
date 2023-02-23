package main

import "math"

/*
 * @lc app=leetcode.cn id=209 lang=golang
 *
 * [209] 长度最小的子数组
 */

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	slow, fast := 0, 0
	min := 0
	sum := 0
	result := math.MaxInt64
	for fast < len(nums) {
		sum = sum + nums[fast]
		for sum >= target {
			min = fast - slow + 1
			if min < result {
				result = min
			}
			sum = sum - nums[slow]
			slow++
		}
		fast++
	}
	if result == math.MaxInt64 {
		return 0
	}
	return result

}
