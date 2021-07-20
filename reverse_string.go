package main

import (
	"strings"
)

func Reverse(str string) string {
	var reversed strings.Builder
	for i := len(str) - 1; i >= 0; i-- {
		reversed.WriteByte(str[i])
	}
	return reversed.String()
}

func ReverseWithOtherChars(str string) string {
	// var reversed strings.Builder
	var reversed string
	for _, char := range str {
		reversed = string(char) + reversed
		// reversed.WriteRune(char)
	}
	return reversed
}
