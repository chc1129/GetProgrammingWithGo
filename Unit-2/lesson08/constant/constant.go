package main

import "fmt"

func main() {
  const distance = 24000000000000000000
  const lightSpeed = 299792
  const secondsPerDay = 86400

  const days = distance / lightSpeed / secondsPerDay

  fmt.Println("アンドロメダ銀河まで光速で", days, "日の距離")
  fmt.Println("アンドロメダ銀河まで光速で", 24000000000000000000/299792/86400, "日の距離")
}
