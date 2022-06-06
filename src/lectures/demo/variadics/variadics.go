package main

import "fmt"

func sum(nums ...int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}
	return sum
}

func main() {
	a := []int{1, 2, 3, 4}
	b := []int{1, 12, 9, 4}

	all := append(a, b...)

	answer := sum(all...)
	fmt.Println(answer)

	answer = sum(1, 2, 3, 4)
	fmt.Println(answer)
}
