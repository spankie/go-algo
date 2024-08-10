package main

import "testing"

func TestOddOccurence(t *testing.T) {
	i := 301
	arr := make([]int, i)
	for i > 0 {
		if i%2 == 0 {
			arr[i-1] = 7
		} else {
			arr[i-1] = 9
		}
		i--
	}
	arr[0] = 42
	r := OddOccurenceCopilot(arr)
	if r != 42 {
		t.Errorf("want %v got %v", 3, r)
	}
}
