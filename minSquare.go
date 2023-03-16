package main

import "fmt"

func main() {
	var k, minX, minY, maxX, maxY int
	var x, y int
	fmt.Scan(&k)
	fmt.Scan(&x, &y)
	minX = x
	maxX = x
	minY = y
	maxY = y
	for i := 1; i < k; i++ {
		fmt.Scan(&x, &y)
		if x > maxX {
			maxX = x
		}
		if x < minX {
			minX = x
		}
		if y > maxY {
			maxY = y
		}
		if y < minY {
			minY = y
		}
	}

	fmt.Println(minX, minY, maxX, maxY)

}
