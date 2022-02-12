package main

type ListNode struct {
	Val int
	Next *ListNode
}

func ans(list *ListNode) *ListNode {
	s := 0
	var tmp *ListNode
	for {
		if s&1 == 0 {
			list.next
		}
	}

}

func main() {
	head := ListNode{Val: 1}
	head.Next = &ListNode{Val: 2}
	head.Next.Next = &ListNode{Val: 3}
	head.Next.Next.Next = &ListNode{Val: 4}
	head.Next.Next.Next.Next = &ListNode{Val: 5}
}
