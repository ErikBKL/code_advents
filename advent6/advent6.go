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
	if err != nil {
		return nil, 0, err
	}

	basePosition, curr := FindGuard(mtx)
	direction := matrix.UP

	counter := 0

	// place new barrier in basePosition
	
	for mtx.IsNextValid(direction, curr) { //exit loop when get out of board
		
		copyMtx := matrix.Matrix[rune]{
			Rows: mtx.Rows,
			Cols: mtx.Cols,
			Data: make([]rune, mtx.Cols*mtx.Rows),
			Curr: curr,
		}
		
		copy(copyMtx.Data, mtx.Data) 		//reset matrix on every re-run
		collisions := map[Collision]int{} 	//reset map on every re-run

		TryMakeLoop(&copyMtx, direction, basePosition, curr, &counter, collisions)

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

func TryMakeLoop(mtx *matrix.Matrix[rune], direction matrix.Direction, basePosition, currNewBlock matrix.Point, counter *int, collisions map[Collision]int) {
	if !mtx.IsNextValid(direction, currNewBlock) { //check out of board
		return
	}

	mtx.Set(currNewBlock.Y, currNewBlock.X, BLOCK) //new block 

	next := basePosition

	for mtx.IsNextValid(direction, next) { //exit loop when get out of board

		for mtx.At(next.Y, next.X) == BLOCK {
			// add Collision{next, direction} to a map of collisions.
			c := Collision{next, direction}
			collisions[c]++
			if collisions[c] > 1 {
				*counter++
				return
			}

			direction = (direction + 2) % 8 //turn 90 degrees clockwise
		}
		
		next = mtx.NextPoint(direction, next)
	}
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
