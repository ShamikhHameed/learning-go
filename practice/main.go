package main

import (
	"fmt"
	"math"
)

func main() {

	i1, i2, i3 := 23, 45, 56
	intSum := i1 + i2 + i3
	fmt.Println("Integer Sum : ", intSum)

	f1, f2, f3 := 23.5, 45.8, 56.8
	floatSum := f1 + f2 + f3
	fmt.Println("Float Sum : ", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("The sum is now ", floatSum)

	circleRadius := 15.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)

}
