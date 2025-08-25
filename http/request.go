package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	response, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("The status received is ::", response.Status) // this reads the status received from the get request
	scanner := bufio.NewScanner(response.Body)                // scanner object reads data line by line
	for i := 0; scanner.Scan() && i < 20; i++ {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
