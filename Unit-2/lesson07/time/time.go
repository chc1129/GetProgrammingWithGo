package main

import (
  "fmt"
  "time"
)

func main() {
  future := time.Unix(12622780800, 0)
  fmt.Println(future) // Go Palygroundでの表示:
                      // 2370-01-01 00:00:00 +0000 UTC
}
