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
	changeBoard("e6", "x", &board)
	changeBoard("f4", "o", &board)
	changeBoard("g3", "x", &board)
	changeBoard("c6", "o", &board)
	legalMoves := findLegalMoves(directions, "x", &board)
	for _, move := range legalMoves {
		fmt.Println(IndToAlg(move[0], move[1]))
	}
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
			//fmt.Printf("i=%d, j=%d, row=%v, square=%s, square==. - %t, legal=%t\n", i, j, row, square, square==".", isLegalMove(directions, i, j, color, board))
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