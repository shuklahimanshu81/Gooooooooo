package main

import "fmt"

func main() {

	//how slices are different than array => arrays are fixed in size, while slices are dynamic.
	fmt.Println("welcome to slices in golang")
	var fruits = []string{"apple"}
	fmt.Println(len(fruits))
	fruits = append(fruits, "lsjflasf", "lfsa;dfja")
	fmt.Println(fruits)

	//remove the item based on index
	var languages = []string{"js", "pythong", "ts", "golang", "C++"}
	var index int = 2
	languages = append(languages[:index], languages[index+1:]...)
	fmt.Println(languages)

}
