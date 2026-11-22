package main

import "fmt"

func main() {

	fmt.Println("Loops in golang")

	days := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	for i := range days {
		fmt.Println(days[i])
	}
	for index, value := range days {
		fmt.Println(index, value)
	}
	i := 0
	for i < 10 {
		if i > 5 {
			break
		} else if i < 5 {
			fmt.Println(i)
			i++
			continue
		} else if i == 5 {
			goto newLabel
		}
	}

newLabel:
	fmt.Println("Got to this label")
}
