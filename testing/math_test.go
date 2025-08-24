package main

import (
	"fmt"
	"testing"
)

type check struct {
	input    int
	expected int
}

<<<<<<< HEAD
func TestCalculate(t *testing.T) { // testing.T is a struct defined in Go built in testing package
=======
func TestCalculate(t *testing.T) {. // testing.T is a struct defined in Go built in testing package
>>>>>>> 536a76cb8750eb19edc7edc2a7dca41a077e0697
	if Calculate(2) != 4 {
		t.Error("The value should be 4")
	}
	t.Log("Success")
}

// t.error marks the test as fail and log the message
// t.fail marks the test as fail without any error message
// t.log logs the message but the test is not failed
// t.fatalf marks the test as fail, logs the message, and ends the testing immediately

func TestArrayCalculate(t *testing.T) {
	var tabul = []check{
		{1, 3},
		{22, 24},
		{25, 27},
		{999, 1001},
	}
<<<<<<< HEAD
	for _, tt := range tabul {
		testName := fmt.Sprintf("%d", tt.input)
		t.Run(testName, func(t *testing.T) {
			ans := Calculate(tt.input)
			if ans != tt.expected {
				t.Errorf("This test doesn't pass\n Input:%d \n Output:%d\nExpected:%d", tt.input, ans, tt.expected)
			}
		})
	}
}

func BenchmarkCalculate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Calculate(2)
=======
	for _, val := range tabul {
		if output := Calculate(val.input); output != val.expected {
			t.Error("Test failed Input:{}\nOutput:{}\nExpected:{}", val.input, output, val.expected)
		}
>>>>>>> 536a76cb8750eb19edc7edc2a7dca41a077e0697
	}
}
