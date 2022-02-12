package main

import "fmt"

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func GetLen(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return length
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	l := GetLen(head)
	dummy := &ListNode{0, head}
	cur := dummy
	for i := 0; i < l-n; i++ {
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return dummy.Next
}

func main() {
	h := ListNode{Val: 2}
	h.Next = &ListNode{Val: 5}
	h.Next.Next = &ListNode{Val: 7}
	h.Next.Next.Next = &ListNode{Val: 1}
	h.Next.Next.Next.Next = &ListNode{Val: 9}
	fmt.Println(GetLen(&h))
	fmt.Println(h.Val)
}
