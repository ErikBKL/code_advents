package advent14
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	
	matrix "erikberman.matrix.com"
)

type Guard struct {
	Pos		matrix.Point
	Vx		int
	Vy		int
}

func SecurityFactor(pathToFile string) (*matrix.Matrix[int]){
	mtx := matrix.New[int]()
	mtx.Resize(7, 11)

	file, err := os.Open(pathToFile)
	if err != nil {
		return nil
	}
	defer file.Close()

	// Read the file and extract a slice of guards
	guards, err := ReadFile(file)
	if err != nil {
		return nil
	}
	// Foreach guard in guards (repeat 100 times):
	
	RelocateGuards(mtx, guards)
	
	PlaceGuardsOnBoard(mtx, guards)

	m := SumQuadrants(mtx)
	
	return mtx
}

func SumQuadrants(mtx *matrix.Matrix[int]) map[int]int {
	ret := map[int]int{}
	rowsInQuad := mtx.Rows/2
	ColsInQuad := mtx.Cols/2

	// 

	// 1st quadrant top-left is 0
	map[0] = SumSpecificQuadrant(mtx, 0,)
	// 2nd quadrant top-left is 1st quadrant + colsInQuad + 1
	// 3rd quadrant top-left is 0 + cols * (rowsInQuad  + 1)
	// 4th quadrant top-left is 3rd quadrant + colsInQuad + 1
}

func SumSpecificQuadrant(mtx *matrix.Matrix[int], idx int ) int {
	for i := 0 ; i < len(mtx.Data)/4
}

func PlaceGuardsOnBoard(mtx *matrix.Matrix[int], guards []*Guard) {
	for _,guard := range guards {
		mtx.Set(guard.Pos.Y, guard.Pos.X, mtx.At(guard.Pos.Y, guard.Pos.X) + 1)
	}

}
func RelocateGuards(mtx *matrix.Matrix[int], guards []*Guard) {
	for i := 0  ; i < 100 ; i++ {
		for _, guard := range guards {
			newPosX := (guard.Pos.X + guard.Vx)
			if newPosX < 0 {
				guard.Pos.X = mtx.Cols + newPosX % mtx.Cols
			} else {
				guard.Pos.X = (guard.Pos.X + guard.Vx) % mtx.Cols
			}

			newPosY := (guard.Pos.Y + guard.Vy)
			if newPosY < 0 {
				guard.Pos.Y = mtx.Rows + newPosY % mtx.Rows
			} else {
				guard.Pos.Y = (guard.Pos.Y + guard.Vy) % mtx.Rows
			}
		}
	}
}


func ReadFile(file *os.File) ([]*Guard, error) {
	scanner := bufio.NewScanner(file)
	guards := []*Guard{}
	for scanner.Scan() {
		line := scanner.Text()
		g, err := ExtractNumbers(line, "p")
		if err != nil {
			return nil, err
		}
		guards = append(guards, g)
	}

	return guards, nil
}


func ExtractNumbers(line, separator string) (*Guard, error) {
	guard := Guard{}
	
	subStrings := strings.Split(line, " ")
	posx, posy, err := ExtractXY(subStrings[0])
	if err != nil {
		return &Guard{}, err
	}

	vx, vy, err := ExtractXY(subStrings[1])
	if err != nil {
		return &Guard{}, err
	}
	
	guard.Pos = matrix.Point{X: posx, Y: posy}
	guard.Vx = vx
	guard.Vy = vy

	return &guard, nil

}

func ExtractXY(exp string) (int, int, error) {
	// p=0,4

	subStrings := strings.Split(exp, "=")

	subStrings = subStrings[1:] //0,4

	subStrings = strings.Split(subStrings[0], ",") //{0,4}

	x, err := strconv.Atoi(subStrings[0])
	if err != nil {
		return -1,-1, err
	}

	y, err := strconv.Atoi(subStrings[1])
	if err != nil {
		return -1,-1, err
	}

	return x, y, nil
}