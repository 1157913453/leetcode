package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}


//func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
//	pre := &ListNode{}
//	result := pre
//	for l1 != nil && l2 != nil {
//		if l1.Val < l2.Val {
//			pre.Next = l1
//			l1 = l1.Next
//		} else {
//			pre.Next = l2
//			l2 = l2.Next
//		}
//		pre = pre.Next
//	}
//	if l1 == nil {
//		pre.Next = l2
//	} else {
//		pre.Next = l1
//	}
//	return result.Next
//}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := &ListNode{}
	prev := p
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			p.Next = l1
			l1 = l1.Next
		} else {
			p.Next = l2
			l2 = l2.Next
		}
        p = p.Next
	}
    if l2 == nil {
        p.Next = l1
    } else {
        p.Next = l2
    }
	return prev.Next
}









func main() {
	a := &ListNode{1,nil}
	a.Next = &ListNode{4,nil}

	b := &ListNode{1,nil}
	b.Next = &ListNode{3,nil}

	c := mergeTwoLists(a,b)
	fmt.Println(c)
}
