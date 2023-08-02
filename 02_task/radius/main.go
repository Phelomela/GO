package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		lenCircle = 35
		r         float32
	)
	r = float32(lenCircle) / math.Pi
	R := &r
	S := *R * *R * math.Pi
	fmt.Println(math.Round(float64(S*100)) / 100)
}
