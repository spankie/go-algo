package main

func NumInList(list []int, num int) bool {
	Nummap := map[int]struct{}{}

	for _, v := range list {
		if _, ok := Nummap[v]; !ok {
			Nummap[v] = struct{}{}
		}
	}
	_, ok := Nummap[num]

	return ok
}
