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
	//makeGrid()

	placeWord(DIR_UP, "hello")
	placeWord(DIR_DOWN, "there")
	placeWord(DIR_RIGHT, "everyone")
	placeWord(DIR_DOWN, "this")
	placeWord(DIR_DOWN, "is")
	placeWord(DIR_DOWN, "a")
	placeWord(DIR_DOWN, "test")
	printGrid()
}

func clearGrid() {
	for y := 0; y < maxSize; y++ {
		for x := 0; x < maxSize; x++ {
			board[x][y] = ' '
		}
	}
}

func makeGrid() {
	for y := 0; y < maxSize; y++ {
		for x := 0; x < maxSize; x++ {
			randNum := rand.Intn(numChars)
			board[x][y] = rune(charList[randNum])
		}
	}
}

func printSep() {
	for x := 0; x < int(boardSize.X*2)+2; x++ {
		fmt.Print("-")
	}
	fmt.Println("")
}

func printGrid() {
	printSep()
	for y := 0; y < int(boardSize.Y); y++ {
		line := ""

		for x := 0; x < int(boardSize.X); x++ {
			line = line + string(board[x][y]) + " "
		}
		fmt.Println("|" + line + "|")
	}
	printSep()
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
	for i, c := range placeWord {
		newPos := XY{X: randPos.X + POS(i), Y: randPos.Y}
		board[newPos.X][newPos.Y] = c
		newSpot := SPOT{Rune: c, Pos: newPos}
		Spots = append(Spots, newSpot)
	}
	newWord.Spot = Spots
	wordList = append(wordList, newWord)

	return nil
}
