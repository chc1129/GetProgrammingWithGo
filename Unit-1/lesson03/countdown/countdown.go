package main

import (
  "fmt"
  "time"
)

func main() {
  var count = 10  // 宣言と初期化

  for count > 0 { // 条件
    fmt.Println(count)
    time.Sleep(time.Second)
    count--  // 永久ループになるからカウントをディクリメントする
  }
  fmt.Println("Liftoff!") // 打ち上げ
}
