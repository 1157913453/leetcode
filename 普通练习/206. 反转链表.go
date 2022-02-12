package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

//func reverseList(head *ListNode) *ListNode {
//	arr := make([]*ListNode, 1)
//	i := 0
//	for head != nil {
//		arr = append(arr, head)
//		head = head.Next
//		i++
//	}
//	for n := i; n > 0; n-- {
//		arr[n].Next = arr[n-1]
//	}
//	head = arr[i]
//	return head
//}

//官方解答
//func reverseList(head *ListNode) *ListNode {
//	var tmp *ListNode
//	for head != nil {
//		next := head.Next
//		head.Next = tmp
//		tmp = head
//		head = next
//	}
//	return tmp
//}
//
//func main() {
//	head := &ListNode{Val: 2}
//	head.Next = &ListNode{Val: 5}
//	head.Next.Next = &ListNode{Val: 8}
//	head.Next.Next.Next = &ListNode{Val: 9}
//	head = reverseList(head)
//	fmt.Println(head.Val, head.Next.Val, head.Next.Next.Val, head.Next.Next.Next.Val)
//}

//-----------2021.8.11----------------

//
//func reverseList(head *ListNode) *ListNode {
//	var tmp *ListNode
//	var next *ListNode
//	for head != nil{
//		next = head.Next
//		head.Next = tmp
//		tmp = head
//		head = next
//	}
//	return tmp
//}





//-----------------------2021.11.16--------------
//func reverseList(head *ListNode) *ListNode {
//	var p *ListNode
//	var next *ListNode
//	for head.Next != nil {
//		next = head.Next
//		head.Next = p
//		p = head
//		head = next
//	}
//	return p
//}

func main(){
	reverseList(1)
}