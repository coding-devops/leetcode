package main

//二分查找，简单版，熟悉下理论基础
//func main() {
//	ints := []int{-1, 0, 3, 5, 9, 12, 13}
//	fmt.Println(binary(ints, 13))
//}

func binary(ints []int, a int) int {
	low, high := 0, len(ints)
	for low < high {
		mid := (low + high) / 2
		if ints[mid] == a {
			return mid
		} else if ints[mid] > a {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}
