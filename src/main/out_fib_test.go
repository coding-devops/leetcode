package main

import (
	"testing"
)

//输出前n位fib
func TestFib(t *testing.T) {
	var a int = 1
	var b int = 0
	for i := 0; i < 10; i++ {
		t.Log(a)
		temp := b
		b = a
		a = a + temp
	}
}

//快速交换两个值
func TestSwitchPara(t *testing.T) {
	a := 1
	b := 2
	a, b = b, a
	t.Log(a)
}

//递归计算前n位fib数的和

// 数组比较
func TestCompareArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	c := [...]int{1, 2, 3, 4, 5}
	t.Log(a == b)
	t.Log(c == c)
}

const (
	Readable = 1 << iota
	Writeable
	R_W
)

//按位清零
func TestBitClear(t *testing.T) {
	a := 7            // 0111
	a = a &^ Readable //  将可读位置零
	t.Log(a)
}

//循环
func TestLoop(t *testing.T) {
	for {
		t.Log("111")
		break
	}
}

//集合
// 数组和切片

func TestArrayAndSlice(t *testing.T) {

	//var a [3]int //默认初始化为0值
	b := [3]int{1, 2, 3} //声明的同时初始化
	b[0] = 1
	b[1] = 2
	b[2] = 3

}

//数组的遍历

func TestArrayTravel(t *testing.T) {
	b := [3]int{1, 2, 3} //声明的同时初始化
	for _, e := range b {
		t.Log(e)
	}
}

//数组截取
func Test_sectionArray(t *testing.T) {
	arr3 := [...]int{1, 2, 3, 4}
	ar_sec := arr3[1:]
	t.Log(ar_sec)

	ar_sec2 := arr3[:]
	t.Log(ar_sec2)
}
