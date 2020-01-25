package logic

import (
	"errors"
	"math/rand"
	"time"
)

type row []int
type matrix []row

type Direction int

const (
	LEFT  Direction = 0
	RIGHT Direction = 1
	UP    Direction = 2
	DOWN  Direction = 3
)

type GameState struct {
	Board matrix
	Score int64
}

func NewGameState() *GameState {
	gs := &GameState{
		Board: matrix{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		Score: 0,
	}

	rand.Seed(time.Now().UnixNano())

	pos1 := rand.Intn(16)
	pos2 := pos1
	for pos2 == pos1 {
		pos2 = rand.Intn(16)
	}

	gs.Board[int(float64(pos1)/4)][pos1%4] = 2
	gs.Board[int(float64(pos2)/4)][pos2%4] = 2

	return gs
}

func (gs *GameState) IsValid() (valid bool) {
	for i, w := range gs.Board {
		for j, v := range w {
			if v == 0 {
				return true
			}

			// check above
			if i != 0 && gs.Board[i-1][j] == v {
				return true
			}

			// check below
			if i != len(gs.Board)-1 && gs.Board[i+1][j] == v {
				return true
			}

			// check left
			if j != 0 && gs.Board[i][j-1] == v {
				return true
			}

			// check right
			if j != len(w)-1 && gs.Board[i][j+1] == v {
				return true
			}
		}
	}
	return
}

func (gs *GameState) IsMovePossible() bool {
	return gs.IsValid() && gs.countEmptyTiles() > 0
}

func (gs *GameState) Move(direction Direction) (finished bool, err error) {
	moved := false

	switch direction {
	case LEFT:
		moved = gs.moveLeft()
	case RIGHT:
		moved = gs.moveRight()
	case UP:
		moved = gs.moveUp()
	case DOWN:
		moved = gs.moveDown()
	}

	if !gs.IsMovePossible() {
		return true, nil
	}

	if moved {
		err = gs.addRandomTile()
		if err != nil {
			return false, err
		}
	}

	return false, nil
}

func (gs *GameState) moveLeft() (moved bool) {
	for i, w := range gs.Board {
		lastValue := 0
		emptyTiles := 0
		for j, v := range w {
			if v == 0 {
				emptyTiles++
				continue
			}
			if j == 0 {
				lastValue = v
				continue
			}

			if lastValue == v {
				emptyTiles++
				gs.Board[i][j-emptyTiles] = 2 * v
				gs.Board[i][j] = 0
				lastValue = 0
				gs.Score += 2 * int64(v)
				moved = true
				continue
			}

			if emptyTiles != 0 {
				gs.Board[i][j-emptyTiles] = v
				gs.Board[i][j] = 0
				moved = true
			}
			lastValue = v
		}
	}
	return
}

func (gs *GameState) moveRight() (moved bool) {

	for i, w := range gs.Board {
		lastValue := 0
		emptyTiles := 0
		for j := len(w) - 1; j >= 0; j-- {
			v := w[j]
			if v == 0 {
				emptyTiles++
				continue
			}
			if j == len(w)-1 {
				lastValue = v
				continue
			}

			if lastValue == v {
				emptyTiles++
				gs.Board[i][j+emptyTiles] = 2 * v
				gs.Board[i][j] = 0
				lastValue = 0
				gs.Score += 2 * int64(v)
				moved = true
				continue
			}

			if emptyTiles != 0 {
				gs.Board[i][j+emptyTiles] = v
				gs.Board[i][j] = 0
				moved = true
			}
			lastValue = v
		}
	}
	return
}

func (gs *GameState) moveUp() (moved bool) {
	boardT := transpose(gs.Board)

	for i, w := range boardT {
		lastValue := 0
		emptyTiles := 0
		for j, v := range w {
			if v == 0 {
				emptyTiles++
				continue
			}
			if j == 0 {
				lastValue = v
				continue
			}

			if lastValue == v {
				emptyTiles++
				boardT[i][j-emptyTiles] = 2 * v
				boardT[i][j] = 0
				lastValue = 0
				gs.Score += 2 * int64(v)
				moved = true
				continue
			}

			if emptyTiles != 0 {
				boardT[i][j-emptyTiles] = v
				boardT[i][j] = 0
				moved = true
			}
			lastValue = v
		}
	}
	gs.Board = transpose(boardT)
	return
}

func (gs *GameState) moveDown() (moved bool) {
	boardT := transpose(gs.Board)

	for i, w := range boardT {
		lastValue := 0
		emptyTiles := 0
		for j := len(w) - 1; j >= 0; j-- {
			v := w[j]
			if v == 0 {
				emptyTiles++
				continue
			}
			if j == len(w)-1 {
				lastValue = v
				continue
			}

			if lastValue == v {
				emptyTiles++
				boardT[i][j+emptyTiles] = 2 * v
				boardT[i][j] = 0
				lastValue = 0
				gs.Score += 2 * int64(v)
				moved = true
				continue
			}

			if emptyTiles != 0 {
				boardT[i][j+emptyTiles] = v
				boardT[i][j] = 0
				moved = true
			}
			lastValue = v
		}
	}
	gs.Board = transpose(boardT)
	return
}

func (gs *GameState) addRandomTile() error {
	value := rand.Intn(10)
	if value < 10 {
		value = 2
	} else {
		value = 4
	}

	pos := rand.Intn(gs.countEmptyTiles())

	err := gs.placeTile(value, pos)
	if err != nil {
		return err
	}

	return nil
}

func (gs *GameState) countEmptyTiles() (emptyTiles int) {
	for _, w := range gs.Board {
		for _, v := range w {
			if v == 0 {
				emptyTiles++
			}
		}
	}
	return
}

func (gs *GameState) placeTile(value int, pos int) (err error) {
	count := 0
	for i, w := range gs.Board {
		for j, v := range w {
			if count == pos && v == 0 {
				gs.Board[i][j] = value
				return nil
			}
			if v == 0 {
				count++
			}
		}
	}
	return errors.New("couldn't place tile")
}

func transpose(m matrix) matrix {
	r := make(matrix, len(m[0]))
	for x, _ := range r {
		r[x] = make(row, len(m))
	}
	for y, s := range m {
		for x, e := range s {
			r[x][y] = e
		}
	}
	return r
}
