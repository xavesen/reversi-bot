package main

import (
	"github.com/xavesen/reversi-bot/config"
	"github.com/xavesen/reversi-bot/games"
	"github.com/xavesen/reversi-bot/utils"
)

func main() {
	// game := models.GameState{
	// 	White: 68853694464,
	// 	Black: 34628173824,
	// 	IsBlack: true,
	// }
	directions := []func(uint64)uint64{utils.ShiftN, utils.ShiftNe, utils.ShiftE, utils.ShiftSe, utils.ShiftS, utils.ShiftSw, utils.ShiftW, utils.ShiftNw}
	// games.LocalGame(&directions, &game)

	conf := config.Config{
		ApiURL: "http://localhost:8000/",
		CreateGameURL: "/reversi/v1/create_game",
		GameListUrl: "/reversi/v1/game_list",
		GameStatusURL: "/reversi/v1/game_status",
		MakeMoveURL: "/reversi/v1/move",
		JoinGameURL: "/reversi/v1/join",
		PlayerUUID: "f328b162-d8be-11ef-85bf-244bfe533006",
		SearchDepth: 4,
		ApiRetries: 3,
	}

	games.ServerGame(&directions, &conf)
}
