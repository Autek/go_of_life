package gui

import(
	"my_module/backend"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

type GridWidget struct {
	widget.BaseWidget
	Zoom float32
	Grid []cellState
}

type cellState bool

type GridWidgetRenderer struct {
	Widget *GridWidget
	Cells []fyne.CanvasObject
}

func (widget *GridWidget) Update(to_update *[]backend.Position){
	for _, pos := range *to_update {
		index := positionToIndex(pos)
		if widget.Grid[index] == true {
			widget.Grid[index] = false
		} else {
			widget.Grid[index] = true
		}
	}
}

func (widget *GridWidget) Scrolled(ev *fyne.ScrollEvent){
	scaleFactor := float32(1.1)
	if ev.Scrolled.DY > 0 {
		widget.Zoom *= scaleFactor
	} else {
		widget.Zoom /= scaleFactor
	}
	widget.Refresh()
}

func NewGridWidget() *GridWidget {
	item := &GridWidget{Zoom: 50, Grid: make([]cellState, 100 * 100)}
	item.ExtendBaseWidget(item)
	return item
}

func (widget *GridWidget) CreateRenderer() fyne.WidgetRenderer {
	cells := make([]fyne.CanvasObject, 0, len(widget.Grid))
	for i:= 0; i < len(widget.Grid); i++{
		rect := canvas.NewRectangle(color.Black)
		rect.Resize(fyne.NewSize(widget.Zoom, widget.Zoom))
		cells = append(cells, rect)
	}
	return &GridWidgetRenderer{Widget: widget, Cells: cells}
}

func (renderer *GridWidgetRenderer) Layout(size fyne.Size) {
	for i := range renderer.Cells {
		zoom := renderer.Widget.Zoom
		pos := indexToPosition(i)
		renderer.Cells[i].Move(fyne.NewPos(float32(pos.X) * zoom, float32(pos.Y) * zoom))
	}
}

func (renderer *GridWidgetRenderer) MinSize() fyne.Size {
	return fyne.NewSize(0, 0)
}

func (renderer *GridWidgetRenderer) Refresh() {
	layout := false
	for i := range renderer.Cells {
		cell := renderer.Cells[i].(*canvas.Rectangle)
		if cell.FillColor == color.White && !renderer.Widget.Grid[i] {
			cell.FillColor = color.Black
			cell.Refresh()
		} else if cell.FillColor == color.Black && renderer.Widget.Grid[i] {
			cell.FillColor = color.White
			cell.Refresh()
		}
		zoom := renderer.Widget.Zoom
		if cell.Size().Width != zoom {
			cell.Resize(fyne.NewSize(zoom, zoom))
			cell.Refresh()
			layout = true
		}
	}
	if layout {
		renderer.Layout(renderer.Widget.Size())
	}
}

func (renderer *GridWidgetRenderer) Objects() []fyne.CanvasObject {
	return renderer.Cells
}

func (renderer *GridWidgetRenderer) Destroy() {
	//nothing to do
}

func indexToPosition(i int) backend.Position {
	x := i % 100
	y := i / 100
	return backend.Position{X: x, Y: y}
}

func positionToIndex(pos backend.Position) int {
	return pos.Y * 100 + pos.X
}

func inRange(i int, length int) bool {
	return i >= 0 && i < length
}

func inRange2D(pos backend.Position, width int, height int) bool {
	return inRange(pos.X, width) && inRange(pos.Y, height)
}
