package main

import (
	"fmt"
	"math/rand"
)

func init() {
	fmt.Println("This is where initialization of the program occurs!!")
}

func main() {
	x := rand.Intn(250)
	fmt.Printf("The variable X is of Type %T and has Value %v", x, x)

	// if x > 0 && x <= 100 {
	// 	fmt.Println("between 0 and 100")
	// } else if x > 100 && x <= 200 {
	// 	fmt.Println("between 101 and 200")
	// } else {
	// 	fmt.Println("between 201 and 250")
	// }

	switch {
	case x < 100:
		fmt.Println("less than 100")
	case x >= 100 && x < 200:
		fmt.Println("between 101 and 200")
	case x >= 200 && x < 250:
		fmt.Println("between 201 and 250")
	default:
		fmt.Println("250 or more")
	}

	checkRandIntervals(3)

	for u := 0; u < 100; u++ {
		handsOn22()
	}

	for v := 0; v <= 42; v++ {
		fmt.Println("This is iteration: ", v)
		handsOn23()
	}

}

func checkRandIntervals(a int) {
	for x := 0; x < 10; x++ {
		fmt.Println(rand.Intn(a))
	}
}

func handsOn22() {
	a := rand.Intn(10)
	b := rand.Intn(10)

	res := ""

	fmt.Printf("A: %v B: %v\n", a, b)

	// if a < 4 && b < 4 {
	// 	res = "both less than 4"
	// } else if a > 6 && b > 6 {
	// 	res = "both greater than 6"
	// } else if a >= 4 && a <= 6 {
	// 	res = "a is greater than or equal to 4 and less than or equal to 6"
	// } else if b != 5 {
	// 	res = "b is not equal to 5"
	// } else {
	// 	res = "none of the previous cases were met"
	// }

	switch {
	case a < 4 && b < 4:
		res = "both less than 4"
	case a > 6 && b > 6:
		res = "both greater than 6"
	case a >= 4 && a <= 6:
		res = "a is greater than or equal to 4 and less than or equal to 6"
	case b != 5:
		res = "b is not equal to 5"
	default:
		res = "none of the previous cases were met"
	}

	fmt.Println(res)
}

func handsOn23() {
	x := rand.Intn(5)
	switch x {
	case 0:
		fmt.Printf("The variable is of type %T and value %v \n", x, x)
	case 1:
		fmt.Printf("The variable is of type %T and value %v \n", x, x)
	case 2:
		fmt.Printf("The variable is of type %T and value %v \n", x, x)
	case 3:
		fmt.Printf("The variable is of type %T and value %v \n", x, x)
	case 4:
		fmt.Printf("The variable is of type %T and value %v \n", x, x)
	default:
		fmt.Println("You're not going ti print me anyway...")
	}
}
