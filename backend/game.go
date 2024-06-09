package backend

type Game struct{
	State State
}

type Position struct{
	X int
	Y int
}

type cellState int

const (
	black cellState = iota
	white
)

type State map[Position]interface{}	// State is set of positions of white cells

func (g *Game) Update() *[]Position {
	nextState := make(State)
	updatePos := make([]Position, 0)
	alreadyUpdated := make(map[Position]struct{})
	for pos := range g.State {
		for posX := pos.X - 1 ; posX <= pos.X + 1; posX++ {
			for posY := pos.Y - 1; posY <= pos.Y + 1; posY++ {
				currentPos := Position{X: posX, Y: posY}
				if _, updated := alreadyUpdated[currentPos]; updated {
					continue
				}
				neighbours := neighboursNb(currentPos, g.State)

				_, exists := g.State[currentPos]
				var currentState cellState
				if exists {
					currentState = white
				} else {
					currentState = black
				}

				newCellState := newCell(currentState, neighbours)
				if newCellState == white {
					nextState[currentPos] = struct{}{}
				}
				if newCellState != currentState {
					updatePos = append(updatePos, currentPos)
				}
				alreadyUpdated[currentPos] = struct{}{}
			}
		}
	}
	g.State = nextState
	return &updatePos
}

func neighboursNb(pos Position, state State) int{
	neighbours := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if _, exists := state[Position{X: pos.X + i, Y: pos.Y + j}]; exists{
				neighbours += 1
			}
		}
	}
	return neighbours
}

func newCell(currentState cellState, neighboursNb int) cellState{
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

