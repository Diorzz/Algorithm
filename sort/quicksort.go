package sort

func QuickSort(A []int, start, end int) {
	if start >= end {
		return
	}

	q := partition(A, start, end)
	QuickSort(A, start, q-1)
	QuickSort(A, q+1, end)
}

func partition(A []int, left, right int) int {
	pivot := A[right]
	i := left
	for j := left; j < right; j++ {
		if A[j] < pivot {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}

	A[i], A[right] = A[right], A[i]
	return i
}
