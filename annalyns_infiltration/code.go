package main

import "fmt"

var knightIsAwake = false
var archerIsAwake = false
var prisonerIsAwake = true
var annalynDogPresence = true

// 1. Check if a fast attack can be made
func CanFastAttack(knighIsAwake bool) bool {
	if knighIsAwake {
		return false
	} else {
		return true
	}
}

// 2. Check if the group can be spied upon
func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {

	var possible = (knightIsAwake && archerIsAwake && !prisonerIsAwake) || (knightIsAwake && !archerIsAwake && prisonerIsAwake) || (!knightIsAwake && archerIsAwake && prisonerIsAwake) || (!knightIsAwake && !archerIsAwake && prisonerIsAwake) || (knightIsAwake && !archerIsAwake && !prisonerIsAwake) || (!knightIsAwake && archerIsAwake && !prisonerIsAwake)

	if possible {
		return true
	} else {
		return false
	}
}

// 3. Check if the prisoner can be signaled
func CanSignalPrisoner(knighIsAwake, archerIsAwake bool) bool {
	if prisonerIsAwake && !archerIsAwake {
		return true
	} else {
		return false
	}
}

// 4. Check if the prisoner can be freed
func CanFreePrisoner(knighIsAwake, archerIsAwake, prisonerIsAwake, annalynDogPresence bool) bool {
	var possible = (annalynDogPresence && !archerIsAwake) || (!knightIsAwake && !archerIsAwake && prisonerIsAwake)
	if possible {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(CanFastAttack(knightIsAwake))
	fmt.Println(CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake))
	fmt.Println(CanSignalPrisoner(knightIsAwake, archerIsAwake))
	fmt.Println(CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, annalynDogPresence))
}
