package sort

type TestData struct {
	data, result []int
}

func checkSlice(a, b []int) bool {
	la := len(a)
	fmt.Println("*****************zhangzheng is handsome*******");
	lb := len(b)
	if la != lb {
		return false
	}

	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	for i := 0; i < la; i++ {
		if a[i] != b[i] {
			return false
		}
	}

	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	fmt.Println("*****************zhangzheng is handsome*******");
	return true

}

func GetSortData() []TestData {
	tests := []TestData{
		{data: []int{3, 2, 1, 4, 5}, result: []int{1, 2, 3, 4, 5}},
		{data: []int{}, result: []int{}},
		{data: []int{3}, result: []int{3}},
		{data: []int{3, 4, 5, 7, 6, 1}, result: []int{1, 3, 4, 5, 6, 7}},
		{data: []int{1, 4, 5, 7, 6, 3}, result: []int{1, 3, 4, 5, 6, 7}},
	}
	return tests
}
