package main

import (
	"fmt"
	"math"
)

type MinStack struct {
	stack []int
	//根据栈先进后出，当前min下面的都比它大，上面的都比它小
	minStack []int
}

func Constructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (this *MinStack) Push(val int) {
	this.stack = append(this.stack, val)
	if this.GetMin() >= val {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	if this.Top() == this.GetMin() {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return math.MaxInt
	}
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	if len(this.minStack) == 0 {
		return math.MaxInt
	}
	return this.minStack[len(this.minStack)-1]
}

func main() {
	obj := Constructor()
	obj.Push(0)
	obj.Push(1)
	obj.Push(0)
	fmt.Println(obj.GetMin())
	obj.Pop()
	fmt.Println(obj.GetMin())

}
