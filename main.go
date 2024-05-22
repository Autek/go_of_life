package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"my_module/utils"
	"my_module/gui"
)

const (
	black cellState = iota
	white
)

type cellState int

type gameState map[utils.Position]struct{}

func neighboursNb(pos utils.Position, game *gameState) int{
	neighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if _, exists := (*game)[utils.Position{X: pos.X + i, Y: pos.Y + j}]; exists{
				neighbours += 1
			}
		}
	}
	return neighbours
}

func newCellState(currentState cellState, neighboursNb int) cellState{
	if currentState == white {
		if neighboursNb == 2 || neighboursNb == 3 {
			return white
		}
		return black
	} else {
		if neighboursNb == 3 {
			return white
		}
		return black
	}
}

//room for improvement here
func updateStates(game * gameState) *[]utils.Position {
	newGame := make(map[utils.Position]struct{})
	updatePos := make([]utils.Position, 0)
	alreadyUpdated := make(map[utils.Position]struct{})
	for pos := range *game {
		for posX := pos.X - 1 ; posX <= pos.X + 1; posX++ {
			for posY := pos.Y - 1; posY <= pos.Y + 1; posY++ {
				currentPos := utils.Position{X: posX, Y: posY}
				if _, updated := alreadyUpdated[currentPos]; updated {
					continue
				}
				neighbours := neighboursNb(currentPos, game)
				_, exists := (*game)[currentPos]
				var currentState cellState
				if exists {
					currentState = white
				} else {
					currentState = black
				}
				newState := newCellState(currentState, neighbours)
				if newState == white {
					newGame[currentPos] = struct{}{}
				}
				if newState != currentState {
					updatePos = append(updatePos, currentPos)
				}
				alreadyUpdated[currentPos] = struct{}{}
			}
		}
	}
	*game = newGame
	return &updatePos
}


func main() {
	a := app.New()
	w := a.NewWindow("Go of Life")

	w.Resize(fyne.NewSize(400, 400)) //useless

	gameState := initGame()
	widget := gui.NewGridWidget()
	w.SetContent(widget)
	go func() {
		time.Sleep(400 * time.Millisecond)
		for pos := range gameState {
			widget.UpdatePos = append(widget.UpdatePos, pos)
		}
		widget.Refresh()
		for range time.Tick(500 * time.Millisecond) {
			updatePos := updateStates(&gameState)
			widget.UpdatePos = append(widget.UpdatePos, *updatePos...)
			widget.Refresh()
		}
		widget.Refresh()
	}()

	w.Show()
	a.Run()
}

func initGame() gameState {
	grid := make(map[utils.Position]struct{})

	//a simple blinker:
	grid[utils.Position{X: 1, Y: 1}] = struct{}{}
	grid[utils.Position{X: 1, Y: 2}] = struct{}{}
	grid[utils.Position{X: 1, Y: 3}] = struct{}{}
	return grid
	
	//a simple glider:
	grid[utils.Position{X: 1, Y: 1}] = struct{}{}
	grid[utils.Position{X: 2, Y: 2}] = struct{}{}
	grid[utils.Position{X: 2, Y: 3}] = struct{}{}
	grid[utils.Position{X: 1, Y: 3}] = struct{}{}
	grid[utils.Position{X: 3, Y: 3}] = struct{}{}
	return grid

	//a blinker with 15 states
	grid[utils.Position{X: 3, Y: 5}] = struct{}{}
	grid[utils.Position{X: 4, Y: 4}] = struct{}{}
	grid[utils.Position{X: 4, Y: 5}] = struct{}{}
	grid[utils.Position{X: 4, Y: 6}] = struct{}{}
	grid[utils.Position{X: 7, Y: 4}] = struct{}{}
	grid[utils.Position{X: 7, Y: 5}] = struct{}{}
	grid[utils.Position{X: 7, Y: 6}] = struct{}{}
	grid[utils.Position{X: 9, Y: 4}] = struct{}{}
	grid[utils.Position{X: 9, Y: 6}] = struct{}{}
	grid[utils.Position{X: 10, Y: 4}] = struct{}{}
	grid[utils.Position{X: 10, Y: 6}] = struct{}{}
	grid[utils.Position{X: 12, Y: 4}] = struct{}{}
	grid[utils.Position{X: 12, Y: 5}] = struct{}{}
	grid[utils.Position{X: 12, Y: 6}] = struct{}{}
	grid[utils.Position{X: 15, Y: 4}] = struct{}{}
	grid[utils.Position{X: 15, Y: 5}] = struct{}{}
	grid[utils.Position{X: 15, Y: 6}] = struct{}{}
	grid[utils.Position{X: 16, Y: 5}] = struct{}{}
	return grid
}
