package advent6

import (
	"testing"
	"erikberman.matrix.com"
	"fmt"

)


func TestCountDistinctPositions(t *testing.T) {

	mtx, got, err := CountDistinctPositions("./test.txt")
	if err != nil {
		t.Fatalf("Unexpected error")
	}

	want := 41

	if got != want {
		t.Errorf("want: %d, got : %d", want, got)
		for r := 0 ; r < mtx.Rows ; r++ {
			for c :=0 ; c <= mtx.Cols ; c++ {
				fmt.Printf("%c", mtx.At(r,c))
			}
			fmt.Printf("\n")
		}
	}
}

// func TestFileToMatrix(t * testing.T) {
// 	_,_,got := FileToMatrix("./test.txt")
// 	want := 130

// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}

// }

func TestFindGuard( t *testing.T) {
	mtx, err := FileToMatrix("./test.txt")
	if err != nil {
		t.Fatalf("Unexpected error")
	}

	got := FindGuard(mtx)
	want := matrix.Point{X: 4 , Y: 6}

	if got.X != want.X || got.Y != want.Y {
		t.Errorf("got %+v want %+v", got, want)
	}
}