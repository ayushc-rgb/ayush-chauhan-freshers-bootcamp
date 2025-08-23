package main

import "fmt"

func Calculate(val int) int {
	return val + 2
}

func main() {
	result := Calculate(2)
	fmt.Println("The value that we got")
	fmt.Println(result)
}
