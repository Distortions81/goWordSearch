package main

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

type Alphabetic []wordData

func (list Alphabetic) Len() int { return len(list) }

func (list Alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func (list Alphabetic) Less(i, j int) bool {
	var si string = list[i].Word
	var sj string = list[j].Word
	return si < sj
}

func (a XY) addXY(b XY) XY {
	return XY{X: a.X + b.X, Y: a.Y + b.Y}
}

func (a XY) multXY(b XY) XY {
	return XY{X: a.X * b.X, Y: a.Y * b.Y}
}

func (pos XY) inBounds() bool {
	if pos.Y >= boardSize.Y || pos.Y < 0 {
		return false
	}
	if pos.X >= boardSize.X || pos.X < 0 {
		return false
	}

	return true
}

func initGrid() {
	wordList = []wordData{}

	for y := 0; y < maxSize; y++ {
		for x := 0; x < maxSize; x++ {
			board[x][y] = SPOT{Used: false}
		}
	}
}

func makeGrid() {
	initGrid()
	for y := 0; y < int(boardSize.X); y++ {
		for x := 0; x < int(boardSize.Y); x++ {
			randNum := rand.Intn(numChars)
			board[x][y] = SPOT{Rune: rune(charList[randNum]), Used: false}
		}
	}
}

func printGrid() {
	for y := 0; y < int(boardSize.Y); y++ {
		line := ""

		for x := 0; x < int(boardSize.X); x++ {
			line = line + string(board[x][y].Rune) + " "
		}
		fmt.Println(" " + line)
	}

	fmt.Println("")
	fmt.Printf("%v words to be found.\n", len(wordList))

	sort.Sort(Alphabetic(wordList))
	for w, word := range wordList {
		if w > 0 {
			fmt.Print(", ")
		}
		fmt.Printf("%v: (%v)", strings.ToLower(word.Word), dirName[word.Dir])
	}
	fmt.Println()

}
