package test

func MergeSort(A []int, start, end int, tmp []int) {
	if start >= end {
		return
	}

	q := (start + end) / 2
	MergeSort(A, start, q, tmp)
	MergeSort(A, q+1, end, tmp)
	merge(A, start, q, end, tmp)

}

func merge(A []int, left, mid, right int, tmp []int) {
	i, j, t := left, mid+1, 0
	for i <= mid && j <= right {
		if A[i] < A[j] {
			tmp[t] = A[i]
			i++
			t++
		} else {
			tmp[t] = A[j]
			j++
			t++
		}
	}

	for i <= mid {
		tmp[t] = A[i]
		i++
		t++
	}
	for j <= right {
		tmp[t] = A[j]
		j++
		t++
	}

	t = 0
	for left <= right {
		A[left] = tmp[t]
		left++
		t++
	}

}
