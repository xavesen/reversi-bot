package main

import (
	"github.com/xavesen/reversi-bot/games"
	"github.com/xavesen/reversi-bot/models"
	"github.com/xavesen/reversi-bot/utils"
)

func main() {
	game := models.GameState{
		White: 68853694464,
		Black: 34628173824,
		IsBlack: true,
	}
	directions := []func(uint64)uint64{utils.ShiftN, utils.ShiftNe, utils.ShiftE, utils.ShiftSe, utils.ShiftS, utils.ShiftSw, utils.ShiftW, utils.ShiftNw}
	games.LocalGame(&directions, &game)
}
