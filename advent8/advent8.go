package advent8

import (
	"bufio"
	"os"
	"erikberman.matrix.com"
)



func UniqueAntinodes(pathToFile string) (int, error) {
	// iterate the matrix once and insert every new char that's not a dot into a map[rune]Point of targets
	mtx, err := FileToMatrix(pathToFile)
	if err != nil {
		return 0, err
	}

	targets := MapAntennas(mtx)
	
	// make a map[Point]bool of antinodes
	antinodes := map[matrix.Point]bool{}

	// for each target in range targets:
	for _, target := range targets { //target is a slice of all points where "a" is present
		MarkAntinodesOfTarget(mtx, target, antinodes)
	}
	
	// return amount of keys in map
	return len(antinodes), nil
}


func MarkAntinodesOfTarget(mtx *matrix.Matrix[rune], target []matrix.Point, validAntinodes map[matrix.Point]bool) {
	// for each point in target:
	antinodes := []matrix.Point{}

	for i := 0 ; i < len(target) - 1 ; i++ {
		for j := i + 1 ; j < len(target) ; j++ {
			antinodes = TwoAntinodes(target[i], target[j])
			RegisterValidAntinodes(mtx, antinodes, validAntinodes)
		}
	}
}

func RegisterValidAntinodes(mtx *matrix.Matrix[rune], antinodes []matrix.Point, validAntinodes map[matrix.Point]bool) {
	for _,v := range antinodes {
		if mtx.IsValidPoint(v) {
			validAntinodes[v] = true
		}
	}
}

func TwoAntinodes(pointA matrix.Point, pointB matrix.Point) []matrix.Point {
	deltaX := Abs(pointB.X - pointA.X)
	deltaY := Abs(pointB.Y - pointA.Y)
	
	antinodes := make([]matrix.Point, 2)
	
	switch {
	case pointA.X < pointB.X && pointA.Y < pointB.Y:
		
		antinodes[0].X = pointA.X - deltaX
		antinodes[0].Y = pointA.Y - deltaY
		antinodes[1].X = pointB.X + deltaX
		antinodes[1].Y = pointB.Y + deltaY
	
	case pointA.X < pointB.X && pointA.Y > pointB.Y:
		antinodes[0].X = pointA.X - deltaX
		antinodes[0].Y = pointA.Y + deltaY
		antinodes[1].X = pointB.X + deltaX
		antinodes[1].Y = pointB.Y - deltaY
	
	case pointA.X > pointB.X && pointA.Y < pointB.Y:
		antinodes[0].X = pointA.X + deltaX
		antinodes[0].Y = pointA.Y - deltaY
		antinodes[1].X = pointB.X - deltaX
		antinodes[1].Y = pointB.Y + deltaY

	case pointA.X > pointB.X && pointA.Y > pointB.Y:
		antinodes[0].X = pointA.X + deltaX
		antinodes[0].Y = pointA.Y + deltaY
		antinodes[1].X = pointB.X - deltaX
		antinodes[1].Y = pointB.Y - deltaY
	
	case pointA.X == pointB.X:
		antinodes[0].X, antinodes[1].X = pointA.X, pointB.X
		
		if pointA.Y < pointB.Y {
			antinodes[0].Y = pointA.Y - deltaY
			antinodes[1].Y = pointB.Y + deltaY
		} else {
			antinodes[0].Y = pointA.Y + deltaY
			antinodes[1].Y = pointB.Y - deltaY
		}
	
	case pointA.Y == pointB.Y:
		antinodes[0].Y, antinodes[1].Y = pointA.Y, pointB.Y
		if pointA.X < pointB.X {
			antinodes[0].X = pointA.X - deltaX
			antinodes[1].X = pointB.X + deltaX
		}else {
			antinodes[0].X = pointA.X + deltaX
			antinodes[1].X = pointB.X - deltaX
		}
	}

	return antinodes
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
func MapAntennas(mtx *matrix.Matrix[rune]) map[rune][]matrix.Point {
	m := map[rune][]matrix.Point{}

	for i,v := range mtx.Data {
		if v != '.' {
			m[v] = append(m[v], mtx.IdxToPoint(i))
		}
	}

	return m
}

func FileToMatrix(pathToFile string) (*matrix.Matrix[rune], error) {

	file, err := os.Open(pathToFile)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	mtx := matrix.New[rune]()
	isMatrixResize := false
	rowToInsert := 0

	for scanner.Scan() {
		line := scanner.Text()
		runeLine := []rune(line)

		if !isMatrixResize {
			mtx.Resize(len(runeLine), len(runeLine))
			isMatrixResize = true
		}

		for col, v := range runeLine {
			mtx.Set(rowToInsert, col, v)
		}
		rowToInsert++
	}

	return mtx, nil
}