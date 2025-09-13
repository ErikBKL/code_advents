package advent4

import (
	"os"
	"bufio"
	"errors"
)

type Matrix struct {
	MaxRow 		int
	MaxCols 	int
	Data 		[]rune
}

type Gps struct {
	X int
	Y int
}

type Direction int

const (
	up Direction = iota
	up_right
	right
	right_down
	down
	down_left
	left
	up_left
)

const invalid = -7

var g_targets = []rune{'X', 'M', 'A', 'S'}


var g_AllPossibleNeighbors = []Gps{
	{X: 0, Y: 1},	//up
	{X: 1, Y: 1},	//up_right
	{X: 1, Y: 0},	//right
	{X: 1, Y: -1},	//right_down
	{X: 0, Y: -1},	//down
	{X: -1, Y: -1},	//down_left
	{X: -1, Y: 0},	//left
	{X: -1, Y: 1},	//upleft
}





func (m *Matrix) At(row, col int) rune {
	return m.Data[row * m.MaxCols + col]
}

func (m *Matrix) Set(row, col int, value rune) {
	m.Data[row * m.MaxCols + col] = value
}



func ReadFileToMatrix(pathToFile string) (Matrix, error) {
	file, err := os.Open(pathToFile)
	if err != nil {
		return Matrix{}, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text()) 
	}

	err = AssertValidLines(lines)
	if err != nil {
		return Matrix{}, err
	}

	mtx := Matrix{
		MaxRow: len(lines),
		MaxCols: len(lines[0]),
		Data: []rune{},
	}

	for _, line := range lines {
		mtx.Data = append(mtx.Data, []rune(line)...)
	}

	return mtx, nil
}

func ProcessBatchStage1(mtx Matrix) int {
	
	matches := 0
	for row := 0 ; row < mtx.MaxRow ; row++ {
		for col := 0 ; col < mtx.MaxCols ; col++ {
			if mtx.At(row, col) == 'X' {
				validNeighbors := GetValidNeighbors(mtx, row, col)
				// Search all valid neighbors for an 'M'
				// Once find M,  search for A and S in same direction in straight line
				ComposeXMAS(mtx, Gps{X: row, Y: col}, validNeighbors, 'M', &matches)
			}
		}
	}

	return matches
}

func ProcessBatchStage2(mtx Matrix) int {
	matches := 0
	for row := 0 ; row < mtx.MaxRow ; row++ {
		for col := 0 ; col < mtx.MaxCols ; col++ {
			if mtx.At(row, col) == 'A' {
				allNeighbors := GetAllNeighbors(mtx, row, col)
				// If any of the cross neighbors are invalid, continue
				if IsCrossValid(mtx, allNeighbors) { 
					matches++
				}
			}
		}
	}

	return matches
}

func IsCrossValid (mtx Matrix, neighbors []Gps) bool {
	topLeft := Gps{neighbors[up_left].X, neighbors[up_left].Y}
	topRight := Gps{neighbors[up_right].X, neighbors[up_right].Y}
	bottomLeft := Gps{neighbors[down_left].X, neighbors[down_left].Y}
	bottomRight := Gps{neighbors[right_down].X, neighbors[right_down].Y}

	topLeftValid := IsValidNeighbor(topLeft.X, topLeft.Y, mtx.MaxRow, mtx.MaxCols)
	topRightValid := IsValidNeighbor(topRight.X, topRight.Y, mtx.MaxRow, mtx.MaxCols)
	bottomLeftValid := IsValidNeighbor(bottomLeft.X, bottomLeft.Y, mtx.MaxRow, mtx.MaxCols)
	bottomRightValid := IsValidNeighbor(bottomRight.X, bottomRight.Y, mtx.MaxRow, mtx.MaxCols)
	
	if ! topLeftValid || ! topRightValid || ! bottomLeftValid || ! bottomRightValid {
		return false
	}

	topLeftValue := mtx.At(topLeft.X, topLeft.Y)
	topRightValue := mtx.At(topRight.X, topRight.Y)
	bottomLeftValue := mtx.At(bottomLeft.X, bottomLeft.Y)
	bottomRightValue := mtx.At(bottomRight.X, bottomRight.Y)

	// case 1 topleft is M and topRight is equal
	if topLeftValue == 'M' && topRightValue == 'M' && bottomLeftValue == 'S' && bottomRightValue == 'S' {
		return true
		// case 3 topLeft is S and topRight is equal
	} else if topLeftValue == 'S' && topRightValue == 'S' && bottomLeftValue == 'M' && bottomRightValue == 'M' {
		return true
		// case 2 topLeft is M and topRight is different
	} else if topLeftValue == 'M' && topRightValue == 'S' && bottomLeftValue == 'M' && bottomRightValue == 'S' {
		return true
		// case 4 topleft is S and topRight is different
	} else if topLeftValue == 'S' && topRightValue == 'M' && bottomLeftValue == 'S' && bottomRightValue == 'M' {
		return true
	} else {
		return false
	}
}


