package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x <= 0 || y <= 0 {
		return
	}
	for liN := 1; liN <= y; liN++ {
		for roW := 1; roW <= x; roW++ {
			if liN == 1 {
				if roW == 1 {
					z01.PrintRune('/')
				} else if roW == x {
					z01.PrintRune('\\')
				} else {
					z01.PrintRune('*')
				}
			} else if liN == y {
				if roW == 1 {
					z01.PrintRune('\\')
				} else if roW == x {
					z01.PrintRune('/')
				} else {
					z01.PrintRune('*')
				}
			} else {
				if roW == 1 {
					z01.PrintRune('*')
				} else if roW == x {
					z01.PrintRune('*')
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
		println("Usage: ./quadB x y")
		return
	}

	x, err1 := strconv.Atoi(os.Args[1])
	y, err2 := strconv.Atoi(os.Args[2])

	if err1 != nil || err2 != nil || x <= 0 || y <= 0 {
		println("Invalid arguments")
		return
	}
	QuadB(x, y)
}
