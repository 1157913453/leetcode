package main

//type ListNode struct {
//	Val  int
//	Next *ListNode
//}

func isPalindrome(head *ListNode) bool {
	arr := []int{}
	for ;head != nil; head = head.Next {
		arr = append(arr, head.Val)
	}
	len := len(arr)
	for i:=0;i<len/2;i++{
		if arr[i] != arr[len-1-i] {
			return false
		}
	}
	return true
}

func main() {
	h := ListNode{Val: 1}
	h.Next = &ListNode{Val: 2}
	h.Next.Next = &ListNode{Val: 2}
	h.Next.Next.Next = &ListNode{Val: 1}
	//fmt.Println(isPalindrome(&h))
}
