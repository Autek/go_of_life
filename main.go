package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"my_module/gui"
	"my_module/backend"
)

func main() {
	a := app.New()
	w := a.NewWindow("Go of Life")
	w.Resize(fyne.NewSize(400, 400)) //useless
	widget := gui.NewGridWidget()
	w.SetContent(widget)
	w.Show()
	positions := []backend.Position{
		{X: 3, Y: 3},
		{X: 1, Y: 3},
		{X: 2, Y: 3},
		{X: 3, Y: 2},
		{X: 2, Y: 1},
	}
	game := backend.Game{State: make(backend.State)}
	game.State[backend.Position{X: 3, Y: 3}] = struct{}{}
	game.State[backend.Position{X: 1, Y: 3}] = struct{}{}
	game.State[backend.Position{X: 2, Y: 3}] = struct{}{}
	game.State[backend.Position{X: 3, Y: 2}] = struct{}{}
	game.State[backend.Position{X: 2, Y: 1}] = struct{}{}
	widget.Update(&positions)
	

	go func() {
		for range time.Tick(500 * time.Millisecond) {
			updatePos := game.Update()
			widget.Update(updatePos)
			widget.Refresh()
		}
	}()

	a.Run()
}
