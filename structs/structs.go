package main

import "fmt"

type EmptyInterface interface {
	displaySomething() string
}

type Calculation interface {
	calculateArea() float32
	calculatePerimeter() float32
}

type Rectange struct {
	length float32
	bredth float32
}

type Empty struct {
}

type Circle struct {
	radius float32
}

func (e Empty) displaySomething() string {
	return "Print Something"
}

func (c *Circle) calculateArea() float32 {
	return (22 / 7.0) * c.radius * c.radius
}

func (r *Rectange) calculateArea() float32 {
	return r.length * r.bredth
}

func (r *Rectange) calculatePerimeter() float32 {
	return 2 * (r.length + r.bredth)
}

func main() {
	shapeOne := Rectange{10, 20}
	circleOne := Circle{10}
	arrayOfInterfaces := []EmptyInterface{Empty{}, Empty{}}
	for _, array := range arrayOfInterfaces {
		fmt.Println(array.displaySomething())
	}
	fmt.Println(shapeOne)
	fmt.Println(shapeOne.calculateArea())
	fmt.Println(shapeOne.calculatePerimeter())
	fmt.Println(circleOne.calculateArea())
}
