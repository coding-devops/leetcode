package main

// 有序数组的平方

//func main() {
//	ints := []int{-12, -1, 12, 13, 134, 4123}
//	nums := sort_arr_squar(ints)
//	for _, v := range nums {
//		fmt.Println(v)
//	}
//}

// 核心思路，有序数组，每个值中 平方后最大的值 一定出现在数组的两端
func sort_arr_squar(nums []int) []int {
	i, j := 0, len(nums)-1

	newarr := make([]int, len(nums))
	len := len(newarr) - 1
	for i <= j {
		if nums[j]*nums[j] > nums[i]*nums[i] {
			newarr[len] = nums[j] * nums[j]
			len--
			j--
		} else {
			newarr[len] = nums[i] * nums[i]
			len--
			i++
		}
	}
	return newarr
}
