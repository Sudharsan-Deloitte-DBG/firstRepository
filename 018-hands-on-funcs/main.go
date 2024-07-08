package main

import (
	"fmt"
	"math"
)

type person struct {
	first string
	age   int
}

type square struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.width
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

type shape interface {
	area() float64
}

func calcShapeArea(sh shape) float64 {
	return sh.area()
}

func (p person) speak() {
	fmt.Println("I am", p.first, "and I am", p.age, "years old")
}

func main() {

	sq := square{
		length: 10,
		width:  10,
	}

	ci := circle{
		radius: 10,
	}

	fmt.Println("Area of square:", calcShapeArea(sq))
	fmt.Println("Area of circle:", calcShapeArea(ci))

	p1 := person{
		first: "James",
		age:   27,
	}

	p1.speak()

	fmt.Println(namedReturn([]int{1, 2, 3, 4, 5}))

	//1st defer
	defer fmt.Println("First defer")

	fres := foo()
	defer fmt.Println("Third defer")
	bres, bresstring := bar()
	defer fmt.Println("Second defer")

	fmt.Println("foo result: ", fres)
	fmt.Println("bar result: ", bres, bresstring)

	fmt.Println(foovariadic([]int{1, 2, 3, 4, 5}...))
	fmt.Println(barsum([]int{1, 2, 3, 4, 5}))

	cnt := func(s string) int {
		count := 0
		for _, _ = range s {
			count++
		}
		return count
	}("This is a sample string")

	fmt.Printf("No of characters in string: %d", cnt)

	fn := func() {
		fmt.Println("This is an example of function expression")
	}

	fn()

	raf := returnsAFunc()
	raf(2)

	printSquare(squarefunc, 5)

	gc1 := provisionGiftCard()
	fmt.Println("Remaining amount in gc1:", gc1(500))
	fmt.Println("Remaining amount in gc1:", gc1(400))

	gc2 := provisionGiftCard()
	fmt.Println("Remaining amount in gc2:", gc2(512))
	fmt.Println("Remaining amount in gc2:", gc2(412))

}

func namedReturn(xs []int) (total int) {
	total = 0
	for _, v := range xs {
		total += v
	}
	return
}

func foo() int {
	return 0
}

func bar() (int, string) {
	return 0, "f"
}

func foovariadic(i ...int) int {
	sum := 0
	for _, val := range i {
		sum += val
	}
	return sum
}

func barsum(i []int) int {
	sum := 0
	for _, val := range i {
		sum += val
	}
	return sum
}

func returnsAFunc() func(int) {
	return func(i int) {
		fmt.Println("I just print it:", i)
	}
}

func printSquare(f func(int) int, i int) {
	fmt.Println(f(i))
}

func squarefunc(i int) int {
	return i * i
}

func provisionGiftCard() func(int) int {
	amount := 1000
	return func(debitAmount int) int {
		amount -= debitAmount
		return amount
	}
}
