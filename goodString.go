package main

import "fmt"

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}

func goodString(n int, counts []int) int {
	//key -  sym,
	//val - количество повторов
	hash := make(map[rune]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&counts[i])
		hash['a'+rune(i)] = counts[i]
	}

	maxGoodStr := 0

	for i := 'a'; i <= 'z'; i++ {
		maxGoodStr += min(hash[i], hash[i+1])

	}
	return maxGoodStr
}

func main() {
	var n int
	fmt.Scan(&n)
	counts := make([]int, n)

	fmt.Println(goodString(n, counts))
}
