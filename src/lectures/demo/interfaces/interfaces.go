package main

import "fmt"

type Preparer interface {
	PreparerDish()
}

type Chicken string
type Salad string

func (c Chicken) PreparerDish() {
	fmt.Println("Cook chicken")
}

func (s Salad) PreparerDish() {
	fmt.Println("chop salad")
	fmt.Println("add dressing")
}

func preparerDishes(dishes []Preparer) {
	fmt.Println("Preparer dishes:")

	for i := 0; i < len(dishes); i++ {
		dish := dishes[i]
		fmt.Printf("---Dish %v---\n", dish)
		dish.PreparerDish()
	}

	fmt.Println()
}

func main() {
	dishes := []Preparer{Chicken("Chicken Wings"), Salad("Iceberg salad")}
	preparerDishes(dishes)
}
