package main

//142. 环形链表 II  leetcode

//func main() {
//
//}

// hash 解法
func is_circle(head *Node) *Node {
	nodetable := make(map[*Node]int)
	if head == nil {
		return nil
	}
	tmp := head
	for head != nil {
		nodetable[tmp] = 1
		tmp = tmp.Next
		if _, ok := nodetable[tmp]; ok {
			return tmp
		}
	}
	return nil
}

func double_pointer(head *Node) *Node {
	slow, fast := head, head

	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			for head != slow {
				head = head.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
