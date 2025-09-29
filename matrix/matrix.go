package matrix

type Point struct {
	X int
	Y int
}

type Matrix[T comparable] struct {
	Rows int
	Cols int
	Data []T
	Curr Point
}

type Direction int

var Neighbors = []Point{
	{X: 0, Y: -1},	//up
	{X: 1, Y: -1},	//up_right
	{X: 1, Y: 0},	//right
	{X: 1, Y: 1},	//right_down
	{X: 0, Y: 1},	//down
	{X: -1, Y: 1},	//down_left
	{X: -1, Y: 0},	//left
	{X: -1, Y: -1},	//upleft
}

const (
	UP Direction = iota
	RIGHT_UP
	RIGHT
	RIGHT_DOWN
	DOWN
	LEFT_DOWN
	LEFT
	LEFT_UP
)

func New[T comparable]() *Matrix[T] {
	return &Matrix[T]{
		Rows: 0,
		Cols: 0,
		Data: []T{},
		Curr: Point{X: 0, Y: 0},
	}
}

func (m *Matrix[T]) At(row, col int) T {
	return m.Data[row*m.Cols+col]
}

func (m *Matrix[T]) Set(row, col int, value T) {
	m.Data[row*m.Cols+col] = value
}

// If new size is smaller than current size, elements will be eliminated
func (m *Matrix[T]) Resize(newRows, newCols int) {
	oldSize := m.Rows * m.Cols

	m.Rows = newRows
	m.Cols = newCols

	newSize := newRows * newCols
	if oldSize > newSize { //shrink
		m.Data = m.Data[:newSize]
	} else {
		newData := make([]T, newSize)
		copy(newData, m.Data)
		m.Data = newData
	}
}

func (m *Matrix[T]) Size() int {
	return cap(m.Data)
}

func (m *Matrix[T]) IsNextValid(dir Direction, origin Point) bool {

	switch dir {
	case UP:
		return m.IsValidNeighbor(Point{origin.X + Neighbors[UP].X, origin.Y + Neighbors[UP].Y})
	case RIGHT_UP:
		return m.IsValidNeighbor(Point{origin.X + Neighbors[RIGHT_UP].X, origin.Y + Neighbors[RIGHT_UP].Y})
	case RIGHT:
		return m.IsValidNeighbor(Point{origin.X + Neighbors[RIGHT].X, origin.Y + Neighbors[RIGHT].Y})
	case RIGHT_DOWN:
		return m.IsValidNeighbor(Point{origin.X + Neighbors[RIGHT_DOWN].X, origin.Y + Neighbors[RIGHT_DOWN].Y})
	case DOWN:
		return m.IsValidNeighbor(Point{origin.X + Neighbors[DOWN].X, origin.Y + Neighbors[DOWN].Y})
	case LEFT_DOWN:
		return m.IsValidNeighbor(Point{origin.X + Neighbors[LEFT_DOWN].X, origin.Y + Neighbors[LEFT_DOWN].Y})
	case LEFT:
		return m.IsValidNeighbor(Point{origin.X + Neighbors[LEFT].X, origin.Y + Neighbors[LEFT].Y})
	case LEFT_UP:
		return m.IsValidNeighbor(Point{origin.X + Neighbors[LEFT_UP].X, origin.Y + Neighbors[LEFT_UP].Y})
	default: //TODO verify
		return false
	}
}

func (m *Matrix[T]) NextPoint(dir Direction, origin Point) Point {
	return Point{origin.X + Neighbors[dir].X, origin.Y + Neighbors[dir].Y}
}


func (m *Matrix[T])IsValidNeighbor( neighbor Point ) bool {
	if neighbor.X >= 0 && neighbor.X < m.Cols && neighbor.Y >= 0 && neighbor.Y < m.Rows {
		return true
	}

	return false
}