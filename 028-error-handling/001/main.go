package main

import (
	"fmt"
	"log"
)

type Errors struct {
	function string
	msg      string
}

func (e Errors) Error() string {
	return e.function + " ; " + e.msg
}

func main() {
	ans, err := functionReturningError()
	if err != nil {
		log.Fatalf("Error returned: %v", err)
	}
	fmt.Println(ans)
}

func functionReturningError() (int, error) {
	return 0, Errors{function: "functionReturningError", msg: "This is the error message"}
}
