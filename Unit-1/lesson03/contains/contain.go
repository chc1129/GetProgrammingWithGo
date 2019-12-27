package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println("君は薄暗い洞窟の中にいる.")
  var command = "walk outside"            // コマンド:外を歩く
  var exit = strings.Contains(command, "outside") //コマンドが外を含む?
  fmt.Println("洞窟を出る:", exit)        // 洞窟を出る:true
}
