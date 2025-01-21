# Rock Paper Scissors Lizard Spock Game

A simple command-line game of Rock Paper Scissors Lizard Spock written in Go.

## How to Play

1. Run the program using `go run main.go`.
2. The computer will randomly select one of the five options: rock, paper, scissors, lizard, or spock.
3. You will be prompted to enter your choice.
4. Type one of the five options (rock, paper, scissors, lizard, or spock) and press enter.
5. The program will determine the winner based on the game's rules.

## Rules

- Rock crushes Scissors
- Scissors cuts Paper
- Paper covers Rock
- Rock crushes Lizard
- Lizard poisons Spock
- Spock smashes Scissors
- Scissors decapitates Lizard
- Lizard eats Paper
- Paper disproves Spock
- Spock vaporizes Rock

## Code Structure

The program consists of two functions: `playGame` and `main`. The `playGame` function takes a slice of strings as an argument, representing the possible choices. It randomly selects the computer's choice and prompts the user for their choice. The `main` function sets up the game options and calls `playGame`.

## Error Handling

The program checks for invalid user input and prints an error message if the input is not one of the five valid options.

## Dependencies

- Go 1.14 or later
- `math/rand` package for random number generation
- `fmt` package for input/output operations
