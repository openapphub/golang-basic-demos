package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal("Error reading file:", err)
	}

	err = ioutil.WriteFile("output.txt", content, 0644)
	if err != nil {
		log.Fatal("Error writing file:", err)
	}

	fmt.Println("File copied successfully")
}
