package games

import (
	"github.com/xavesen/reversi-bot/models"
	"github.com/xavesen/reversi-bot/utils"
)

func minimax(depth int, isMaximizing bool, alpha int, beta int, directions *[]func(uint64)uint64, game *models.GameState) (int, uint64) {
	result := game.FindLegalMoves(directions)
	switch result {
	case models.GameOverBlackWon:
		return 10000, 0
	case models.GameOverWhiteWon:
		return -10000, 0
	case models.GameOverDraw:
		return 0, 0
	case models.GameContinues:
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
