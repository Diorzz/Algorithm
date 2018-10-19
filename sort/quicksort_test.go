package sort

import "testing"

func TestQuickSort(t *testing.T) {
	tests := GetSortData()

	for _, test := range tests {
		if QuickSort(test.data, 0, len(test.data)-1); !checkSlice(test.data, test.result) {
			t.Errorf("error")
		}
	}
}
