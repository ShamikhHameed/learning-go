package main

import (
	"fmt"
	"sort"
)

func main() {
	
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println(colors)
	colors = append(colors, "Purple")
	fmt.Println(colors)

	colors = append(colors[1:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	numbers := make([]int, 5)
	numbers[0] = 89
	numbers[1] = 34
	numbers[2] = 53
	numbers[3] = 83
	numbers[4] = 93
	fmt.Println(numbers)

	numbers = append(numbers, 345)
	fmt.Println(numbers)

	sort.Ints(numbers)
	fmt.Println(numbers)

}
