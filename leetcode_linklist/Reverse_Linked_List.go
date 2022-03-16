package main

//  反转链表
func reverse_linked_list(head *Node) *Node {
	pre := &Node{}
	cur := &Node{}
	pre.Next = nil
	cur.Next = head
	//for cur.next != nil {
	//	tmp := cur.next
	//	cur.next = cur.next.next
	//	tmp.next = pre.next
	//	pre.next = tmp
	//}
	method2(pre, cur)
	return pre.Next
}

func method2(pre *Node, cur *Node) {
	if cur.Next == nil {
		return
	}
	tmp := cur.Next
	cur.Next = cur.Next.Next
	tmp.Next = pre.Next
	pre.Next = tmp
	method2(pre, cur)
}
