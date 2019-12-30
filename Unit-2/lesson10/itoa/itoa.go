package main

import (
  "fmt"
  "strconv"
)

func main() {
  countdown := 10
  str := "発射まで" + strconv.Itoa(countdown) + "秒."
  fmt.Println(str)
}
