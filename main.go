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

	printBoard(board)
	changeBoard("f5", "x", &board)
	changeBoard("f6", "o", &board)
	changeBoard("e6", "x", &board)
	changeBoard("f4", "o", &board)
	changeBoard("e3", "x", &board)
	changeBoard("d2", "o", &board)
	changeBoard("g7", "x", &board)
	changeBoard("h8", "o", &board)
	changeBoard("d3", "x", &board)
	changeBoard("c2", "o", &board)
	changeBoard("g6", "x", &board)
	changeBoard("c6", "o", &board)
	changeBoard("f2", "x", &board)
	changeBoard("h7", "o", &board)
	changeBoard("c1", "x", &board)
	changeBoard("d7", "o", &board)
	changeBoard("h6", "x", &board)
	changeBoard("d1", "o", &board)
	changeBoard("b1", "x", &board)
	changeBoard("a1", "o", &board)
	changeBoard("g8", "x", &board)
	changeBoard("b2", "o", &board)
	changeBoard("c8", "x", &board)
	changeBoard("h5", "o", &board)
	changeBoard("b5", "x", &board)
	changeBoard("e2", "o", &board)
	changeBoard("a2", "x", &board)
	changeBoard("g1", "o", &board)
	changeBoard("g2", "x", &board)
	changeBoard("a3", "o", &board)
	changeBoard("g4", "x", &board)
	changeBoard("h2", "o", &board)
	changeBoard("f1", "x", &board)
	changeBoard("f8", "o", &board)
	changeBoard("h1", "x", &board)
	changeBoard("b7", "o", &board)
	changeBoard("a8", "x", &board)
	changeBoard("g5", "o", &board)
	changeBoard("h3", "x", &board)
	changeBoard("e1", "o", &board)
	changeBoard("h4", "x", &board)
	changeBoard("c5", "o", &board)
	changeBoard("c4", "x", &board)
	changeBoard("a5", "o", &board)
	changeBoard("a6", "x", &board)
	changeBoard("a7", "o", &board)
	changeBoard("f3", "x", &board)
	changeBoard("f7", "o", &board)
	changeBoard("d6", "x", &board)
	changeBoard("b3", "o", &board)
	changeBoard("a4", "x", &board)
	changeBoard("d8", "o", &board)
	changeBoard("c3", "x", &board)
	changeBoard("g3", "o", &board)
	changeBoard("e8", "x", &board)
	changeBoard("b8", "o", &board)
	changeBoard("c7", "x", &board)
	changeBoard("e7", "o", &board)
	fmt.Println("e7")
	printBoard(board)
	changeBoard("b6", "o", &board)
	fmt.Println("b6")
	printBoard(board)
	changeBoard("b4", "o", &board)
	fmt.Println("b4")
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
	

	// up
	if i != 0 {
		for k := i-1; k >= 0; k-- {
			recolor, moveOn := checkSquare(k, j, color, board)
			if moveOn {
				continue
			} else if recolor {
				// do recolor
				fmt.Println("recolor up")
				recolorVertical(k, i, j, color, board)
				break
			} else {
				break
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
				break
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
				break
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
				break
			} else {
				break
			}
		}
	}

	// diagonal left up
	if i !=0 && j != 0 {
		l := j-1
		for k := i-1; k >= 0 && l >= 0 ; k-- {
			recolor, moveOn := checkSquare(k, l, color, board)
			if moveOn {
				l--
				continue
			} else if recolor {
				// do recolor
				fmt.Println("recolor left up")
				recolorLeftDiagonal(k, l, i, j, color, board)
				break
			} else {
				break
			}
		}
	}

	// diagonal right down
	if i != 7 && j != 7 {
		l := j+1
		for k := i+1; k <= 7 && l <= 7;  k++ {
			recolor, moveOn := checkSquare(k, l, color, board)
			if moveOn {
				l++
				continue
			} else if recolor {
				// do recolor
				fmt.Println("recolor right down")
				recolorLeftDiagonal(i, j, k, l, color, board)
				break
			} else {
				break
			}
		}
	}

	// diagonal right up
	if i != 0 && j != 7 {
		l := j+1
		for k := i-1; k >=0 && l <=7 ; k-- {
			recolor, moveOn := checkSquare(k, l, color, board)
			if moveOn {
				l++
				continue
			} else if recolor {
				// do recolor
				fmt.Println("recolor right up")
				recolorRightDiagonal(k, l, i, j, color, board)
				break
			} else {
				break
			}
		}
	}

	// diagonal left down
	if i != 7 && j != 0 {
		l := j-1
		for k := i+1; k <= 7 && l >= 0; k++ {
			recolor, moveOn := checkSquare(k, l, color, board)
			if moveOn {
				l--
				continue
			} else if recolor {
				// do recolor
				fmt.Println("recolor left down")
				recolorRightDiagonal(i, j, k, l, color, board)
				break
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

 func recolorLeftDiagonal(fromI int, fromJ int, toI int, toJ int, color string, board *[][]string) {
	l := fromJ + 1
	for k := fromI+1; k < toI; k++ {
		if l >= toJ {
			break
		}
		(*board)[k][l] = color
		l++
	}
 }

 func recolorRightDiagonal(fromI int, fromJ int, toI int, toJ int, color string, board *[][]string) {
	l := fromJ - 1
	for k := fromI+1; k < toI; k++ {
		if l <= toJ {
			break
		}
		(*board)[k][l] = color
		l--
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