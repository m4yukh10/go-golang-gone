package main

import (
	"fmt"
	"maps"
)

func main(){
	fmt.Println("hello world")
	m := make(map[string]string)

	m["name"] = "mayukh"
	m["hello"] = "world"
	m["age"] = "20"
	fmt.Println(m)
	
	// delete function

	delete(m, "age")
	fmt.Println(m)

	fmt.Println(len(m))

	y1 := map[string]int{"name":108, "age":900}
	y2 := map[string]int{"name":108, "age":900}
	_, ok := y1["age"]
	if ok {
		fmt.Println("all ok")
	} else {
		fmt.Println("not ok")
	}

	fmt.Println(maps.Equal(y1, y2))


}
