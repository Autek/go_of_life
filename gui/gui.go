package gui

import(
	"my_module/utils"

	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"fyne.io/fyne/v2/canvas"
	"image/color"
)

type GridWidget struct {
	widget.BaseWidget
	ScaleFactor float32
	CellsState map[utils.Position]struct{}
	UpdatePos []utils.Position
}

type GridWidgetRenderer struct {
	Widget *GridWidget
	Cells []fyne.CanvasObject
}

func NewGridWidget() *GridWidget {
	item := &GridWidget{ScaleFactor: 50, CellsState: make(map[utils.Position]struct{}), UpdatePos: make([]utils.Position, 0)}
	item.ExtendBaseWidget(item)
	return item
}

func (widget *GridWidget) CreateRenderer() fyne.WidgetRenderer {
	cells := make([]fyne.CanvasObject, 0, len(widget.CellsState))
	for _ = range widget.CellsState {
		rect := canvas.NewRectangle(color.White)
		rect.Resize(fyne.NewSize(widget.ScaleFactor, widget.ScaleFactor))
		cells = append(cells, rect)
	}
	return &GridWidgetRenderer{Widget: widget, Cells: cells}
}

func (renderer *GridWidgetRenderer) Layout(size fyne.Size) {
	for i := range renderer.Cells {
		scaleFactor := renderer.Widget.ScaleFactor
		pos := utils.IndexToPosition(i, int(math.Ceil(float64(size.Width / scaleFactor))))

		renderer.Cells[i].Move(fyne.NewPos(float32(pos.X) * scaleFactor, float32(pos.Y) * scaleFactor))
	}
}

func (renderer *GridWidgetRenderer) MinSize() fyne.Size {
	return fyne.NewSize(0, 0)
}

func (renderer *GridWidgetRenderer) Refresh() {
	for _, position := range renderer.Widget.UpdatePos {
		columns := int(math.Ceil(float64(renderer.Widget.Size().Width / renderer.Widget.ScaleFactor)))
		cell_index := utils.PositionToIndex(position, columns)
		if !utils.InRange(cell_index, len(renderer.Cells)) { // the index is out of the drawing bounds
			continue
		}
		cell := renderer.Objects()[cell_index].(*canvas.Rectangle)
		if cell.FillColor == color.White {
			cell.FillColor = color.Black
			delete(renderer.Widget.CellsState, position)

		} else if cell.FillColor == color.Black {
			renderer.Widget.CellsState[position] = struct{}{}
			cell.FillColor = color.White
		}
		cell.Refresh()
	}
	renderer.Widget.UpdatePos = renderer.Widget.UpdatePos[:0]
}

func (renderer *GridWidgetRenderer) Objects() []fyne.CanvasObject {
	scaleFactor := renderer.Widget.ScaleFactor
	columns := int(math.Ceil(float64(renderer.Widget.Size().Width / scaleFactor)))
	rows := int(math.Ceil(float64(renderer.Widget.Size().Height / scaleFactor)))
	cellsNb := rows * columns
	if len(renderer.Cells) < cellsNb {
		for len(renderer.Cells) < cellsNb {
			c := canvas.NewRectangle(color.Black)
			c.Resize(fyne.NewSize(scaleFactor, scaleFactor))
			renderer.Cells = append(renderer.Cells, c)
		}
	} else if len(renderer.Cells) > cellsNb {
		renderer.Cells = renderer.Cells[:cellsNb]
	}
	return renderer.Cells
}

func (renderer *GridWidgetRenderer) Destroy() {
	//nothing to do
}

