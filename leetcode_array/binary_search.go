package main

//二分查找，简单版，熟悉下理论基础
//func main() {
//	ints := []int{-1, 0, 3, 5, 9, 12, 13}
//	fmt.Println(binary(ints, 13))
//}

func binary(nums []int, target int) int {
	low, high := 0, len(nums)-1
	tmp := -1
	for low < high {
		middle := (low + high + 1) / 2
		if target > nums[middle] {
			low = middle
		} else if target < nums[middle] {
			high = middle - 1
		} else {
			return middle
		}
	}
	return tmp
}
