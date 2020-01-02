package main

import "fmt"

type celsius float64
type kelvin float64

// kelvinToCelsiusは,Kを摂氏に変換する
func kelvinToCelsius(k kelvin) celsius {
  return celsius(k - 273.15) // 型変換が必要
}

func main() {
  var k kelvin = 294.0 // 引数は必ずkelvin型にする
  c := kelvinToCelsius(k)
  fmt.Print(k, "°K is ", c, "°Cです.")
  // 294°Kは,20.850000000000023°Cです
}
