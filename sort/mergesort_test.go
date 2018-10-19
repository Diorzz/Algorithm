package sort

import (
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := GetSortData()

	for _, test := range tests {
		tmp := make([]int, len(test.data))
		if Mergesort(test.data, 0, len(test.data)-1, tmp); !checkSlice(test.data, test.result) {
			t.Errorf("error")
		}
	}
}
