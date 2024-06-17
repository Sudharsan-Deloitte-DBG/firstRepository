package main

import "fmt"

type Person struct {
	first string
}

func (p Person) Speak() {
	fmt.Println("I am ", p.first)
}
