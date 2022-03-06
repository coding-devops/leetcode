package main

// 递归求斐波那契数列
//func main() {
//	fmt.Println(recursion(2, 20))
//}

func recursion(x int, n int) int {
	if n == 0 {
		return 1
	}

	t := recursion(x, n/2)
	if n%2 == 1 {
		return t * t * x
	}
	return t * t
}
