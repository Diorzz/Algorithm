package sort

import (
	"testing"
)

func TestSelectSort(t *testing.T) {
	tests := GetSortData()

	for _, test := range tests {
		if SelectSort(test.data, len(test.data)-1); !checkSlice(test.data, test.result) {
			t.Errorf("error")
		}
	}
}
