package main

import "fmt"

//设计一个单链表存储
func main() {
	a := InitMyLinkList()
	a.addAtIndex(0, 10)
	a.addAtIndex(0, 20)
	a.addAtIndex(1, 30)
	fmt.Println(a.Get(0))
}

type MyLinkedList struct {
	cell *node
}
type node struct {
	Val  int
	Next *node
}

func InitMyLinkList() MyLinkedList {
	return MyLinkedList{}
}

// 获取 第 index 个链表元素的值
func (this *MyLinkedList) Get(index int) int {
	if this.cell == nil {
		return -1
	}
	node := this.cell
	for node.Next != nil && index > 0 {
		node = node.Next
		index--
	}
	// index 小于0 或者 当前节点到链表结尾了同时计数index 还没减到0
	if index < 0 || (node.Next == nil && index > 0) {
		return -1
	}
	return node.Val
}

// 头插法
func (this *MyLinkedList) AddAtHead(val int) {
	head := &node{}
	head.Val = val
	head.Next = this.cell
	this.cell = head
}

// 尾插法
func (this *MyLinkedList) AddAtTail(val int) {
	tmp := &node{}
	tmp = this.cell
	if tmp == nil {
		this.AddAtHead(val)
		return
	}
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tailnode := &node{}
	tailnode.Val = val
	tmp.Next = tailnode
}

// 插到 index 前
func (this *MyLinkedList) addAtIndex(index int, val int) {
	nodemiddle := this.cell
	insertNode := &node{
		Val: val,
	}
	if index <= 0 {
		this.AddAtHead(val)
		return
	}
	if this.cell == nil {
		return
	}
	for nodemiddle.Next != nil {
		if index == 1 {
			insertNode.Next = nodemiddle.Next
			nodemiddle.Next = insertNode
			return
		}
		nodemiddle = nodemiddle.Next
		index--
	}
	if index == 1 && nodemiddle.Next == nil { //这种情况刚好是 index值等于 链表的长度
		this.AddAtTail(val)
		return
	}
}

//删除链表节点
func (this *MyLinkedList) deleteLinkList_node(index int) {
	if this.cell == nil {
		return
	}
	if index < 0 {
		return
	}
	if this.cell.Next == nil && index == 0 {
		this.cell = nil
		return
	}
	tmp := &node{}
	tmp.Next = this.cell
	move := tmp
	for move.Next != nil {
		if index == 0 {
			move.Next = move.Next.Next
			this.cell = tmp.Next
			return
		}
		move = move.Next
		index--
	}
}
