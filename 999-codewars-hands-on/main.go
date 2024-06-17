package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println(NbDig(5750, 0))

}

func NbDig(n int, d int) int {
	var xs string
	var count int
	for i := 0; i <= n; i++ {
		xs += strconv.Itoa(i * i)
	}
	for _, v := range xs {
		if int(v-'0') == d {
			count++
		}
	}
	return count
}
