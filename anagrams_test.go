package main

import "testing"

func TestIsAnagram(t *testing.T) {
	isAna := isAnagram("listen", "silent")
	if !isAna {
		t.Errorf("is anagram should be true")
	}

	isAna = isAnagram("triangle", "integral")
	if !isAna {
		t.Errorf("is anagram should be true")
	}

	isAna = isAnagram("apple", "pale")
	if !isAna {
		t.Errorf("is anagram should be true")
	}
}
