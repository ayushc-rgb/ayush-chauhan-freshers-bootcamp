package main

import (
	"fmt"
	"maps"
)

func main() {
	m := make(map[string]int)
	m["ayush"] = 7
	m["anythingElse"] = 14
	fmt.Println(m)
	theFirstValue := m["ayush"]
	fmt.Println(theFirstValue)
	// to erase an element in the map we use delete
	delete(m, "ayush")
	fmt.Println(m)

	// to erase an element we use clear
	clear(m)
	fmt.Println(m)
	//THIS IS VERY IMPORTANT
	//When we access any element in the map it also returns an optional second value which tells us whether the key is present in
	// the map or not
	m["ayush"] = 26
	val, whetherPresent := m["ayush"]
	fmt.Println(val, whetherPresent) // will return true because the value is present
	val1, whetherPresent1 := m["ankit"]
	fmt.Println(val1, whetherPresent1) // becuase the key is not present the value is returned as 0, but using whetherPresent(the second variable)
	// we can check if the key is present in the map or not

	n := map[string]int{"shubham": 2, "utkarsh": 56} // another way to create a map
	fmt.Println(n)

	fmt.Println(maps.Equal(m, n)) // to check whether the key value pairs are equal in two maps
}
