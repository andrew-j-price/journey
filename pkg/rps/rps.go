package rps

import (
	"math/rand"
	"time"

	"github.com/andrew-j-price/journey/pkg/logger"
)

func init() {
	logger.Logger()
	rand.Seed(time.Now().UnixNano())
}

type TheScore struct {
	UserWins int
	CompWins int
	Draws    int
}

var choices = []string{"rock", "paper", "scissors"}

func RandomChoice() string {
	/* // NOTE: Longer way
	randomIndex := rand.Intn(len(choices))
	choice := choices[randomIndex]
	return choice
	*/
	return choices[rand.Intn(len(choices))]
}

func PlayGame(score *TheScore, player_choice string, computer_choice string) string {
	// logger.Debug.Printf("User plays: %v, Computer plays: %v\n", player_choice, computer_choice)
	response := []string{"draw", "userWin", "compWin"}
	if player_choice == computer_choice {
		logger.Info.Printf("Draw... Both threw: %v\n", player_choice)
		score.Draws++
		return response[0]
	} else if player_choice == "rock" && computer_choice == "scissors" {
		logger.Info.Printf("Player wins! User played: %v, Computer played: %v\n", player_choice, computer_choice)
		score.UserWins++
		return response[1]
	} else if player_choice == "scissors" && computer_choice == "paper" {
		logger.Info.Printf("Player wins! User played: %v, Computer played: %v\n", player_choice, computer_choice)
		score.UserWins++
		return response[1]
	} else if player_choice == "paper" && computer_choice == "rock" {
		logger.Info.Printf("Player wins! User played: %v, Computer played: %v\n", player_choice, computer_choice)
		score.UserWins++
		return response[1]
	} else {
		logger.Info.Printf("Computer won :( User played: %v, Computer played: %v\n", player_choice, computer_choice)
		score.CompWins++
		return response[2]
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
