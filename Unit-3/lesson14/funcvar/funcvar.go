package main

import "fmt"

func main() {
	f := func(message string) { // 無名関数を変数に代入する
		fmt.Println(message)
	}
	f("Go to the party.") // Go to the party.
}
