package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Message struct {
	str  string
	wait chan bool
}

func main() {
	c := make(chan string)
	go boring("boring!", c)
	sc := boringToo("boring Too!")
	for i := 0; i < 5; i++ {
		fmt.Println("Message received:", <-c)
		fmt.Println("Message received:", <-sc)
	}
	time.Sleep(5 * time.Second)
	fmt.Println("You're boring, I'm leaving!!")

	fc := fanin(boringToo("joe"), boringToo("ann"))
	for i := 0; i < 10; i++ {
		msg1 := <-fc
		fmt.Println(msg1.str)
		msg2 := <-fc
		fmt.Println(msg2.str)
		msg1.wait <- true
		msg2.wait <- true
	}
	fmt.Println("You're boring!")

}

func boring(msg string, c chan<- string) {
	for i := 0; ; i++ {
		c <- msg
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
	}
}

var waitForIt chan bool

func boringToo(msg string) <-chan Message {
	c := make(chan Message)

	go func() {
		for i := 0; ; i++ {
			c <- Message{fmt.Sprintf("%s %d", msg, i), waitForIt}
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			<-waitForIt
		}
	}()

	return c
}

func fanin(input1, input2 <-chan Message) <-chan Message {
	c := make(chan Message)
	go func() {
		for {
			c <- <-input1
		}
	}()
	go func() {
		for {
			c <- <-input2
		}
	}()
	return c
}
