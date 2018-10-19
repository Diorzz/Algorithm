package search

func BinarySearch(A []int, key int) int {
	left := 0
	right := len(A) - 1
	for left <= right {
		mid := (left + right) / 2
		if A[key] < A[mid] {
			right = mid - 1
		} else if A[key] > A[mid] {
			left = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
