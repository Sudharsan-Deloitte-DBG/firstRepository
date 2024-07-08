package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	xs := []int{1, 3, 5, 7, 9, 10}
	res := variadic(xs...)
	fmt.Println("the sum of all numbers in this slice is: ", res)

	p1 := Person{
		first: "firstname",
	}

	p1.Speak()

	//below example is an anonymous function that returns a string
	str := func(s string) string {
		return s
	}("Hello")

	//we could also assign a func to a variable as below and use that variable identifier to invloke/call the func by providing necessary arguments if any

	fexp := func(a int) int {
		return a * a
	}

	i := fexp(10)
	fmt.Println(i)

	fmt.Println(str)

	y := returnsFunc(10)
	fmt.Println(y())

	addition := doMath(10, 20, add)
	fmt.Println("Sum: ", addition)

	subtraction := doMath(20, 10, subtract)
	fmt.Println("Difference: ", subtraction)

	f := increment()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	g := increment()
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())
	fmt.Println(g())

	fact := factorial(5)
	fmt.Println(fact)

	xb, err := readFile("sample.txt")
	if err != nil {
		log.Fatalf("error occured when calling readFile from main: %v", err)
	}
	fmt.Println(string(xb))

}

func variadic(i ...int) int {
	n := 0
	for _, v := range i {
		n += v
	}
	return n
}

// this is a function that returns a function. When calling the returnsFunc, the returned function can be assigned to a variable and that variable can be used to invoke the returned function from returnsFunc
func returnsFunc(a int) func() int {
	return func() int {
		return a
	}
}

func add(a int, b int) int {
	return a + b
}

func subtract(a int, b int) int {
	return a - b
}

func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

func increment() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func readFile(filename string) ([]byte, error) {
	xb, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("error occured when reading file inside readFile func: %v", err)
	}
	return xb, nil
}
