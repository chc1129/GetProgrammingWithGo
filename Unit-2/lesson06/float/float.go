package main

import "fmt"

func main() {
  third := 1.0 / 3.0
  fmt.Println(third + third + third) // 1

  piggyBank := 0.1 // 貯金箱の残高は0.1ドル(10セント)
  piggyBank += 0.2 // 20セント追加
  fmt.Println(piggyBank) // 0.30000000000000004
}
