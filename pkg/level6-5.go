package main

import (
	"fmt"
	"math"
)

// Hands-on exercise #5
// create a type SQUARE
// create a type CIRCLE
// attach a method to each that calculates AREA and returns it
// circle area= π r 2
// square area = L * W
// create a type SHAPE that defines an interface as anything that has the AREA method
// create a func INFO which takes type shape and then prints the area
// create a value of type square
// create a value of type circle
// use func info to print the area of square
// use func info to print the area of circle

type square struct {
	sideLength float64
}

func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{sideLength: 12.5}

	ci := circle{radius: 1.05}

	info(sq)
	info(ci)

}
