package main

import "fmt"

func main() {
	fmt.Println(Add(10, 15))
}

func Add(a int, b int) int {
	return a - b
}
