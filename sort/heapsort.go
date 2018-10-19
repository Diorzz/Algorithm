package sort

func maxheapify(A []int, i, n int) {
	l := 2*i + 1
	r := 2*i + 1 + 1

	largest := i
	if l <= n && A[l] > A[i] {
		largest = l
	}

	if r <= n && A[r] > A[i] {
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
	//n := len(A) - 1
	for i := len(A) - 1; i >= 1; i-- {
		A[0], A[i] = A[i], A[0]
		//n--
		maxheapify(A, 0, i-1)
	}
}
