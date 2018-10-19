package test

import "testing"

func TestQuickSort(t *testing.T) {
	tests := []struct{ data, result []int }{
		{data: []int{3, 2, 1, 4, 5}, result: []int{1, 2, 3, 4, 5}},
		{data: []int{}, result: []int{}},
		{data: []int{3}, result: []int{3}},
	}

	for _, test := range tests {
		if QuickSort(test.data, 0, len(test.data)-1); !checkSlice(test.data, test.result) {
			t.Errorf("error")
		}
	}
}
