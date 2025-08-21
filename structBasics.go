package main

import "fmt"

type person struct {
	name string
	age  int
}
type vertex struct {
	abscissa int
	ordinate int
}

// go doesn't have classes but you an create methods on type, the methods will have the receiver arguments
// so like if you want to associate a function with any type you can do it with methods with special receiver arguments

func (v vertex) sumOfCoordinates() int {
	return v.abscissa + v.ordinate
}

func newPerson(name string) *person {
	p := person{name: name, age: 18}
	return &p
	// although this is returning a pointer to the local variable which should have been destroyed when the function ends
	// but in Go language, a variable will only be cleaned by a garbage collected only when there are no active references to it
}

func main() {
	fmt.Println(person{name: "abc", age: 22})
	fmt.Println(person{age: 50}) // default string value is an empty string

	fmt.Println(&person{name: "something", age: 99}) // when we try to access the address of the struct, a '&' is displayed on the screeen

	p := person{} // an instance of the structure is created and is saved in variable p
	fmt.Println(p)
	p.name = "firstName"
	p.age = 34
	fmt.Println(p)

	fmt.Println(newPerson("tango"))
	// newPerson is returning a pointer so this will print the person struct with an ampersand sign in front

	// pointer to person p
	pp := &p
	// we can also use '.' with pointers to structs also they are deferenced automatically

	// so to get the age of the person p we can do
	fmt.Println(pp.age)
	// also we can change the field values using pointers

	pp.age = 89
	fmt.Println(pp.age)

	//if we have to instantiate the struct just once, we can just do it anonymously

	vehicle := struct {
		name        string
		vehicleType int
	}{
		"honda",
		4,
	}
	fmt.Println(vehicle)

	coordinates := vertex{2, 3}
	pointerToCoordinates := &coordinates
	fmt.Println(pointerToCoordinates)
	(*pointerToCoordinates).abscissa = 10. // this is the classic way to change the value through the pointer
	fmt.Println(coordinates)
	pointerToCoordinates.abscissa = 15 // but I can also do something like this
	fmt.Println(coordinates)

	// calling a method of the vertex struct
	fmt.Println(coordinates.sumOfCoordinates())

}
