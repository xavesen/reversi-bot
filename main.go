package main

import (
	"fmt"
	"strconv"
)

const (
	aInANCII = 97
	green = "\033[32m"
	yellow = "\033[33m"
	blue = "\033[34m"
	reset = "\033[0m"
	notAFile uint64 = 0xfefefefefefefefe
	notHFile uint64 = 0x7f7f7f7f7f7f7f7f
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

func main() {
	game := GameState{
		White: 68853694464,
		Black: 34628173824,
		IsBlack: true,
	}
	directions := []func(uint64)uint64{shiftN, shiftNe, shiftE, shiftSe, shiftS, shiftSw, shiftW, shiftNw}
	localGame(&directions, &game)
}

func localGame(directions *[]func(uint64)uint64, game *GameState) {
	printBoard(game, false)
	gameOver := false

	for !gameOver {
		switch result := findLegalMoves(directions, game); result {
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
		changeBoard(move, directions, game)
		printBoard(game, true)
	}
}

func printBoard(game *GameState, withMoveHighlight bool) {
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

func findFirstSetBit(n uint64) uint64 {
	var currBit uint64 = 1
	for {
		if currBit & n != 0 {
			return currBit
		}
		currBit <<= 1
	}
}

func findLegalMoves(directions *[]func(uint64)uint64, game *GameState) uint64 {
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
		if currPos & emptySquares != 0 && checkRecolor(currPos, currColor, oppColor, directions) != 0 {
			legalMoves |= currPos
		}
		currPos <<= 1
	}

	if legalMoves == 0 {
		currPos = 1
		for currPos != leftBitSet {
			if currPos & emptySquares != 0 && checkRecolor(currPos, oppColor, currColor, directions) != 0 {
				return gameContinues
			}
			currPos <<= 1
		}
		gameResult := evaluateGameState(game)
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

func evaluateGameState(game *GameState) int {
	return countSetBits(game.Black) - countSetBits(game.White)
}

func minimax(depth int, isMaximizing bool, alpha int, beta int, directions *[]func(uint64)uint64, game *GameState) (int, uint64) {
	result := findLegalMoves(directions, game)
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
		return evaluateGameState(game), 0
	}

	bestScore := -100000
	bestOrigScore := 0
	localAlpha := alpha
	localBeta := beta
	var bestMove uint64 = 0
	var score int

	for result != 0 {
		move := findFirstSetBit(result)
		result &= ^move
		newGame := *game
		changeBoard(move, directions, &newGame)
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

func algToBit(alg string) uint64 {
	var result uint64 = 1
	file := int(alg[0]) - aInANCII
	rank, _ := strconv.Atoi(string(alg[1]))

	bitsToMove := (rank-1) * 8 + file
	return result << uint64(bitsToMove)
}

func bitToAlg(move uint64) string {
	var curr uint64 = 1
	ind := 0

	for {
		if curr == move {
			break
		}
		ind += 1
		curr <<= 1
	}

	fileI := ind % 8
	rankI := ind / 8

	rank := strconv.Itoa(int(rankI) + 1)
	file :=  string(rune(aInANCII + fileI))

	return file+rank
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

func changeBoard(move uint64, directions *[]func(uint64)uint64, game *GameState) {
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
	allRecolorBits := checkRecolor(move, *currColor, *oppColor, directions)
	*currColor |= allRecolorBits
	*oppColor ^= allRecolorBits
	game.LastRecolors = allRecolorBits
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