package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const (
	maxSize = 128

	defaultSize   = 8
	minLenDefault = 2
	maxLenDefault = 50
	maxDepth      = 10000
	charList      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numChars      = len(charList)
)

var (
	boardSize XY
	board     [maxSize][maxSize]SPOT
	wordList  []wordData
	minLength int = minLenDefault
	maxLength int = maxLenDefault
)

func main() {

	//Init
	boardSize = XY{X: defaultSize, Y: defaultSize}

	if maxLenDefault > (boardSize.X + boardSize.Y) {
		maxLength = boardSize.X + boardSize.Y
		fmt.Printf("Limiting word size to %v (board size)\n", maxLength)
	}

	fixDict()
	limitDict()
	initGrid()

	makeGrid()

	found := false
	for c := 0; c < maxDepth; c++ {
		for i := 0; i < newDictLen; i++ {
			randWord := newDict[i]
			for _, word := range wordList {
				if strings.EqualFold(randWord, word.Word) {
					found = true
					//fmt.Printf("Word already present: %v\n", randWord)
					break
				}
			}
			if !found {
				randWord = strings.ToUpper(randWord)
				placeWord(DIR_ANY, randWord, 0)
			}
		}
	}

	printGrid()
}

func placeWord(inDir int, pWord string, d int) error {
	if d > maxDepth {
		return nil
	}

	//fmt.Printf("inDir: %v\n", dirName[inDir])
	dir := inDir
	if inDir == DIR_ANY {
		dir = rand.Intn(DIR_ANY)
	}
	if inDir == DIR_NORMAL {
		num := rand.Intn(3)

		switch num {
		case 0:
			dir = DIR_DOWN
		case 1:
			dir = DIR_DOWN_RIGHT
		case 2:
			dir = DIR_RIGHT
		}
	}
	//fmt.Printf("outDir: %v\n", dirName[dir])

	randPos := XY{X: rand.Intn(int(boardSize.X)), Y: rand.Intn(int(boardSize.Y))}
	for i := range pWord {
		newPos := randPos.addXY(dirMap[dir].multXY(XY{X: i + 1, Y: i + 1}))
		if !newPos.inBounds() {
			//fmt.Printf("Word %v went off the edge, dir: %v\n", pWord, dirName[dir])
			placeWord(inDir, pWord, d+1)
			return nil
		}
	}

	newWord := wordData{Word: pWord, Dir: dir}
	Spots := []SPOT{}
	for i, c := range pWord {
		newPos := randPos.addXY(dirMap[dir].multXY(XY{X: i + 1, Y: i + 1}))
		if !newPos.inBounds() {
			continue
		}
		if board[newPos.X][newPos.Y].Used {

			if board[newPos.X][newPos.Y].Rune != c {
				placeWord(inDir, pWord, d+1)
				return nil
			} /* else {
				fmt.Printf("Crossed word: %v at %v,%v. Dir: %v, Letter: %v\n", pWord, newPos.X, newPos.Y, dirName[dir], string(c))
			} */
		}
		newSpot := SPOT{Rune: c, Pos: newPos}
		Spots = append(Spots, newSpot)
	}
	newWord.Spot = Spots
	wordList = append(wordList, newWord)
	for _, c := range newWord.Spot {
		board[c.Pos.X][c.Pos.Y] = SPOT{Rune: c.Rune, Used: true}
	}

	return nil
}
