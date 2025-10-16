package advent13

import (
	"testing"
	"slices"
	"os"
)

func Test(t *testing.T) {
	got, _ := LowestPriceForMaxPrizes("./test.txt")
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestReadFile(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		t.Fatalf("Unexpected error")
	}

	got, err := ReadFile(file)
	if err != nil {
		t.Fatalf("Unexpected error")
	}

	want := []Challange{
		{
			A: 		Point{X: 94, Y: 34},
			B: 		Point{X: 22, Y: 67},
			Prize: 	Point{X: 8400, Y: 5400},
		},
		{
			A: 		Point{X: 26, Y: 57},
			B: 		Point{X: 51, Y: 13},
			Prize: 	Point{X: 15496, Y: 17815},
		},
		{
			A: 		Point{X: 20, Y: 81},
			B: 		Point{X: 66, Y: 14},
			Prize: 	Point{X: 13062, Y: 10478},
		},
	}

	if !slices.Equal(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}