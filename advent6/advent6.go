package advent6

import (
	// "slices"
	"bufio"
	"os"

	// "fmt"

	matrix "erikberman.matrix.com"
)

const (
	OCCUPIED = 'X'
	GUARD    = '^'
	BLOCK    = '#'
)

type Collision struct {
	Blockade  matrix.Point
	Direction matrix.Direction
}

func CountLoopMakingBlocks(pathToFile string) (*matrix.Matrix[rune], int, error) {
	mtx, err := FileToMatrix(pathToFile)
	basePosition := FindGuard(mtx)
	mtx, _, err = CountDistinctPositions(pathToFile)
	if err != nil {
		return nil, 0, err
	}

	// traverse the matrix and mark all points of X
	s := []matrix.Point{}
	for i,v := range mtx.Data {
		if v == OCCUPIED {
			s = append(s, mtx.IdxToPoint(i))
		}
	}
	
	counter := 0

	// for mtx.IsNextValid(direction, newBlockade) { //exit loop when get out of board
	for _, blockade := range s {
		collisions := map[Collision]int{}
		copyMtx := matrix.Matrix[rune]{
			Rows: mtx.Rows,
			Cols: mtx.Cols,
			Data: make([]rune, mtx.Cols*mtx.Rows),
			Curr: basePosition,
		}

		copy(copyMtx.Data, mtx.Data) //reset matrix on every re-run

		LookForLoop(&copyMtx, basePosition, blockade, collisions, &counter)
	}

	return mtx, counter, nil
}

func LookForLoop(mtx *matrix.Matrix[rune], basePosition matrix.Point, newBlockade matrix.Point, collisions map[Collision]int, counter *int) {
	// Set the newBlockade on the mtx
	mtx.Set(newBlockade.Y, newBlockade.X, BLOCK)

	direction := matrix.UP
	c := Collision{}
	// for is next valid
	for mtx.IsNextValid(direction, basePosition) { //exit loop when get out of board
		next := mtx.NextPoint(direction, basePosition)
		for mtx.At(next.Y, next.X) == BLOCK {

			// add Collision{next, direction} to a map of collisions.
			c = Collision{next, direction}
			collisions[c]++

			if collisions[c] > 1 {
				*counter++
				return
			}

			direction = (direction + 2) % 8 //turn 90 degrees clockwise
			next = mtx.NextPoint(direction, basePosition)
		}

		basePosition = next
	}
}

func NextValidPoint(mtx *matrix.Matrix[rune], direction matrix.Direction, curr matrix.Point) (matrix.Point, matrix.Direction) {
	next := mtx.NextPoint(direction, curr)

	for mtx.At(next.Y, next.X) == BLOCK {
		direction = (direction + 2) % 8 //turn 90 degrees clockwise
		next = mtx.NextPoint(direction, curr)
	}
	curr = next

	return curr, direction
}

func CountDistinctPositions(pathToFile string) (*matrix.Matrix[rune], int, error) {
	// read file into matrix
	mtx, err := FileToMatrix(pathToFile)
	if err != nil {
		return nil, 0, err
	}

	curr := FindGuard(mtx)
	direction := matrix.UP

	counter := 0
	for {

		TryCaptureSpot(mtx, curr, &counter)

		if !mtx.IsNextValid(direction, curr) { //check out of board
			break
		}

		next := mtx.NextPoint(direction, curr)

		for mtx.At(next.Y, next.X) == BLOCK {
			direction = (direction + 2) % 8 //turn 90 degrees clockwise
			next = mtx.NextPoint(direction, curr)
		}
		curr = next
	}

	return mtx, counter, nil
}

func TryCaptureSpot(mtx *matrix.Matrix[rune], position matrix.Point, counter *int) {
	if mtx.At(position.Y, position.X) != OCCUPIED {
		mtx.Set(position.Y, position.X, OCCUPIED)
		*counter++
	}
}

func FileToMatrix(pathToFile string) (*matrix.Matrix[rune], error) {

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

		if !isMatrixResize {
			mtx.Resize(len(runeLine), len(runeLine))
			isMatrixResize = true
		}

		for col, v := range runeLine {
			mtx.Set(rowToInsert, col, v)
		}
		rowToInsert++
	}
	// fmt.Printf("matrix is: %+v", mtx)

	return mtx, nil
}

func FindGuard(mtx *matrix.Matrix[rune]) matrix.Point {
	for r := 0; r < mtx.Rows; r++ {
		for c := 0; c < mtx.Cols; c++ {
			if mtx.At(r, c) == GUARD {
				return matrix.Point{X: c, Y: r}
			}
		}
	}

	return matrix.Point{}
}
