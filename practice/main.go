package main

import (
	"fmt"
)

const aConst int = 99

func main() {
	var aString string = "This is Go"
	fmt.Println(aString)
	fmt.Printf("The variable's type is %T\n\n", aString)

	var aInt int = 42
	fmt.Println(aInt)
	fmt.Printf("The variables type is %T\n\n", aInt)

	var defaultInt int
	fmt.Println(defaultInt)
	fmt.Printf("The variables type is %T\n\n", defaultInt)

	var anotherString = "This another string"
	fmt.Println(anotherString)
	fmt.Printf("The variables type is %T\n\n", anotherString)

	var anotherInt = 53
	fmt.Println(anotherInt)
	fmt.Printf("The variables type is %T\n\n", anotherInt)

	myString := "String with colon"
	fmt.Println(myString)
	fmt.Printf("The variables type is %T\n\n", myString)

	myInt := 67
	fmt.Println(myInt)
	fmt.Printf("The variables type is %T\n\n", myInt)

	fmt.Println(aConst)
	fmt.Printf("The variables type is %T\n\n", aConst)
}
