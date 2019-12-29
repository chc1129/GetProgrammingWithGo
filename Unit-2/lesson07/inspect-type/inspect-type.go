package main

import "fmt"

func main() {
  a := "text"
  fmt.Printf("%T型: %[1]v\n", a) // string型: text

  b := 42
  fmt.Printf("%T型: %[1]v\n", b) // int型: 42

  c := 3.14
  fmt.Printf("%T型: %[1]v\n", c) // float64型: 3.14

  d := true
  fmt.Printf("%T型: %[1]v\n", d) // bool型: true
}
