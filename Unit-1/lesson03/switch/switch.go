package main

import "fmt"

func main() {
  var room = "lake"

  switch {
  case room == "cave":
    fmt.Println("君は薄暗い洞窟の中にいる.")
  case room == "lake":
    fmt.Println("堅そうに氷が張っている.")
    fallthrough // 次のcaseに落ちる
  case room == "underwater": // 水面下
    fmt.Println("水は凍るくらいに冷たい.")
  }
}
