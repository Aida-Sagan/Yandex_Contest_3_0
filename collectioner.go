package main

import (
	"fmt"
	"sort"
)

func removeDuplicates(arr []int) []int {
	newArr := []int{}
	prevEl := 0
	for idx, el := range arr {
		if el != prevEl || idx == 0 {
			newArr = append(newArr, el)
		}
		prevEl = el
	}
	return newArr
}
func floor(L int, R int) int {
	if ((L+R)/2)*2 > L+R {
		return ((L + R) / 2) - 1
	}
	return (L + R) / 2
}
func binarySearch(arr []int, target int) int {
	n := len(arr)
	L := 0
	R := n
	m := 0
	for {
		if L >= R {
			break
		}
		m = floor(L, R)
		if arr[m] < target {
			L = m + 1
		} else {
			R = m
		}

	}
	return L
}
func main() {
	var N, K, newEl int
	collection := []int{}

	fmt.Scan(&N)
	for i := 0; i < N; i++ {
		fmt.Scan(&newEl)
		collection = append(collection, newEl)
	}
	sort.Ints(collection)
	readyCollection := removeDuplicates(collection)
	fmt.Scan(&K)
	for i := 0; i < K; i++ {
		fmt.Scan(&newEl)
		fmt.Println(binarySearch(readyCollection, newEl))
	}

}
