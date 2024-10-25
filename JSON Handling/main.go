package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData := `{"name": "Alice", "age": 30}`

	var person Person
	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		log.Fatal("Error parsing JSON:", err)
	}

	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
