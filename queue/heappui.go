package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Heap []int

func (h *Heap) Insert(elem int) {
	*h = append(*h, elem)
	//меняет местами элементы, если добавленный элемент больше своего родителя
	(*h).FromElemToRoot(len(*h) - 1)
}

func (h *Heap) FromElemToRoot(ind int) {
	if ind == 0 {
		return
	}
	parent := (ind - 1) / 2
	//меняет местами элементы, если добавленный элемент больше своего родителя
	if (*h)[ind] > (*h)[parent] {
		(*h).swap(ind, parent)
		(*h).FromElemToRoot(parent)
	}
}

func (h *Heap) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Extract() int {
	/*
		наибольший эелмент -- это элемент в вершине кучи (корень)
			меняет местами корневой элемент с последним элементом в срезе
			 и затем "просматривает" путь от корня к листьям и меняет местами элементы, если корневой элемент меньше своих детей.
	*/
	last_ind := len(*h) - 1
	maximum := (*h)[0]

	(*h).swap(0, last_ind)
	*h = (*h)[:last_ind]
	(*h).FromRootToElem(0)

	return maximum
}

func (h *Heap) FromRootToElem(ind int) {
	//элменты опускаем вниз
	for ind < len(*h)-1 {
		left := 2*ind + 1
		right := 2*ind + 2
		max_ind := ind

		if left < len(*h) && (*h)[left] > (*h)[max_ind] {
			max_ind = left
		}

		if right < len(*h) && (*h)[right] > (*h)[max_ind] {
			max_ind = right
		}

		if max_ind != ind {
			(*h).swap(max_ind, ind)
			ind = max_ind
		} else {
			break
		}
	}
}

func scanInt(scanner *bufio.Scanner) int {
	scanner.Scan()
	num, _ := strconv.Atoi(scanner.Text())
	return num
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	n := scanInt(scanner)
	heap := Heap{}

	for i := 0; i < n; i++ {
		cmd := scanInt(scanner)

		switch cmd {
		case 0:
			num := scanInt(scanner)
			heap.Insert(num)
		case 1:
			max_val := heap.Extract()
			fmt.Println(max_val)
		}
	}
}
