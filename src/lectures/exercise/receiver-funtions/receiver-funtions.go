//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Name string

type Health struct {
	health    int
	maxHealth int
}

type Energy struct {
	energy    int
	maxEnergy int
}

type Player struct {
	name   Name
	health Health
	energy Energy
}

func (energy *Energy) changeEnergy(newEnergy int, key string) {
	if key == "energy" {
		energy.energy = newEnergy
	} else {
		energy.maxEnergy = newEnergy
	}
}

func (health *Health) changeHealth(newHealth int, key string) {
	health.health = newHealth
	if key == "health" {
		health.health = newHealth
	} else {
		health.maxHealth = newHealth
	}
}

func main() {
	player := Player{name: "Max"}

	fmt.Println("Initial statistic:", player)

	player.energy.changeEnergy(10, "maxEnergy")
	fmt.Println("New energy:", player)
	player.health.changeHealth(10, "health")
	fmt.Println("New health:", player)
}
