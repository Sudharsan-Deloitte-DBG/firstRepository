package main

import (
	"fmt"
	"math/rand"
)

func idioms() {
	for i := 0; i < 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Println("X is 3")
		}
	}
}
