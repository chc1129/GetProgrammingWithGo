package main

import "fmt"

func main() {
  fmt.Println("洞窟の入り口だ. 東へ進む道もある.")
  var command = "go inside"

  switch command {  // 各caseの値をcommandを比較
  case "go east":
    fmt.Println("君は, さらに山を登る.")
  case "enter cave", "go inside": // カンマで区切った複数の値(リスト)のどれか
    fmt.Println("きみは薄暗い洞窟の中にいる.")
  case "read sign":
    fmt.Println("「未成年立入り禁止」と書いてある.")
  default:
    fmt.Println("なんだか,よくわからない.")
  }
}
