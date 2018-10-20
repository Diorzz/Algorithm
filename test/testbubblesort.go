package test

func BubbleSort(A []int, n int) {

	for i := 0; i <= n; i++ {
		flag := false
		for j := 0; j < n-i; j++ {
			if A[j] > A[j+1] {
				A[j], A[j+1] = A[j+1], A[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}
}
