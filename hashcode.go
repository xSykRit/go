package main

import "fmt"

func HashCode(dec string) string {
	var result string
	size := len(dec)
	for x := 0; x < size; x++ {
		ascii := int(dec[x])
		hash := (ascii + size) % 127
		if ascii < 32 || ascii > 126 {
			ascii += 33
		}
		done := rune(hash)
		result += string(done)
	}
	return result
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
