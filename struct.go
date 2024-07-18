package main

import "image/color"

type XY struct {
	X, Y int
}

type SPOT struct {
	Rune rune
	Pos  XY
	Used bool
}

type wordData struct {
	Word  string
	Dir   int
	Spot  []SPOT
	Found bool
	Color color.RGBA64
}
