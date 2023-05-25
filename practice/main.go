package main

import (
	"fmt"
)

func main() {
	
	doSomething()
	sum := addValue(3, 5)
	fmt.Println("The sum is:", sum)

	multiSum, count := addMultiValue(3, 5, 6, 7)
	fmt.Println("Multi sum value:", multiSum)
	fmt.Println("Multi sum count:", count)
}

func doSomething() {
	fmt.Println("This prints a text")
}

func addValue(val1, val2 int) int {
	return val1 + val2
}

func addMultiValue(values ...int) (int, int) {
	total := 0
	for _, v := range values {
		total += v
	}
	return total, len(values)

}