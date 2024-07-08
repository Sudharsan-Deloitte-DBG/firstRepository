package main

import "testing"

func TestAdd(t *testing.T) {
	addRes := add(10, 15)
	if addRes != 25 {
		t.Errorf("Addition failed. Expected %d and got %d", 25, addRes)
	}
}

func TestSubtract(t *testing.T) {
	subRes := subtract(15, 10)
	if subRes != 5 {
		t.Errorf("Subtraction failed. Expected %d and got %d", 5, subRes)
	}
}

func TestDoMath(t *testing.T) {
	addRes := doMath(10, 5, add)
	subRes := doMath(10, 5, subtract)

	if addRes != 15 {
		t.Errorf("Addition failed in doMath func. Expected %d and got %d", 15, addRes)
	}
	if subRes != 5 {
		t.Errorf("Subtraction failed in doMath func. Expected %d and got %d", 5, subRes)
	}
}
