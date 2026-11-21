package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of pointers")

	// var ptr *int
	// fmt.Println(ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println(ptr, *ptr)
	*ptr = *ptr * 3
	// var newPtr = ptr
	fmt.Println(myNumber, *ptr)
}
