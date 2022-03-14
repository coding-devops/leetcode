package main

type ListNode struct {
	val  int
	next *ListNode
}

//  反转链表
func reverse_linked_list(head *ListNode) *ListNode {
	pre := &ListNode{}
	cur := &ListNode{}
	pre.next = nil
	cur.next = head
	//for cur.next != nil {
	//	tmp := cur.next
	//	cur.next = cur.next.next
	//	tmp.next = pre.next
	//	pre.next = tmp
	//}
	method2(pre, cur)
	return pre.next
}

func method2(pre *ListNode, cur *ListNode) {
	if cur.next == nil {
		return
	}
	tmp := cur.next
	cur.next = cur.next.next
	tmp.next = pre.next
	pre.next = tmp
	method2(pre, cur)
}
