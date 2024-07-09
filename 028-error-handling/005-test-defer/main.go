package main

import "fmt"

func main() {
	i := 0
	defer fmt.Println(i)

	for {
		if i == 20 {
			break
		}
		i++
	}

	fmt.Println(i)

	fmt.Println("Testing defer for named returns:", testDefer(200))
}

func testDefer(i int) (j int) {
	j = i
	defer func() { j += 10 }()
	return
}
