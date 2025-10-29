package main

import (
	"fmt"
)

func counter() func() int {  // -> function returns a function of return type int
	var count int = 0

	// here logic is written...
	return func() int {         
		count = count + 1       //retains the value of the outer scope. 							
		return count            // happens in closures
	}
}

// another implementation of closures in golang
func exponent(b int) func() int {
	var num int = 1
	return func() int {
		num = num * b
		return num
	}
}


func main(){
	increment := counter()
	ex := exponent(2)      // argument given: 2
	fmt.Println(increment())  // 1
	fmt.Println(increment()) // 2

	fmt.Println(ex())  // 2
	fmt.Println(ex())  // 4
	fmt.Println(ex())  // 8
}
