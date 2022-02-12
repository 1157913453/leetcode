//@Time : 2021/7/2 19:14
//@Author : zr
//@File : 设计循环队列.go
//@Software : GoLand

package main

type MyCircularQueue struct {
	queue      []int
	head, tail int
}

func Constructor(k int) MyCircularQueue {
	return MyCircularQueue{
		queue: make([]int, k+1),
		head:  0,
		tail:  0,
	}
}

func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull() {
		return false
	}
	this.queue[this.tail] = value
	this.tail = (this.tail + 1) % cap(this.queue)
	return true
}

func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}
	this.head = (this.head+1)%cap(this.queue)
	return true
}

func (this *MyCircularQueue) Front() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.head]
}

func (this *MyCircularQueue) Rear() int {
	if this.IsEmpty() {
		return -1
	}
	return this.queue[(this.tail-1+cap(this.queue))%cap(this.queue)]
}

func (this *MyCircularQueue) IsEmpty() bool {
	return this.head == this.tail
}

func (this *MyCircularQueue) IsFull() bool {
	return (this.tail+1)%cap(this.queue) == this.head
}

func main() {
	k := 3
	value := 888
	obj := Constructor(k)
	param_1 := obj.EnQueue(value)
	param_1 = obj.EnQueue(1)
	param_1 = obj.EnQueue(2)
	println(param_1)
	param_1 = obj.EnQueue(3)

	param_2 := obj.DeQueue()
	param_3 := obj.Front()
	param_4 := obj.Rear()
	param_5 := obj.IsEmpty()
	param_6 := obj.IsFull()
	println(param_1)
	println(param_2)
	println(param_3)
	println(param_4)
	println(param_5)
	println(param_6)

}
