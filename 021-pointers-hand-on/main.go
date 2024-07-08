package main

import "fmt"

var (
	a, b, c *string
	d       *int
)

func init() {
	// x := 42
	// fmt.Println(&x)

	p := "Drop by drop, the bucket gets filled."
	q := "Persistently, patiently, you are bound to succeed."
	r := "The purpose of life is to just live it to the fullest."
	n := 42
	a = &p
	b = &q
	c = &r
	d = &n

}

func main() {
	fmt.Println("+++++++++++++++Values+++++++++++++++")
	fmt.Println(a, b, c, d)
	fmt.Println("+++++++++++++++Types++++++++++++++++")
	fmt.Printf("Type of a is: %T and value is: %v\n", a, a)
	fmt.Printf("Type of b is: %T and value is: %v\n", b, b)
	fmt.Printf("Type of c is: %T and value is: %v\n", c, c)
	fmt.Printf("Type of d is: %T and value is: %v\n", d, d)
	fmt.Println("+++++++++++Dereferencing++++++++++++")
	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
	fmt.Println(*d)
}
