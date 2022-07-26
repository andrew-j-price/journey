package rps

import (
	"math/rand"
	"time"

	"github.com/andrew-j-price/journey/logger"
)

func init() {
	// logger.Info.Printf("rps package initialized") // CAUSES: panic: runtime error: invalid memory address or nil pointer dereference
	// log.Printf("rps package initialized")
	rand.Seed(time.Now().UnixNano())
}

type TheScore struct {
	UserWins int
	CompWins int
	Draws    int
}

var choices [3]string = [3]string{"rock", "paper", "scissors"}

func RandomChoice() string {
	/* // NOTE: Longer way
	randomIndex := rand.Intn(len(choices))
	choice := choices[randomIndex]
	return choice
	*/
	return choices[rand.Intn(len(choices))]
}

func PlayGame(score *TheScore, player_choice string, computer_choice string) {
	// logger.Info.Printf("User plays: %v, Computer plays: %v\n", player_choice, computer_choice)
	if player_choice == computer_choice {
		logger.Info.Printf("Draw... Both threw: %v\n", player_choice)
		score.Draws++
	} else if player_choice == "rock" && computer_choice == "scissors" {
		logger.Info.Printf("Player wins! User played: %v, Computer played: %v\n", player_choice, computer_choice)
		score.UserWins++
	} else if player_choice == "scissors" && computer_choice == "paper" {
		logger.Info.Printf("Player wins! User played: %v, Computer played: %v\n", player_choice, computer_choice)
		score.UserWins++
	} else if player_choice == "paper" && computer_choice == "rock" {
		logger.Info.Printf("Player wins! User played: %v, Computer played: %v\n", player_choice, computer_choice)
		score.UserWins++
	} else {
		logger.Info.Printf("Computer won :( User played: %v, Computer played: %v\n", player_choice, computer_choice)
		score.CompWins++
	}
}

func playRandomGame(score *TheScore) {
	PlayGame(score, RandomChoice(), RandomChoice())
}

func LocalTesting() {
	gameScore := TheScore{UserWins: 0, CompWins: 0, Draws: 0}
	logger.Info.Printf("Score: %v\n", gameScore)
	logger.Info.Printf("Choices: %v\n", choices)
	logger.Info.Printf("Choice: %v\n", RandomChoice())
	logger.Info.Printf("Choice: %v\n", RandomChoice())
	logger.Info.Printf("Choice: %v\n", RandomChoice())
	playRandomGame(&gameScore)
	playRandomGame(&gameScore)
	playRandomGame(&gameScore)
	playRandomGame(&gameScore)
	playRandomGame(&gameScore)
	logger.Info.Printf("Score: %v\n", gameScore)
}
