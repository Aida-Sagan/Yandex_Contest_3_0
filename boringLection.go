package main

import "fmt"

func countLetters(favWord string) {
	lettersCount := make(map[byte]int)
	n := len(favWord)
	for i := 0; i < 26; i++ {
		lettersCount[byte(97+i)] = 0
	}
	for i := 0; i < len(favWord); i++ {
		lettersCount[favWord[i]] += (n - i) * (i + 1)
	}
	for i := 0; i < 26; i++ {
		if lettersCount[byte(97+i)] != 0 {
			fmt.Printf("%c: %d\n", byte(97+i), lettersCount[byte(97+i)])
		}
	}
}

func main() {
	var favWord string
	fmt.Scan(&favWord)
	countLetters(favWord)

}
