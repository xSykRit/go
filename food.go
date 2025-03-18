package main

import "fmt"

type food struct {
	name      string
	prix      int
	time      int
	avaliable bool
}

func main() {
	//menu := map[string]food
	food1 := food{"Burger", 15, 123, true}
	food2 := food{"Nuggets", 17, 44, false}

	fmt.Println(food1.name, food1.time, food1.avaliable)
	fmt.Println(food2.prix)
}
