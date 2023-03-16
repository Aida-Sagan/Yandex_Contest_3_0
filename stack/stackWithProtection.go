package main

import (
	"fmt"
)

type Stack []int32

func (s *Stack) push(n int32) {
	*s = append(*s, n)
}

func (s *Stack) pop() int32 {

	elem := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return elem
}

func (s *Stack) back() int32 {

	res := (*s)[len(*s)-1]
	return res
}

func (s *Stack) size() int {
	return len(*s)
}

func (s *Stack) clear() {
	*s = []int32{}
}

func main() {
	var stack Stack

	for {
		var cmd string
		fmt.Scan(&cmd)

		switch cmd {
		case "push":
			var n int32
			fmt.Scan(&n)
			stack.push(n)
			fmt.Println("ok")

		case "pop":
			if len(stack) == 0 {
				fmt.Println("error")
			} else {
				fmt.Println(stack.back())
				stack.pop()
			}

		case "back":
			if len(stack) == 0 {
				fmt.Println("error")
			} else {
				fmt.Println(stack.back())
			}

		case "size":
			fmt.Println(stack.size())

		case "clear":
			stack.clear()
			fmt.Println("ok")

		case "exit":
			fmt.Println("bye")
			return
		}
	}
}
