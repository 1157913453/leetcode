package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 *
 * @param pHead ListNode类
 * @return ListNode类
 */
func ReverseList(pHead *ListNode) *ListNode {
	fmt.Println("传入中",&pHead)
	arr := make([]*ListNode, 1)

	i := 0
	for pHead != nil {
		arr = append(arr, pHead)
		pHead = pHead.Next
		i++
	}
	for n := i; n > 0; n-- {
		arr[n].Next = arr[n-1]
	}

	pHead = arr[i]
	//fmt.Println(11111,pHead,pHead.Next,pHead.Next.Next)
	return pHead
}

func main() {
	head := &ListNode{Val: 5}
	next := head
	next.Next = &ListNode{Val: 3}
	next = head.Next
	next.Next = &ListNode{Val: 2}
	fmt.Println("传入前",&head)
	ReverseList(head)
	fmt.Println("传入后",&head)
	//for i := 0; i < 3; i++ {
	//	head = head.Next
	//}
}
