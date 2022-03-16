package main

//func main() {
//	n1 := &Node{}
//	n1.Val = 1
//
//	//n2 := &listNode{}
//	//n2.val = 2
//
//	n3 := &Node{}
//	n3.Val = 3
//
//	n4 := &Node{}
//	n4.Val = 4
//	n1.Next = n3
//	//n2.next = n3
//	n3.Next = n4
//	n4.Next = nil
//
//	newhead := switchtwo(n1)
//
//	fmt.Println(newhead.Val)
//	fmt.Println(newhead.Next.Val)
//	fmt.Println(newhead.Next.Next.Val)
//	//fmt.Println(newhead.next.next.next.val)
//}

func switchtwo(head *Node) *Node {
	dum := &Node{}
	dum.Next = head
	pre := dum
	for head != nil && head.Next != nil {
		pre.Next = head.Next
		tmp := head.Next.Next
		pre.Next.Next = head
		head.Next = tmp
		pre = head
		head = tmp
	}
	return dum.Next
}
