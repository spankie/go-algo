package main

import "testing"

func TestBinaryGap(t *testing.T) {
	testCases := []struct {
		desc string
		N    int
		want int
	}{
		{
			desc: "9",
			N:    9,
			want: 2,
		},
		{
			desc: "32",
			N:    32,
			want: 0,
		},
		{
			desc: "529",
			N:    529,
			want: 4,
		},
		{
			desc: "20",
			N:    20,
			want: 1,
		},
		{
			desc: "15",
			N:    15,
			want: 0,
		},
		{
			desc: "1041",
			N:    1041,
			want: 5,
		},
		{
			desc: "32",
			N:    32,
			want: 0,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := binaryGap(tC.N)
			if got != tC.want {
				t.Errorf("binaryGap(%d) = %d; want %d", tC.N, got, tC.want)
			}
		})
	}
}
