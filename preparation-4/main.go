package main

import (
	"fmt"
	"math"
)

type CentiMeters float64

func (c CentiMeters) IsTooLong() {
	if c > 100 {
		fmt.Println("The value is too long!")
	} else {
		fmt.Println("Nicee")
	}
}

type shape interface {
	getArea() float64
	getPerimeter() float64
}

type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) getArea() float64 {
	return r.width * r.height
}

func (r rectangle) getPerimeter() float64 {
	return 2 * (r.width + r.height)
}

type circle struct {
	radius float64
}

func (c circle) getArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) getPerimeter() float64 {
	return 2 * math.Pi * c.radius
}

type square struct {
	width float64
}

func (s square) getArea() float64 {
	return s.width * s.width
}

func (s square) getPerimeter() float64 {
	return 4 * s.width
}

func measureShape(s shape) {
	fmt.Printf("Shape has an area of %f\n", s.getArea())
	fmt.Printf("Shape has a perimeter of %f\n", s.getPerimeter())
}

func main() {
	myVar := CentiMeters(1000.9)
	fmt.Printf("Type : %T and value : %v\n", myVar, myVar)
	// myVar.IsTooLong()

	myRectangle := rectangle{width: 10, height: 20}
	myCircle := circle{radius: 10}
	mySquare := square{width: 10}
	// fmt.Printf("Type : %T and value : %+v\n", myRectangle, myRectangle)
	fmt.Printf("Rectangle : %+v\n", myRectangle)
	measureShape(myRectangle)
	fmt.Printf("Circle : %+v\n", myCircle)
	measureShape(myCircle)
	fmt.Printf("Square : %+v\n", mySquare)
	measureShape(mySquare)
}
