package main

import (
	"flag"
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"sync/atomic"

	"github.com/remeh/sizedwaitgroup"
)

const (
	maxSize = 128

	defaultSize     = 8
	minLenDefault   = 2
	maxLenDefault   = 64
	bestOfDefault   = 1024
	defaultMaxDepth = 10000

	charList = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	numChars = len(charList)
)

var (
	boardSize XY
	minLength = minLenDefault
	maxLength = maxLenDefault
	maxDepth  = defaultMaxDepth
	bestOfAtt = bestOfDefault
)

func main() {

	//Flags
	var sqSize, xSize, ySize, maxWordLen, minWordLen, maxSearchDepth *int
	sqSize = flag.Int("squareSize", defaultSize, "set board x and y")
	xSize = flag.Int("xSize", defaultSize, "set board width")
	ySize = flag.Int("ySize", defaultSize, "set board height")
	maxWordLen = flag.Int("maxWordLen", maxLenDefault, "max number of letters for words")
	minWordLen = flag.Int("minWordLen", minLenDefault, "min number of letters for words")
	maxSearchDepth = flag.Int("maxSearchDepth", defaultMaxDepth, "(advanced) max search depth when constructing the board (affects speed)")
	flag.Parse()

	/* Auto default max attempts */
	if *sqSize != defaultSize {
		boardSize = XY{X: *sqSize, Y: *sqSize}
	} else {
		boardSize = XY{X: *xSize, Y: *ySize}
	}
	bestOfAtt = boardSize.X * boardSize.Y * DIR_ANY * 2

	if *maxWordLen != maxLenDefault {
		maxLength = *maxWordLen
	}

	/* Calc max word size */
	diagsize := int(math.Ceil(float64(boardSize.X+boardSize.Y) / 2.0))
	if maxLength > diagsize {
		maxLength = diagsize
		fmt.Printf("Limiting word size to %v (board size)\n", maxLength)
	}

	if *minWordLen != minLenDefault {
		minLength = *minWordLen
	}
	if *maxSearchDepth != defaultMaxDepth {
		maxDepth = *maxSearchDepth
	}

	/* Don't overflow array */
	if boardSize.X >= maxSize {
		boardSize.X = maxSize - 1
	}
	if boardSize.Y >= maxSize {
		boardSize.Y = maxSize - 1
	}

	//fixDict()
	limitDict()

	numThreads := runtime.NumCPU()
	wg := sizedwaitgroup.New(numThreads)

	var topScore atomic.Int32
	for c := 0; c < bestOfAtt; c++ {
		wg.Add()
		go func(c int) {
			local := localWork{}
			local.shuffleDict()
			local.makeBoard()
			for i := 0; i < newDictLen; i++ {
				randWord := local.dict[i]
				found := false
				for _, word := range local.words {
					if randWord == word.Word {
						found = true
						//fmt.Printf("Word already used: %v\n", word.Word)
						break
					}
				}
				if !found {
					for depth := 0; depth < maxDepth; depth++ {
						if local.placeWord(DIR_ANY, randWord) {
							break
						}
					}
				}
			}
			score := len(local.words)

			if score > int(topScore.Load()) {
				topScore.Store(int32(score))
				fmt.Printf("\nAttempt %v of %v.\n", c, bestOfAtt)
				local.printBoard()
			}

			wg.Done()
		}(c)
	}
	wg.Wait()
}

func (local *localWork) placeWord(inDir int, pWord string) bool {

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

	randPos := local.randPos(dir, pWord)
	newPos := randPos

	Spots := []SPOT{}
	for i, c := range pWord {

		/*
			//Sanity check
			if !newPos.inBounds() {
				fmt.Printf("\nWent out of bounds... %v,%v: %v at %v in dir: %v\n", randPos.X, randPos.Y, pWord, i, dirName[dir])
				return false
			}
		*/

		if local.board[newPos.X][newPos.Y].Used {
			if local.board[newPos.X][newPos.Y].Rune != c {
				return false
			}
		}
		newSpot := SPOT{Rune: c, Pos: newPos}
		Spots = append(Spots, newSpot)
		newPos = randPos.addXY(dirMap[dir].multXY(XY{X: i + 1, Y: i + 1}))
	}

	newWord := wordData{Word: pWord, Dir: dir}
	newWord.Spot = Spots
	local.words = append(local.words, newWord)
	for _, c := range newWord.Spot {
		local.board[c.Pos.X][c.Pos.Y] = SPOT{Rune: c.Rune, Used: true}
	}

	return true
}
