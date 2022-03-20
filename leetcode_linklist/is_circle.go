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
	fast := head
	slow := head
	for fast != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			tmp := head
			for tmp != fast {
				fast = fast.Next
				tmp = tmp.Next
			}
			return fast
		}
	}
	return nil
}
