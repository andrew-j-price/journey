package rps

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

var choices [3]string = [3]string{"rock", "paper", "scissors"}

type TheScore struct {
	UserWins int
	CompWins int
	Draws    int
}

func RandomChoice() string {
	/* // NOTE: Longer way
	randomIndex := rand.Intn(len(choices))
	choice := choices[randomIndex]
	return choice
	*/
	return choices[rand.Intn(len(choices))]
}

func PlayGame(score *TheScore, player_choice string, computer_choice string) {
	// fmt.Printf("User plays: %v, Computer plays: %v\n", player_choice, computer_choice)
	if player_choice == computer_choice {
		fmt.Printf("Draw.  Both threw: %v\n", player_choice)
		score.Draws++
	} else if player_choice == "rock" && computer_choice == "scissors" {
		fmt.Printf("Player wins!  User played: %v, Computer played: %v\n", player_choice, computer_choice)
		score.UserWins++
	} else if player_choice == "scissors" && computer_choice == "paper" {
		fmt.Printf("Player wins!  User played: %v, Computer played: %v\n", player_choice, computer_choice)
		score.UserWins++
	} else if player_choice == "paper" && computer_choice == "rock" {
		fmt.Printf("Player wins!  User played: %v, Computer played: %v\n", player_choice, computer_choice)
		score.UserWins++
	} else {
		fmt.Printf("Computer wins :(  User played: %v, Computer played: %v\n", player_choice, computer_choice)
		score.CompWins++
	}
}

func playRandomGame(score *TheScore) {
	PlayGame(score, RandomChoice(), RandomChoice())
}

func init() {
	log.Printf("rps package initialized")
	rand.Seed(time.Now().UnixNano())
}

func Main() {
	gameScore := TheScore{UserWins: 0, CompWins: 0, Draws: 0}
	fmt.Printf("Score: %v\n", gameScore)
	fmt.Printf("Choices: %v\n", choices)
	fmt.Printf("Choice: %v\n", RandomChoice())
	fmt.Printf("Choice: %v\n", RandomChoice())
	fmt.Printf("Choice: %v\n", RandomChoice())
	playRandomGame(&gameScore)
	playRandomGame(&gameScore)
	playRandomGame(&gameScore)
	playRandomGame(&gameScore)
	playRandomGame(&gameScore)
	fmt.Printf("Score: %v\n", gameScore)

}
