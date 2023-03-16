package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	stack := make([]int, 0)
	expectedWagon := 1

	for i := 0; i < n; i++ {
		var wagon int
		fmt.Scan(&wagon)

		for len(stack) != 0 && stack[len(stack)-1] == expectedWagon {
			stack = stack[:len(stack)-1]
			expectedWagon++
		}

		if wagon == expectedWagon {
			expectedWagon++
			continue
		}

		stack = append(stack, wagon)
	}

	for len(stack) > 0 && stack[len(stack)-1] == expectedWagon {
		stack = stack[:len(stack)-1]
		expectedWagon++
	}

	if len(stack) == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
