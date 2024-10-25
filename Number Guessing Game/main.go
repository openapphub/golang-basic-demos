package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	target := rand.Intn(100) + 1
	var guess int

	fmt.Print("Guess a number between 1 and 100")

	for {
		fmt.Print("Enter you guess:")
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Your guess was too low!")
		} else if guess > target {
			fmt.Println("Your guess was too high!")
		} else {
			fmt.Println("Your guess was correct!")
			break
		}
	}
}
