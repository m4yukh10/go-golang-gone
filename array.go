package main

import (
	"fmt"
)

func main(){
	nums := [3]int{1,2,3}

	for i := range 3 {
		fmt.Println(nums[i])
	}

	var x = make([]string, 2)
	x[0] = "mayukh"
	x[1] = "hello"

	// out
	for i := range 2 {
		fmt.Println(x[i])
	}

	//cap -> capacity. maap bole deye
	fmt.Println(cap(x))
}
