package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadA(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for iLine := 1; iLine <= y; iLine++ {
		for iAlph := 1; iAlph <= x; iAlph++ {
			if iLine == 1 || iLine == y {
				if iAlph == 1 || iAlph == x {
					z01.PrintRune('o')
				} else {
					z01.PrintRune('-')
				}
			} else {
				if iAlph == 1 || iAlph == x {
					z01.PrintRune('|')
				} else {
					z01.PrintRune(' ')
				}
			}
		}
		z01.PrintRune('\n')
	}
}

func main() {
	if len(os.Args) != 3 {
		println("Usage: ./quadA x y")
		return
	}

	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || x <= 0 || y <= 0 {
		println("Invalid arguments")
		return
	}
	QuadA(x, y)
}
