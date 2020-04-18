package objects

import (
	"fmt"
)

type Board struct {
	Row int
	Col int
}

func (b Board) MaximumMove() int{
	return b.Row * b.Col
}

func (b Board) GenerateBoardArray() [][]int {
	array2d := make([][]int, b.Row)
	for i := range array2d {
		array2d[i] = make([]int, b.Col)
	}
	return array2d
}

func (b Board) GenerateBoard(boardArray [][]int, player_1_mark string, player_2_mark string) {

	for i := range boardArray {
		for j := range boardArray[i] {
			switch boardArray[i][j] {
			case 1: fmt.Print(" " + player_1_mark + " ")
			case 2: fmt.Print(" " + player_2_mark + " ")
			default: fmt.Print(" . ")
			}
		}
		fmt.Println()
	}
}

