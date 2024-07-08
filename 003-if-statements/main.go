package main

import (
	"fmt"
	"math/rand"
)

func main() {
	//SEQUENTIAL
	fmt.Println("This is the first line to be printed.")
	fmt.Println("This is the second line to be printed.")

	x := 41
	y := x - 10

	//Simple if
	if x > y {
		fmt.Println("First if conditional statement with a comparison operator ", x)
	}

	z := 1

	//if..else if..else statement
	if x > z {
		fmt.Println("This if conditional statement as long as z is greater than x")
	} else if x > y {
		fmt.Println("This else if conditional statement is executed as long as x is greater than y")
	} else {
		fmt.Println("This statement executes as long as x is greater than z and y is greater than x")
	}

	//Logical comparison operators
	if x > z && x > y {
		fmt.Println("Compared using logical operators to compare more than one condition")
	} else if z > x || z > y {
		fmt.Println("This doesn't get printed")
	} else if x != 42 {
		fmt.Println("What else is the meaning of life??!!")
	} else {
		fmt.Println("No satisfying condition")
	}

	//Statement; statement idiom
	if st := 2 * rand.Intn(x); st >= x {
		fmt.Println("This gets printed when this statement;statement condition is satisfied!!")
	}
}
