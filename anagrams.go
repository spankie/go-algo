package main

func isAnagram(s1, s2 string) bool {
	// create an array of ascii character
	ascii := make([]int, 128)
	for _, v := range s1 {
		ascii[v] += 1
	}

	for _, v := range s2 {
		ascii[v] -= 1
	}

	for _, v := range ascii {
		if v > 0 {
			return false
		}
	}
	return true
}
