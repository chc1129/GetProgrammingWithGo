package main

import "fmt"

// kelvinToCelsiusは,ケルビンを摂氏に変換する
// 1個のパラメータを受け取り,1個の結果を返す関数を宣言
func kelvinToCelsius(k float64) float64 {
  k -= 273.15
  return k
}

func main() {
  kelvin := 294.0
  celsius := kelvinToCelsius(kelvin) // kelvinを引数として関数を呼び出す
  fmt.Print(kelvin, "°Kは," celsius, "°Cです.")
  // 294°Kは, 20.850000000000023°Cです.
}
