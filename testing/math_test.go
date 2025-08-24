package main

import (
	"fmt"
	"testing"
)

type check struct {
	input    int
	expected int
}

func TestCalculate(t *testing.T) { // testing.T is a struct defined in Go built in testing package
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
	}
}
