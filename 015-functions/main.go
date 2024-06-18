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

}

func variadic(i ...int) int {
	n := 0
	for _, v := range i {
		n += v
	}
	return n
}

//this is a function that returns a function. When calling the returnsFunc, the returned function can be assigned to a variable and that variable can be used to invoke the returned function from returnsFunc
func returnsFunc(a int) func() int {
	return func() int {
		return a
	}
}
