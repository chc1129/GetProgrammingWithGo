package main

import "fmt"

func main() {
	canada := "Canada"

	var home *string
	fmt.Printf("homeは %Tです.\n", home)

	home = &canada
	fmt.Println(*home) // Canada
}
