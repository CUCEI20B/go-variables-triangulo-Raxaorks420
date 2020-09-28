package main

import "fmt"

func main() {
	var base uint64
	fmt.Scanln(&base)

	var altura uint64
	fmt.Scanln(&altura)

	area := (base * altura) / 2
	fmt.Println(area)
}
