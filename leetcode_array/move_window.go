package main

// 双指针法（滑动窗口）判断  长度最小的子数组
// leetcode_array ：209
//func main() {
//	ints := []int{2, 3, 1, 2, 4, 7}
//	fmt.Println(move_window(ints, 7))
//}

func move_window(nums []int, a int) int {

	i := 0
	sum := 0
	min := len(nums) + 1
	for j := 0; j < len(nums); j++ {
		sum = sum + nums[j]
		for sum >= a {
			len := j - i + 1
			if len < min {
				min = len
			}
			sum = sum - nums[i]
			i++
		}
	}
	if min == len(nums)+1 {
		return 0
	}
	return min
}
