package main

import "fmt"

func main() {
  launch := false // 打ち上げ準備完了?
  launchText := fmt.Sprintf("%v", launch)
  fmt.Println("Ready for launch", launchText) // Ready for launch: false

  var yesNo string
  if launch {
    yesNo = "yes"
  } else {
    yesNo = "no"
  }
  fmt.Println("Ready for launch", yesNo)
}
