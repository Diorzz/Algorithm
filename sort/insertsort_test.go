package sort

import (
	"testing"
)

func TestInsertSort(t *testing.T) {
	tests := GetSortData()

	for _, test := range tests {
		if InsertSort(test.data, len(test.data)-1); !checkSlice(test.data, test.result) {
			t.Errorf("error")
		}
	}
}
