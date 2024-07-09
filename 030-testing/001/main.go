package main

import "fmt"

func main() {
	fmt.Println(add(1, 5))
	fmt.Println(add(2, 6, 8))
	fmt.Println(add(1, 5, 9, 43))
	fmt.Println(add(-11, -45))
}

func add(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}
