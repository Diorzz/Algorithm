package sort
import "fmt"
func SelectSort(A []int, n int) {
	for i := 0; i <= n; i++ {
		min := i
		for j := i + 1; j <= n; j++ {
			if A[i] > A[j] {
				min = j
			}
		}
		fmt.Println("test")
		A[i], A[min] = A[min], A[i]

	}
}
