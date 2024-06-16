package main

import "fmt"

func iterateMaps() {
	mi := map[string]int{
		"Name1": 1,
		"Name2": 2,
		"Q":     3,
	}
	// for k, v := range mi {
	// 	fmt.Println("Key: ", k, "Value: ", v)
	// }

	// val := mi["Name1"]
	// fmt.Println(val)

	// for k, v := range mi {
	// 	if mi[k] == 'Q' {
	// 		fmt.Println("Key: ", k, "Value: ", v)
	// 	}
	// }

	// for k, v := range mi {
	// 	if mi[k] == 'Q' {
	// 		fmt.Println("Key: ", k, "Value: ", v)
	// 	}
	// }

	// if valu, ok := mi["Q"]; ok {
	// 	fmt.Println(valu)
	// }

	if val, ok := mi["Q"]; ok {
		fmt.Println(val)
	}
}
