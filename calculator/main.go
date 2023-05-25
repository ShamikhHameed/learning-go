package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	
	getInputs()

}

func getInputs() {
	
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Value 1: ")
	input1, _ := reader.ReadString('\n')
	float1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}

	fmt.Print("Value 2: ")
	input2, _ := reader.ReadString('\n')
	float2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}

	fmt.Print("Enter operation (+ - / *): ")
	op, _ := reader.ReadString('\n')
	op1 := strings.TrimSpace(op)

	var result float64

	switch op1 {
	case "+":
		result = add(float1, float2)
	case "-":
		result = substract(float1, float2)
	case "*":
		result = multiply(float1, float2)
	case "/":
		result = divide(float1, float2)
	}

	result = math.Round(result*100) / 100
	fmt.Println("The result is", result)

}

func add(v1, v2 float64) float64 {
	return v1 + v2
}

func multiply(v1, v2 float64) float64 {
	return v1 * v2
}

func substract(v1, v2 float64) float64 {
	return v1 - v2
}

func divide(v1, v2 float64) float64 {
	return v1 / v2
}
