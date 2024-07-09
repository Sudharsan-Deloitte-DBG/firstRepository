package main

import "fmt"

// customErr is a struct that provides a custom error definition
type customErr struct {
	errorName    string
	errorMessage string
}

func (c customErr) Error() string {
	return c.errorName + " ; " + c.errorMessage
}

//the main function
func main() {
	fmt.Println("Creating a custom error message")
	err := customErr{
		errorName:    "Custom Error",
		errorMessage: "Custom Error Message",
	}

	foo(err)
}

// foo accepts an error type and prints it
func foo(err error) {
	if err != nil {
		fmt.Println(err.(customErr).errorMessage)
		// the above line where customErr is written within brackets is called assertion. When customErr is passed as a parameter to a function that accept error type, although
		// customErr implements the Error() method, the errorMessage field is not directly accessible as err.errorMessage inside foo. We need to assert that this error type is
		// also of type customErr by saying err.(customErr).errorMessage
	}
}
