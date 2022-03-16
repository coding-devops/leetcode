package main

//func main() {
//	n1 := &Node{}
//	n1.val = 1
//
//	n2 := &Node{}
//	n2.val = 2
//
//	n3 := &Node{}
//	n3.val = 3
//
//	n4 := &Node{}
//	n4.val = 2
//
//	n1.Next = n2
//	n2.Next = n3
//	n3.Next = n4
//	n4.Next = nil
//	head := non_v_head(n1, 2)
//	fmt.Println(head.val)
//	fmt.Println(head.Next.val)
//	//linklist(n1, 2)
//}

//1->4->3->2->4 // 虚拟头节点写法
func linklist(head *Node, val int) *Node {
	tmp := &Node{}
	tmp.Next = head
	move := tmp
	for move.Next != nil {
		if move.Next.Val == val {
			move.Next = move.Next.Next
		} else {
			move = move.Next
		}
	}
	return tmp.Next
}

// 非虚拟头节点写法

func non_v_head(head *Node, val int) *Node {

	if head == nil {
		return nil
	}
	for head != nil && head.Val == val {
		head = head.Next

	}
	tmp := &Node{}
	tmp.Next = head
	for tmp.Next != nil {
		if tmp.Next.Val == val {
			tmp.Next = tmp.Next.Next
		} else {
			tmp = tmp.Next
		}
	}
	return head
}
