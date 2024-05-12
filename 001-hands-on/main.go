package main

import "fmt"

//package scope
var a int = 20

const b = 30

func main() {

	//block scope
	x := 50

	fmt.Println(a, b, x)

}
