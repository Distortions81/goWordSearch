package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"strings"
)

const (
	maxSize = 128

	defaultSize     = 8
	minLenDefault   = 2
	maxLenDefault   = 64
	defaultMaxDepth = 1000

	charList = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numChars = len(charList)
)

var (
	boardSize XY
	board     [maxSize][maxSize]SPOT
	wordList  []wordData
	minLength = minLenDefault
	maxLength = maxLenDefault
	maxDepth  = defaultMaxDepth
)

func main() {

	var sqSize, xSize, ySize, maxWordLen, minWordLen, maxSearchDepth *int
	sqSize = flag.Int("squareSize", defaultSize, "set board x and y")
	xSize = flag.Int("xSize", defaultSize, "set board width")
	ySize = flag.Int("ySize", defaultSize, "set board height")
	maxWordLen = flag.Int("maxWordLen", maxLenDefault, "max number of letters for words")
	minWordLen = flag.Int("minWordLen", minLenDefault, "min number of letters for words")
	maxSearchDepth = flag.Int("maxSearchDepth", maxDepth, "(advanced) max search depth when constructing the board (affects speed)")

	flag.Parse()
	if *sqSize != defaultSize {
		xSize = sqSize
		ySize = sqSize
	}

	//Init
	boardSize = XY{X: *xSize, Y: *ySize}

	diagsize := int(math.Ceil(float64(boardSize.X+boardSize.Y) / 2.0))
	if maxLenDefault > diagsize {
		maxLength = diagsize
		fmt.Printf("Limiting word size to %v (board size)\n", maxLength)
	}
	if *maxWordLen != maxLenDefault {
		maxLength = *maxWordLen
	}
	if *minWordLen != minLenDefault {
		minLength = *minWordLen
	}
	if *maxSearchDepth != defaultMaxDepth {
		maxDepth = *maxSearchDepth
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
