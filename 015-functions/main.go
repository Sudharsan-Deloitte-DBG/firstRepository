package main

import "fmt"

func main() {
	xs := []int{1, 3, 5, 7, 9, 10}
	res := variadic(xs...)
	fmt.Println("the sum of all numbers in this slice is: ", res)

	p1 := Person{
		first: "firstname",
	}

	p1.Speak()

}

func variadic(i ...int) int {
	n := 0
	for _, v := range i {
		n += v
	}
	return n
}
