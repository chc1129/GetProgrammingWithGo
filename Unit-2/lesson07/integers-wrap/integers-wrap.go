package main

import "fmt"

func main() {
  var red uint8 = 255
  red++
  fmt.Println(red) // 0

  var number int8 = 127
  number++
  fmt.Println(number) // -128
}
