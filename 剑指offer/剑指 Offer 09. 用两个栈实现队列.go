//@Time : 2021/7/25 21:01
//@Author : zr
//@File : 剑指 Offer 09. 用两个栈实现队列.go
//@Software : GoLand

package main

import (
	"container/list"
	"fmt"
)

type CQueue struct {
	stack1, stack2 *list.List
}

func Constructor() CQueue {
	return CQueue{
		stack1: list.New(),
		stack2: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.stack1.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.stack2.Len() == 0 {
		if this.stack1.Len() == 0 {
			return -1
		}
		for this.stack1.Len() != 0 {
			this.stack2.PushBack(this.stack1.Remove(this.stack1.Back()))
		}
	}
	e := this.stack2.Back()
	this.stack2.Remove(e)
	return e.Value.(int)
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */

func main() {
	obj := Constructor()
	obj.AppendTail(3)
	obj.AppendTail(5)
	fmt.Println("1111",obj.DeleteHead())
	fmt.Println(222222,obj.DeleteHead())
	obj.AppendTail(6)
	fmt.Println(33333,obj.DeleteHead())
	obj.AppendTail(3)
	obj.AppendTail(3)
}
