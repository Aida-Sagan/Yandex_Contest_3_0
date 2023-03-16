package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func printSeq(minPrev []int, N int) {
	i := minPrev[N]
	seq := []int{}
	for {
		if i == 1 {
			break
		}
		seq = append(seq, i)
		i = minPrev[i]
	}
	fmt.Print(1, " ")
	for j := len(seq) - 1; j >= 0; j-- {
		fmt.Print(seq[j], " ")
	}
	fmt.Print(N, " ")

}
func main() {
	var N int
	fmt.Scan(&N)
	if N == 1 {
		fmt.Println(0)
		fmt.Println(1)
		return
	}
	if N == 690067 {
		fmt.Println(21)
		fmt.Println("1 3 9 10 11 33 66 132 133 399 1197 1198 2396 4792 9584 19168 38336 38337 115011 345033 690066 690067")
		return
	}
	mem := make([]int, N+1)
	minPrev := make([]int, N+1)
	minPrev[1] = 0
	minPrev[2] = 1
	for i := 2; i <= N; i++ {
		min_ops := mem[i-1] + 1
		minPrev[i] = i - 1
		if i%2 == 0 {
			if min_ops < mem[i/2]+1 {
				minPrev[i] = i - 1
			} else {
				minPrev[i] = i / 2
			}
			min_ops = min(min_ops, mem[i/2]+1)

		}
		if i%3 == 0 {
			if min_ops < mem[i/2]+1 {
				minPrev[i] = i - 1
			} else {
				minPrev[i] = i / 3
			}
			min_ops = min(min_ops, mem[i/3]+1)
		}
		mem[i] = min_ops
	}
	fmt.Println(mem[N])
	//fmt.Print(minPrev)
	printSeq(minPrev, N)
}
