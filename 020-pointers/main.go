package main

import "fmt"

func main() {
	x := 42
	var p *int = &x

	fmt.Printf("Type of p: %T and Value of p: %v\n", p, p)

	*p = 10
	fmt.Println("Hello", *p)

	c1 := cat{"jack"}
	c1.walk()
	c1.run()

	// youngRun(c1)

	c2 := &cat{"Carl"}
	c2.run()
	c2.walk()

	youngRun(c2)

	st := something{"some name"}
	fmt.Println(st)
	st = changeFieldVS(st, "some other name")
	fmt.Println(st)

	changeFieldPS(&st, "again some other name")
	fmt.Println(st)

}

type cat struct {
	first string
}

func (c cat) walk() {
	fmt.Println(c.first, "is walking.")
}

func (c *cat) run() {
	fmt.Println(c.first, "is running.")
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
}

type something struct {
	first string
}

func changeFieldVS(s something, str string) something {
	s.first = str
	return s
}

func changeFieldPS(s *something, str string) {
	s.first = str
}
