package main

import "fmt"

type Stack []rune

func (s *Stack) IsEmpty() string {
	if len(*s) == 0 {
		return "yes"
	}
	return "no"
}

func isSequence(s string) string {
	var stack Stack

	for _, ch := range s {
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)

		} else if stack.IsEmpty() == "no" {
			top := stack[len(stack)-1]

			if (top == '(' && ch == ')') || (top == '[' && ch == ']') || (top == '{' && ch == '}') {
				stack = stack[:len(stack)-1]
			} else {
				return "no"
			}
		} else {
			return "no"
		}
	}
	return stack.IsEmpty()
}

func main() {
	var s string
	fmt.Scan(&s)
	fmt.Println(isSequence(s))
}
