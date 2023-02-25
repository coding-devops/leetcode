package main

func main() {
	generateMatrix(3)
}
test
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
func generateMatrix(n int) [][]int {

	nums := make([][]int, n)
	for i := 0; i < len(nums); i++ {
		nums[i] = make([]int, n)
	}
	currentNum := 1
	final := n * n
	left, right := 0, n-1
	top, button := 0, n-1

	for currentNum <= final {
		// 纵坐标不变，横坐标移动
		for i := left; i <= right; i++ {
			nums[top][i] = currentNum
			currentNum++
		}
		top++
		for i := top; i <= button; i++ {
			nums[i][right] = currentNum
			currentNum++
		}
		right--
		for i := right; i >= left; i-- {
			nums[button][i] = currentNum
			currentNum++
		}
		button--
		for i := button; i >= top; i-- {
			nums[i][left] = currentNum
			currentNum++
		}
		left++
	}
	return nums
}
