package main

import "fmt"

func main() {

	fmt.Println("Structs in golang")
	// there is no inheritence, no super or parent

	himanshu := User{"himanshu", "himans@gmail.com", true, 16}
	fmt.Printf("Name is %v and email is %v", himanshu.Name, himanshu.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
