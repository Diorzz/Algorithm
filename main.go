package main

import (
	"fmt"
	"go/search"
)

func main() {
	a := []int{4, 1, 2, 7, 9, 8, 0}
	r := search.BinarySearch(a, 3)
	fmt.Println(r)
}
