package main

import "fmt"

func iterateSlices() {
	xi := []int{1, 2, 3, 4, 5}
	for i, v := range xi {
		fmt.Println("Index: ", i, "Value: ", v)
	}
}
