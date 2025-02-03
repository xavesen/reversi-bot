package games

import (
	"fmt"

	"github.com/xavesen/reversi-bot/models"
)

func LocalGame(directions *[]func(uint64)uint64, game *models.GameState) {
	game.PrintBoard(false)
	gameOver := false

	for !gameOver {
		switch result := game.FindLegalMoves(directions); result {
		case models.GameOverBlackWon:
			fmt.Println("Game over, black won")
			gameOver = true
			continue
		case models.GameOverWhiteWon:
			fmt.Println("Game over, white won")
			gameOver = true
			continue
		case models.GameOverDraw:
			fmt.Println("Game over, draw")
			gameOver = true
			continue
		case models.GameContinues:
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