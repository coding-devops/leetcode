package main

import (
	"fmt"
)

// 找出n个字符串中相同的两个字符串
// 暴力解法
// O(n^2)
//func main() {
//	a := "stringa"
//	b := "stringb"
//	c := "stringc"
//	d := "stringa"
//
//	var stringss []string
//	stringss = append(append(append(append(stringss, a), b), c), d)
//	flag := 0
//	for i := 0; i < len(stringss); i++ {
//		x := stringss[i]
//		for j := i + 1; j < len(stringss); j++ {
//			y := stringss[j]
//			if x == y {
//				fmt.Println("找到了", x, y)
//				flag = 1
//				break
//			} else {
//				fmt.Println("不相等")
//			}
//		}
//		if flag == 1 {
//			break
//		}
//	}
//}

// 优化解法
// 先把若干字符串按照字典序排序
// O（nlogn）
//

func main() {
	ints := []int{7, 212312, 32, 4, 5, 6}
	quick_sort_for_number(ints, 0, len(ints)-1) // 0，5
	for _, v := range ints {
		fmt.Println(v)
	}
}

func quick_sort_for_number(ints []int, low int, high int) {
	i, j := low, high
	if i >= j {
		return
	}
	tmp := ints[i]
	for i < j {
		for i < j {
			if ints[j] < tmp {
				ints[i] = ints[j]
				i++
				break
			} else {
				j--
			}
		}
		for i < j {
			if ints[i] > tmp {
				ints[j] = ints[i]
				j--
				break
			} else {
				i++
			}
		}
	}
	ints[i] = tmp
	quick_sort_for_number(ints, low, i-1)
	quick_sort_for_number(ints, i+1, high)
}
