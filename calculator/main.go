package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"math"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Value 01 :")
	v1, _ := reader.ReadString('\n')
	v1Float, errV1 := strconv.ParseFloat(strings.TrimSpace(v1), 64)
	if errV1 != nil {
		fmt.Println(errV1)
		panic("Value 1 must be a number")
	}

	fmt.Print("Value 02 :")
	v2, _ := reader.ReadString('\n')
	v2Float, errV2 := strconv.ParseFloat(strings.TrimSpace(v2), 64)
	if errV2 != nil {
		fmt.Println(errV2)
		panic("Value 2 must be a number")
	}

	sumFloat := v1Float + v2Float
	sumFloat = math.Round(sumFloat*100) / 100 

	fmt.Printf("The sum of %v and %v is %v", v1Float, v2Float, sumFloat)

}
