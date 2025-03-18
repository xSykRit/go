package main

import "github.com/01-edu/z01"

func main() {
	n := 5 // You can change this number to test other values
	if n%2 == 0 {
		printStr("Even")
	} else {
		printStr("Odd")
	}
}

func printStr(s string) {
	for _, c := range s {
		z01.PrintRune(c)
	}
	z01.PrintRune('\n')
}
