package main

import "testing"

func TestAdd(t *testing.T) {
	total := Add(10, 15)

	if total != 25 {
		t.Errorf("Add function failed to produce expected value - 25 and instead gave - %v", total)
	}
}
