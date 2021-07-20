package main

import "testing"

func TestNumInList(t *testing.T) {
	vals := []struct {
		Name     string
		List     []int
		Expected bool
		Target   int
	}{
		{
			Name:     "sometest",
			List:     []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			Expected: false,
			Target:   20,
		},
	}
	for _, v := range vals {
		isPresent := NumInList(v.List, v.Target)
		if isPresent != v.Expected {
			t.Errorf("Expected value %v; Got: %v", v.Expected, isPresent)
		}
	}

}
