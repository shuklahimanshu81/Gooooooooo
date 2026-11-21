package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golangs")

	var arr [4]string
	arr[0] = "apple"
	// arr[1] = "anthropic"
	// arr[2] = "openai"
	arr[3] = "sarvam"
	fmt.Println("AI companies ", arr)
	fmt.Println("Lenght is ", len(arr))

	fmt.Printf("Type of arr is %T\n", arr)

	var vegies = [3]string{"potato", "beans", "mushroom"}
	fmt.Println(len(vegies))
}
