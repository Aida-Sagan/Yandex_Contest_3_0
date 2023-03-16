package main

import "fmt"

func f(n int) int{
	dp := []int{2, 4, 7}
	if n == 1 {
		return dp[0]
	} else if n == 2 {
		return dp[1]
	} else if n == 3 {
		return dp[2]
	}
	return f(n-1)+f(n-2)+f(n-3)
	
}

func main() {
	var n int 
	fmt.Scan(&n)

	fmt.Println(f(n))
}