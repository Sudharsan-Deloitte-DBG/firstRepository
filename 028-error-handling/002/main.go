package main

import (
	"log"
	"os"
)

func main() {
	f1, err := os.Create("log.txt")
	if err != nil {
		log.Println("Unable to create file:", err)
	}
	defer f1.Close()

	log.SetOutput(f1)

	f2, err := os.Open("sample1.txt")
	if err != nil {
		log.Println("Unable to open file:", err)
	}
	defer f2.Close()
}
