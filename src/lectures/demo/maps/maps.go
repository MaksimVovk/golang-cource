package main

import "fmt"

func main() {
	shoppingList := make(map[string]int)
	shoppingList["eggs"] = 11
	shoppingList["milk"] = 10
	shoppingList["bread"] += 1

	shoppingList["eggs"] += 1
	fmt.Println(shoppingList)

	delete(shoppingList, "milk")
	fmt.Println("Milk deleted", shoppingList)

	fmt.Println("Eggs need:", shoppingList["eggs"], "eggs")

	cereal, found := shoppingList["cereal"]

	fmt.Println("Need cereal?")
	if !found {
		fmt.Println("No")
	} else {
		fmt.Println("Yes, need", cereal, "boxes")
	}

	totalCountItems := 0

	for _, amount := range shoppingList {
		totalCountItems += amount
	}

	fmt.Println("Total item count:", totalCountItems)
}
