package main

import "fmt"

func main() {
  year := 2018
  fmt.Printf("%T型: %v\n", year, year) // int型:2018
  days := 365.2425
  fmt.Printf("%T型 %[1]v\n", days) // float64型:365.2425
}
