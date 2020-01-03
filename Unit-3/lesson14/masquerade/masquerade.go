package main

import "fmt"

var f = func() { // 無名関数を変数に代入する
	fmt.Println("Dress up for the masquerade.")
}

func main() {
	f() // Dress up for the masquerade.
}
