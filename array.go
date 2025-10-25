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
	x[1] = "lamine"

	// out
	for i := range 2 {
		fmt.Println(x[i])
	}

	//cap -> capacity. maap bole deye
	fmt.Println(cap(x))

	var y = make([]int, 7)
	y = append(y, 10)
	y = append(y, 20)
	fmt.Println(y)

	fmt.Println(cap(y)) // returns 14. increases the size to accommodate more elements. a self-expanding party hall where more people means bigger hall.
	
}
