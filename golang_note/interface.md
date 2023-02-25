**为什么nil类型接口变量  !=nil**
```
package main
import (
	"errors"
)

type MyError struct{ error }

var ErrBad = MyError{error: errors.New("bad things happened")}

func main() {
	var p *MyError = nil
	var a interface{} = p
	// var err2 interface{} = nil
	// var p *MyError = nil
	// println("err1:", err1)
	println("p:", p)
	println("a:", a)
	// output:
	// p: 0x0
	// a: (0x102fb18a0,0x0)
	// 这里的 a != nil
	// 因为a这个接口类型低层是由两部分组成的动态类型
	//
	// 对于非空接口类型低层，itab信息里除了保存当前赋值给此接口类型的具体的动态类型的信息（_type），还要保存当前当前接口自身的信息，自己所包含的接口方法声明等信息，
	// unsafe.Pointer 则表示赋值给当前接口的动态类型的实际数据
	// $GOROOT/src/runtime/runtime2.go
	// type iface struct {
	//     tab  *itab
	//     data unsafe.Pointer
	// }

	// 对应空接口类型低层，空接口就比较简单了，_type就直接表示当前赋值给此接口类型的具体的动态类型的信息
	// unsafe.Pointer 则表示赋值给当前接口的动态类型的实际数据，
	// type eface struct {
	//     _type *_type
	//     data  unsafe.Pointer
	// }
}
```

**还有一点需要注意**
```
func printEmptyInterfaceAndNonEmptyInterface() {
  var eif interface{} = T(5)
  var err error = T(5)
  println("eif:", eif)
  println("err:", err)
  println("eif = err:", eif == err)

  err = T(6)
  println("eif:", eif)
  println("err:", err)
  println("eif = err:", eif == err)
}

输出结果：

eif: (0x10b3b00,0x10eb4d0)
err: (0x10ed380,0x10eb4d8)
eif = err: true
eif: (0x10b3b00,0x10eb4d0)
err: (0x10ed380,0x10eb4e0)
eif = err: false
```
空接口类型变量和非空接口类型变量内部表示的结构有所不同（第一个字段：_type vs. tab)，两者似乎一定不能相等。但 Go 在进行等值比较时，类型比较使用的是 eface 的 _type 和 iface 的 tab._type，因此就像我们在这个例子中看到的那样，当 eif 和 err 都被赋值为T(5)时，两者之间是划等号的。

<br>
<br>
<br>

**<h1>小结**
<h2>要更好地理解 Go 接口的这两种特性，我们需要深入到 Go 接口在运行时的表示层面上去。接口类型变量在运行时表示为 eface 和 iface，eface 用于表示空接口类型变量，iface 用于表示非空接口类型变量。只有两个接口类型变量的类型信息（eface._type/iface.tab._type）相同，且数据指针（eface.data/iface.data）所指数据相同时，两个接口类型变量才是相等的。</h2>



