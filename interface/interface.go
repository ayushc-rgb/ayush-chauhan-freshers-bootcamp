package main

import (
	"fmt"
	"math"
)

type shape interface { // interface is used to group the data entities based on the all the methods they implement
	area() float32 // so in any type if it implements both area and perimeter then it can be considered of type shape
	perimeter() float32
}

// so as soon as the type implements all the methods present in the shape interface, the type automatically implements an interface
type square struct {
	side float32
}
type circle struct {
	radius float32
}

func (s square) area() float32 {
	return (s.side * s.side)
}
func (s square) perimeter() float32 {
	return (4 * s.side)
}
func (c circle) area() float32 {
	return (math.Pi * c.radius * c.radius)
}
func (c circle) perimeter() float32 {
	return (2 * math.Pi * c.radius)
}

func generic(s shape) { // so this is the generic function which will work on any type of shape, be it square or circle
	fmt.Println("The area of the shape is:", s.area())
	fmt.Println("The perimeter of the shape is:", s.perimeter())
}

func main() {
	var inter shape
	inter = square{6} // because square type is implementing all the functions present in the interface(shape) so this can be stored
	// in a variable of type shape
	fmt.Println(inter)
	generic(inter)

	// now let's make inter(type shape) to be equal to the type circle
	inter = circle{3.4}
	generic(inter)
}
