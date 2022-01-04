package main

import (
	"fmt"
	"os"
)

type A struct {
}

func (a A) show() {
	fmt.Println("A call by show()")
}

func (a *A) ptShow() {
	fmt.Println("A call by ptShow()")
}

type B struct {
	a int
	A //结构
}

type C struct {
	*A // 结构的引用 这种定义的方式情况下，通过C （外部对象）创建的对象 就只能访问建立在*A的 方法上了 比如 ptShow
}

//func main() {
//	//这段执行没毛病
//	fmt.Println("test a,pa show:")
//	//a, pa := A{}, new(A)
//	//a.show()
//	//a.ptShow()
//	//pa.show()
//	//pa.ptShow()
//	//fmt.Printf("type of a(%T), pa(%T)\n", a, pa)
//	//这段执行没毛病
//
//	//结构内嵌套 "结构"  也没毛病
//	//b, bp := B{}, new(B)
//	//b.show()
//	//b.ptShow()
//	//
//	//bp.show()
//	//bp.ptShow()
//	//结构内嵌套 "结构"  也没毛病
//
//	//结构内嵌套 "结构引用" 时 只能访问 建立在结构引用上的方法 比如：ptShow()
//	c := C{}
//	//c.show()
//	c.ptShow()
//	//
//	//cp.show()
//	//cp.ptShow()
//
//}

type Programer interface {
	WriteResponseTest() string
}

type Goprogramer struct {
}

func (g Goprogramer) WriteResponseTest() string {
	fmt.Println("123123")
	return "12123"
}

//func main() {
//
//	var p Programer = new(Goprogramer)
//
//	p.WriteResponseTest()
//
//}

//go语言对象在内存中的分配 猜测

type object struct {
	sex int
	age int
}

type test struct {
	object
}

//func main() {
//	o := object{1, 2}
//
//	fmt.Printf("object type : %x", &o.age)
//	test_obj := test{o}
//	//值发生了复制
//	fmt.Printf("test_object type : %x", &test_obj.age)
//
//}

func main() {
	for index, v := range os.Args {
		fmt.Println(index, v)
	}
}

type Phone struct {
	UserName string
}

func NewPhone(userName string) Phone {
	return Phone{UserName: userName}
}

func (p *Phone) Call() {
	fmt.Println(p.UserName, " Call...")
}

type PhoneFeatures interface {
	Call()
}
