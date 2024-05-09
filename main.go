package main

import (
	"fmt"
)

const(GAME_SIZE = 10)
type cellState int

const (
	black cellState = iota
	white
)

type cell struct {
	state cellState
}

func (c *cell) String() string {
	if c.state == black {
		return "@"
	} else if c.state == white {
		return " "
	}
	return "ERROR"
}

type gameGrid [GAME_SIZE][GAME_SIZE]cell

func main() {
	grid := initGrid()
	printGrid(grid)
}

func initGrid() gameGrid {
	var grid gameGrid
	view := grid[GAME_SIZE/2][GAME_SIZE/2: GAME_SIZE/2 + 4]
	for i := range view{
		view[i].state = white
	}
	return grid
}

func printGrid(grid gameGrid) {
	str := ""
	for _, row := range grid {
		rowStr := ""
		for _, c := range row {
			rowStr += c.String()
		}
		str += rowStr+ "\n"
	}
	fmt.Println(str)
}
