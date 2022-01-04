package main

import "testing"

// Go语言不支持隐式类型转换

type MyInt int64

func TestImplicit(t *testing.T) {

	var a int32 = 1
	var b int64
	b = int64(a)
	t.Log(a, b)
}
