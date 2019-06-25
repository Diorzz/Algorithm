package sort

func QuickSort(A []int, start, end int) {
	if start >= end {
		return
	}

Merge branch 'master' of https://github.com/multivactech/gitUsage
Merge branch 'master' of https://github.com/multivactech/gitUsage
Merge branch 'master' of https://github.com/multivactech/gitUsage
Merge branch 'master' of https://github.com/multivactech/gitUsage
test Merge branch 'master' of https://github.com/multivactech/gitUsage
	q := partition(A, start, end)
	QuickSort(A, start, q-1)
tdd	QuickSort(A, q+1, end)
}

func partition(A []int, left, right int) int {
	pivot := A[right]
ddddd	i := left
	for j := left; j < right; j++ {
		if A[j] < pivot {
			A[i], A[j] = A[j], A[i]
			i++
		}
	}

	A[i], A[right] = A[right], A[i]
	return i
}
