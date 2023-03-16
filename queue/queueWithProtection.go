package main

import "fmt"

type Queue []int

func (q *Queue) push(n int) {
	*q = append(*q, n)
}

func (q *Queue) pop() int {
	el := (*q)[0]
	*q = (*q)[1:]
	return el
}

func (q *Queue) front() int {
	el := (*q)[0]
	return el
}

func (q *Queue) size() int {
	return len(*q)
}

func (q *Queue) clear() {
	*q = []int{}
}

func main() {
	var queue Queue
	var cmd string

	for {
		fmt.Scan(&cmd)

		switch cmd {
		case "push":
			var n int
			fmt.Scan(&n)
			queue.push(n)
			fmt.Println("ok")

		case "pop":
			if len(queue) == 0 {
				fmt.Println("error")
			} else {
				el := queue.front()
				fmt.Println(el)
				queue.pop()
			}

		case "front":
			if len(queue) == 0 {
				fmt.Println("error")
			} else {
				fmt.Println(queue.front())
			}

		case "size":
			fmt.Println(queue.size())

		case "clear":
			queue.clear()
			fmt.Println("ok")

		case "exit":
			fmt.Println("bye")
			return
		}

	}

}
