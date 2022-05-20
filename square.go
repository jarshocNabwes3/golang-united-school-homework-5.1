package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (square Square) End() Point {
	var x, y int

	x = square.start.x + int(square.a)
	y = square.start.y + int(square.a)

	// implement me

	return Point{x, y}
}

func (square Square) Area() uint {
	var area uint
	// implement me

	return area
}

func (square Square) Perimeter() uint {
	var prmtr uint
	// implement me

	return prmtr
}
