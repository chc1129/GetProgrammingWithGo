package main

import "fmt"

func main() {
  var haveTorch = true // トーチを持ってる?
  var litTorch = false // トーチが点灯している?
  if !haveTorch || !litTorch {
    fmt.Println("ここには何も見えない.")
  }
}
