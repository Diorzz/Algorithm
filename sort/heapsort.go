package sort

import (
	"fmt"
)

func maxheapify(A []int, i, n int) {
	l := 2*i + 1
	r := 2*i + 1 + 1
	var largest int

	if l <= n && A[l] > A[i] {
		largest = l
	} else {
		largest = i
	}

	if r <= n && A[r] > A[largest] {
		largest = r
	}

	if largest != i {
		A[i], A[largest] = A[largest], A[i]
		maxheapify(A, largest, n)
	}

}

func buildmaxheap(A []int) {
	n := len(A) - 1
	for i := (n - 1) / 2; i >= 0; i-- {
		maxheapify(A, i, n)
	}
}

func HeapSort(A []int) {
	if len(A) <= 1 {
		return
	}
	buildmaxheap(A)
	for i := len(A) - 1; i >= 1; i-- {
		A[0], A[i] = A[i], A[0]
		maxheapify(A, 0, i-1)
	}
}

func printArray(A []int) {
	for k := 0; k < len(A); k++ {
		fmt.Printf("%d ", A[k])
	}
	fmt.Println("")
}
