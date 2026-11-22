package main

import "fmt"

func main() {
	fmt.Println("welcome to functions & methods in go")
	adder(3, 4)
	result, mesage := multiAdder(3, 4, 5, 5, 2)
	defer fmt.Println(result, mesage)

	himanshu := User{"himanshu", "himanshu@go.dev", false, 27}

	defer himanshu.GetStatus()

	himanshu.UpdateMail()

}

func adder(a int, b int) int {
	return a + b
}

func multiAdder(values ...int) (int, string) {
	total := 0
	for _, value := range values {
		total = total + value
	}

	return total, "multi adder has replied"
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) UpdateMail() {
	u.Email = "newhimanshu@go.dev"
	fmt.Println("New email is ", u.Email)
}
func (u User) GetStatus() {
	fmt.Println("Status of the user is ", u.Status)
}
