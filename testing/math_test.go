package main

import "testing"

type check struct {
	input    int
	expected int
}

func TestCalculate(t *testing.T) {. // testing.T is a struct defined in Go built in testing package
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
	for _, val := range tabul {
		if output := Calculate(val.input); output != val.expected {
			t.Error("Test failed Input:{}\nOutput:{}\nExpected:{}", val.input, output, val.expected)
		}
	}
}
