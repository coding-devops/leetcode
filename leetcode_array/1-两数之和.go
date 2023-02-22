package main

import "fmt"

func main() {
	nums := []int{1, 2, 4, 5, 7} // target = 6
	resultArray := twoSum(nums, 7)
	fmt.Println(resultArray[0], resultArray[1])
}

// 时间复杂度 O(n²)
func twoSum(nums []int, target int) []int {
	var resultArray []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target && i != j {
				resultArray = append(resultArray, i, j)
			}
		}
	}
	return resultArray
}
