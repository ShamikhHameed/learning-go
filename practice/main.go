package main

import (
	"fmt"
)

func main() {
	
	anInt := 43
	var p = &anInt
	fmt.Println("The value is ", *p)

	value1 := 43.7
	pointer1 := &value1
	fmt.Println("The value is :", *pointer1)

	*pointer1 = *pointer1 / 3
	fmt.Println("Value of pointer is :", *pointer1)
	fmt.Println("Value is :", value1)

}