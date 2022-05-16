//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:
//* Create a structure to store items and their security tag state
//  - Security tags have two states: active (true) and inactive (false)
//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

const (
	Active   = true
	Inactive = false
)

type SecurityTag bool

type Item struct {
	name string
	tag  SecurityTag
}

func activate(tag *SecurityTag) {
	*tag = Active
}

func diativate(tag *SecurityTag) {
	*tag = Inactive
}

func checkout(items []Item) {
	fmt.Println("checking out...")

	for i := 0; i < len(items); i++ {
		diativate(&items[i].tag)
	}
}

func main() {
	shirt := Item{name: "shirt", tag: Active}
	pants := Item{name: "pants", tag: Active}
	watch := Item{name: "watch", tag: Active}
	purse := Item{name: "purse", tag: Active}

	storeList := []Item{shirt, pants, watch, purse}
	fmt.Println(storeList)

	diativate(&storeList[3].tag)
	fmt.Println(storeList)

	checkout(storeList)
	fmt.Println(storeList)
}
