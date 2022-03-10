package main

import (
	"fmt"
)

// 输出螺旋矩阵
// leetcode_array ： 59
//func main() {
//	v := Spiral_Matrix(3)
//	for i := 0; i < len(v); i++ {
//		for j := 0; j < len(v[i]); j++ {
//			fmt.Println(v[i][j])
//		}
//	}
//}

// 1 2 3
// 4 5 6
// 7 8 9

func Spiral_Matrix(n int) [][]int {
	arr := make([][]int, n)

	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}

	fmt.Println(len(arr))
	currentNum := 1

	left, right := 0, n-1
	top, button := 0, n-1
	count := n * n
	for currentNum <= count {
		// 上边移动循环 [左闭右开)
		for i := left; i <= right; i++ {
			arr[top][i] = currentNum
			currentNum++
		}
		top++
		for i := top; i <= button; i++ {
			arr[i][right] = currentNum
			currentNum++
		}
		right--
		for i := right; i >= left; i-- {
			arr[button][i] = currentNum
			currentNum++
		}
		button--
		for i := button; i >= top; i-- {
			arr[i][left] = currentNum
			currentNum++
		}
		left++
	}
	return arr
}

// 给定一个螺旋矩阵 并非正方形，按顺时针输出 leetcode_array:54
func main() {
	twoarr := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	arr := print_SM(twoarr)
	for _, v := range arr {
		fmt.Println(v)
	}
}

func print_SM(matrix [][]int) []int {
	left := 0
	right := len(matrix[0]) - 1
	top := 0
	button := len(matrix) - 1
	var result []int
	sum := len(matrix[0]) * len(matrix)
	for sum > 0 {

		for i := left; i <= right; i++ {
			if matrix[top][i] == 101 {
				break
			}
			result = append(result, matrix[top][i])
			matrix[top][i] = 101
			sum--
		}
		top++

		for i := top; i <= button; i++ {
			if matrix[i][right] == 101 {
				break
			}
			result = append(result, matrix[i][right])
			matrix[i][right] = 101
			sum--
		}
		right--

		for i := right; i >= left; i-- {
			if matrix[button][i] == 101 {
				break
			}
			result = append(result, matrix[button][i])
			matrix[button][i] = 101
			sum--
		}
		button--

		for i := button; i >= top; i-- {
			if matrix[i][left] == 101 {
				break
			}
			result = append(result, matrix[i][left])
			matrix[i][left] = 101
			sum--
		}
		left++
	}

	return result
}
