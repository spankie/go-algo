package main

import (
	"testing"
)

func TestReverse(t *testing.T) {
	vals := []struct {
		Name     string
		Expected string
	}{
		{Name: "", Expected: ""},
		{Name: "a", Expected: "a"},
		{Name: "ab", Expected: "ba"},
		{Name: "abc", Expected: "cba"},
		{Name: "abcd", Expected: "dcba"},
		{Name: "abcde", Expected: "edcba"},
		{Name: "abcdef", Expected: "fedcba"},
		{Name: "abcdefg", Expected: "gfedcba"},
		{Name: "abcdefgh", Expected: "hgfedcba"},
		{Name: "abcdefghi", Expected: "ihgfedcba"},
		{Name: "abcdefghij", Expected: "jihgfedcba"},
		{Name: "abcdefghijk", Expected: "kjihgfedcba"},
		{Name: "abcdefghijkl", Expected: "lkjihgfedcba"},
		{Name: "abcdefghijklm", Expected: "mlkjihgfedcba"},
		{Name: "abcdefghijklmn", Expected: "nmlkjihgfedcba"},
		{Name: "abcdefghijklmnopqrstuvwxyzabcdefghijk", Expected: "kjihgfedcbazyxwvutsrqponmlkjihgfedcba"}}
	for _, v := range vals {
		rev := Reverse(v.Name)
		if rev != v.Expected {
			t.Errorf("Expected: %v; Got: %v\n", v.Expected, rev)
		}
	}
}

func TestReverseWithOtherChars(t *testing.T) {
	vals := []struct {
		Name     string
		Expected string
	}{
		{Name: "", Expected: ""},
		{Name: "a", Expected: "a"},
		{Name: "ab", Expected: "ba"},
		{Name: "abc", Expected: "cba"},
		{Name: "abcd", Expected: "dcba"},
		{Name: "abcde", Expected: "edcba"},
		{Name: "abcdef", Expected: "fedcba"},
		{Name: "abcdefg", Expected: "gfedcba"},
		{Name: "abcdefgh", Expected: "hgfedcba"},
		{Name: "abcdefghi", Expected: "ihgfedcba"},
		{Name: "abcdefghij", Expected: "jihgfedcba"},
		{Name: "abcdefghijk", Expected: "kjihgfedcba"},
		{Name: "abcdefghijkl", Expected: "lkjihgfedcba"},
		{Name: "abcdefghijklm", Expected: "mlkjihgfedcba"},
		{Name: "abcdefghijklmn", Expected: "nmlkjihgfedcba"},
		{Name: "abcdefghijklmnopqrstuvwxyzabcdefghijk", Expected: "kjihgfedcbazyxwvutsrqponmlkjihgfedcba"}}
	for _, v := range vals {
		rev := ReverseWithOtherChars(v.Name)
		if rev != v.Expected {
			t.Errorf("Expected: %v; Got: %v\n", v.Expected, rev)
		}
	}
}
