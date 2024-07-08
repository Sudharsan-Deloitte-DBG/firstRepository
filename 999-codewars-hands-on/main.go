package main

import (
	"fmt"
	"slices"
	"strconv"
)

func main() {

	//s := TwiceAsOld(55, 30)
	//fmt.Println(s)

	xs := []int{11, 5, 16}

	res := Gimme(xs)

	fmt.Println(res)

	//fmt.Println(NbDig(5750, 0))

	str := DNAStrand("AAAA")
	fmt.Println(str)

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

func TwiceAsOld(d int, s int) int {
	if d == s*2 {
		return 0
	} else if d-s > d/2 {
		return goToFuture(d, s)
	} else {
		return goToPast(d, s)
	}
}

func goToPast(d int, s int) int {
	count := 0
	for {
		d--
		s--
		count++
		if d == s*2 {
			return count
		}
	}
}

func goToFuture(d int, s int) int {
	count := 0
	for {
		d++
		s++
		count++
		if d == s*2 {
			return count
		}
	}
}

func Gimme(array []int) int {
	for idx, val := range array {
		if val != slices.Min(array) && val != slices.Max(array) {
			return idx
		}
	}
	return 0
}

func DNAStrand(dna string) string {
	mp := map[string]string{
		"A": "T",
		"T": "A",
		"G": "C",
		"C": "G",
	}
	var result string
	for _, v := range dna {
		result += mp[string(v)]
	}
	return result
}
