package main

import "fmt"

type kelvin float64

// sensor関数型
type sensor func() kelvin

func realSensor() kelvin {
	return 0 // To-do: 実数値で実装
}

func calibrate(s sensor, offset kelvin) sensor {
	return func() kelvin { // 無名関数を宣言して返す
		return s() + offset
	}
}

func main() {
	sensor := calibrate(realSensor, 5)
	fmt.Println(sensor()) // 5
}
