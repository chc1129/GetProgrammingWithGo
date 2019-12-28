package main

import "fmt"

func main() {
  var command = "go east"

  // もしコマンドが"go east"と等しければ
  if command == "go east" {
    fmt.Println("きみは,さらに山を登る")
  // あるいは,もしコマンドが"go inside"と等しければ
  } else if command == "go inside" {
    fmt.Println("きみは洞窟に入り,そこで一生を過ごす")
  // あるいあは,他のコマンドならば
  } else {
    fmt.Println("なんだか,よくわからない")
  }
}
