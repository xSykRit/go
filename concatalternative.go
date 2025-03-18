package main

import "fmt"

func ConcatAlternate(slice1, slice2 []int) []int {
	rzlt := []int{}
	if len(slice1) < len(slice2) {
		slice1, slice2 = slice2, slice1
	}
	for i, x := range slice1 {
		rzlt = append(rzlt, x)
		if i < len(slice2) {
			rzlt = append(rzlt, slice2[i])
		}
	}
	return rzlt
}

func main() {
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6}))
	fmt.Println(ConcatAlternate([]int{2, 4, 6, 8, 10}, []int{1, 3, 5, 7, 9, 11}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{4, 5, 6, 7, 8, 9}))
	fmt.Println(ConcatAlternate([]int{1, 2, 3}, []int{}))
}
