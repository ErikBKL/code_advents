package advent12

import (
	"unicode"

	matrix "erikberman.matrix.com"
)

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








func TotalFencePricePt2(pathToFile string) (int, error) {
	mtx, err := matrix.FileToMatrix(pathToFile)
	if err != nil {
		return 0, err
	}

	ret := 0
	for idx := range mtx.Data {
		current := mtx.IdxToPoint(idx)
		area := 0
		sides := 0
		RecGetRegionPt2(mtx, current, &area, &sides)
		
		ret += area * sides
	}

	return ret, nil
}

// check if is corner increment sides
// if I've already been here return
// Mark that I've been here and increment area
// from current pos for all 4 directions:
	// if next is out of board don't check this direction
	// Now that I know next is valid, get the next point and the next value
	// if the next value is different than me - it's not a part of my region, so don't check this direction
	// Now that I know next is valid and it is just like me - re-enter the function.
func RecGetRegionPt2(mtx *matrix.Matrix[rune], current matrix.Point, area, sides *int){

	dirs := []matrix.Direction{matrix.UP, matrix.RIGHT, matrix.DOWN, matrix.LEFT}
	
	currValue := mtx.At(current.Y, current.X)
	
	if unicode.IsUpper(currValue) {
		
		if IsExternCorner(mtx, current, matrix.UP, matrix.LEFT ) {
			(*sides)++
		} 
		if IsExternCorner(mtx, current, matrix.DOWN, matrix.LEFT) {
			(*sides)++
		} 
		if IsExternCorner(mtx, current, matrix.UP, matrix.RIGHT) {
			(*sides)++
		} 
		if IsExternCorner(mtx, current, matrix.DOWN, matrix.RIGHT) {
			(*sides)++
		} 
		if IsInternCorner(mtx, current, matrix.UP, matrix.LEFT, matrix.LEFT_UP){
			(*sides)++
		} 
		if IsInternCorner(mtx, current, matrix.DOWN, matrix.LEFT, matrix.LEFT_DOWN) {
			(*sides)++
		} 
		if IsInternCorner(mtx, current, matrix.UP, matrix.RIGHT, matrix.RIGHT_UP) {
			(*sides)++
		} 
		if IsInternCorner(mtx, current, matrix.DOWN, matrix.RIGHT, matrix.RIGHT_DOWN) {
			(*sides)++
		}
	} else {
		return
	}
	
	mtx.Set(current.Y, current.X, unicode.ToLower(currValue))
	(*area)++

	for _,dir := range dirs {

		if ! mtx.IsNextValid(dir, current) {			
			continue
		}

		next := mtx.NextPoint(dir, current)
		nextValue := mtx.At(next.Y, next.X)

		if nextValue != currValue {
			continue
		} 

		RecGetRegionPt2(mtx, next, area, sides)
	}

}



// foreach idx in mtx.slice:

//outern top left corner 
// if !up && !left
func IsExternCorner(mtx *matrix.Matrix[rune], current matrix.Point, dir1, dir2 matrix.Direction) bool {
	currValue := mtx.At(current.Y, current.X)
	nextValueDir1 := ' '
	nextValueDir2 := ' '

	if mtx.IsNextValid(dir1, current) {
		nextValueDir1 = mtx.At(mtx.NextPoint(dir1, current).Y, mtx.NextPoint(dir1, current).X)
	}
	if mtx.IsNextValid(dir2, current) {
		nextValueDir2 = mtx.At(mtx.NextPoint(dir2, current).Y, mtx.NextPoint(dir2, current).X)
	}

	if 	(!mtx.IsNextValid(dir1, current) || nextValueDir1 !=  currValue && nextValueDir1 != unicode.ToLower(currValue) ) &&
		(!mtx.IsNextValid(dir2, current) || nextValueDir2 !=  currValue && nextValueDir2 != unicode.ToLower(currValue)) {
			return true
		}

		return false
}

func IsInternCorner(mtx *matrix.Matrix[rune], current matrix.Point, dir1, dir2, diagonal matrix.Direction) bool {
	currValue := mtx.At(current.Y, current.X)
	nextValueDir1 := ' '
	nextValueDir2 := ' '
	nextValueDiagonal := ' '

	if mtx.IsNextValid(dir1, current) {
		nextValueDir1 = mtx.At(mtx.NextPoint(dir1, current).Y, mtx.NextPoint(dir1, current).X)
	}
	if mtx.IsNextValid(dir2, current) {
		nextValueDir2 = mtx.At(mtx.NextPoint(dir2, current).Y, mtx.NextPoint(dir2, current).X)
	}
	if mtx.IsNextValid(diagonal, current) {
		nextValueDiagonal = mtx.At(mtx.NextPoint(diagonal, current).Y, mtx.NextPoint(diagonal, current).X)
	}
	if 	(mtx.IsNextValid(dir1, current) && unicode.ToLower(nextValueDir1) == unicode.ToLower(currValue)) &&
		(mtx.IsNextValid(dir2, current) && unicode.ToLower(nextValueDir2) == unicode.ToLower(currValue)) && 
		(!mtx.IsNextValid(diagonal, current) || nextValueDiagonal !=  currValue && nextValueDiagonal != unicode.ToLower(currValue)) {
			return true
		}

		return false
}