package main

import (
	"fmt"
	"strconv"
)

const (
	aInANCII = 97
	//green = "\033[32m"
	//yellow = "\033[33m"
	blue = "\033[34m"
	reset = "\033[0m"
	notAFile uint64 = 0xfefefefefefefefe
	notHFile uint64 = 0x7f7f7f7f7f7f7f7f
)

type GameState struct {
	White uint64
	Black uint64
	IsBlack bool
	LastMove uint64
	LastRecolors uint64
}

func main() {
	game := GameState{
		White: 68853694464,
		Black: 34628173824,
		IsBlack: true,
	}
	directions := []func(uint64)uint64{shiftN, shiftNe, shiftE, shiftSe, shiftS, shiftSw, shiftW, shiftNw}
	changeBoard(algToBit("f5"), true, &directions, &game)
	printBoard(&game)
}

func printBoard(game *GameState) {
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
		if bitMask & game.Black != 0 {
			result += "x "
		} else if bitMask & game.White != 0 {
			result += "o "
		} else {
			result += ". "
		}
		bitMask <<= 1
		j++
	}

	result += blue + "8\n" + "  a b c d e f g h" + reset

	fmt.Println(result)
}

func countSetBits(n uint64) int {
	count := 0
	curr := n

	for curr != 0 {
		if curr & 1 != 0 {
			count++
		}
		curr = curr >> 1
	}

	return count
}

func findLegalMoves(isBlack bool, directions *[]func(uint64)uint64, game *GameState) uint64 {
	var currColor uint64
	var oppColor uint64
	var legalMoves uint64 = 0

	if isBlack {
		currColor = game.Black
		oppColor = game.White
	} else {
		currColor = game.White
		oppColor = game.Black
	}

	emptySquares := ^(currColor|oppColor)
	
	var currPos uint64 = 1
	for currPos != 0x8000000000000000 {
		if currPos & emptySquares != 0 && checkRecolor(currPos, currColor, oppColor, directions) != 0 {
			legalMoves |= currPos
		}
		currPos <<= 1
	}

	return legalMoves
}

func evaluateGameState(game *GameState) int {
	return countSetBits(game.White) - countSetBits(game.Black)
}

func algToBit(alg string) uint64 {
	var result uint64 = 1
	file := int(alg[0]) - aInANCII
	rank, _ := strconv.Atoi(string(alg[1]))

	bitsToMove := (rank-1) * 8 + file
	return result << uint64(bitsToMove)
}

func checkRecolor(move uint64, currColor uint64, oppColor uint64, directions *[]func(uint64)uint64) uint64 {
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
		}
	}

	return allRecolorBits
}

func changeBoard(move uint64, isBlack bool, directions *[]func(uint64)uint64, game *GameState) {
	var currColor *uint64
	var oppColor *uint64
	
	if isBlack {
		currColor = &game.Black
		oppColor = &game.White
	} else {
		currColor = &game.White
		oppColor = &game.Black
	}

	game.LastMove = move
	*currColor |= move
	allRecolorBits := checkRecolor(move, *currColor, *oppColor, directions)
	*currColor |= allRecolorBits
	*oppColor ^= allRecolorBits
	game.IsBlack = !game.IsBlack
}

func shiftN(bit uint64) uint64 {
	return bit << 8
}

func shiftNe(bit uint64) uint64 {
	return bit << 9 & notAFile
}

func shiftE(bit uint64) uint64 {
	return bit << 1 & notAFile
}

func shiftSe(bit uint64) uint64 {
	return bit >> 7 & notAFile
}

func shiftS(bit uint64) uint64 {
	return bit >> 8
}

func shiftSw(bit uint64) uint64 {
	return bit >> 9 & notHFile
}

func shiftW(bit uint64) uint64 {
	return bit >> 1 & notHFile
}

func shiftNw(bit uint64) uint64 {
	return bit << 7 & notHFile
}