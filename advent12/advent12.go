package advent12

import (
	matrix "erikberman.matrix.com"
)

const VISITED = '.'

func TotalFencePrice(pathToFile string) (int, error) {
	mtx, err := matrix.FileToMatrix(pathToFile)
	if err != nil {
		return 0, err
	}

	ret := 0
	for idx := range mtx.Data {
		current := mtx.IdxToPoint(idx)
		area := 0
		perimeter := 0
		RecGetRegion(mtx, current, &area, &perimeter)
		
		ret += area * perimeter
	}

	return ret, nil
}

func RecGetRegion(mtx *matrix.Matrix[rune], current matrix.Point, area, perimeter *int){

	dirs := []matrix.Direction{matrix.UP, matrix.RIGHT, matrix.DOWN, matrix.LEFT}
	
	currValue := mtx.At(current.Y, current.X)
	
	if currValue == VISITED {
		return
	}
	
	mtx.Set(current.Y, current.X, VISITED)
	(*area)++

	for _,dir := range dirs {
		if ! mtx.IsNextValid(dir, current) {
			(*perimeter)++			
			continue
		}

		next := mtx.NextPoint(dir, current)
		nextValue := mtx.At(next.Y, next.X)

		if nextValue != currValue{
			// do ++ to permiter and continue
			(*perimeter)++
			continue
		} 

		RecGetRegion(mtx, next, area, perimeter)
	}
}