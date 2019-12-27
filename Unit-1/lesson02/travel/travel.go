package main

import "fmt"

func main() {
  const hourPerDay = 24

  var speed = 100800 // km/h
  var distance = 9630000 // km

  fmt.Println(distance/speed/hourPerDay, "day")
}
