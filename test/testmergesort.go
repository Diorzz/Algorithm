package test

func MergeSort(A []int, start, end int, tmp []int) {
	if start >= end {
		return
	}

	mid := (start + end) / 2
	MergeSort(A, start, mid, tmp)
	MergeSort(A, mid+1, end, tmp)
	merge(A, start, mid, end, tmp)
}

func merge(A []int, left, mid, right int, tmp []int) {
	i, j, t := left, mid+1, 0
	for i <= mid && j <= right {
		if A[i] < A[j] {
			tmp[t] = A[i]
			i++
			t++
		}
		if A[j] < A[i] {
			tmp[t] = A[j]
			j++
			t++
		}
	}

	for i <= mid {
		tmp[t] = A[i]
		t++
		i++
	}
	for j <= right {
		tmp[t] = A[j]
		t++
		j++
	}

	t = 0
	for left <= right {
		A[left] = tmp[t]
		left++
		t++
	}
}
