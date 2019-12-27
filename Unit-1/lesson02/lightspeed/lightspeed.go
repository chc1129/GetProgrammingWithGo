// 火星まで, 光速で何秒かかるか?
package main

import "fmt"

func main() {
  const lightSpeed = 299792 // km/sec
  var distance = 56000000   // km

  fmt.Println(distance/lightSpeed, "seconds")

  distance = 401000000
  fmt.Println(distance/lightSpeed, "seconds")
}
