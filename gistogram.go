package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"sort"
)

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

func gistogramma(text []string) [][]string {
	hash := make(map[byte]int)
	height := 0

	for _, str := range text {
		for _, ch := range str {
			if ch != ' ' && ch != '\n' {
				hash[byte(ch)]++
			}
		}

	}
	height = maxValue(hash) + 1
	table := make([][]string, height)

	keys := make([]string, 0)
	for k := range hash {
		keys = append(keys, string(k))

	}
	sort.Strings(keys)

	for _, key := range keys {
		table[height-1] = append(table[height-1], string(key))
	}

	for i := height - 2; i >= 0; i-- {
		for j := 0; j < len(hash); j++ {

			if hash[byte(table[height-1][j][0])] > 0 {
				table[i] = append(table[i], string('#'))
			} else {
				table[i] = append(table[i], string(' '))
			}
			hash[byte(table[height-1][j][0])]--

		}

	}

	return table
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var strSlice []string
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			panic(err)
		}
		strSlice = append(strSlice, text)

	}
	//fmt.Println(strSlice)
	ans := gistogramma(strSlice)

	for _, str := range ans {
		for _, ch := range str {
			fmt.Print(ch)

		}
		fmt.Println()
	}

}
