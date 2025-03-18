package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . <brainfuck_code>")
		return
	}
	code := os.Args[1]
	data := make([]byte, 2048)
	ptr := 0
	var loopStack []int
	for x := 0; x < len(code); x++ {
		switch code[x] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			data[ptr]++
		case '-':
			data[ptr]--
		case '.':
			fmt.Printf("%c", data[ptr])
		case '[':
			if data[ptr] == 0 {
				openLoops := 1
				for openLoops > 0 {
					x++
					if code[x] == '[' {
						openLoops++
					} else if code[x] == ']' {
						openLoops--
					}
				}
			} else {
				loopStack = append(loopStack, x)
			}
		case ']':
			if data[ptr] != 0 {
				x = loopStack[len(loopStack)-1]
			} else {
				loopStack = loopStack[:len(loopStack)-1]
			}
		}
	}
}
