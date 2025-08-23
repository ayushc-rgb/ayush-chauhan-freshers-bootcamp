package main

import "testing"

func TestCalculate(t *testing.T) {
	if Calculate(2) != 4 {
		t.Error("The value should be 4")
	}
	t.Log("Success")

}
