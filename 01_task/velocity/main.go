package main

import (
	"fmt"
	"math"
)

func main() {
	const mile float32 = 1.609 //1 nile = 1.609 km/h
	var (
		firstVelocity    float32 = 120.4 //m/s
		secondVelocity   float32 = 130   //m/s
		AmericanVelocity float32
		EuropeanVelocity float32
	)
	EuropeanVelocity = firstVelocity * 3600 / 1000 //km/h
	AmericanVelocity = secondVelocity * 3600 / (mile * 1000)
	fmt.Println("EuropeanVelocity for", firstVelocity, "m/s is", math.Round(math.Pow(10, 2)*float64(EuropeanVelocity))/math.Pow(10, 2), "km/h.")
	fmt.Println("AmericanVelocity for", secondVelocity, "m/s is", math.Round(math.Pow(10, 2)*float64(AmericanVelocity))/math.Pow(10, 2), "km/h.")

}
