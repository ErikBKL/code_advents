package advent15

import (
	"bufio"
	"os"
	"unicode/utf8"
	"slices"
	"strings"
	"fmt"
	matrix "erikberman.matrix.com"
)

type  Challange  struct {
	mtx 	*matrix.Matrix[rune]
	moves	[]rune
}

func SumBoxGps(pathToFile string) (int,error) {
	file, err := os.Open("./test.txt")
	if err != nil {
		return 0, err
	}
	ch, err := ReadFile(file)
	if err != nil {
		return 0, err
	}

	// foreach v in ch.moves:
	for _,move := range ch.moves {
		RelocateElements(&ch, move)
	}


	// sum all boxes
	ret := 0
	for i,v := range ch.mtx.Data {
		if v == 'O' {
			ret += BoxGpsCoordinate(ch.mtx.IdxToPoint(i))
		}
	}

	fmt.Printf("%q", ch.mtx.Data)
	return ret, nil
}



func ReadFile(file *os.File) (Challange, error) {
	scanner := bufio.NewScanner(file)
	ch := Challange{}
	ch.mtx = matrix.New[rune]()

	cols := 0
	rows := 0
	for scanner.Scan() {
		line := scanner.Text()
		
		// if line contains # append to matrix.Data
		switch {
		case strings.Contains(line, "#"):
			rows++
			cols = utf8.RuneCountInString(line)
			for _,v := range line {
				ch.mtx.Data = append(ch.mtx.Data, v)
			}
		case strings.ContainsAny(line, "<>^v"):
			for _,v := range line {
				ch.moves = append(ch.moves, v)
			}
		default:
			continue
		}

		ch.mtx.Cols = cols
		ch.mtx.Rows = rows

		ch.mtx.Curr = ch.mtx.IdxToPoint(FindRobot(ch.mtx))
	}

	return ch, nil
}

func FindRobot(mtx *matrix.Matrix[rune]) int {
	return slices.Index(mtx.Data, '@')
}

func BoxGpsCoordinate(p matrix.Point) int {
	return p.X + p.Y * 100
}

const (
	LEFT = iota
	UP
	RIGHT
	DOWN
)

func RelocateElements(ch *Challange, move rune) {

	moves := []matrix.Direction{
		matrix.UP,
		matrix.RIGHT,
		matrix.DOWN,
		matrix.LEFT,
	}

	commands := []rune{'^', '>', 'v', '<'}

	next := matrix.Point{}

	for i,v := range commands {
		
		if v != move {
			continue
		}
		current := ch.mtx.Curr
		isRelocate := true

		next = ch.mtx.NextPoint(moves[i], current)

		if ch.mtx.At(next.Y, next.X) == '#' {
			// a movement should happen, if and only if there is not a continuos line of boxes up to the wall
			break
		}
	
		// while next is a 0:
		for ch.mtx.At(next.Y, next.X) == 'O' {
			// next = next->next
			next = ch.mtx.NextPoint(moves[i], next)

			if ch.mtx.At(next.Y, next.X) == '#' {
				isRelocate = false
				break
			}
		}

		if isRelocate {
			ch.mtx.Set(current.Y, current.X, '.')
			ch.mtx.Set(next.Y, next.X, 'O')
			newGuard := ch.mtx.NextPoint(moves[i], current)
			ch.mtx.Set( newGuard.Y, newGuard.X , '@')
			ch.mtx.Curr = newGuard
		}
	}
}



