package main

import "fmt"

// coordinateは,東西南北の半球ごとの度分秒で地理座標を示す
type coordinate struct {
	d, m, s float64
	h       rune // N/S/E/W
}

// decimalはDMS座標を10進数に変換する
func (c coordinate) decimal() float64 {
	sign := 1.0
	switch c.h {
	case 'S', 'W', 's', 'w':
		sign = -1
	}
	return sign * (c.d + c.m/60 + c.s/3600)
}

func main() {
	// ブラッドベリ着陸地点:4°35'22.2"S, 137°26'30.1"E
	lat := coordinate{4, 35, 22.2, 'S'}
	long := coordinate{137, 26, 30.12, 'E'}

	fmt.Println(lat.decimal(), long.decimal())
}
