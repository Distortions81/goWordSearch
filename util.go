package main

import (
	"fmt"
	"math/rand"
	"os"
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

// Return aXY plus bXY
func (a XY) addXY(b XY) XY {
	return XY{X: a.X + b.X, Y: a.Y + b.Y}
}

// Return aXY multiplied by bXY
func (a XY) multXY(b XY) XY {
	return XY{X: a.X * b.X, Y: a.Y * b.Y}
}

/*
// Position fits on the board
func (pos XY) inBounds() bool {
	if pos.Y >= boardSize.Y || pos.Y < 0 {
		return false
	}
	if pos.X >= boardSize.X || pos.X < 0 {
		return false
	}

	return true
}
*/

// Make game board (random characters)
func (local *localWork) makeBoard() {
	for y := 0; y < int(boardSize.X); y++ {
		for x := 0; x < int(boardSize.Y); x++ {
			randNum := rand.Intn(numChars)
			local.board[x][y] = SPOT{Rune: rune(charList[randNum]), Used: false}
		}
	}
}

// Print the game board
func (local *localWork) printBoard() {
	for y := 0; y < int(boardSize.Y); y++ {
		line := ""

		for x := 0; x < int(boardSize.X); x++ {
			line = line + string(local.board[x][y].Rune) + " "
		}
		fmt.Println(" " + line)
	}

	fmt.Println("")
	fmt.Printf("%v words to be found.\n", len(local.words))

	sort.Sort(Alphabetic(local.words))
	for w, word := range local.words {
		if w > 0 {
			fmt.Print(", ")
		}
		if shortNames {
			fmt.Printf("%v %v", strings.ToLower(word.Word), dirNameShort[word.Dir])
		} else {
			fmt.Printf("%v: %v", strings.ToLower(word.Word), dirName[word.Dir])
		}

		/*
			//Sanity check
			for _, item := range word.Spot {
				if !item.Pos.inBounds() {
					fmt.Printf("Error: %v,%v", item.Pos.X+1, item.Pos.Y+1)
					os.Exit(0)
					return
				}
			}
		*/
	}
	fmt.Println()
}

// Random position, where the word will fit on the board in (dir)
func (local *localWork) randPos(dir int, word string) XY {
	wl := len(word) - 1
	bsx := boardSize.X
	bsy := boardSize.Y

	switch dir {
	case DIR_RIGHT:
		return XY{X: rand.Intn((bsx - wl)), Y: rand.Intn((bsy))}
	case DIR_LEFT:
		return XY{X: wl + rand.Intn(bsx-wl), Y: rand.Intn((bsy))}
	case DIR_UP:
		return XY{X: rand.Intn((bsx)), Y: wl + rand.Intn((bsy - wl))}
	case DIR_DOWN:
		return XY{X: rand.Intn((bsx)), Y: rand.Intn((bsy - wl))}
	case DIR_UP_LEFT:
		return XY{X: wl + rand.Intn((bsx - wl)), Y: wl + rand.Intn((bsy - wl))}
	case DIR_UP_RIGHT:
		return XY{X: rand.Intn((bsx - wl)), Y: wl + rand.Intn((bsy - wl))}
	case DIR_DOWN_LEFT:
		return XY{X: wl + rand.Intn((bsx - wl)), Y: rand.Intn((bsy - wl))}
	case DIR_DOWN_RIGHT:
		return XY{X: rand.Intn((bsx) - wl), Y: rand.Intn((bsy - wl))}
	default:
		fmt.Println("randPos: Invalid direction!")
		os.Exit(1)
	}
	return XY{}
}
