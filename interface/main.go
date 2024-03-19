package main

import "fmt"

type rect struct {
	length int
	breath int
}
type circle struct {
	radius int
}

func (a circle) circleArea() {
	fmt.Println(a.radius * a.radius)
}

func (a rect) rectarea() {
	area := a.breath * a.length
	fmt.Println(area)

}

type area interface {
	rectarea()
	circleArea()
}

func main() {
	a := rect{10, 20}
	a.rectarea()

}

func areacal(a area) {
	a.rectarea()
	a.circleArea()

}
