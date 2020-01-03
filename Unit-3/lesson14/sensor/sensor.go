package main

import (
	"fmt"
	"math/rand"
)

type kelvin float64

func fakeSensor() kelvin {
	return kelvin(rand.Intn(151) + 150)
}

func realSensor() kelvin {
	return 0 // To-do: 本物のセンサを実装
}

func main() {
	sensor := fakeSensor // 関数を変数に代入する
	fmt.Println(sensor())

	sensor = realSensor
	fmt.Println(sensor())
}
