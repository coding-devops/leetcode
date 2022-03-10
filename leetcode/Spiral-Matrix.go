package main

// 螺旋矩阵
// leetcode ： 59

func main() {
	//Spiral_Matrix(3)
}

// 1 2 3
// 4 5 6
// 7 8 9

func Spiral_Matrix(a int) [][]int {
	var arr [][]int
	currentNum := 0

	left, right := 0, a-1
	top, button := 0, a-1
	count := a * a
	offset := 1
	for currentNum <= count {
		// 上边移动循环 [左闭右开)
		for i := left; i < a-offset; i++ {
			arr[top][i] = currentNum
			currentNum++
		}

		for i := top; i < a-offset; i++ {
			arr[i][right] = currentNum
			currentNum++
		}

		for i := right; i > left; i-- {
			arr[right][i] = currentNum
			currentNum++
		}

		for i := button; i < top; i++ {
			arr[i][left] = currentNum
			currentNum++
		}
		top++
		right--
		left++
		button--
		offset = offset - 2
	}

	return [][]int{}
}
