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

func RecGetRegionPt2(mtx *matrix.Matrix[rune], current matrix.Point, area, sides *int){

	dirs := []matrix.Direction{matrix.UP, matrix.RIGHT, matrix.DOWN, matrix.LEFT}
	
	currValue := mtx.At(current.Y, current.X)
	
	if unicode.IsUpper(currValue) {
		
		if IsTopLeftCorner(mtx, current) /* || IsInnerTopLeftCorner(mtx, current)  */{
			(*sides)++
		} 
		if IsTopRightCorner(mtx, current) /* ||  IsInnerTopRightCorner(mtx, current) */ {
			(*sides)++
		} 
		if IsBottomLeftCorner(mtx, current)/*  || IsInnerBottomLeftCorner(mtx, current) */{
			(*sides)++
		} 
		if IsBottomRightCorner(mtx, current) /* || IsInnerBottomRightCorner(mtx, current) */ {
			(*sides)++
		} 
		if IsInnerBottomLeftCorner(mtx, current){
			(*sides)++
		} 
		if IsInnerBottomRightCorner(mtx, current) {
			(*sides)++
		} 
		if IsInnerTopLeftCorner(mtx, current) {
			(*sides)++
		} 
		if IsInnerTopRightCorner(mtx, current) {
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

		RecGetRegionPt2(mtx, next, area, sides)
	}

}



// foreach idx in mtx.slice:

//outern top left corner 
// if !up && !left
func IsTopLeftCorner(mtx *matrix.Matrix[rune], current matrix.Point) bool {
	currValue := mtx.At(current.Y, current.X)
	nextUpValue := ' '
	nextLeftValue := ' '

	if mtx.IsNextValid(matrix.UP, current) {
		nextUpValue = mtx.At(mtx.NextPoint(matrix.UP, current).Y, mtx.NextPoint(matrix.UP, current).X)
	}
	if mtx.IsNextValid(matrix.LEFT, current) {
		nextLeftValue = mtx.At(mtx.NextPoint(matrix.LEFT, current).Y, mtx.NextPoint(matrix.LEFT, current).X)
	}

	if 	(!mtx.IsNextValid(matrix.UP, current) || nextUpValue !=  currValue && nextUpValue != unicode.ToLower(currValue) ) &&
		(!mtx.IsNextValid(matrix.LEFT, current) || nextLeftValue !=  currValue && nextLeftValue != unicode.ToLower(currValue)) {
			return true
		}

		return false
}



// //outern top right corner
// if !up && !right 
func IsTopRightCorner(mtx *matrix.Matrix[rune], current matrix.Point) bool {
	currValue := mtx.At(current.Y, current.X)
	nextUpValue := ' '
	nextRightValue := ' '

	if mtx.IsNextValid(matrix.UP, current) {
		nextUpValue = mtx.At(mtx.NextPoint(matrix.UP, current).Y, mtx.NextPoint(matrix.UP, current).X)
	}
	if mtx.IsNextValid(matrix.RIGHT, current) {
		nextRightValue = mtx.At(mtx.NextPoint(matrix.RIGHT, current).Y, mtx.NextPoint(matrix.RIGHT, current).X)
	}

	if 	(!mtx.IsNextValid(matrix.UP, current) || nextUpValue !=  currValue && nextUpValue != unicode.ToLower(currValue) ) &&
		(!mtx.IsNextValid(matrix.RIGHT, current) || nextRightValue !=  currValue && nextRightValue != unicode.ToLower(currValue)) {
			return true
		}

		return false
}



// //outern bottom right corner
// if !down && !right
func IsBottomRightCorner(mtx *matrix.Matrix[rune], current matrix.Point) bool {
	currValue := mtx.At(current.Y, current.X)
	nextDownValue := ' '
	nextRightValue := ' '

	if mtx.IsNextValid(matrix.DOWN, current) {
		nextDownValue = mtx.At(mtx.NextPoint(matrix.DOWN, current).Y, mtx.NextPoint(matrix.DOWN, current).X)
	}
	if mtx.IsNextValid(matrix.RIGHT, current) {
		nextRightValue = mtx.At(mtx.NextPoint(matrix.RIGHT, current).Y, mtx.NextPoint(matrix.RIGHT, current).X)
	}

	if 	(!mtx.IsNextValid(matrix.DOWN, current) || nextDownValue !=  currValue && nextDownValue != unicode.ToLower(currValue) ) &&
		(!mtx.IsNextValid(matrix.RIGHT, current) || nextRightValue !=  currValue && nextRightValue != unicode.ToLower(currValue)) {
			return true
		}

		return false
}

// // outern bottom left corner
// if !down && !left
func IsBottomLeftCorner(mtx *matrix.Matrix[rune], current matrix.Point) bool {
	currValue := mtx.At(current.Y, current.X)
	nextDownValue := ' '
	nextLeftValue := ' '

	if mtx.IsNextValid(matrix.DOWN, current) {
		nextDownValue = mtx.At(mtx.NextPoint(matrix.DOWN, current).Y, mtx.NextPoint(matrix.DOWN, current).X)
	}
	if mtx.IsNextValid(matrix.LEFT, current) {
		nextLeftValue = mtx.At(mtx.NextPoint(matrix.LEFT, current).Y, mtx.NextPoint(matrix.LEFT, current).X)
	}

	if 	(!mtx.IsNextValid(matrix.DOWN, current) || nextDownValue !=  currValue && nextDownValue != unicode.ToLower(currValue) ) &&
		(!mtx.IsNextValid(matrix.LEFT, current) || nextLeftValue !=  currValue && nextLeftValue != unicode.ToLower(currValue)) {
			return true
		}

		return false
}

// // inner bottom right corner
// if right && down && !bottom-right
func IsInnerBottomRightCorner(mtx *matrix.Matrix[rune], current matrix.Point) bool {
	currValue := mtx.At(current.Y, current.X)
	nextDownValue := ' '
	nextRightValue := ' '
	nextBottomRightValue := ' '

	if mtx.IsNextValid(matrix.DOWN, current) {
		nextDownValue = mtx.At(mtx.NextPoint(matrix.DOWN, current).Y, mtx.NextPoint(matrix.DOWN, current).X)
	}
	if mtx.IsNextValid(matrix.RIGHT, current) {
		nextRightValue = mtx.At(mtx.NextPoint(matrix.RIGHT, current).Y, mtx.NextPoint(matrix.RIGHT, current).X)
	}
	if mtx.IsNextValid(matrix.RIGHT_DOWN, current) {
		nextBottomRightValue = mtx.At(mtx.NextPoint(matrix.RIGHT_DOWN, current).Y, mtx.NextPoint(matrix.RIGHT_DOWN, current).X)
	}
	if 	(mtx.IsNextValid(matrix.DOWN, current) && nextDownValue == currValue) &&
		(mtx.IsNextValid(matrix.RIGHT, current) && nextRightValue ==  currValue) && 
		(!mtx.IsNextValid(matrix.RIGHT_DOWN, current) || nextBottomRightValue !=  currValue && nextBottomRightValue != unicode.ToLower(currValue)) {
			return true
		}

		return false
}
// // inner top right corner
// if right && up && !top-right
func IsInnerTopRightCorner(mtx *matrix.Matrix[rune], current matrix.Point) bool {
	currValue := mtx.At(current.Y, current.X)
	nextUpValue := ' '
	nextRightValue := ' '
	nextTopRightValue := ' '

	if mtx.IsNextValid(matrix.UP, current) {
		nextUpValue = mtx.At(mtx.NextPoint(matrix.UP, current).Y, mtx.NextPoint(matrix.UP, current).X)
	}
	if mtx.IsNextValid(matrix.RIGHT, current) {
		nextRightValue = mtx.At(mtx.NextPoint(matrix.RIGHT, current).Y, mtx.NextPoint(matrix.RIGHT, current).X)
	}
	if mtx.IsNextValid(matrix.RIGHT_UP, current) {
		nextTopRightValue = mtx.At(mtx.NextPoint(matrix.RIGHT_UP, current).Y, mtx.NextPoint(matrix.RIGHT_UP, current).X)
	}
	if 	(mtx.IsNextValid(matrix.UP, current) && nextUpValue == currValue) &&
		(mtx.IsNextValid(matrix.RIGHT, current) && nextRightValue ==  currValue) && 
		(!mtx.IsNextValid(matrix.RIGHT_UP, current) || nextTopRightValue !=  currValue && nextTopRightValue != unicode.ToLower(currValue)) {
			return true
		}

		return false
}

// // inner top left corner
// if left && up && !top-left
func IsInnerTopLeftCorner(mtx *matrix.Matrix[rune], current matrix.Point) bool {
	currValue := mtx.At(current.Y, current.X)
	nextUpValue := ' '
	nextLeftValue := ' '
	nextTopLeftValue := ' '

	if mtx.IsNextValid(matrix.UP, current) {
		nextUpValue = mtx.At(mtx.NextPoint(matrix.UP, current).Y, mtx.NextPoint(matrix.UP, current).X)
	}
	if mtx.IsNextValid(matrix.LEFT, current) {
		nextLeftValue = mtx.At(mtx.NextPoint(matrix.LEFT, current).Y, mtx.NextPoint(matrix.LEFT, current).X)
	}
	if mtx.IsNextValid(matrix.LEFT_UP, current) {
		nextTopLeftValue = mtx.At(mtx.NextPoint(matrix.LEFT_UP, current).Y, mtx.NextPoint(matrix.LEFT_UP, current).X)
	}
	if 	(mtx.IsNextValid(matrix.UP, current) && nextUpValue == currValue) &&
		(mtx.IsNextValid(matrix.LEFT, current) && nextLeftValue ==  currValue) && 
		(!mtx.IsNextValid(matrix.LEFT_UP, current) || nextTopLeftValue !=  currValue && nextTopLeftValue != unicode.ToLower(currValue)) {
			return true
		}

		return false
}




// // inner bottom left corner
// if left && bottom && !bottom-left

func IsInnerBottomLeftCorner(mtx *matrix.Matrix[rune], current matrix.Point) bool {
	currValue := mtx.At(current.Y, current.X)
	nextDownValue := ' '
	nextLeftValue := ' '
	nextBottomLeftValue := ' '

	if mtx.IsNextValid(matrix.DOWN, current) {
		nextDownValue = mtx.At(mtx.NextPoint(matrix.DOWN, current).Y, mtx.NextPoint(matrix.DOWN, current).X)
	}
	if mtx.IsNextValid(matrix.LEFT, current) {
		nextLeftValue = mtx.At(mtx.NextPoint(matrix.LEFT, current).Y, mtx.NextPoint(matrix.LEFT, current).X)
	}
	if mtx.IsNextValid(matrix.LEFT_DOWN, current) {
		nextBottomLeftValue = mtx.At(mtx.NextPoint(matrix.LEFT_DOWN, current).Y, mtx.NextPoint(matrix.LEFT_DOWN, current).X)
	}
	if 	(mtx.IsNextValid(matrix.DOWN, current) && nextDownValue == currValue) &&
		(mtx.IsNextValid(matrix.LEFT, current) && nextLeftValue ==  currValue) && 
		(!mtx.IsNextValid(matrix.LEFT_DOWN, current) || nextBottomLeftValue !=  currValue && nextBottomLeftValue != unicode.ToLower(currValue)) {
			return true
		}

		return false
}
