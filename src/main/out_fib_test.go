package main

import (
	"fmt"
	"testing"
	"unsafe"
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

	ar_sec2 := arr3[:] //这里返回的其实是一个切片
	t.Log(ar_sec2)
}

// 切片 : 共享存储结构
//	长度已知的情况可以用数组。
//	切片的方便之处主要来自于自动增长
//	但切片的自动增长会导致内存分配和数据复制，以及未来相关的GC开销。

// cap存在的意义是什么？
//cap是指后面对应的连续存储空间的大小，你可以向slice中append元素，没有超过cap及不会引发slice扩容。
func TestSlice(t *testing.T) {

	var s0 []int

	t.Log(len(s0), cap(s0)) // 0 0  // len 初始化元素个数。和最大长度

	s0 = append(s0, 1)

	t.Log(len(s0), cap(s0))

	s1 := make([]int, 3, 7)

	t.Log(len(s1), cap(s1))

	t.Log(s1[0], s1[1], s1[2])

	s1 = append(s1, 1)
	t.Log(s1[0], s1[1], s1[2], s1[3])

}

// 切片中的共享存储
func TestSliceShareMemory(t *testing.T) {

	year := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S"}

	slice_01 := year[3:7] //D E F G

	t.Log(slice_01, "长度:", len(slice_01), "容量:", cap(slice_01))

	slice_02 := year[5:11] //F G H I J K

	t.Log(slice_02, "长度:", len(slice_02), "容量:", cap(slice_02))

	slice_02[0] = "OJBK"

	t.Log("slice_02修改后的 slice_01信息", slice_01, "长度:", len(slice_01), "容量:", cap(slice_01))
}

//通过 切片变量 和 nil的比较，来判断切片变量是否被初始化了

func TestSliceInit(t *testing.T) {
	//var a []int
	//var b = make([]int,0,0)
	//c := []int{}
	//t.Log(a, len(a), cap(a))
	//t.Log(b, len(b), cap(b))
	//t.Log(c, len(c), cap(c))
	//t.Log(a==nil, b==nil, c==nil)

	a := []int{1, 2, 3, 4, 5, 6}
	a = append(a, 3)
	t.Log(a, len(a), cap(a))
}

//可变长参数

func Test_arg(t *testing.T) {
	Sum_arg(1, 2, 3, 4)
}
func Sum_arg(ok ...int) {

	for _, v := range ok {
		fmt.Println(v)
	}
}

//延迟执行
func Test_defer(t *testing.T) {
	defer func() {
		t.Log("Sssss")
	}()

	t.Log("1")
}

type Employee struct {
}

type First_Object struct {
	id             string
	age            string
	String_pointer func() string
}

func Test_object(t *testing.T) {

	e := First_Object{"1", "2", func() string {
		return "method in struct"
	}}

	t.Logf("e 的内存地址: %x", unsafe.Pointer(&e.age))
	t.Log(e.String_pointer())
	//t.Log("*********************")
	//e.String_value()

}

// pointer
//func (test_struct First_Object) String_pointer() string {
//	fmt.Printf("test_struct pointer 的内存地址: %x\n", unsafe.Pointer(&test_struct.age))
//	return "pointer is running"
//}

//func (test_struct First_Object) String_value() string {
//	fmt.Printf("test_struct value 的内存地址: %x\n", unsafe.Pointer(&test_struct))
//	return "value is running"
//}
