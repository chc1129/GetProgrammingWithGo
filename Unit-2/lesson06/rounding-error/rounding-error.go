package main

import "fmt"

func main() {
  celsius := 21.0
  fmt.Println((celsius/5.0*9.0)+32, "ºF\n")
  fmt.Println((9.0/5.0*celsius)+32, "ºF\n")
}