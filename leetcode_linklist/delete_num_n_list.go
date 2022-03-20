package main

//func main() {
//	n1 := &Node{}
//	n1.Val = 1
//	n3 := &Node{}
//	n3.Val = 3
//	n4 := &Node{}
//	n4.Val = 4
//	n1.Next = n3
//	n3.Next = n4
//	n4.Next = nil
//
//	head := delete(n1, 3)
//	fmt.Println(head.Val)
//}

func delete(head *Node, n int) *Node {

	pre := &Node{
		Next: head,
	}
	bih := &Node{
		Next: head,
	}
	tmp := bih
	for pre != nil {
		if n >= 0 {
			pre = pre.Next
			n--
			continue
		}
		pre = pre.Next
		tmp = tmp.Next
	}
	tmp.Next = tmp.Next.Next
	return bih.Next

}
