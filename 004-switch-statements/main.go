package main

import "fmt"

func main() {
	x, y := 4, 20

	switch {
	case x > 10:
		fmt.Println("The value is 10")
	case y == 20:
		fmt.Println("This value is 20")
		fallthrough
	default:
		fmt.Println("This is a default having a fallthrough")
	}
}
