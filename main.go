package main

import (
	"fmt"
	"image/color"
	"math/rand"
	"strings"
)

const (
	DIR_RIGHT = iota
	DIR_LEFT
	DIR_UP
	DIR_DOWN
	DIR_UP_LEFT
	DIR_UP_RIGHT
	DIR_DOWN_LEFT
	DIR_DOWN_RIGHT
)

const (
	maxSize     = 254
	defaultSize = 20
)

type DIR uint8
type POS uint8

type XY struct {
	X, Y POS
}

type wordData struct {
	Word  string
	Dir   DIR
	Pos   []XY
	Found bool
	Color color.RGBA64
}

var (
	boardSize XY
	board     [maxSize][maxSize]rune
	wordList  []wordData
)

func main() {
	boardSize = XY{X: defaultSize, Y: defaultSize}

	placeWord(DIR_UP, "HORK")
	placeWord(DIR_DOWN, "BORK")
	placeWord(DIR_RIGHT, "DORK")
}

func placeWord(dir DIR, word string) error {
	wLen := len(word)
	word = strings.ToUpper(word)

	if wLen > int(boardSize.X)+int(boardSize.Y) {
		return fmt.Errorf("Word too large for the game board.", word)
	}

	for c := 0; c < 10000000; c++ {
		randPos := XY{X: POS(rand.Intn(int(boardSize.X))), Y: POS(rand.Intn(int(boardSize.Y)))}

		if board[randPos.X][randPos.Y] != 0 {
			continue
		}
	}
}
