package main

import "fmt"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func maxValue(hash map[byte]int) int {
	num := 0

	for _, val := range hash {
		num = max(num, val)
	}
	return num
}

func beautyString(k int, s string) int {
	left := 0
	maxLen := 0

	hash := make(map[byte]int)

	for right := range s {
		hash[s[right]]++

		windowLen := right - left + 1

		statement := windowLen - maxValue(hash)

		if statement <= k {
			maxLen = max(maxLen, windowLen)
		} else {
			hash[s[left]]--
			left++
		}
	}

	return maxLen
}

func main() {
	var k int
	var s string
	fmt.Scan(&k)
	fmt.Scan(&s)

	fmt.Println(beautyString(k, s))
}
