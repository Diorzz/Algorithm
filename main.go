package main

import (
	"fmt"
	"go/sort"
)

func main() {
	a := []int{4, 1, 2, 7, 9, 8, 0}
        b:=[]int{0,0}
	sort.HeapSort(a)
	for _, i := range a {
		fmt.Printf("the value is %d ", i)
	}
	fmt.Println("*****************zhangzheng is handsome*******");
}
