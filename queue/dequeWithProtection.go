package main

import "fmt"

type Deque struct {
	nums []int
}

func (d *Deque) push_front(n int) {
	d.nums = append([]int{n}, d.nums...)
}

func (d *Deque) push_back(n int) {
	(d.nums) = append((d.nums), n)
}

func (d *Deque) pop_front() int {
	el := (d.nums)[0]
	(d.nums) = (d.nums)[1:]
	return el
}

func (d *Deque) pop_back() int {
	el := (d.nums)[len(d.nums)-1]
	(d.nums) = (d.nums)[:len(d.nums)-1]
	return el
}

func (d *Deque) Front() int {
	el := (d.nums)[0]
	return el
}

func (d *Deque) back() int {
	el := (d.nums)[len(d.nums)-1]
	return el
}

func (d *Deque) size() int {
	return len(d.nums)
}

func (d *Deque) clear() {
	(d.nums) = []int{}
}

func main() {
	var deque Deque
	var cmd string

	for {
		fmt.Scan(&cmd)

		switch cmd {
		case "push_front":
			var n int
			fmt.Scan(&n)
			deque.push_front(n)
			fmt.Println("ok")
		case "push_back":
			var n int
			fmt.Scan(&n)
			deque.push_back(n)
			fmt.Println("ok")
		case "pop_front":
			if len(deque.nums) == 0 {
				fmt.Println("error")
			} else {
				fmt.Println(deque.pop_front())
			}

		case "pop_back":
			if len(deque.nums) == 0 {
				fmt.Println("error")
			} else {
				fmt.Println(deque.pop_back())
			}
		case "front":
			if len(deque.nums) == 0 {
				fmt.Println("error")
			} else {
				fmt.Println(deque.Front())
			}
		case "back":
			if len(deque.nums) == 0 {
				fmt.Println("error")
			} else {
				fmt.Println(deque.back())
			}
		case "size":
			fmt.Println(deque.size())
		case "clear":
			deque.clear()
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			return
		}
	}
}
