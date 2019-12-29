package main

import (
  "fmt"
  "math/rand"
)

var era = "AD" // era(紀元)は,このパッケージ全体で使える

func main() {
  year := 2018 // eraとyearがスコープに入っている

  switch month := rand.Intn(12) + 1; month { // era,year,monthがスコープ内
  case 2:
    day := rand.Intn(28)
    fmt.Println(era, year, month, day)

  case 4, 6, 9, 11:
    day := rand.Intn(30) + 1 // これは新たなday
    fmt.Println(era, year, month, day)

  default:
    day := rand.Intn(31) + 1 // これも新たなday
    fmt.Println(era, year, month, day)
  } // monthとdayはスコープ外
} // yearはスコープ外
