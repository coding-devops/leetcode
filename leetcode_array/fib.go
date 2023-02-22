package main

import "fmt"

//斐波那契数列 递归算法

func main() {
	//fmt.Println(fib(5)) // 输出第n位斐波那契数
	print_n_fib(8) //输出前n位斐波那契数
}

func print_n_fib(n int) {

	a, b := 1, 1
	fmt.Println(b)
	for i := 0; i < n; i++ {
		fmt.Println(b)
		tmp := a
		a = b
		b = tmp + b
	}
}

func fib(i int) int {

	if i <= 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fib(i-1) + fib(i-2)

}
