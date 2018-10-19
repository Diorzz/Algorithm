package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	test := []struct {
		data   []int
		key    int
		result int
	}{
		{data: []int{0, 1, 2, 3, 4, 5, 6, 7}, key: 2, result: 2},
		{data: []int{}, key: 2, result: -1},
	}

	for _, v := range test {
		if r := BinarySearch(v.data, v.key); r != v.result {
			t.Log(r)
			t.Errorf("error")
		}
	}
}
