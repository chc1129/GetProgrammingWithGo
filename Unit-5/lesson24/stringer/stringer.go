package main

import "fmt"

// 緯度と経度(10進法)によるlocation
type location struct {
	lat, long float64
}

// Stringは,緯度と経度によるlocationの整形を行う
func (l location) String() string {
	return fmt.Sprintf("%v, %v", l.lat, l.long)
}

func main() {
	curiosity := location{-4.5895, 137.4417}
	fmt.Println(curiosity)
	// -4.5895, 137.4417
}
