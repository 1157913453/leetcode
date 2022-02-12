package main

func Merge( pHead1 *ListNode ,  pHead2 *ListNode ) *ListNode {
    // write code here
	p1, p2, head := pHead1, pHead2, &ListNode{}
	ans := head
	for p1 != nil && p2 != nil {
		if p1.Val < p2.Val {
			head.Next = p1
			p1 = p1.Next
		} else {
			head.Next = p2
			p2 = p2.Next
		}
	}
	if p1 == nil {
		head.Next = p2
	} else {
		head.Next = p1
	}
	return ans.Next
}
