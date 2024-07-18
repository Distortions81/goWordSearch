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
	charList    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numChars    = len(charList)
)

type DIR uint8
type POS uint8

type XY struct {
	X, Y POS
}

type SPOT struct {
	Rune rune
	Pos  XY
}

type wordData struct {
	Word  string
	Dir   DIR
	Spot  []SPOT
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

	clearGrid()
	makeGrid()
	printGrid()

	placeWord(DIR_UP, "HORK")
	placeWord(DIR_DOWN, "BORK")
	placeWord(DIR_RIGHT, "DORK")
}

func clearGrid() {
	for x := 0; x < maxSize; x++ {
		for y := 0; y < maxSize; y++ {
			board[x][y] = ' '
		}
	}
}

func makeGrid() {
	for x := 0; x < maxSize; x++ {
		for y := 0; y < maxSize; y++ {
			randNum := rand.Intn(numChars)
			board[x][y] = rune(charList[randNum])
		}
	}
}

func printGrid() {
	for x := 0; x < int(boardSize.X); x++ {
		line := ""

		for y := 0; y < int(boardSize.Y); y++ {
			line = line + string(board[x][y]) + " "
		}
		fmt.Println(line)
	}
}

func placeWord(dir DIR, placeWord string) error {
	wLen := len(placeWord)
	placeWord = strings.ToUpper(placeWord)

	if wLen > int(boardSize.X)+int(boardSize.Y) {
		return fmt.Errorf("word too large for the game board size (%v,%v): %v", boardSize.X, boardSize.Y, placeWord)
	}

	randPos := XY{X: POS(rand.Intn(int(boardSize.X))), Y: POS(rand.Intn(int(boardSize.Y)))}

	for _, word := range wordList {
		for _, spot := range word.Spot {
			if spot.Pos == randPos {
				fmt.Printf("Collsion from word: %v with word %v at %v,%v.\n", word.Word, placeWord, randPos.X, randPos.Y)
				break
			}
		}
	}

	newWord := wordData{Word: placeWord, Dir: dir}
	Spots := []SPOT{}
	for _, c := range placeWord {
		board[randPos.X][randPos.Y] = c
		newSpot := SPOT{Rune: c, Pos: randPos}
		Spots = append(Spots, newSpot)
	}
	newWord.Spot = Spots
	wordList = append(wordList, newWord)

	return nil
}
