package main

import "fmt"

func main() {
	fmt.Println("maps in golang")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["TS"] = "TypeScript"
	languages["PY"] = "Python"

	fmt.Println(languages)
	fmt.Println(languages["JS"])

	delete(languages, "TS")
	fmt.Println(languages)

	// loops are interesting in golang

	for key, value := range languages {
		fmt.Println(key, value)
	}

	for _, value := range languages {
		fmt.Println(value)
	}
}
