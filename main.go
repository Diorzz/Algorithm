package main

import (
	"fmt"
	"go/sort"
)

func main() {
	fmt.Println("*****************zhangzheng is a cute boy*******");
	a := []int{4, 1, 2, 7, 9, 8, 0}
	fmt.Println("*****************zhangzheng is not handsome*******");
	fmt.Println("*****************zhangzheng is not handsome*******");
	fmt.Println("*****************zhangzheng is not handsome*******");
	fmt.Println("*****************zhangzheng is not handsome*******");
	fmt.Println("*****************zhangzheng is not handsome*******");
        b:=[]int{0,0}
	sort.HeapSort(a)
	for _, i := range a {
		fmt.Printf("the value is %d ", i)
	}
	fmt.Println("*****************zhangzheng is handsome*******");
}
