package main

import "fmt"

func main() {
	// arr := []int{1, 2, 3, 4, 5}

	// // indices: [0, 1, 2, 3, 4, 5]
	// // values : [1, 2, 3, 4, 5, 6]

	// //arr1 := make([]int, 10)

	// fmt.Println("Array:", arr)

	// for _, val := range arr {
	// 	fmt.Printf("The value is: %v and type is: %T\n", val, val)
	// }

	// for i := 0; i <= 10; i-- {
	// 	//code to be written here
	// 	break
	// }

	// 0	1	2	3	4	5	6	7	8	9
	// 42	43	44	45	46	47	48	49	50	51

	// xs := createSlice()

	// fmt.Printf("%v\n", xs[:5])
	// fmt.Printf("%v\n", xs[5:])
	// fmt.Printf("%v\n", xs[2:7])
	// fmt.Printf("%v\n", xs[1:6])

	//appendToSlice()

	//deleteFromSlice()

	//states()

	//multidslice()

	createMap()

}

func createSlice() []int {
	x := []int{}

	for i := 42; i < 52; i++ {
		x = append(x, i)
	}

	for idx, val := range x {
		fmt.Printf("%v - %T - %v\n", idx, val, val)
	}

	return x

}

func appendToSlice() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	x = append(x, 52)

	fmt.Printf("%v", x)

	x = append(x, []int{53, 54, 55}...)

	fmt.Printf("%v", x)

	y := []int{56, 57, 58, 59, 60}

	x = append(x, y...)

	fmt.Printf("%v", x)
}

func deleteFromSlice() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// 0	1	2	3	4	5	6	7	8	9
	// 42	43	44	45	46	47	48	49	50	51

	x = append(x[:3], x[6:]...)

	fmt.Printf("%v", x)

}

func states() {
	st := make([]string, 0)

	st = append(st, "Alabama", " Alaska", " Arizona", " Arkansas", " California", " Colorado", " Connecticut", "Delaware", " Florida", " Georgia", " Hawaii", " Idaho", " Illinois", " Indiana", " Iowa", " Kansas", "Kentucky", " Louisiana", " Maine", " Maryland", " Massachusetts", " Michigan", " Minnesota", "Mississippi", " Missouri", " Montana", " Nebraska", " Nevada", " New Hampshire", " New Jersey", " New Mexico", " New York", " North Carolina", " North Dakota", " Ohio", " Oklahoma", " Oregon", " Pennsylvania", " Rhode Island", " South Carolina", " South Dakota", " Tennessee", " Texas", "Utah", " Vermont", " Virginia", " Washington", " West Virginia", " Wisconsin", " Wyoming")

	fmt.Println(st)

	fmt.Printf("Len: %v\t Cap: %v", len(st), cap(st))

	for i := 0; i < 50; i++ {
		fmt.Printf("%v - %v\n", i, st[i])
	}

}

func multidslice() {

	jb := []string{"James", "bond", "Shaken, not stirred"}
	mp := []string{"Miss", "Moneypenny", "I'm 008"}
	xxs := [][]string{jb, mp}

	for _, val := range xxs {
		for jdx, jal := range val {
			fmt.Printf("%v - %T - %v\n", jdx, jal, jal)
		}
	}
}

func createMap() {
	mp := map[string]int{}

	mp["string1"] = 1

	fmt.Println(mp)
}
