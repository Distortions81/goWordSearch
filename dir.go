package main

const (
	DIR_RIGHT = iota
	DIR_LEFT
	DIR_UP
	DIR_DOWN
	DIR_UP_LEFT
	DIR_UP_RIGHT
	DIR_DOWN_LEFT
	DIR_DOWN_RIGHT

	DIR_ANY
	DIR_NORMAL
)

var dirMap []XY = []XY{
	{X: 1, Y: 0},   //right
	{X: -1, Y: 0},  //left
	{X: 0, Y: -1},  //up
	{X: 0, Y: 1},   //down
	{X: -1, Y: -1}, //up left
	{X: 1, Y: -1},  //up right
	{X: -1, Y: 1},  //down left
	{X: 1, Y: 1},   //down right

	{X: 0, Y: 0}, //any
	{X: 0, Y: 0}, //normal
}

var dirName []string = []string{
	"Right",
	"Left",
	"Up",
	"Down",
	"Up & left",
	"Up & right",
	"Down & left",
	"Down & right",

	"Any",
	"Normal",
}

var dirNameShort []string = []string{
	"R",
	"L",
	"U",
	"D",
	"U&L",
	"U&R",
	"D&L",
	"D&R",

	"*",
	"+",
}
