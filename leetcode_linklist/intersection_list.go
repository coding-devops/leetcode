package main

import "fmt"

// 相交链表 leetcode：面试题 0207
// 思路：先获取到两个链表的长度
// 记录较短长度的链表的长度 y
// 然后从较长链表的 x = len - y 的位置开始同时和较短链表一起向后移动，比较指针是否相同
func main() {

	x2 := &Node{
		Val:  2123,
		Next: nil,
	}
	x1 := &Node{
		Val:  1,
		Next: x2,
	}
	x0 := &Node{
		Val:  1241243123,
		Next: x1,
	}

	a2 := &Node{
		Val:  2,
		Next: x0,
	}
	a1 := &Node{
		Val:  1,
		Next: a2,
	}
	a0 := &Node{
		Val:  0,
		Next: a1,
	}
	//
	//b1 := &Node{
	//	Val:  1,
	//	Next: x0,
	//}
	//b0 := &Node{
	//	Val:  0,
	//	Next: b1,
	//}

	xx := intersection_list(x2, a0)
	fmt.Println(xx.Val)

}

func intersection_list(a *Node, b *Node) *Node {
	len_a, len_b := 0, 0

	tmpa := a
	tmpb := b
	for tmpa != nil {
		tmpa = tmpa.Next
		len_a++
	}
	for tmpb != nil {
		tmpb = tmpb.Next
		len_b++
	}
	fmt.Println(" a len : ", len_a)
	fmt.Println(" b len : ", len_b)

	step := 0
	if len_a > len_b {
		step = len_a - len_b
		for i := 0; i < step; i++ {
			a = a.Next
		}
	} else {
		step = len_b - len_a
		for i := 0; i < step; i++ {
			b = b.Next
		}
	}
	for a != nil {
		if a == b {
			return a
		}
		a = a.Next
		b = b.Next
	}
	return nil
}

//func a_b_1(a *Node, b *Node, ib int) *Node {
//}
