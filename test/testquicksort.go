package test

func QuickSort(A []int, start, end int) {
	if start >= end {
		return
	}

	q := partition(A, start, end)
	QuickSort(A, start, q-1)
	QuickSort(A, q+1, end)

}

func partition(A []int, left, right int) int {
	i := left
	pivot := A[right]
	for j := left; j <= right; j++ {
		if A[j] < pivot {
			A[j], A[i] = A[i], A[j]
			i++
		}
	}

	A[i], A[right] = A[right], A[i]
	return i
}
