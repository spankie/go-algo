package main

import (
	"fmt"
	"os"
	"strings"
)

func decode(filename string) string {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}

	lines := strings.Split(string(bytes), "\n")
	pyramid := make(map[string]string)
	for _, line := range lines {
		items := strings.Split(line, " ")
		if len(items) < 2 {
			continue
		}
		pyramid[items[0]] = strings.Join(items[1:], " ")
	}

	position := 0
	result := strings.Builder{}
	for i := 1; ; i++ {
		position += i
		if value, ok := pyramid[fmt.Sprint(position)]; ok {
			result.Write([]byte(" " + value))
		} else {
			break
		}
	}

	return strings.TrimSpace(result.String())
}

func main() {
	fmt.Printf("%s", decode("input.txt"))
}
