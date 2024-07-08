package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type alias int

func main() {
	var x alias = 3
	sum := addG(1, x)
	fmt.Println(sum)
	sumF := add(1.1, 2.2)
	fmt.Println(sumF)

	sumG := addG("a", "b")
	fmt.Println(sumG)
}

func add[T int | float64](a, b T) T {
	return a + b
}

func addG[T numbers](a ...T) T {
	var sum T
	for _, v := range a {
		sum += v
	}
	return sum
}

type numbers interface {
	constraints.Integer | constraints.Float | ~string
}
