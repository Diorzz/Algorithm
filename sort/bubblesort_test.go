package sort

import (
	"testing"
)

func TestBubbleSort(t *testing.T) {
	tests := GetSortData()

	for _, test := range tests {
		if BubbleSort(test.data, len(test.data)-1); !checkSlice(test.data, test.result) {
			t.Errorf("error")
		}
	}
}
