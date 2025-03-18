package main

import (
	"os"
)

func main() {
	program := "++++++++++[>+++++++>++++++++++>+++<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>."
	tape := [30000]byte{}
	ptr, loop := 0, 0

	for i := 0; i < len(program); i++ {
		switch program[i] {
		case '>':
			ptr++
		case '<':
			ptr--
		case '+':
			tape[ptr]++
		case '-':
			tape[ptr]--
		case '.':
			os.Stdout.Write([]byte{tape[ptr]})
		case ',':
			os.Stdin.Read(tape[ptr : ptr+1])
		case '[':
			if tape[ptr] == 0 {
				loop = 1
				for loop > 0 {
					i++
					if program[i] == '[' {
						loop++
					} else if program[i] == ']' {
						loop--
					}
				}
			}
		case ']':
			if tape[ptr] != 0 {
				loop = 1
				for loop > 0 {
					i--
					if program[i] == ']' {
						loop++
					} else if program[i] == '[' {
						loop--
					}
				}
			}
		}
	}
}
