package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Print(piscine.PrintIfNot("azerty"))
	fmt.Print(piscine.PrintIfNot("234"))
	fmt.Print(piscine.PrintIfNot(""))
	fmt.Print(piscine.PrintIfNot("World!"))
	fmt.Print(piscine.PrintIfNot("123"))
	fmt.Print(piscine.PrintIfNot("0000000012345"))
	fmt.Print(piscine.PrintIfNot("234"))
}
