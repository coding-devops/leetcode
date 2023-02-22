package main

import "fmt"

// 简单直接插入排序
func main() {
	var a = []int{1, 2, 4, 7, 17, 101, 1}
	InsectionSort(a)
	for _, v := range a {
		fmt.Println(v)
	}
}

func InsectionSort(a []int) {
	for i := 1; i < len(a); i++ {
		tmp := a[i]
		for j := i - 1; j > 0; j-- {
			if tmp < a[j] {
				a[j+1] = a[j]
				a[j] = tmp
			}
		}
	}
}
