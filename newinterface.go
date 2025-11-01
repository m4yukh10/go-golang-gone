package main

import (
	"fmt"
)

type circle struct {
	radius float32
}
type rectangle struct{
	length, breadth float32
}

func (r rectangle) area(){
	areaR := r.length*r.breadth
	fmt.Println("area of rectangle:", areaR)
}

func (c circle) area(){
	areaC := 3.14 * c.radius * c.radius
	fmt.Println("area of circle:", areaC)
}

type measure interface {
	area()
}

type calculateArea struct {
	figure measure
}


func main() {
	fmt.Println("")
	rectangleSpecs := rectangle{length:16.4, breadth: 12}
	x := calculateArea{
		rectangleSpecs,
	}

	x.figure.area()

	

	

}
