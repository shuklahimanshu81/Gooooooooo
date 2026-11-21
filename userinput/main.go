package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to the go programming"

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please tell us your name")
	input, _ := reader.ReadString('\n')

	fmt.Println(welcome, input)

}
