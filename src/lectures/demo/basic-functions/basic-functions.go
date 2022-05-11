package main

import "fmt"

func double(x int) int {
	return x + x
}

func add(x, y int) int {
	return x + y
}

func greaet() {
	fmt.Println("Hello from greet function!")
}

func main() {
	greaet()

	dozen := double(26)
	fmt.Println("dozen =", dozen)

	bakersDozen := add(dozen, 2)
	fmt.Println("bakersDozen =", bakersDozen)

	anotherBakersDozen := add(double(6), 2)
	fmt.Println("anotherBakersDozen =", anotherBakersDozen)
}
