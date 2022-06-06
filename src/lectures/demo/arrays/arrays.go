package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliess(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]

		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is not clean")
		}
	}
}
func main() {
	rooms := [...]Room{
		{name: "Office", cleaned: false},
		{name: "Warehouse", cleaned: false},
		{name: "Reception", cleaned: false},
		{name: "Ops", cleaned: false},
	}

	checkCleanliess(rooms)

	fmt.Println("Performing cleaning...")

	rooms[2].cleaned = true
	rooms[3].cleaned = true

	checkCleanliess(rooms)
}
