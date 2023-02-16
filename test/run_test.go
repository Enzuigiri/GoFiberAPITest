package test

import "testing"

func TestPrint(t *testing.T) {
	result := Print("enzu")
	if result != "Hello enzu" {
		panic("Result is not Hello enzu")
	}
}
