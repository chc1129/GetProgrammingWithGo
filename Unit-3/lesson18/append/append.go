package main

import "fmt"

func main() {
	dwarfs := []string{"Ceres", "Pluto", "Haumea", "Makemake", "Eris"}

	dwarfs = append(dwarfs, "Orcus")
	fmt.Println(dwarfs)
	dwarfs = append(dwarfs, "Salacisa", "Quaoar", "Sedna")
	fmt.Println(dwarfs)
}
