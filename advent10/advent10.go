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
		visitedPoints := map[matrix.Point]bool{}
		TrailheadScore(mtx, trailhead, &score, visitedPoints)
	}
	return score, nil
}


func TrailheadScore(mtx *matrix.Matrix[rune], current matrix.Point, score *int, visitedPoints map[matrix.Point]bool) {
	// foreach direction of 4 directions:
	cur := matrix.ASCIIToInt(mtx.At(current.Y, current.X))
    if cur == 9 {
		_, ok := visitedPoints[current]
		if !ok {
			visitedPoints[current] = true
			(*score)++
		}
    }
	
	dirs := []matrix.Direction{matrix.UP, matrix.RIGHT, matrix.DOWN, matrix.LEFT}
	
	for _, dir := range dirs{
		// get next point in dir
		// if next is not valid:
		if !mtx.IsNextValid(dir, current) {
			continue
		}
		
		next := mtx.NextPoint(dir, current)
		
		// if next + 1 != current:
		if matrix.ASCIIToInt(mtx.At(next.Y, next.X)) - 1 != matrix.ASCIIToInt(mtx.At(current.Y, current.X)){
			continue
		}

		TrailheadScore(mtx, next, score, visitedPoints)
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
