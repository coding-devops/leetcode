package main

import "fmt"

func main() {
	ints := []int{-1, 0, 3, 5, 9, 12, 13}
	fmt.Println(binary(ints, 13))
}

func binary(nums []int, target int) int {
	low, high := 0, len(nums)-1
	for low < high || low == high {
		if target > nums[(low+high)/2] {
			low = (low+high)/2 + 1
		} else if target < nums[(low+high)]/2 {
			high = (low+high)/2 - 1
		} else {
			return (low + high) / 2
		}
	}
	return -1
}
