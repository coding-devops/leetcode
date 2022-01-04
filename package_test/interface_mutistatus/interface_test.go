package interface_mutistatus

import (
	"fmt"
	"testing"
)

type S string

type Programer interface {
	//WriteResponse() S
	WriteResponseWithPt() S
}

func forTest(p Programer) {
	fmt.Printf("p的类型是: %T", p)
}

type Goprogramer struct {
	age int
}

func (g *Goprogramer) WriteResponseWithPt() S {
	fmt.Println("123")
	return "---"
}

func Test_muti_status(t *testing.T) {
	//var goProg Programer = &Goprogramer{} //这个问题的关键我理解在于判定当前这个 Goprogramer 对象 是不是Programer的实现类

}

//断言判断传入类型
func DoSomething(p interface{}) {
	if v, ok := p.(int); ok {
		fmt.Println("转换成功，value ", v)
	}
}

func Test_interface_arg(t *testing.T) {
	var a string
	a = "13"
	DoSomething(a)
}
