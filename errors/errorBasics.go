package main

import (
	"errors"
	"fmt"
)

func returnError(arg int) (int, error) {
	if arg == 2 {
		return -1, fmt.Errorf("This is not a valid argument")
	}
	return arg, nil
}

func main() {
	displayError := fmt.Errorf("This is the new error") // this is just to return a new error
	fmt.Println(displayError)
	// this is how to return an error from a function
	arr := []int{1, 2}
	for _, val := range arr {
		if val, err := returnError(val); err != nil {
			fmt.Println("The function failed: ", err)
		} else {
			fmt.Println("The functions return no error and the value returned is:", val)
		}
	}
	//we can also wrap the errors with other errrors to give them wide meanings
	error1 := fmt.Errorf("This is the inner error")
	error2 := fmt.Errorf("This is the outer error and -> %w", error1)
	fmt.Println(error2)

	// now suppose we have the 'error2' now we need to check if the error1 is also present in 'error2'
	// here we can't use equality operator but we can use error.Is(errro2,error1) to check if the errro1 is present in error 2

	if errors.Is(error2, error1) {
		fmt.Println("The inner error is present in the outer error") // if we would have used "==" this would have been false
	} else {
		fmt.Println("The inner error is not present")
	}

}
