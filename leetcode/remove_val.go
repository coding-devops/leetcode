package main

import "fmt"

//双指针法 移除指定数组元素

func main() {

	ints := []int{1, 3, 3, 5, 5, 7, 7}
	fmt.Println(remove_same_val(ints))
	//fmt.Println(remove_val(ints, 3))
}

// leetcode ： 27
func remove_val(nums []int, target int) int {
	slowindex := 0
	for fastindex := 0; fastindex < len(nums); fastindex++ {
		if nums[fastindex] != target {
			nums[slowindex] = nums[fastindex]
			slowindex++
		}
	}
	return slowindex
}

//leetcode ：26 删除重复元素 返回结果数组长度
func remove_same_val(nums []int) int {

	slowindex := 0

	for fastindex := 1; fastindex < len(nums); fastindex++ {

		if nums[fastindex] != nums[slowindex] {
			slowindex++
			nums[slowindex] = nums[fastindex]
		}
	}
	return slowindex + 1
}
