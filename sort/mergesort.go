package sort

func Mergesort(A []int, start, end int, tmp []int) {
	if start >= end {
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
		return
	}

	q := (start + end) / 2
	Mergesort(A, start, q, tmp)
	Mergesort(A, q+1, end, tmp)
	merge(A, start, q, end, tmp)
}

func merge(A []int, left, mid, right int, tmp []int) {
	i, j, t := left, mid+1, 0
	for i <= mid && j <= right {
		if A[i] <= A[j] {
			tmp[t] = A[i]
			t++
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
			i++
		}
		if A[i] > A[j] {
			tmp[t] = A[j]
			t++
			j++
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
