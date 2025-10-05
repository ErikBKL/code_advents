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
		score += TrailheadScore(mtx, trailhead)
	}
	return score, nil
}


func TrailheadScore(mtx *matrix.Matrix[rune], trailhead matrix.Point) int {
	dir := matrix.UP

	score := 0
	for dir <= 6 {
		if mtx.IsNextValid(dir, trailhead) {
			RecTrailheadScore(mtx, mtx.NextPoint(dir, trailhead), &score, matrix.UP)
		}
		dir += 2
	}
	
	return score
}

func RecTrailheadScore(mtx *matrix.Matrix[rune], current matrix.Point, score *int, direction matrix.Direction) {

	for direction <= 6 {

		if !mtx.IsNextValid(direction, current) {
			return
		}

		next := mtx.NextPoint(direction, current)
		if mtx.At(next.Y, next.X) + 1 != mtx.At(current.Y, current.X) {
			return
		}
		
		//hapy path
		if mtx.At(current.Y, current.X) == 9 {
			*score++
			return
		}
	
		RecTrailheadScore(mtx, current, score, direction + 2)
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
