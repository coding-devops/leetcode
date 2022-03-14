package main

import "fmt"

func main() {
	n1 := &listNode{}
	n1.val = 1

	//n2 := &listNode{}
	//n2.val = 2

	n3 := &listNode{}
	n3.val = 3

	n4 := &listNode{}
	n4.val = 4
	n1.next = n3
	//n2.next = n3
	n3.next = n4
	n4.next = nil

	newhead := switchtwo(n1)

	fmt.Println(newhead.val)
	fmt.Println(newhead.next.val)
	fmt.Println(newhead.next.next.val)
	//fmt.Println(newhead.next.next.next.val)
}

type listNode struct {
	val  int
	next *listNode
}

func switchtwo(head *listNode) *listNode {
	dum := &listNode{}
	dum.next = head
	pre := dum
	move := head
	// 1、当前需要交换的两个数字中的第一个数字的下一个节点不为空   2、当前需要交换的两个数字中的第一个数字不为空
	for move != nil && move.next != nil {
		pre.next = move.next
		tmp := move.next.next
		pre.next.next = move
		move.next = tmp
		pre = move
		move = tmp
	}
	return dum.next
}
