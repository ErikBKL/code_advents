package advent12

import (
	"unicode"

	matrix "erikberman.matrix.com"
)

const VISITED = '.'

// func TotalFencePrice(pathToFile string) (int, error) {
// 	mtx, err := matrix.FileToMatrix(pathToFile)
// 	if err != nil {
// 		return 0, err
// 	}

// 	ret := 0
// 	for idx := range mtx.Data {
// 		current := mtx.IdxToPoint(idx)
// 		area := 0
// 		perimeter := 0
// 		RecGetRegion(mtx, current, &area, &perimeter)
		
// 		ret += area * perimeter
// 	}

// 	return ret, nil
// }

// func RecGetRegion(mtx *matrix.Matrix[rune], current matrix.Point, area, perimeter *int){

// 	dirs := []matrix.Direction{matrix.UP, matrix.RIGHT, matrix.DOWN, matrix.LEFT}
	
// 	currValue := mtx.At(current.Y, current.X)
	
// 	if unicode.IsLower(currValue) {
// 		return
// 	}
	
// 	mtx.Set(current.Y, current.X, unicode.ToLower(currValue))
// 	(*area)++

// 	for _,dir := range dirs {
// 		if ! mtx.IsNextValid(dir, current) {
// 			(*perimeter)++			
// 			continue
// 		}

// 		next := mtx.NextPoint(dir, current)
// 		nextValue := mtx.At(next.Y, next.X)

// 		if nextValue != currValue && nextValue != unicode.ToLower(currValue){
// 			// do ++ to permiter and continue
// 			(*perimeter)++
// 			continue
// 		} 

// 		RecGetRegion(mtx, next, area, perimeter)
// 	}
// }


type Side struct {
	coordinate	int
	dir 		matrix.Direction
}

func TotalFencePricePt2(pathToFile string) (int, error) {
	mtx, err := matrix.FileToMatrix(pathToFile)
	if err != nil {
		return 0, err
	}

	ret := 0
	for idx := range mtx.Data {
		current := mtx.IdxToPoint(idx)
		area := 0
		mapX := map[Side]int{}
		mapY := map[Side]int{}

		RecGetRegionPt2(mtx, current, &area, mapX, mapY)
		
		ret += area * (len(mapX) + len(mapY))
	}

	return ret, nil
}

func RecGetRegionPt2(mtx *matrix.Matrix[rune], current matrix.Point, area *int, mapX, mapY map[Side]int){

	dirs := []matrix.Direction{matrix.UP, matrix.RIGHT, matrix.DOWN, matrix.LEFT}
	
	currValue := mtx.At(current.Y, current.X)
	
	if unicode.IsLower(currValue) {
		return
	}
	
	mtx.Set(current.Y, current.X, unicode.ToLower(currValue))
	(*area)++

	for _,dir := range dirs {
		if ! mtx.IsNextValid(dir, current) {
			switch {
			case dir == matrix.UP || dir == matrix.DOWN:
				mapY[Side{current.Y, dir}]++
			case dir == matrix.RIGHT || dir == matrix.LEFT:
				mapX[Side{current.X, dir}]++
			}

			continue
		}



		next := mtx.NextPoint(dir, current)
		nextValue := mtx.At(next.Y, next.X)

		if nextValue != currValue && nextValue != unicode.ToLower(currValue){
			// do ++ to permiter and continue
			switch {
			case dir == matrix.UP || dir == matrix.DOWN:
				mapY[Side{current.Y, dir}]++
			case dir == matrix.RIGHT || dir == matrix.LEFT:
				mapX[Side{current.X, dir}]++
			}

			continue
		} 

		RecGetRegionPt2(mtx, next, area, mapX, mapY)
	}
}

