package main

import "fmt"

func main() {
	nums := []int{1, 2, 4, 5, 7} // target = 7
	resultArray := twoSum(nums, 7)
	fmt.Println(resultArray[0], resultArray[1])
}

// 时间复杂度 = O(n²)
// 空间复杂度 = ？
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

// 时间复杂度  O(N)
// 1,0
// 2,1
func twoSumWithO1(nums []int, target int) []int {
	hashmap := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if value, ok := hashmap[target-nums[i]]; ok {
			return []int{value, i}
		} else {
			hashmap[nums[i]] = i
		}
	}
	return nil
}
