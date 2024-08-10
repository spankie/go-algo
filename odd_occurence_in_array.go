package main

import "log"

func OddOccurence(A []int) int {
	if len(A) == 1 {
		return A[0]
	}
	Amap := make(map[int]int)
	for _, v := range A {
		Amap[v] += 1
	}

	for i, v := range Amap {
		log.Printf("i = %v - v = %v", i, v)
		if v == 1 {
			return i
		}
	}
	return 0
}

// from Copilot
func OddOccurenceCopilot(A []int) int {
	result := 0
	for _, v := range A {
		result ^= v
	}
	return result
}
