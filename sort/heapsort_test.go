package sort

import (
	"testing"
)

func TestHeapSort(t *testing.T) {
	tests := GetSortData()

	for _, test := range tests {
		if HeapSort(test.data); !checkSlice(test.data, test.result) {
			t.Errorf("error")
		}
	}
}
