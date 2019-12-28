package main

import "fmt"

func main() {
  var room = "lake"

  switch room {
  case "cave":
    fmt.Println("君は薄暗い洞窟の中にいる.")
  case "lake":
    fmt.Println("堅そうに氷が張っている.")
    fallthrough
  case "underwater":
    fmt.Println("水は凍るくらいに冷たい.")
  }
}
