package main

/*
func SumList(list []int) int {
	total := 0
	length := len(list)

	for i := 0; i < length; i += 2 {
		if i+2 >= length {
			total += list[i]
			break
		}
		total += list[i] + list[i+1]
	}
	return total
}
*/

func SumList(list []int) int {
	var total int
	for _, v := range list {
		total += v
	}
	return total
}
