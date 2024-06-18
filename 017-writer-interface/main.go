package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

type person struct {
	first string
}

func (p person) writeOut(w io.Writer) {
	w.Write([]byte(p.first))
}

func main() {
	p := person{
		first: "James",
	}

	f, err := os.Create("sample.txt")
	if err != nil {
		log.Fatalf("Couldn't create file: %v", err)
	}

	defer f.Close()

	bs := []byte("Hello Gophers!")

	_, err = f.Write(bs)
	if err != nil {
		log.Fatalf("Couldn't write to file: %v", err)
	} else {
		fmt.Println("Wrote to file: ", f.Name())
	}

	b := bytes.NewBufferString("")
	b.WriteString("new line")
	fmt.Println(b.String())
	//resetting the buffer, which means all data in the buffer is cleared
	//b.Reset()
	fmt.Println("Unused bytes: ", b.Available())

	p.writeOut(f)
	p.writeOut(b)
	fmt.Println(b.String())

}
