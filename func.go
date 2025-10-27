package main

import (
	"fmt"
)

func add (a int, b int) int {     // int outside parentheses -> return type
	return a + b
}

func getStrings()(string, string, string) {
	return "hello", "hello ji", "hello sir"
}

func getNums(a int) (int, bool) {
	if a%2 == 0{
		return a, true
	} else {
		return a, false
	}
}

func main(){

	x := add(3,4)
	fmt.Println(x)

	fmt.Println(getStrings())

	l1, l2, _ := getStrings() // underscore shows value is stored but is okay if not used
	
	fmt.Println("hello,", l1)
	fmt.Println("hello,", l2)

	num, is := getNums(56)
	fmt.Println(num)
	fmt.Println(is)


}
