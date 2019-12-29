package main

import "fmt"

func main() {
  var blue uint8 = 255
  fmt.Printf("%08b\n", blue) // 11111111

  blue++
  fmt.Printf("%08b\n", blue) // 00000000
}
