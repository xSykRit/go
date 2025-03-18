package main

import "github.com/01-edu/z01"

func quadA(x, y int) {
	/* z01.PrintRune('o')
	for i := 1; i <= x-2; i++ {
		z01.PrintRune('-')
	}
	z01.PrintRune('o')
	z01.PrintRune('\n')
	for i := 1; i <= y-2; i++ {
		z01.PrintRune('|')
		for i := 1; i <= x-2; i++ {
			z01.PrintRune(' ')
		}
		z01.PrintRune('|')
		z01.PrintRune('\n')
	} */
	for i := 1; i <= x; i++ {
		for j := 1; j <= y; j++ {
			if i == 1 && i == x {
				z01.PrintRune('o')
			}
			if i != 1 && i != x && j != 1 && j != y {
				z01.PrintRune('*')
			} else {
				z01.PrintRune('|')
			}
			z01.PrintRune('\n')
		}
	}
}

func main() {
	quadA(3, 5)
}
