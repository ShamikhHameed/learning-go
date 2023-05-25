package main

import (
	"fmt"
)

func main() {

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	fmt.Printf("\n")

	for i := range colors {
		fmt.Println(colors[i])
	}

	fmt.Printf("\n")

	for _, color := range colors {
		fmt.Println(color)
	}

	fmt.Printf("\n")

	value := 1
	for value < 10 {
		fmt.Println("Value:", value)
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum > 300 {
			goto theEnd
		}
	}

	theEnd: fmt.Println("The end of program")
	
}
