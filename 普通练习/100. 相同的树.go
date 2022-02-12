package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if q == nil || p == nil{
		return false
	}else if q.Val != p.Val {
		return false
	}
	return isSameTree(q.Left,p.Left) && isSameTree(q.Right, p.Right)
}

func main() {
	p := &TreeNode{1,nil,nil}
	p.Left = &TreeNode{2,nil,nil}
	p.Right = &TreeNode{3,nil,nil}
	q := &TreeNode{1,nil,nil}
	q.Left = &TreeNode{2,nil,nil}
	q.Right = &TreeNode{3,nil,nil}
	fmt.Println(isSameTree(p,q))
	//fmt.Println(p)
}
