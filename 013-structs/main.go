package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   30,
	}

	p2 := person{
		first: "Jenny",
		last:  "Moneypenny",
		age:   27,
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   30,
		},
		ltk: true,
	}

	fmt.Println(p1, p2)

	fmt.Println(p1.first, p2.last, p2.age)

	fmt.Println(sa1)

	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk, sa1.person.age)
}
