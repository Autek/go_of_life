package main

import (
	"fmt"
	"time"
	"image/color"
)

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
    "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

const(GAME_SIZE = 25)
type cellState int

const (
	black cellState = iota
	white
)

type cell struct {
	state cellState
}

func in_range(row int, col int, tab * gameGrid) bool{
	return row >= 0 && row < len(tab) && col >= 0 && col < len(tab[0])
}

func nb_neighbours(row int, col int, tab * gameGrid) int{
	neighbours := 0
	for i := row-1; i <= row + 1; i++ {
		for j := col-1; j <= col + 1; j++ {
			if in_range(i, j, tab) && tab[i][j].state == white {
				neighbours += 1
			}
		}
	}
	if tab[row][col].state == white {
		neighbours -= 1;
	}
	return neighbours
}

func nb_neighbours_tab(tab * gameGrid) neighbours_grid{
	var neighbours_tab [len(tab)][len(tab[0])]int
	for i, row := range neighbours_tab {
		for j := range row{
			neighbours_tab[i][j] = nb_neighbours(i, j, tab)
		}
	}
	return neighbours_tab
}


func new_cell_state(currentState cellState, nb_neighbours int) cellState{
	if currentState == white {
		if nb_neighbours == 2 || nb_neighbours == 3 {
			return white
		}
		return black
	} else {
		if nb_neighbours == 3 {
			return white
		}
		return black
	}
}

func update_states(tab * gameGrid) {
	neighbours_tab := nb_neighbours_tab(tab)
	for i, row := range tab {
		for j := range row {
			tab[i][j].state = new_cell_state(tab[i][j].state, neighbours_tab[i][j])
		}
	}
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
type neighbours_grid [GAME_SIZE][GAME_SIZE]int

func main() {
	a := app.New()
	w := a.NewWindow("Go of Life")

	w.Resize(fyne.NewSize(200, 200)) //useless

	grid := initGrid()
	container, cells := init_grid()
	w.SetContent(container)

	go func() {
		for range time.Tick(200 * time.Millisecond) {
			update_states(&grid)
			draw(&grid, &cells)
			container.Refresh()
		}
	}()
	w.Show()
	a.Run()
}

func draw(grid * gameGrid, cells *[][]*canvas.Rectangle){
	for i, rows := range grid {
		for j := range rows {
			var cell_color color.Color
			if grid[i][j].state == black {
				cell_color = color.Black
			} else if grid[i][j].state == white {
				cell_color = color.White
			}
			(*cells)[i][j].FillColor = cell_color
		}
	}
}

func init_grid() (*fyne.Container,  [][]*canvas.Rectangle) {
	cells := make([][]*canvas.Rectangle, GAME_SIZE)
	container := container.New(layout.NewGridLayout(GAME_SIZE))
	for i := 0; i < GAME_SIZE; i++{
		for j := 0; j < GAME_SIZE; j++{
			cell := canvas.NewRectangle(color.White)
			cell.SetMinSize(fyne.NewSize(25, 25))
			cells[i] = append(cells[i], cell)
			container.Add(cell)
		}
	}
	return container, cells
}

func initGrid() gameGrid {
	var grid gameGrid

	// a glider
	grid[0][1].state = white
	grid[1][2].state = white
	grid[2][0].state = white
	grid[2][1].state = white
	grid[2][2].state = white
	
	// a blinker
	grid[20][20].state = white
	grid[20][21].state = white
	grid[20][22].state = white


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
