package main

import (
	"fmt"
	"strconv"
	"os"
	"alive/caro/objects"
)

const (
	winPoint = 2
	blockPoint = 2
)

func main () {
	StartGame()
}

func StartGame() {

	var board = objects.Board{Row: 10, Col: 10}
	var player_1 = objects.Player{ID: 1, Mark: "x"}
	var player_2 = objects.Player{ID: 2, Mark: "o"}

	boardArray := board.GenerateBoardArray()
	moveCount := board.MaximumMove()

	board.GenerateBoard(boardArray, player_1.Mark, player_2.Mark)
	for i := 0; i < moveCount; i++ {
		PlayerInput(board, player_1, player_1.Mark, player_2.Mark, boardArray)
		PlayerInput(board, player_2, player_1.Mark, player_2.Mark, boardArray)
	}
	Draw()
	Reset()
}

func Draw(){
	fmt.Println("It's a Draw!")
}

func Reset() {
	var reset string
	fmt.Println("Do you want to try again? Press Y to continue")
	fmt.Scanf("%s", &reset)
	if reset == "Y" {
		StartGame()
	} else {
		os.Exit(0)
	}
}

func PlayerInput(board objects.Board, player objects.Player, player_1_mark string, player_2_mark string, boardArray [][]int) {

	recentRow, recentCol, playerID := RecentMove(player, boardArray)
	board.GenerateBoard(boardArray, player_1_mark, player_2_mark)
	CheckWinCondition(boardArray, recentRow, recentCol, playerID)
}

func RecentMove(player objects.Player, boardArray [][]int) (int, int, int){

	var row int
	var col int

	fmt.Println("Player " + strconv.Itoa(player.ID) + " turn!")
	fmt.Print("Row: ")
	fmt.Scanf("%d", &row)
	if row > len(boardArray) {
		fmt.Println("Selected row must be smaller than the number of rows")
		RecentMove(player, boardArray)
	} else if row < 0 {
		fmt.Println("Selected row cannot be negative")
		RecentMove(player, boardArray)
	}

	fmt.Print("Col: ")
	fmt.Scanf("%d", &col)
	if col > len(boardArray[0]) {
		fmt.Println("Selected column must be smaller than the number of columns")
		RecentMove(player, boardArray)
	} else if row < 0 {
		fmt.Println("Selected column cannot be negative")
		RecentMove(player, boardArray)
	}

	if boardArray[row][col] == 0 {
		boardArray[row][col] = player.ID
	} else {
		fmt.Println("This position is already used")
		RecentMove(player, boardArray)
	}

	return row, col, player.ID
}

func CheckWinCondition(boardArray [][]int, recentRow int, recentCol int, playerID int) {
	CheckRow(boardArray, recentRow, recentCol, playerID)
	CheckCol(boardArray, recentRow, recentCol, playerID)
	//CheckDiag(boardArray, recentRow, recentCol, playerID)
	CheckAntiDiag(boardArray, recentRow, recentCol, playerID)
}

func CheckRow(boardArray [][]int, recentRow int, recentCol int, playerID int){

	currentWinPoint := 1
	currentBlockPoint := 0

	// Check left
	for i := recentCol; i > 0; i-- {
		if boardArray[recentRow][i-1] == playerID {
			currentWinPoint++
		} else if boardArray[recentRow][i-1] != 0 {
			currentBlockPoint++
			break
		}
	}

	// Check right
	for i := recentCol; i < len(boardArray[0])-1; i++ {
		if boardArray[recentRow][i+1] == playerID {
			currentWinPoint++
		} else if boardArray[recentRow][i+1] != 0 {
			currentBlockPoint++
			break
		}
	}

	if currentWinPoint >= winPoint && currentBlockPoint < blockPoint {
		Win(playerID)
	}
}

func CheckCol(boardArray [][]int, recentRow int, recentCol int, playerID int){

	currentWinPoint := 1
	currentBlockPoint := 0

	// Check top
	for i := recentRow; i > 0; i-- {
		if boardArray[i-1][recentCol] == playerID {
			currentWinPoint++
		} else if boardArray[i-1][recentCol] != 0 {
			currentBlockPoint++
			break
		}
	}

	// Check bottom
	for i := recentRow; i < len(boardArray)-1; i++ {
		if boardArray[i+1][recentCol] == playerID {
			currentWinPoint++
		} else if boardArray[i+1][recentCol] != 0 {
			currentBlockPoint++
			break
		}
	}
	
	if currentWinPoint >= winPoint && currentBlockPoint < blockPoint {
		Win(playerID)
	}
}

// func CheckDiag(boardArray [][]int, recentRow int, recentCol int){

// 	currentWinPoint := 1
// 	currentBlockPoint := 0

// 	// Check upper diag
// 	maxUpperDiagMoves := 0
// 	upperIncrementIndex := 1

// 	for {
// 		if recentRow - upperIncrementIndex - maxUpperDiagMoves >= 0 && recentCol - upperIncrementIndex  - maxUpperDiagMoves >= 0 {
// 			maxUpperDiagMoves++
// 		} else {
// 			break
// 		}
// 	}

// 	fmt.Println(maxUpperDiagMoves)
// 	for i := 1; i <= maxUpperDiagMoves; i++ {
// 		if boardArray[recentRow - i][recentCol - i] == 1 {
// 			currentWinPoint++
// 		} else if boardArray[recentRow - i][recentCol - i] == 2 {
// 			currentBlockPoint++
// 		} else {
// 			break
// 		}
		
// 		fmt.Println(currentWinPoint)
// 		fmt.Println(currentBlockPoint)
// 		if currentWinPoint >= winPoint && currentBlockPoint < blockPoint {
// 			return true
// 		}
// 	}

// 	// Check lower diag
// 	maxLowerDiagMoves := 0
// 	lowerIncrementIndex := 1
// 	for {
// 		if recentRow + lowerIncrementIndex + maxLowerDiagMoves < len(boardArray) && recentCol + lowerIncrementIndex + maxLowerDiagMoves < len(boardArray[0]) {
// 			maxLowerDiagMoves++
// 		} else {
// 			break
// 		}
// 	}

// 	fmt.Println("|||||||||")
// 	fmt.Println(maxLowerDiagMoves)
// 	for i := 1; i <= maxLowerDiagMoves; i++ {
// 		if boardArray[recentRow + i][recentCol + i] == 1 {
// 			currentWinPoint++
// 		} else if boardArray[recentRow + i][recentCol + i] == 2 {
// 			currentBlockPoint++
// 		} else {
// 			break
// 		}
		
// 		fmt.Println(currentWinPoint)
// 		fmt.Println(currentBlockPoint)
// 		if currentWinPoint >= winPoint && currentBlockPoint < blockPoint {
// 			return true
// 		}
// 	}
// 	return false
// }

func CheckAntiDiag(boardArray [][]int, recentRow int, recentCol int, playerID int) {
	// currentWinPoint := 0
	// currentBlockPoint := 0
	fmt.Println("GO here")
}

func Win(playerID int) {
	fmt.Println("The Winner is Player " + strconv.Itoa(playerID))
	Reset()
}