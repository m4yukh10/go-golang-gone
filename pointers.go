package main

import (
	"fmt"
)

// func changeNum (num int) {
// 	num = 5
// 	fmt.Println("in changeNum", num)
// }

// by reference

func changeNum (num *int) {        // x *int is giving a pointer (memory address) to the function as an argument
	*num = 5  //dereferencing to get actual value and update it
	fmt.Println("in changeNum", num)
}

func sumToN (n *int) int {
	var sum int = 0
	var x int = *n    // accessed the value in the memory address by dereferencing it
	for val := range x+1 {
		sum =  sum + val
	}
	return sum

}

func main() {
	num := 1
	// changeNum(num)

	changeNum(&num)

	hola := 10
	x := sumToN(&hola)   // &<some_variable_name> basically means the memory address of it
	
	// fmt.Println("memory address:", &num)
	fmt.Println("after changeNum", num)

	fmt.Println("sum of N", x)
} 
