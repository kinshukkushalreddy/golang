package main

import "fmt"

func main() {
	str := "kinshuk reddy"
	count := make(map[rune]int)

	for _, c := range str {
		count[c]++
	}

	for c, n := range count {
		fmt.Printf("%c: %d\n", c, n)
	}
}
