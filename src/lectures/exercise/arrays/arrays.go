//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Product struct {
	price float64
	name  string
}

func lastItem(products []Product) string {
	if products[len(products)-1].name != "" {
		return products[len(products)-1].name
	} else {
		return products[len(products)-2].name
	}
}

func totalNumberOfItems(products []Product) int {
	if products[len(products)-1].name != "" {
		return len(products)
	} else {
		return len(products) - 1
	}
}

func totalCoast(products []Product) float64 {
	sum := 0.0

	for i := 0; i < len(products); i++ {
		product := products[i]
		sum += product.price
	}

	return sum
}

func printInformationAboutProducts(products []Product) {
	fmt.Println("The last item on the list", lastItem(products))
	fmt.Println("The total number of items", totalNumberOfItems(products))
	fmt.Println("The total cost of the items", totalCoast(products))
}

func main() {
	products := [4]Product{
		{price: 1.1, name: "apple"},
		{price: 2.5, name: "Orange"},
		{price: 1.3, name: "Panapple"},
	}

	printInformationAboutProducts(products[:])

	products[3] = Product{name: "Blackberry", price: 3.4}

	printInformationAboutProducts(products[:])
}
