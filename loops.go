package main

import (
	"fmt"
)

func main(){
	const (
		port = 5000
		host = "localhost"
	)
	fmt.Println(port, host)

	i := 0

	for i < 5 {
		fmt.Println(i)
		i = i + 1
	}

	for i:=0; i<=10; i++ {
		if i%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}

	for j := range 3 {
		fmt.Println(j)
	}


}
