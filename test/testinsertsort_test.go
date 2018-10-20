package test

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	tests := []TestData{
		{data: []int{3, 2, 1, 4, 5}, result: []int{1, 2, 3, 4, 5}},
		{data: []int{}, result: []int{}},
		{data: []int{3}, result: []int{3}},
		{data: []int{3, 4, 5, 7, 6, 1}, result: []int{1, 3, 4, 5, 6, 7}},
		{data: []int{1, 4, 5, 7, 6, 3}, result: []int{1, 3, 4, 5, 6, 7}},
	}

	for _, test := range tests {
		if InsertSort(test.data, len(test.data)-1); !checkSlice(test.data, test.result) {
			t.Errorf("error")
		}
	}
}
