package main

import "fmt"

func main() {
	var planets [8]string
	planets[0] = "水星" // インデックス0の位置に惑星を代入
	planets[1] = "金星"
	planets[2] = "地球"

	earth := planets[2] // インデックス2の惑星を取り出す
	fmt.Println(earth)  // 地球
}
