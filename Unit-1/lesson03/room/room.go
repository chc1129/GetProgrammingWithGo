package main

import "fmt"

func main() {
  var room = "cave"

  if room == "cave" {
    fmt.Println("君は薄暗い洞窟の中にいる.")
  } else if room == "entrance" {
    fmt.Println("洞窟の入り口だ.東へ進む道もある.")
  } else if room == "mountain" {
    fmt.Println("崖があるぞ.西への道は山を下りる.")
  } else {
    fmt.Println("何もかも真っ白だ.")
  }
}
