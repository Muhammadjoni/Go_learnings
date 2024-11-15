package main

import "fmt"

//1. Welcome
func Welcome(name string) string {
	return fmt.Sprintf("Welcome %s", name)
}

//2.HappyBirthday

func HappyBirthday(name string, age int) string {
	return fmt.Sprintf("Welcome to the party %s, You are ", name)
}

//Assign direction
func AssignTable(name string, tn int, seatmate string, drct string, dtnc float64) string {
	return fmt.Sprintf("Welcome to my party, %s\n You have been assigned to table %d. Your table is %s, exactly %f meters from here.\n You will be sitting next to %s", name, tn, drct, dtnc, seatmate)
}

func main() {
	fmt.Println(HappyBirthday("Frank", 58))
	fmt.Println(Welcome("Joni"))
	fmt.Println(AssignTable("Christiane", 27, "Frank", "on the left", 23.7834298))
}
