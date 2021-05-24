package main

// import (
// 	"fmt"
// )

type MinStack struct {
    arr    []int
	length int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(val int)  {
    this.arr = append(this.arr, val)
	this.length ++
}


func (this *MinStack) Pop()  {
	if this.length == 0 {
		panic("empty")
	}

	// v := this.arr[this.length - 1]

	// w/ this strategy, the using space will grow, but time won't
	// this.arr = this.arr[0: this.length - 1]

	// w/ this strategy, the using time will grow, but size won't
	newArr := make([]int, 0, this.length - 1)
	for i := 0; i < this.length - 1; i ++ {
		newArr = append(newArr, this.arr[i])
	}
	this.arr = newArr

	this.length --
	// return v
}


func (this *MinStack) Top() int {
	if this.length  == 0 {
		panic("empty")
	}

	v := this.arr[this.length - 1]
	return v
}


func (this *MinStack) GetMin() int {
    if this.length == 0 {
		panic("panic")
	}

	min := this.arr[0]
	for _, val := range(this.arr) {
		if val < min {
			min = val
		}
	}

	// fmt.Println(min)
	return min
}

func main() {
	obj := MinStack{}
	obj.Push(-2)
	obj.Push(0)
	obj.Push(-3)
	obj.GetMin() // return -3
	obj.Pop()
	obj.Top()    // return 0
	obj.GetMin() // return -2
}
