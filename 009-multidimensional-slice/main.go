package main

import "fmt"

func main() {
	hdr := []string{"Firstname", "Lastname", "School", "Course"}

	snv := []string{"Sudharsan", "Venkatesh", "School of Computing", "Information and Communication Technology"}
	bk := []string{"BK", "BK", "SOM", "PE"}

	xxs := [][]string{hdr, snv, bk}

	fmt.Println(xxs)
}
