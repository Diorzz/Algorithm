package main

import (
	"fmt"
	"go/sort"
)

func main() {
	a := []int{4, 1, 2, 7, 9, 8, 0}
	sort.HeapSort(a)
	for _, i := range a {
		fmt.Printf("%d ", i)
	}
}
