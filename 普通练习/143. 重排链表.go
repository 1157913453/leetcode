package main

func reorderList(head *ListNode)  {
	p := head
	middle := middleNode1(p)
	next := middle.Next
	middle.Next = nil
	tem := middle
	tem.Next = next
	re := reverseList1(tem)
	var next1 *ListNode
	var next2 *ListNode
	for re != nil {
		next1 = head.Next
		head.Next = re
		next2 = re.Next
		re.Next = next1
		head = next1
		re = next2
	}
}

func middleNode1(head *ListNode) *ListNode {
	var slow *ListNode
	var fast *ListNode
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow.Next
}

func reverseList1(head *ListNode) *ListNode {
	var p *ListNode
	var next *ListNode
	for head != nil {
		next = head.Next
		head.Next = p
		p = head
		head = next
	}
	return p
}
