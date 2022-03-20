package main

// 双指针法（滑动窗口）判断  长度最小的子数组
// leetcode_array ：209
//func main() {
//	ints := []int{1, 1, 1, 1, 1, 1}
//	fmt.Println(move_window(ints, 6))
//}

func move_window(nums []int, a int) int {
	i, j := 0, 0
	min := len(nums) + 1
	sum := 0
	for j < len(nums) {
		sum = sum + nums[j]
		for sum >= a {
			len := j - i + 1
			if len < min {
				min = len
			}
			sum = sum - nums[i]
			i++
		}
		j++
	}

	if min == len(nums)+1 {
		return 0
	}
	return min
}
