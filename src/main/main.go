package main

import "fmt"

type A struct {
}

func (a A) show() {
	fmt.Println("A call by show()")
}

func (a *A) ptShow() {
	fmt.Println("A call by ptShow()")
}

type B struct {
	A
}

type C struct {
	*A
}

func main() {
	fmt.Println("test a,pa show:")

	a, pa := A{}, new(A)

	a.show()
	a.ptShow()

	pa.show()
	pa.ptShow()

	fmt.Printf("type of a(%T), pa(%T)\n", a, pa)
}
