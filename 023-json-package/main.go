package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Sample struct {
	First string
	Last  string
}

func main() {
	s := Sample{
		First: "first",
		Last:  "last",
	}

	m, err := json.Marshal(s)
	if err != nil {
		log.Fatalf("Unexpected scenario: %v", err)
	}

	fmt.Println(string(m))

	err = json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		log.Fatalf("Unmarshal failed: %v", err)
	}
	fmt.Printf("%+v", animals)
}

var jsonBlob = []byte(`[{"Name": "Octopus", "Category": "Sea animal"},{"Name": "Cheetah", "Category": "Land animal"}]`)

type Animal struct {
	Name     string
	Category string
}

var animals []Animal
