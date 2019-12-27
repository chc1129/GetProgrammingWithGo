package main

import "fmt"

func main() {
  fmt.Printf("火星の表面で,私の体重は, %v ポンド, ", 149.0 * 0.3783)
  fmt.Printf("年齢は, %v歳になるでしょう\n", 41 * 365 / 687)
  // 火星の表面で, 私の体重は, 56.3667 ポンド, 年齢は, 21歳になるでしょう
  fmt.Printf("私の体重は, %v の表面で %v ポンドです\n", "Earth", 149.0)
  // 私の体重は,地球の表面で %v 149 ポンドです
  fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
  fmt.Printf("%-15v $%4v\n", "Virgin Gaiactic", 100)
}
