package main

import (
	"fmt"
	"math/rand"
)

func main() {
	min := 1
	max := 100
	secretNumber := rand.Intn(max-min+1) + min
	maxAttempts := 10

	fmt.Printf("Welcome to the Number Guessing Game!\n")
	fmt.Printf("I'm thinking of a number between %d and %d.\n", min, max)
	fmt.Printf("You have %d attempts to guess it correctly.\n\n", maxAttempts)

	// Game loop
	for attempts := 1; attempts <= maxAttempts; attempts++ {
		var guess int
		fmt.Printf("Attempt %d: Enter your guess: ", attempts)
		fmt.Scan(&guess)

		if guess < secretNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > secretNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed it in %d attempts.\n", attempts)
			return
		}
	}

	fmt.Printf("\nSorry, you've used all your attempts. The number was %d.\n", secretNumber)
}
