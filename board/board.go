package board

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Board [10][10]int

func NewBoard() *Board {
	return &Board{}
}

func (board *Board) PlaceShips() {
	pieces := [5]int{2, 3, 3, 4, 5}

	for _, piece := range pieces {
		for {
			source := rand.NewSource(time.Now().UnixNano())
			newRand := rand.New(source)

			x := newRand.Intn(10)
			y := newRand.Intn(10)

			direction, err := checkLocation(x, y, piece, board)
			if err == nil {
				placePiece(x, y, direction, piece, board)
				break
			}
		}
	}
}

func placePiece(x, y, direction, pieceSize int, board *Board) {
	// directions => 0: left, 1: up, 2: right, 3: down
	board[x][y] = 2

	switch direction {
	case 0:
		for i := 1; i < pieceSize; i++ {
			board[x][y-i] = 2
		}
		break
	case 1:
		for i := 1; i < pieceSize; i++ {
			board[x-i][y] = 2
		}
		break
	case 2:
		for i := 1; i < pieceSize; i++ {
			board[x][y+i] = 2
		}
		break
	case 3:
		for i := 1; i < pieceSize; i++ {
			board[x+i][y] = 2
		}
		break
	}
}

func checkLocation(x, y, pieceSize int, board *Board) (int, error) {
	// directions => 0: left, 1: up, 2: right, 3: down
	directions := [4]int{0, 0, 0, 0}

	if board[x][y] != 0 {
		return 100, errors.New("bad location")
	}

	// check left
	for i := 0; i < pieceSize; i++ {
		if y-i < 0 || board[x][y-i] != 0 {
			directions[0] = -1
		}
	}

	// check up
	for i := 0; i < pieceSize; i++ {
		if x-i < 0 || board[x-i][y] != 0 {
			directions[1] = -1
		}
	}

	// check right
	for i := 0; i < pieceSize; i++ {
		if y+i > 9 || board[x][y+i] != 0 {
			directions[2] = -1
		}
	}

	// check down
	for i := 0; i < pieceSize; i++ {
		if x+i > 9 || board[x+i][y] != 0 {
			directions[3] = -1
		}
	}

	var j int
	// choose random valid direction
	for {
		source := rand.NewSource(time.Now().UnixNano())
		newRand := rand.New(source)
		idx := newRand.Intn(4)

		if directions[idx] != -1 {
			j = idx
			break
		}
	}

	return j, nil
}

func (board *Board) RegisterShot(x, y int) {
	if board[x][y] < 4 && board[x][y] != 1 {
		board[x][y] += 1
	}
}

func printBoard(board Board) {
	fmt.Println("     0 1 2 3 4 5 6 7 8 9")
	fmt.Println("--+---------------------")

	for i, v := range board {
		fmt.Printf("%d | ", i)
		for _, c := range v {
			if c == 0 {
				fmt.Print(" .")
			}
			if c == 2 {
				fmt.Print(" *")
			}
			if c == 1 {
				fmt.Print(" o")
			}
			if c == 3 {
				fmt.Print(" x")
			}
		}
		fmt.Print("\n")
	}
	fmt.Print("\n")
}

func validateCoords(coords string) (int, int, error) {
	if len(coords) != 2 {
		return 0, 0, errors.New("invalid coordinates")
	}

	x, err1 := strconv.Atoi(string(coords[0]))
	if err1 != nil {
		return 0, 0, err1
	}

	y, err2 := strconv.Atoi(string(coords[1]))
	if err2 != nil {
		return 0, 0, err2
	}

	return x, y, nil
}
