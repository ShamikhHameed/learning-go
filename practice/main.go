package main

import (
	"fmt"
)

func main() {

	theAnswer := 43
	var result string

	if theAnswer < 0 {
		result = "Less than 0"
	} else if theAnswer == 0 {
		result = "Zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)

	if x := -34; x < 0 {
		result = "Less than 0"
	} else if theAnswer == 0 {
		result = "Equal to Zero"
	} else {
		result = "Greater than zero"
	}

	fmt.Println(result)
	
}
