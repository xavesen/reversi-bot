package games

import (
	"fmt"

	"github.com/xavesen/reversi-bot/api"
	"github.com/xavesen/reversi-bot/config"
	"github.com/xavesen/reversi-bot/models"
)

func ServerGame(config *config.Config) error {
	var gameId string = ""
	var color string = ""

	gameList, err := getGameList(config)
	if err != nil {
		return err
	}

	if len(gameList) != 0 {
		color, gameId = joinGame(gameList, config)
	}
	
	if gameId == "" { // game id will be unset if joining games failed or game list if empty
		gameInfo, err := createGame(config)
		if err != nil {
			return err
		}
		gameId = gameInfo.GameId
		color = gameInfo.Color
	}

	fmt.Println(gameId, color)

	return nil
}

func getGameList(config *config.Config) ([]models.GameInfo, error) {
	gameListResp, err := api.GetGameList(config.ApiURL+config.GameListUrl, config.PlayerUUID)
	if err != nil {
		fmt.Printf("Encountered an error getting gamelist. Making %d more attempts.\n", config.ApiRetries)
		for i := 0; i < config.ApiRetries; i++ {
			gameListResp, err = api.GetGameList(config.ApiURL+config.GameListUrl, config.PlayerUUID)
			if err != nil {
				continue
			} else {
				break
			}
		}
		if err != nil {
			fmt.Println("Failed getting gamelist after retrying")
			return nil, err
		}
	}

	return gameListResp.Result, nil
}

func createGame(config *config.Config) (*models.CreateGameResult, error) {
	createGameResp, err := api.CreateGame(config.ApiURL+config.CreateGameURL, config.PlayerUUID)
	if err != nil {
		fmt.Printf("Encountered an error creating game. Making %d more attempts.\n", config.ApiRetries)
		for i := 0; i < config.ApiRetries; i++ {
			createGameResp, err = api.CreateGame(config.ApiURL+config.CreateGameURL, config.PlayerUUID)
			if err != nil {
				continue
			} else {
				break
			}
		}
		if err != nil {
			fmt.Println("Failed creating game after retrying")
			return nil, err
		}
	}

	return &createGameResp.Result, nil
}

func joinGame(gameList []models.GameInfo, config *config.Config) (string, string) {
	for _, game := range gameList {
		joinGameResp, err := api.JoinGame(config.ApiURL+config.JoinGameURL, config.PlayerUUID, game.GameId)
		if err != nil {
			fmt.Printf("Encountered an error joining game. Making %d more attempts.\n", config.ApiRetries)
			for i := 0; i < config.ApiRetries; i++ {
				joinGameResp, err = api.JoinGame(config.ApiURL+config.JoinGameURL, config.PlayerUUID, game.GameId)
				if err != nil || !joinGameResp.Result.Result {
					continue
				} else {
					break
				}
			}
			if err != nil {
				fmt.Printf("Failed joining game with id %s after retrying, trying another one if any\n", game.GameId)
				continue
			}
		}

		return joinGameResp.Result.Color, game.GameId
	}
	
	return "", ""
}