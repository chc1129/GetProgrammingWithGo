package main

import "fmt"

func main() {
	answer := 42
	address := &answer

	fmt.Printf("addressの型は %Tです.\n", address)
}
