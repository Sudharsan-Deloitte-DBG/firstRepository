package main

import "fmt"

func iterateSlices() {
	xi := []int{1, 2, 3, 4, 5}
	for i, v := range xi {
		fmt.Println("Index: ", i, "Value: ", v)
	}

	fmt.Println(xi[0])
	fmt.Println(xi[1])
	fmt.Println(xi[2])
	fmt.Println(xi[3])
	fmt.Println(xi[4])
}

func appendToASlice() {
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(xi)

	// appending to a slice
	xi = append(xi, 00, 9)

	fmt.Println(xi)

}

func slicingASlice() {
	xi := []int{1, 2, 3, 4, 5, 7, 8, 9, 10}

	fmt.Println("==============================")

	// [inclusive:exclusive]
	fmt.Println(xi[1:8])

	//[:exclusive]
	fmt.Println(xi[:7])

	//[inclusive:]
	fmt.Println(xi[2:])

	//[:]
	fmt.Println(xi[:])

	//lets try this
	fmt.Println(xi[:2:8])
}

func makeSlice() {
	//using the make method
	xi := make([]string, 0, 20)
	xi = append(xi, "1", "2", "3")

	fmt.Println(len(xi))
	fmt.Println(cap(xi))

}
