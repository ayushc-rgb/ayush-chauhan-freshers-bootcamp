package main

import (
	"testing"
)

type Database interface {
	Get(key int) (string, error)
	Set(kye int, value string) error
}

// in this technique of creating a mockdatabase we implement the interface directly using the set and get method
type MockDatabase struct {
	GetFunc func(key string) (string, error)
	SetFunc func(key, value string) error
}

func (db *MockDatabase) Get(key string) (string, error) {
	return db.GetFunc(key)
}
func (db *MockDatabase) Set(key, value string) error {
	return db.SetFunc(key, value)
}

// this is involving function callbacks
// when we deal with the function that is not present in the interface or that is not part of the interface
func ProcessData(data string, f func(string) string) string {
	result := f(data)
	return result
}
func TestProcessData(t *testing.T) {
	mockFunc := func(data string) string {
		return "This is a fixed string"
	}
	result := ProcessData("something", mockFunc)
	if result != "This is a fixed string" {
		t.Error("There is some error in this function")
	}
}