func ComposeXMAS(mtx Matrix, origin Gps, validNeighbors []Gps, target rune, matches *int) {

	for _, neighbor_m := range validNeighbors {
		if mtx.At(neighbor_m.X, neighbor_m.Y) == 'M' {//found match
			// Understand the direction we're going down and check the valid neighbors only for that direction
			m := Gps{X: neighbor_m.X, Y: neighbor_m.Y}
			direction := GetDirection(origin, neighbor_m)
			neighbors_a_s := GetNextNeighbors(mtx, m, direction)
			if neighbors_a_s != nil && neighbors_a_s[0] == 'A' && neighbors_a_s[1] == 'S' {
				//happy path - found a complete match
				*matches++
			}
		}
	}
}

func GetNextNeighbors(mtx Matrix, m Gps, direction Direction) []rune{
	firstNeighbor := GetNeighborGps(mtx, m, direction)
	secondNeighbor := GetNeighborGps(mtx, firstNeighbor, direction)
	
	if IsValidNeighbor(secondNeighbor.X, secondNeighbor.Y, mtx.MaxRow, mtx.MaxCols) {
		
		return []rune{	mtx.At(firstNeighbor.X, firstNeighbor.Y),
						mtx.At(secondNeighbor.X, secondNeighbor.Y),
					}

	}

	return nil
}

func GetNeighborGps(mtx Matrix, origin Gps, direction Direction) Gps {
	return Gps{origin.X+ g_AllPossibleNeighbors[direction].X, origin.Y + g_AllPossibleNeighbors[direction].Y}
}
func GetDirection(origin, destination Gps) Direction {
	delta_x := destination.X - origin.X
	delta_y := destination.Y - origin.Y

	if delta_x == 0 && delta_y > 0 {
		return up
	} else if delta_x > 0 && delta_y > 0 {
		return up_right
	}else if delta_x > 0 && delta_y == 0 {
		return right
	}else if delta_x > 0 && delta_y < 0 {
		return right_down
	}else if delta_x == 0 && delta_y < 0 {
		return down
	}else if delta_x < 0 && delta_y < 0 {
		return down_left
	}else if delta_x < 0 && delta_y == 0 {
		return left
	}else {
		return up_left
	}
}

func GetValidNeighbors(mtx Matrix, row, col int)  []Gps {
	validNeighbors := []Gps{}

	for _, neighbor := range g_AllPossibleNeighbors {
		neighborRow := row + neighbor.X
		neighborCol := col + neighbor.Y
		if IsValidNeighbor(neighborRow, neighborCol, mtx.MaxRow, mtx.MaxCols) {
			validNeighbors = append(validNeighbors, Gps{X: neighborRow, Y: neighborCol})
		}
	}
	return validNeighbors
}

func GetAllNeighbors(mtx Matrix, row, col int) []Gps {
	AllNeighbors := []Gps{}
	for _, neighbor := range g_AllPossibleNeighbors {
		neighborRow := row + neighbor.X
		neighborCol := col + neighbor.Y
		AllNeighbors = append(AllNeighbors, Gps{X: neighborRow, Y: neighborCol})
	}
	return AllNeighbors
}

func IsValidNeighbor( row, col, maxRow, maxCol int) bool {
	if row >= 0 && row < maxRow && col >= 0 && col < maxCol {
		return true
	}

	return false
}

func AssertValidLines(lines []string) error {
	if len(lines) == 0 {
		return errors.New("no lines found")
	}

	maxCols := len(lines[0])
	for _, line := range lines {
		if len(line) != maxCols {
			return errors.New("lines have different lengths")
		}
	}

	return nil
}