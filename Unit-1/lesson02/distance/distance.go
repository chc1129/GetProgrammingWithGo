package main

import (
  "fmt"
  "math/rand"
)

func main() {
  // 火星までのランダムな距離(km)
  var distance = rand.Intn(345000001) + 56000000
  fmt.Println(distance)
}
