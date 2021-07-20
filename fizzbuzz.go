package main

import (
	"fmt"
	"strings"
)

func FizzBuzz(N int) {
	res := []string{}
	for i := 1; i <= N; i++ {
		if i%3 == 0 && i%5 == 0 {
			res = append(res, "Fizz Buzz")
		} else if i%3 == 0 {
			res = append(res, "Fizz")
		} else if i%5 == 0 {
			res = append(res, "Buzz")
		} else {
			res = append(res, fmt.Sprintf("%d", i))
		}
	}
	fmt.Println(strings.Join(res, ", "))
}

func FizzBuzzWithoutJoin(N int) {
	for i := 1; i < N; i++ {
		printFizzBuzz(i)
		fmt.Print(", ")
	}
	printFizzBuzz(N)
	fmt.Println()
}

func printFizzBuzz(i int) {
	if i%3 == 0 && i%5 == 0 {
		fmt.Print("Fizz Buzz")
	} else if i%3 == 0 {
		fmt.Print("Fizz")
	} else if i%5 == 0 {
		fmt.Print("Buzz")
	} else {
		fmt.Print(i)
	}
}
