package interface_mutistatus

import (
	"fmt"
	"testing"
)

type S string

type Programer interface {
	WriteResponse() S
}

func forTest(p Programer) {
	fmt.Printf("p的类型是: %T, p实际是:%v", p, p.WriteResponse())
}

type Goprogramer struct {
}

type Javaprogramer struct {
}

func (g *Goprogramer) WriteResponse() S {
	return "go programer say hello!"
}

//再次强调一下这种定义方式的区别，
//建立在值类型的结构 之上的接口方法，会在方法被调用的时候 实例（struct）的成员会复制一份 占用多余内存空间
func (g Goprogramer) WriteResponseWithValue() S {
	return "---"
}

func (g *Javaprogramer) WriteResponse() S {
	return "java programer say hello!"
}

func Test_muti_status(t *testing.T) {
	var goProg Programer = &Goprogramer{} //等价 goProg := &Goprogramer{}

	//javaProg := &Javaprogramer{}

	forTest(goProg)
	fmt.Println()
	//forTest(javaProg)
}
