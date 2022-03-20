package main

//func main() {
//
//	d := &Node{}
//	d.Val = 123123
//	c := &Node{}
//	c.Val = 3
//	b := &Node{}
//	b.Val = 2
//	a := &Node{}
//	a.Val = 1
//	a.Next = b
//	b.Next = c
//	c.Next = d
//	d.Next = nil
//
//	new_head := reverse_linked_list(a)
//	fmt.Println(new_head.Val)
//
//}

//  反转链表
func reverse_linked_list(head *Node) *Node {
	vhead := &Node{
		Next: nil,
	}
	move := head
	for move != nil {
		tmp := move.Next
		move.Next = vhead.Next
		vhead.Next = move
		move = tmp
	}
	return vhead.Next
}
