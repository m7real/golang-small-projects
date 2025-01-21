package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func getTrimSpacedLower(s string) string {
	return strings.ToLower(strings.TrimSpace(s))
}

func contains(items []string, target string) bool {
	for _, v := range items {
		if v == target {
			return true
		}
	}
	return false
}

func playAgain() bool {
	var response string
	for {
		fmt.Println("Do you want to play again? (yes/no): ")
		fmt.Scanln(&response)
		response = getTrimSpacedLower(response)
		if response == "yes" || response == "no" {
			break
		}
		fmt.Println("Invalid input, please chose between 'yes' and 'no'")
	}
	return response == "yes"
}

func getWinner(com string, user string) {
	rules := map[string][]string{
		"rock":     {"scissors", "lizard"},
		"paper":    {"rock", "spock"},
		"scissors": {"paper", "lizard"},
		"lizard":   {"paper", "spock"},
		"spock":    {"scissors", "rock"},
	}

	fmt.Println("You chose:", user)
	fmt.Println("Computer chose:", com)
	if com == user {
		fmt.Println("It's a Tie!")
	} else if contains(rules[user], com) {
		fmt.Println("You Won!")
	} else {
		fmt.Println("You Lose!")
	}
}

func playGame(options []string) {
	computerChoice := options[rand.Intn(len(options))]
	var userChoice string
	for {
		fmt.Print("Enter your choice (rock, paper, scissors, lizard, spock): ")
		fmt.Scanln(&userChoice)
		userChoice = getTrimSpacedLower(userChoice)
		if contains(options, userChoice) {
			break
		}
		fmt.Println("Invalid choice, please try again")

	}
	getWinner(computerChoice, userChoice)

}

func main() {
	options := []string{"rock", "paper", "scissors", "lizard", "spock"}
	for {
		playGame(options)
		if playAgain() == false {
			break
		}
	}

}
