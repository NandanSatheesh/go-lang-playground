package main

import "fmt"

type Rectange struct {
	length float32
	bredth float32
}

func (r *Rectange) calculateArea() float32 {
	return r.length * r.bredth
}

func (r *Rectange) calculatePerimeter() float32 {
	return 2 * (r.length + r.bredth)
}

func main() {
	shapeOne := Rectange{10, 20}
	fmt.Println(shapeOne)
	fmt.Println(shapeOne.calculateArea())
	fmt.Println(shapeOne.calculatePerimeter())
}
