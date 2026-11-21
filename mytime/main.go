package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to time study of goland")

	presentTime := time.Now()
	fmt.Println(presentTime.Format("Monday"))

	createdDate := time.Date(2020, time.March, 23, 12, 12, 12, 12, time.UTC)
	fmt.Println(createdDate.Format("01-02-2006 Monday"))

}
