package main

import "fmt"

func NestedLoops() {
	for i := 0; i < 5; i++ {
		fmt.Println("Loop A - Iteration: ", i)
		fmt.Println("========================")
		for j := 0; j < 5; j++ {
			fmt.Println("Loop B - Iteration: ", j)
		}
	}
}

// func prevExercise() {
// 	x := 42
// 	for x > 40 {
// 		fmt.Println(x)
// 		x--
// 	}

// 	for {
// 		fmt.Println(x)
// 		if x > 54 {
// 			break
// 		}
// 		x++
// 	}

// 	a := 1
// 	for {
// 		if a%2 == 0 {
// 			a++
// 			continue
// 		}
// 		fmt.Println("Odd: ", a)
// 		a++
// 		fmt.Println(a)
// 		if a > 15 {
// 			break
// 		}
// 	}
// }
