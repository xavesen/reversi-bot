package games

import (
	"errors"
	"fmt"
	"time"

	"github.com/xavesen/reversi-bot/api"
	"github.com/xavesen/reversi-bot/config"
	"github.com/xavesen/reversi-bot/models"
	"github.com/xavesen/reversi-bot/utils"
)

func ServerGame(directions *[]func(uint64)uint64, conf *config.Config) error {
	var gameId string = ""
	var color string = ""
	var gameStatus *models.GameStatusResult
	var isMaximizing bool

	gameList, err := api.GetGameList(conf)
	if err != nil {
		return err
	}

	if len(gameList) != 0 {
		color, gameId = joinGame(gameList, conf)
	}
	
	if gameId == "" { // game id will be unset if joining games failed or game list if empty
		gameInfo, err := api.CreateGame(conf)
		if err != nil {
			return err
		}
		gameId = gameInfo.GameId
		color = gameInfo.Color
	}

	if color == "black" {
		isMaximizing = true
	} else {
		isMaximizing = false
	}

	game := models.GameState{
		White: 68853694464,
		Black: 34628173824,
		IsBlack: true,
	}

	for {
		gameStatus, err = api.GetGameStatus(gameId, conf)
		if err != nil {
			return err
		}

		switch gameStatus.Status {
		case "pending":
			fmt.Println("waiting for player to join")
		case "white_won", "black_won":
			fmt.Println(gameStatus.Status)
			return nil
		case color:
			//s
		default:
			fmt.Println("Waiting for opp move")
		}

		time.Sleep(2 * time.Second)
	}
}

func joinGame(gameList []models.GameInfo, config *config.Config) (string, string) {
	for _, game := range gameList {
		joinGameRes, err := api.JoinGame(game.GameId, config)
		if err != nil {
			fmt.Printf("Failed joining game with id %s after retrying, trying another one if any\n", game.GameId)
			continue
		}

		return joinGameRes.Color, game.GameId
	}
	
	return "", ""
}

func makeOppMove(gameStatus *models.GameStatusResult, game *models.GameState, directions *[]func(uint64)uint64) {
	
	if gameStatus.LastMove != "" {
		bitMove := utils.AlgToBit(gameStatus.LastMove)
		game.ApplyMove(bitMove, directions)
		fmt.Printf("Opp move: %s\n", gameStatus.LastMove)
		game.PrintBoard(true)
	}
}

func makeMyMove(gameId string, gameStatus *models.GameStatusResult, game *models.GameState, isMaximizing bool, conf *config.Config, directions *[]func(uint64)uint64) error {
	

	currMoves := game.FindLegalMoves(directions)
	switch currMoves{
	case models.GameContinues:
		api.MakeMove(gameId, "pass", conf)
		game.IsBlack = !game.IsBlack
	case models.GameOverBlackWon, models.GameOverWhiteWon, models.GameOverDraw:
		fmt.Println("Error: server didn't end the game but locally game is over")
		api.MakeMove(gameId, "resign", conf)
		return errors.New("local game is over but not server one")
	default:
		alpha := -1000000
		beta := 1000000

		_, myMove := minimax(conf.SearchDepth, isMaximizing, alpha, beta, directions, game)
		algMove := utils.BitToAlg(myMove)
		moveResult, err := api.MakeMove(gameId, algMove, conf)
		if err != nil {
			return err
		} else if !moveResult.Ok {
			fmt.Printf("Error making move %s, server responded with false as a move result\n", algMove)
			return errors.New("unable to apply move")
		}

		game.ApplyMove(myMove, directions)
		fmt.Printf("My move: %s\n", algMove)
		game.PrintBoard(true)

		if !moveResult.Continue {
			fmt.Printf("Game over, winner: %s\n", moveResult.Winner)
			return nil
		}
	}
}