package main

import "fmt"

func main() {
	func() { // 無名関数の宣言
		fmt.Println("Functions anonymous")
	}() // 関数呼び出し
}
