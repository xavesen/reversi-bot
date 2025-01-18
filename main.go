package main

import (
	"fmt"
	"strconv"
)

 const aInANCII = 97

 func main() {
	board := [][]string{
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", "o", "x", ".", ".", "."},
		{".", ".", ".", "x", "o", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
	}
	directions := [][]int{
		{0, 1}, // right
		{0, -1}, // left
		{1, 0}, // up
		{-1, 0}, // down
		{1, 1}, // up right
		{1, -1}, // up left
		{-1, 1}, // down right
		{-1, -1}, // down left
	}

	printBoard(board)
	changeBoard("f5", "x", directions, &board)
	printBoard(board)
 }

 func printBoard(board [][]string) {
	fmt.Println("  0 1 2 3 4 5 6 7")
	for i, line := range board {
		fmt.Print(i, " ")
		for _, square := range line {
			fmt.Print(square)
			fmt.Print(" ")
		}
		fmt.Println()
	}
 }

 func findLegalMoves(directions [][]int, color string, board *[][]string) [][]int {
	var legalMoves [][]int

	for i, row := range *board {
		for j, square := range row {
			if square == "." && isLegalMove(directions, i, j, color, board){
				legalMoves = append(legalMoves, []int{i, j})
			}
		}
	}

	return legalMoves
 }

 func isLegalMove(directions [][]int, ind1 int, ind2 int, color string, board *[][]string) bool {
	var iPlus, jPlus, i, j int
	var opponentInBetween bool

	for _, direction := range directions {
		iPlus = direction[0]
		jPlus = direction[1]
		i = ind1 + iPlus
		j = ind2 + jPlus
		opponentInBetween = false
		
		for i >= 0 && i <= 7 && j >= 0 && j <= 7 {
			square := (*board)[i][j]
			if square == "." {
				break
			} else if square != color {
				opponentInBetween = true
				i += iPlus
				j += jPlus
			} else if opponentInBetween {
				return true
			} else {
				break
			}
		}
	}

	return false
 }

 func changeBoard(move string, color string, directions [][]int, board *[][]string) {
	fmt.Println(move)
	var i, j, iPlus, jPlus int
	var opponentInBetween bool
	var square string

	indexes := algToInd(move)
	ind1 := indexes[0]
	ind2 := indexes[1]
	(*board)[ind1][ind2] = color

	for _, direction := range directions {
		iPlus = direction[0]
		jPlus = direction[1]
		i = ind1 + iPlus
		j = ind2 + jPlus
		opponentInBetween = false

		for i >= 0 && i <= 7 && j >= 0 && j <= 7 {
			square = (*board)[i][j]
		
			if square == "." {
				break
			} else if square != color {
				opponentInBetween = true
				i += iPlus
				j += jPlus
			} else if opponentInBetween {
				recolorSquares(indexes, []int{i, j}, direction, color, board)
				break
			} else {
				break
			}
		}
	}
 }

func recolorSquares(from []int, to []int, direction []int, color string, board* [][]string) {
	fmt.Printf("recoloring from %v to %v, dir %v\n", from, to, direction)
	currI := from[0]
	currJ := from[1]

	for currI != to[0] || currJ != to[1] {
		(*board)[currI][currJ] = color
		currI += direction[0]
		currJ += direction[1]
	}
}

 func algToInd(alg string) []int {
	letter := alg[0]
	number := string(alg[1])

	intNumber, _ := strconv.Atoi(number)
	firstInd := intNumber - 1
	secondInd := letter - aInANCII

	return []int{firstInd, int(secondInd)}
 }

 func IndToAlg(ind1 int, ind2 int) string {
	number := strconv.Itoa(ind1 + 1)
	letter := string(rune(aInANCII + ind2))

	return letter + number
 }