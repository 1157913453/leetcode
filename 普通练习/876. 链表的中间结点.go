package main

//解法一
//func middleNode(head *ListNode) *ListNode {
//	var arr []*ListNode
//	for head != nil {
//		arr = append(arr, head)
//		head = head.Next
//	}
//	l := len(arr)
//	for i:= l/2; i<l-1;i++ {
//		arr[i].Next = arr[i+1]
//	}
//	return arr[l/2]
//}

////解法2，两次反转
//func middleNode(head *ListNode) *ListNode {
//	list1, l := reverseList(head)
//	head, _ = reverseList(list1)
//	for i := 0; i < (l/2); i++ {
//		head = head.Next
//	}
//	return head
//}
//
//func reverseList(head *ListNode) (*ListNode, int) {
//	var p, next *ListNode
//	n := 0
//	for head != nil {
//		next = head.Next
//		head.Next = p
//		p = head
//		head = next
//		n++
//	}
//	return p, n
//}


func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = slow.Next.Next
	}
	return slow
}


