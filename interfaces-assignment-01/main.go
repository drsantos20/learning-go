package main

import (
	"fmt"
)

type shape interface {
	getArea() float64
}

type square struct {
	sideLength float64
}
type triangle struct {
	height, base float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printShape(s shape) {
	fmt.Println(s.getArea())
}

func main() {
	sqr := square{sideLength: 10}
	trgl := triangle{height: 10, base: 10}

	printShape(sqr)
	printShape(trgl)

}
