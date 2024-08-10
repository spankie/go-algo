package main

import (
	"reflect"
	"testing"
)

func TestRotateArray(t *testing.T) {
	testCases := []struct {
		desc string
		A    []int
		K    int
		want []int
	}{
		{
			desc: "Rotate 1",
			A:    []int{3, 8, 9, 7, 6},
			K:    1,
			want: []int{6, 3, 8, 9, 7},
		},
		{
			desc: "Rotate 7",
			A:    []int{3, 8, 9, 7, 6},
			K:    7,
			want: []int{7, 6, 3, 8, 9},
		},
		{
			desc: "Rotate 3",
			A:    []int{6, 3, 8, 9, 7},
			K:    3,
			want: []int{8, 9, 7, 6, 3},
		},
		{
			desc: "Rotate 4",
			A:    []int{9, 7, 6, 3, 8},
			K:    4,
			want: []int{7, 6, 3, 8, 9},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := RotateArray(tC.A, tC.K)
			if !reflect.DeepEqual(got, tC.want) {
				t.Errorf("RotateArray(%v, %d) = %v; want %v", tC.A, tC.K, got, tC.want)
			}
		})
	}
}
