package main

import "fmt"

func main() {
	// y := 41
	// x := 11

	// //like a C for
	// for y := 10; y < 20; y++ {
	// 	fmt.Println("print me till the first loop ends")
	// }

	// //like a C while
	// for y > 42 && y < 45 {
	// 	fmt.Println("print me until the second loop ends")
	// 	y++
	// }

	// //like a C for(;;)
	// for {
	// 	if y > 42 {
	// 		fmt.Println("Print the if condition in this loop")
	// 		break
	// 	}
	// 	y++
	// }

	// for {
	// 	if x > 10 && x < 200 {
	// 		fmt.Println(x)
	// 		if x%2 != 0 {
	// 			x++
	// 			continue
	// 		}
	// 		fmt.Println("Print odd: ", x)
	// 	}
	// 	x++
	// }

	xi := []int{1, 3, 4, 6, 7}
	for i, v := range xi {
		fmt.Println(i, v)
	}

	x2 := map[string]int{
		"Name":   10,
		"Noname": 20,
	}
	for k, v := range x2 {
		fmt.Println(k, v)
	}
}
