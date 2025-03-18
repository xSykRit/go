package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func Itoi(n int) string {
	result := ""
	for n != 0 {
		m := n % 10
		str := string('0' + m)
		result = str + result
		n = n / 10
	}
	return result
}

func main() {
	Tab := [][]string{}
	var output []byte
	Alpha := "ABCDE"
	data, _ := io.ReadAll(os.Stdin)
	if len(data) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	x, y := 0, 0
	for _, e := range data {
		if e != 10 {
			x++
		} else {
			y++
		}
	}
	if y != 0 {
		x /= y
	}
	for _, v := range Alpha {
		cmd := exec.Command("./quad"+string(v), Itoi(x), Itoi(y))
		output, _ = cmd.Output()
		if string(output) == string(data) {
			Tab = append(Tab, []string{"quad" + string(v)}, []string{Itoi(x)}, []string{Itoi(y)})
		}
	}
	if len(Tab) == 0 {
		fmt.Println("Not a quad function")
		return
	}
	for i, e := range Tab {
		fmt.Print(e)
		if (i+1)%3 == 0 && i != len(Tab)-1 {
			fmt.Printf(" || ")
		} else if i != len(Tab)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
