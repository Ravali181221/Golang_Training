package main

import "testing"

func TestHello(t *testing.T) {
	input := Hello()

	checkWithValue := "Hello, World"

	if input != checkWithValue {
		t.Errorf("input is mismatching with value %s", input)
	}
}
func Test_Ravali(t *testing.T) {
	input := Hello_Ravali()

	checkWithValue := "Hello, world."

	if input != checkWithValue {
		t.Errorf("input is mismatching with value %s", input)
	}
}
