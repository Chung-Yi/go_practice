package main

import "fmt"

type shape interface {
	getArea() float64
}
type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {
	tr := triangle{height: 5.0, base: 2.0}
	sq := square{sideLength: 4.0}
	printArea(tr)
	printArea(sq)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func (t triangle) getArea() float64 {
	a := t.base * t.height * 0.5
	return a
}

func (s square) getArea() float64 {
	a := s.sideLength * s.sideLength
	return a
}
