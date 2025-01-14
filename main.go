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
		{".", ".", ".", "x", ".", ".", ".", "."},
		{".", ".", ".", "o", "o", ".", ".", "."},
		{"x", "o", "o", "o", "o", "o", "o", "."},
		{"x", "o", "o", ".", "o", "x", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", "."},
	}
	printBoard(board)
	changeBoard("d6", "x", &board)
	printBoard(board)
 }

 func printBoard(board [][]string) {
	for _, line := range board {
		for _, square := range line {
			fmt.Print(square)
			fmt.Print(" ")
		}
		fmt.Println()
	}
 }

 func changeBoard(move string, color string, board *[][]string) {
	indexes := algToInd(move)
	i := indexes[0]
	j := indexes[1]
	(*board)[i][j] = color
	

	// up + diagonal
	if i != 0 {
		for k := i-1; k >= 0; k-- {
			recolor, moveOn := checkSquare(k, j, color, board)
			if moveOn {
				continue
			} else if recolor {
				// do recolor
				fmt.Println("recolor up")
				recolorVertical(k, i, j, color, board)
			} else {
				break
			}
		}

		// left up
		if j != 0 {
			l := j - 1
			for k := i-1; k >= 0; k-- {
				recolor, moveOn := checkSquare(k, l, color, board)
				if moveOn {
					continue
				} else if recolor {
					// do recolor
					fmt.Println("recolor left up")
					recolorVertical(k, i, j, color, board)
				} else {
					break
				}
			}
		}
	}
	
	// down
	if i != 7 {
		for k := i+1; k <= 7; k++ {
			recolor, moveOn := checkSquare(k, j, color, board)
			if moveOn {
				continue
			} else if recolor {
				// do recolor
				fmt.Println("recolor2")
				recolorVertical(i, k, j, color, board)
			} else {
				break
			}
		}
	}

	// left
	if j != 0 {
		for l := j-1; l >= 0; l-- {
			recolor, moveOn := checkSquare(i, l, color, board)
			if moveOn {
				continue
			} else if recolor {
				// do recolor
				fmt.Println("recolor4")
				recolorHorizontal(l, j, i, color, board)
			} else {
				break
			}
		}
	}

	// right
	if j != 7 {
		for l := j+1; l <= 7; l++ {
			recolor, moveOn := checkSquare(i, l, color, board)
			if moveOn {
				continue
			} else if recolor {
				// do recolor
				fmt.Println("recolor3")
				recolorHorizontal(j, l, i, color, board)
			} else {
				break
			}
		}
	}
 }

 func recolorVertical(fromI int, toI int, j int, color string, board *[][]string) {
	for k := fromI+1 ; k < toI; k++ {
		(*board)[k][j] = color
	}
 }

 func recolorHorizontal(fromJ int, toJ int, i int, color string, board *[][]string) {
	for k := fromJ+1 ; k < toJ; k++ {
		(*board)[i][k] = color
	}
 }

 func checkSquare(ind1 int, ind2 int, color string, board *[][]string) (bool, bool) {
	squareColor := (*board)[ind1][ind2]
	recolor := false
	moveOn := false

	if squareColor == "." {
		return recolor, moveOn
	}

	if squareColor == color {
		recolor = true
		return recolor, moveOn
	} else {
		moveOn = true
		return recolor, moveOn
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