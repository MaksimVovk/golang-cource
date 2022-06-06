//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
package main

import "testing"

func TestMaxHealthAndMaxEnergy(t *testing.T) {
	player := Player{
		name:      "knight",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}

	player.addHealth(10)

	if player.health > player.maxHealth {
		t.Errorf("health:%v should not be more than max health: %v", player.health, player.maxHealth)
	}

	player.addEnergy(10)
	if player.energy > player.maxEnergy {
		t.Errorf("energy:%v should not be more than max energy: %v", player.energy, player.maxEnergy)
	}
}

func TestHealthAndEnergy(t *testing.T) {
	player := Player{
		name:      "knight",
		health:    100,
		maxHealth: 100,
		energy:    500,
		maxEnergy: 500,
	}

	player.applyDamage(player.health + 1)

	if player.health < 0 {
		t.Errorf("health:%v should not be less than 0", player.health)
	}

	player.consumeEnergy(player.energy + 1)
	if player.energy < 0 {
		t.Errorf("energy:%v should not be less than 0", player.energy)
	}
}
