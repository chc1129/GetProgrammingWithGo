package main

import "fmt"

func main() {
  cipherText := "CSOITEUIWUUZNSROCKNKFD"
  keyword := "GOLAND"
  message := ""
  keyIndex := 0

  for i := 0; i < len(cipherText); i++ {
    // A=0, B=1, ... Z=25
    c := cipherText[i] - 'A'
    k := keyword[keyIndex] - 'A'

    // 暗号の文字c-キーの文字k
    c = (c-k+26)%26 + 'A'
    message += string(c)

    // keyIndexをインクリメント
    keyIndex++
    keyIndex %= len(keyword)
  }

  fmt.Println(message)
}
