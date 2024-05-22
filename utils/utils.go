package utils

type Position struct{
	X int
	Y int
}

func IndexToPosition(i int, columnsNb int) Position {
	x := i % columnsNb
	y := i / columnsNb
	return Position{X: x, Y: y}
}

func PositionToIndex(pos Position, columnsNb int) int {
	return pos.Y * columnsNb + pos.X
}

func InRange(i int, length int) bool {
	return i >= 0 && i < length
}

func InRange2D(pos Position, width int, height int) bool {
	return InRange(pos.X, width) && InRange(pos.Y, height)
}
