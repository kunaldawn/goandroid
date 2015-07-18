package geometry

type Point struct {
	X int // X Coordinate of the point
	Y int // Y Coordinate of the point
}

type Points []Point

type Rect struct {
	TopLeft     Point // Top left coordinate of the rectangle
	BottomRight Point // Bottom right coordinate of the rectangle
}
