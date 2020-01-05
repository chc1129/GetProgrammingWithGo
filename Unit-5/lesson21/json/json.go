package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type location struct {
		// フィールド名の先頭を大文字にする必要あり
		Lat, Long float64
	}

	curiosity := location{-4.5895, 137.4417}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes))
}

// exitOnErrorは,エラーがあれば表示して終了する
func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
