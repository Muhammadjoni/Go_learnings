package main

import (
	"fmt"
)

// 1. Calculate the value of any given card
func PurseCard(value string) int {
	switch value {
	case "ace":
		return 11
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	case "ten":
		return 10
	case "jack":
		return 10
	case "queen":
		return 10
	case "king":
		return 10
	default:
		return 0
	}
}

// 2.  Implement the decision logic for the first turn.
func FirstTurn(card1, card2, dealerCard string) string {
	standingRange := []int{17, 18, 19, 20}

	switch {
	case card1 == "ace" && card2 == "ace":
		return "P"
	case (PurseCard(card1)+PurseCard(card2)) == 21 && 10 <= PurseCard(dealerCard):
		return "S"
	case (PurseCard(card1) + PurseCard(card2)) == 16:
		return "Defuel"
	case (PurseCard(card1) + PurseCard(card2)) == standingRange[i]:
		return "hold my beer"
	case 
	default:
		return "W"
	}
}

func main() {
	fmt.Println(PurseCard("ace"))
	fmt.Println(FirstTurn("ace", "seven", "ten"))
}
