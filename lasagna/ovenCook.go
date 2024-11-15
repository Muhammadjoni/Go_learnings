/* Gopher's Gorgeous Lasagna */
package main

import (
	"fmt"
)

var ovenTime = 40
var layerCookedTime = 2

func RemainingOvenTime(actual int) int {
	// TODO
	return ovenTime - actual
}

func PreparationTime(numberOfLayers int) int {
	// TODO
	return numberOfLayers * layerCookedTime
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	// TODO
	return layerCookedTime*numberOfLayers + actualMinutesInOven
}

func main() {

	fmt.Printf("The oventime is %d\n", ovenTime)
	fmt.Println("The remaining oventime is ", RemainingOvenTime(30))
	fmt.Println("The preparation time according to layers is ", PreparationTime(2))
	fmt.Println("The elapsed preparation time is ", ElapsedTime(3, 23))
}
