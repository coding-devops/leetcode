package main

// 有序数组的平方

func main() {
	ints := []int{-4, -3, -1, 0, 3, 5, 9, 12, 13}
	sort_arr_squar(ints)
}

// 核心思路，有序数组，每个值中 平方后最大的值 一定出现在数组的两端
func sort_arr_squar(nums []int) []int {
	a_low_index, a_high_index := 0, len(nums)-1
	var new_arr = make([]int, len(nums))
	new_arr_len := len(new_arr) - 1
	for a_low_index <= a_high_index {
		if nums[a_low_index]*nums[a_low_index] < nums[a_high_index]*nums[a_high_index] {
			new_arr[new_arr_len] = nums[a_high_index] * nums[a_high_index]
			new_arr_len--
			a_high_index--
		} else {
			new_arr[new_arr_len] = nums[a_low_index] * nums[a_low_index]
			new_arr_len--
			a_low_index++
		}
	}
	return new_arr
}
