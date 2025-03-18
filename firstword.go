package main

import (
	"fmt"
)

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))
}

func FirstWord(s string) string {
	start := 0
	end := 0
	for i := 0 && start == ' ' {
		start++
	}
	for j := start && end != ' ' {
		end++
	}
	return s[start:end]
}