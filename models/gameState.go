package models

import (
	"fmt"
	"strconv"
	"github.com/xavesen/reversi-bot/utils"
)

const (
	green = "\033[32m"
	yellow = "\033[33m"
	blue = "\033[34m"
	reset = "\033[0m"
	leftBitSet uint64 = 0x8000000000000000
	maxUint64 uint64 = 0xffffffffffffffff
	gameContinues uint64 = maxUint64
	gameOverDraw uint64 = maxUint64 - 1
	gameOverBlackWon uint64 = maxUint64 - 2
	gameOverWhiteWon uint64 = maxUint64 - 3
)

type GameState struct {
	White uint64
	Black uint64
	IsBlack bool
	LastMove uint64
	LastRecolors uint64
}

func (game *GameState) PrintBoard(withMoveHighlight bool) {
	var bitMask uint64 = 1
	var j = 0
	var i = 0
	result := blue + "  0 1 2 3 4 5 6 7\n0 " + reset
	
	for bitMask != 0 {
		if j == 8 {
			j = 0
			i ++
			result += blue + strconv.Itoa(i) + "\n" + strconv.Itoa(i) + reset + " "
		}

		colored := false
		if withMoveHighlight {
			if bitMask == game.LastMove {
				result += green
				colored = true
			} else if bitMask & game.LastRecolors != 0 {
				result += yellow
				colored = true
			}
		}
		
		if bitMask & game.Black != 0 {
			result += "x "
		} else if bitMask & game.White != 0 {
			result += "o "
		} else {
			result += ". "
		}

		if withMoveHighlight && colored {
			result += reset
		}

		bitMask <<= 1
		j++
	}

	result += blue + "8\n" + "  a b c d e f g h" + reset

	fmt.Println(result)
}

func (game *GameState) EvaluateGameState() int {
	return utils.CountSetBits(game.Black) - utils.CountSetBits(game.White)
}

func (game *GameState) FindLegalMoves(directions *[]func(uint64)uint64) uint64 {
	var currColor uint64
	var oppColor uint64
	var legalMoves uint64 = 0

	if game.IsBlack {
		currColor = game.Black
		oppColor = game.White
	} else {
		currColor = game.White
		oppColor = game.Black
	}

	emptySquares := ^(currColor|oppColor)
	
	var currPos uint64 = 1
	for currPos != leftBitSet {
		if currPos & emptySquares != 0 && checkRecolor(currPos, currColor, oppColor, directions, true) != 0 {
			legalMoves |= currPos
		}
		currPos <<= 1
	}

	if legalMoves == 0 {
		currPos = 1
		for currPos != leftBitSet {
			if currPos & emptySquares != 0 && checkRecolor(currPos, oppColor, currColor, directions, true) != 0 {
				return gameContinues
			}
			currPos <<= 1
		}
		gameResult := game.EvaluateGameState()
		if gameResult > 0 {
			return gameOverBlackWon
		} else if gameResult < 0 {
			return gameOverWhiteWon
		} else {
			return gameOverDraw
		}
	} else {
		return legalMoves
	}
}

func (game *GameState) ApplyMove(move uint64, directions *[]func(uint64)uint64) {
	var currColor *uint64
	var oppColor *uint64
	
	if game.IsBlack {
		currColor = &game.Black
		oppColor = &game.White
	} else {
		currColor = &game.White
		oppColor = &game.Black
	}

	game.LastMove = move
	*currColor |= move
	allRecolorBits := checkRecolor(move, *currColor, *oppColor, directions, false)
	*currColor |= allRecolorBits
	*oppColor ^= allRecolorBits
	game.LastRecolors = allRecolorBits
	game.IsBlack = !game.IsBlack
}

func checkRecolor(move uint64, currColor uint64, oppColor uint64, directions *[]func(uint64)uint64, justCheck bool) uint64 {
	var recolorBit uint64
	var allRecolorBits uint64
	var currRecolorBits uint64

	for _, shift := range *directions {
		recolorBit = shift(move)
		currRecolorBits = 0

		for oppColor & recolorBit != 0 {
			currRecolorBits |= recolorBit
			recolorBit = shift(recolorBit)
		}

		if currRecolorBits != 0 && currColor & recolorBit != 0 {
			allRecolorBits |= currRecolorBits
			if justCheck {
				return allRecolorBits
			}
		}
	}

	return allRecolorBits
}