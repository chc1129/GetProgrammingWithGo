package main

import "fmt"

func main() {
	neptune := "Neptune"
	tune := neptune[3:]

	fmt.Println(tune)
	neptune = "Poseidon"
	fmt.Println(tune)

	question := "¿Cómo estás?"
	fmt.Println(question[:6])

}
