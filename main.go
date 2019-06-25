package main

import (
	"fmt"
	"go/sort"
)

func main() {
	fmt.Println("*****************zhangzheng is handsome*******");
	a := []int{4, 1, 2, 7, 9, 8, 0}
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
        b:=[]int{0,0}
	sort.HeapSort(a)
	for _, i := range a {
		fmt.Printf("the value is %d ", i)
	}
	fmt.Println("*****************zhangzheng is handsome*******");
}
