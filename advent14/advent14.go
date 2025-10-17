package advent14
import (
	"bufio"
	"os"
	"strings"
	"strconv"
	// "fmt"
	matrix "erikberman.matrix.com"
)

type Guard struct {
	Pos		matrix.Point
	Vx		int
	Vy		int
}

func SecurityFactor(pathToFile string) (int){
	mtx := matrix.New[int]()
	mtx.Resize(103, 101)

	file, err := os.Open(pathToFile)
	if err != nil {
		return -1
	}
	defer file.Close()

	// Read the file and extract a slice of guards
	guards, err := ReadFile(file)
	if err != nil {
		return -1
	}
	// Foreach guard in guards (repeat 100 times):
	
	RelocateGuards(mtx, guards)
	
	PlaceGuardsOnBoard(mtx, guards)

	// fmt.Printf("mtx: %+v", mtx.Data)
	m := SumQuadrants(mtx)
	
	ret := 1
	for _,v := range m {
		ret *= v
	}
	return ret
}

func SumQuadrants(mtx *matrix.Matrix[int]) map[int]int {
	ret := map[int]int{}

	// foreach idx from 0 to cols * rows/2:
	for i := 0 ; i < mtx.Cols * (mtx.Rows/2) ; i++ {
		// if idx % cols == cols/2
		switch {
		case i % mtx.Cols == mtx.Cols/2: //blank line
			continue
		case i % mtx.Cols < mtx.Cols/2 : //left quadrant
			p := mtx.IdxToPoint(i)
			ret[0] += mtx.At(p.Y, p.X)
		default: //right quadrant
			p := mtx.IdxToPoint(i)
			ret[1] += mtx.At(p.Y, p.X)
		}
	}

	// fmt.Printf("map part 1 is: %+v", ret)

	for i := mtx.Cols * (mtx.Rows/2 + 1) ; i < len(mtx.Data) ; i++ {

		switch {
		case i % mtx.Cols == mtx.Cols/2: //blank line
			continue
		case i % mtx.Cols < mtx.Cols/2 : //left quadrant
			p := mtx.IdxToPoint(i)
			ret[2] += mtx.At(p.Y, p.X)
		default: //right quadrant
			p := mtx.IdxToPoint(i)
			ret[3] += mtx.At(p.Y, p.X)
		}
	}

	return ret
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