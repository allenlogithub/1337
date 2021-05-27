package main

import (
	"fmt"
	"strconv"
)

type Stack struct {
	arr []int
	length int
}

func Constructor() Stack {
	return Stack{}
}

func (this *Stack) Push(val int) {
	this.arr = append(this.arr, val)
	this.length ++

	// fmt.Println("Push: ", val)
}

func (this *Stack) Pop() int {
	if this.length == 0 {
		panic("empty")
	}

	popVal := this.arr[this.length - 1]

	// opt1
	this.arr = this.arr[0: this.length - 1]
	
	// opt2
	// newArr := make([]int, 0, this.length - 1)
	// for i := 0; i < this.length - 1; i ++ {
	// 	newArr = append(newArr, this.arr[i])
	// }	
	// this.arr = newArr
	// end

	this.length --

	return popVal
}

func evalRPN(tokens []string) int {
	obj := Stack{}
	opMap := make(map[string]string)
	opMap["+"] = "+"
	opMap["-"] = "-"
	opMap["*"] = "*"
	opMap["/"] = "/"
	for _, token := range(tokens) {
		op, ok := opMap[token]; if ok {
			obj2 := obj.Pop()
			obj1 := obj.Pop()
			// fmt.Println("Pop: ", obj1, obj2)
			if op == "+" {
				obj.Push(obj1 + obj2)
			} else if op == "-" {
				obj.Push(obj1 - obj2)
			} else if op == "*" {
				obj.Push(obj1 * obj2)
			} else if op == "/" {
				obj.Push(obj1 / obj2)
			}
		} else {
			num, _ := strconv.Atoi(string(token))
			obj.Push(num)
			// fmt.Println("Push: ", num)
		}
	}
	return obj.Pop()
}

func main() {
	// tokens := []string{"2","1","+","3","*"}

	tokens := []string{"10","6","9","3","+","-11","*","/","*","17","+","5","+"}

	fmt.Println(evalRPN(tokens))
}