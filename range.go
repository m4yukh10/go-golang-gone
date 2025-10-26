package main

import (
	"fmt"
)

func main(){

	var x = []int{2,3,4}
	sum := 0

	for i, num := range x {
		sum = sum + num
		fmt.Println(i)      // here i gives index, num gives value

	}
	fmt.Println(sum)

	// using a map to try range function

	m := map[string]string{"fname":"mayukh", "lname":"sen"}

	for k, e := range m {
		fmt.Println(k, e)    // here k gives the key, e (should be v) gives the value
	}

	for u,i := range "helloladies" {
		fmt.Println(u, i)                  // i gives the unicode value of a character

	}

	for j, k := range "holaculers" {
		fmt.Println(j, string(k))            //gives the actual string value  
	}



}
