package main

import "testing"

func TestCalculate(t *testing.T) {

	if calculate(2) != 4 {
		t.Error("Expected 2 + 2 = 4")
	}

	if calculate(4) == 4 {
		t.Error("Expected 2 + 2 = 4")
	}

}
