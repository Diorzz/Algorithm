package test

import "testing"

func checkSlice(a, b []int) bool {
	la := len(a)
	lb := len(b)
	if la != lb {
		return false
	}

	for i := 0; i < la; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true

}

func TestMergeSort(t *testing.T) {
	tests := []struct{ data, result []int }{
		{data: []int{3, 2, 1, 4, 5}, result: []int{1, 2, 3, 4, 5}},
		{data: []int{}, result: []int{}},
		{data: []int{3}, result: []int{3}},
	}

	for _, test := range tests {
		tmp := make([]int, len(test.data))
		if MergeSort(test.data, 0, len(test.data)-1, tmp); !checkSlice(test.data, test.result) {
			t.Errorf("error")
		}
	}
}
