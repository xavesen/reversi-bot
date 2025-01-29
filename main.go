package main

import (
	"fmt"

	"github.com/xavesen/reversi-bot/models"
	"github.com/xavesen/reversi-bot/utils"
)

const (
	maxUint64 uint64 = 0xffffffffffffffff
	gameContinues uint64 = maxUint64
	gameOverDraw uint64 = maxUint64 - 1
	gameOverBlackWon uint64 = maxUint64 - 2
	gameOverWhiteWon uint64 = maxUint64 - 3
)

func main() {
	game := models.GameState{
		White: 68853694464,
		Black: 34628173824,
		IsBlack: true,
	}
	directions := []func(uint64)uint64{utils.ShiftN, utils.ShiftNe, utils.ShiftE, utils.ShiftSe, utils.ShiftS, utils.ShiftSw, utils.ShiftW, utils.ShiftNw}
	localGame(&directions, &game)
}

func localGame(directions *[]func(uint64)uint64, game *models.GameState) {
	game.PrintBoard(false)
	gameOver := false

	for !gameOver {
		switch result := game.FindLegalMoves(directions); result {
		case gameOverBlackWon:
			fmt.Println("Game over, black won")
			gameOver = true
			continue
		case gameOverWhiteWon:
			fmt.Println("Game over, white won")
			gameOver = true
			continue
		case gameOverDraw:
			fmt.Println("Game over, draw")
			gameOver = true
			continue
		case gameContinues:
			game.IsBlack = !game.IsBlack
			continue
		}

		alpha := -1000000
		beta := 1000000
		_, move := minimax(6, game.IsBlack, alpha, beta, directions, game)
		game.ApplyMove(move, directions)
		game.PrintBoard(true)
	}
}

func minimax(depth int, isMaximizing bool, alpha int, beta int, directions *[]func(uint64)uint64, game *models.GameState) (int, uint64) {
	result := game.FindLegalMoves(directions)
	switch result {
	case gameOverBlackWon:
		return 10000, 0
	case gameOverWhiteWon:
		return -10000, 0
	case gameOverDraw:
		return 0, 0
	case gameContinues:
		newGame := game
		newGame.IsBlack = !newGame.IsBlack
		score, _ := minimax(depth, !isMaximizing, alpha, beta, directions, newGame)
		return score, 0
	}
	if depth == 0 {
		return game.EvaluateGameState(), 0
	}

	bestScore := -100000
	bestOrigScore := 0
	localAlpha := alpha
	localBeta := beta
	var bestMove uint64 = 0
	var score int

	for result != 0 {
		move := utils.FindFirstSetBit(result)
		result &= ^move
		newGame := *game
		newGame.ApplyMove(move, directions)
		origScore, _ := minimax(depth - 1, !isMaximizing, localAlpha, localBeta, directions, &newGame)
		if !isMaximizing {
			score = -1 * origScore
		} else {
			score = origScore
		}
		if score > bestScore {
			bestScore = score
			bestOrigScore = origScore
			bestMove = move

			if !isMaximizing {
				if origScore < localAlpha {
					return origScore, move
				} else {
					localBeta = origScore
				}
			} else {
				if origScore > localBeta {
					return origScore, move
				} else {
					 localAlpha = origScore
				}
			}
		}
	}

	return bestOrigScore, bestMove
}
