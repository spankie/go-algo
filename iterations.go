package main

import (
	"fmt"
	"strconv"
	"strings"
)

func binaryGap(N int) int {
	// convert N to binary
	binary := strconv.FormatInt(int64(N), 2)
	split := strings.Split(binary, "1")
	fmt.Println(binary, split)
	// start := 0
	// end := 0
	l := len(split)
	highest := 0
	for i, v := range split {
		if i == 0 || i == l-1 {
			continue
		}
		if len(v) > highest {
			highest = len(v)
		}
		fmt.Printf("v: %v, type: %T\n", v, v)
	}
	return highest
}

func main() {
	binaryGap(22)
	fmt.Println("-----------------------------")
	binaryGap(32)
}
