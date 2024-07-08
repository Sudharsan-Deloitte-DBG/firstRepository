package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type person struct {
	first string
}

func (p *person) speak() {
	fmt.Printf("I am %v speaking", p.first)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {
	p1 := person{"Firstname"}
	saySomething(&p1)
	p1.speak()

	wg.Add(2)
	go func1()
	go func2()

	wg.Wait()
}

func func1() {
	fmt.Println("Print something from func1!")
	wg.Done()
}

func func2() {
	fmt.Println("Print something from func2!")
	wg.Done()
}
