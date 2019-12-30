package main

import "fmt"

func main() {
  message := "uv vagreangvbany fcnpr fgngvba"

  for i := 0; i < len(message); i++ { // ASCII文字を1つずつ反復処理
    c := message[i]
    if c >= 'a' && c <= 'z' { // スペースや記号は.そのまま残す
      c = c + 13
      if c > 'z' {
        c = c - 26
      }
    }
    fmt.Printf("%c", c)
  }
}
