package advent10

import (
	// "fmt"

	matrix "erikberman.matrix.com"
)


func TotalTrailheadScore(pathToFile string) (int, error) {
	mtx, err := matrix.FileToMatrix(pathToFile)
	if err != nil {
		return 0, err
	}

	trailheads := SpotAllZeroes(mtx)

	score := 0
	for _,trailhead := range trailheads {
		TrailheadScore(mtx, trailhead, &score, matrix.UP)
	}
	return score, nil
}


func TrailheadScore(mtx *matrix.Matrix[rune], current matrix.Point, score *int, dir matrix.Direction) {
	// foreach direction of 4 directions:
	for i := 0 ; i < 4 ; i++ {
		// get next point in dir
		next := mtx.NextPoint(dir, current)
		// if next is not valid:
		if !mtx.IsNextValid(dir, current) {
			// continue
			dir = (dir + 2 ) % 8
			continue
		}
		
		
		// if next + 1 != current:
		if matrix.ASCIIToInt(mtx.At(next.Y, next.X)) - 1 != matrix.ASCIIToInt(mtx.At(current.Y, current.X)) && matrix.ASCIIToInt(mtx.At(current.Y, current.X)) != 9 {
			// continue
			dir = (dir + 2 ) % 8
			continue
		}
	
		if matrix.ASCIIToInt(mtx.At(current.Y, current.X)) == 9 {
			*score++
			return
		}

		TrailheadScore(mtx, next, score, matrix.UP)
	}
}


func SpotAllZeroes(mtx *matrix.Matrix[rune]) []matrix.Point {
	ret := []matrix.Point{}
	
	// foreach v in range mtx 
	for i,v := range mtx.Data {
		if matrix.ASCIIToInt(v) == 0 {
			ret = append(ret, mtx.IdxToPoint(i))
		}
	}
	return ret	
}
