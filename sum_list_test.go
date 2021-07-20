package main

import "testing"

func TestSumList(t *testing.T) {
	vals := []struct {
		Name     string
		List     []int
		Expected int
	}{
		{
			Name:     "Correct Values",
			List:     []int{1, 2, 3, 4, 5},
			Expected: 15,
		},
		{
			Name:     "Correct Values with even length",
			List:     []int{1, 2, 3, 4, 5, 6},
			Expected: 21,
		},
	}
	for _, v := range vals {
		sum := SumList(v.List)
		if sum != v.Expected {
			t.Errorf("Expected: %v; Got: %v\n", v.Expected, sum)
		}
	}
}
