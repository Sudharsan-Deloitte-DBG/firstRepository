package main

import "fmt"

func main() {

	mp := map[string]int{
		"string1": 1,
		"string2": 2,
		"string3": 3,
		"string4": 4,
	}

	//fmt.Println(mp["string2"])

	// fmt.Println(mp)

	// delete(mp, "string3")

	// fmt.Println(mp)

	// if value, ok := mp["string7"]; ok {
	// 	fmt.Println(value)
	// } else {
	// 	fmt.Println("no key exists")
	// }

	if val, ok := mp["string4"]; ok {
		fmt.Println(val)
	} else {
		fmt.Println("no key exists")
	}

}
