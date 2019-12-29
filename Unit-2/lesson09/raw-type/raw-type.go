package main

import "fmt"

func main() {
  fmt.Printf("%vは%[1]T型です\n", "文字列リテラル")
  fmt.Printf("%vは%[1]T型です\n", `生の文字列リテラル`)
}
