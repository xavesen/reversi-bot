package main

import (
	"fmt"
	"strconv"
)

const (
	//green = "\033[32m"
	//yellow = "\033[33m"
	blue = "\033[34m"
	reset = "\033[0m"
)

type GameState struct {
	White uint64
	Black uint64
}

func main() {
	game := GameState{
		White: 62976,
		Black: 246,
	}
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
			result += "\n" + blue + strconv.Itoa(i) + reset + " "
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

	fmt.Println(result)
}