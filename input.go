package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	hello := "welcome"
	fmt.Println(hello)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter rating for pizza: ")

	//comma ok syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating, ", input)
	fmt.Printf("type: %T", input)

	//conversions
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	
	
	if err != nil{
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Println("added 1 to rating", numRating + 1)

	}


}
