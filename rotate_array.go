package main

/*
[3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
[6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
[7, 6, 3, 8, 9] -> [9, 7, 6, 3, 8]
[9, 7, 6, 3, 8] -> [8, 9, 7, 6, 3]
[8, 9, 7, 6, 3] -> [3, 8, 9, 7, 6]
[3, 8, 9, 7, 6] -> [6, 3, 8, 9, 7]
[6, 3, 8, 9, 7] -> [7, 6, 3, 8, 9]
handle the following:
- one element, 0 <= K <= 5
- small functional tests, K >= N
- maximal N and K
*/
func RotateArray(A []int, K int) []int {
	l := len(A)
	if l < 2 {
		return A
	}
	result := make([]int, l)
	for i, v := range A {
		result[(i+K)%l] = v
		// if K+i >= l {
		// 	result[(K+i)-l] = v
		// } else {
		// 	result[K+i] = v
		// }
	}
	return result
}

// result[(i+K)%len(A)] = v
