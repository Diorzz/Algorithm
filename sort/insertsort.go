package sort

func InsertSort(A []int, n int) {
	for i := 1; i <= n; i++ {
		tmp := A[i]
		j := i - 1
		for ; j >= 0; j-- {
			if A[j] > tmp {
				A[j+1] = A[j]
			} else {
				break
			}
		}
		A[j+1] = tmp
	}
}
