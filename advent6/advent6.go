package advent6

import (
	// "slices"
	"os"
	"bufio"

	"erikberman.matrix.com"
)
const (
	OCCUPIED 	=	'X'
	GUARD		=	'^'
)

func CountDistinctPositions(pathToFile string) (int, error) {
	// read file into matrix
	mtx, err := FileToMatrix(pathToFile)
	if err != nil {
		return 0, err
	}

	curr := FindGuard(mtx)
	direction := matrix.UP

	counter := 0
	for {

		TryCaptureSpot(mtx, curr, &counter)

		if ! mtx.IsNextValid(direction, curr) { //check out of board
			break
		}

		next := mtx.NextPoint(direction, curr)
		
		if mtx.At(next.Y, next.X) == GUARD {
			direction = (direction + 2) % 8 //turn 90 degrees clockwise
			next = mtx.NextPoint(direction, curr)
		}
		curr = next
	}

	return counter, nil
}

func TryCaptureSpot( mtx *matrix.Matrix[rune], position matrix.Point, counter *int) {
	if mtx.At(position.Y, position.X) != OCCUPIED {
		mtx.Set(position.Y, position.X, OCCUPIED)
		*counter++
	}
}

func FileToMatrix (pathToFile string) (*matrix.Matrix[rune], error) {

	file, err := os.Open(pathToFile)
	if err != nil { 
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mtx := matrix.New[rune]()
	isMatrixResize := false
	rowToInsert := 0
	
	for scanner.Scan() {
		line := scanner.Text()
		runeLine := []rune(line)
		
		if ! isMatrixResize {
			mtx.Resize(len(runeLine), len(runeLine))
			isMatrixResize = true
		}
		
		colToInsert := 0
		for _, v := range runeLine {
			mtx.Set(rowToInsert, colToInsert, v)
			colToInsert++
		}
		rowToInsert++
	}

	return mtx, nil
}

func FindGuard (mtx *matrix.Matrix[rune]) matrix.Point {
	for r := 0 ; r < mtx.Rows ; r++ {
		for c := 0 ; c < mtx.Cols ; c++ {
			if mtx.At(r,c) == GUARD {
				return matrix.Point{c,r}
			}
		}
	}

	return matrix.Point{}
}