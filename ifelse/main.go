package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("If else in golangs")

	var loginCount int = 23

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount >= 10 {
		result = "More than regular user"
	}
	fmt.Println(result)

	diceNumber := rand.Intn(7)

	switch diceNumber {
	case 1:
		fmt.Println("Let's start the game")
	case 2:
		fmt.Println("Move two spaces")
	case 3:
		fmt.Println("Move 3 spaces")
	case 4:
		fmt.Println("Move 4 spaces")
		fallthrough // this will run the next case automatically for you
	case 5:
		fmt.Println("Move 5 spaces")
	case 6:
		fmt.Println("You got another chance but first move six spaces")
	default:
		fmt.Println("How did that happen!")
	}
}
