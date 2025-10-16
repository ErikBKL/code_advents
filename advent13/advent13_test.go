package advent13

import (
	"testing"
	"slices"
	"os"
)

// func Test(t *testing.T) {
// 	got, _ := LowestPriceForMaxPrizes("./test.txt")
// 	want := 1

// 	if got != want {
// 		t.Errorf("got %d want %d", got, want)
// 	}
// }

func TestReadFile(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		t.Fatalf("Unexpected error")
	}
	defer file.Close()


	challanges, err := ReadFile(file)
	if err != nil {
		t.Fatalf("Unexpected error")
	}

	want := []Challange{
		// {
		// 	A: 		Point{X: 2, Y: 4},
		// 	B: 		Point{X: 4, Y: 2},
		// 	Prize: 	Point{X: 4, Y: 8},
		// },
		{
			A: 		Point{X: 94, Y: 34},
			B: 		Point{X: 22, Y: 67},
			Prize: 	Point{X: 8400, Y: 5400},
		},
		{
			A: 		Point{X: 26, Y: 66},
			B: 		Point{X: 67, Y: 21},
			Prize: 	Point{X: 12748, Y: 12176},
		},
		{
			A: 		Point{X: 17, Y: 86},
			B: 		Point{X: 84, Y: 37},
			Prize: 	Point{X: 7870, Y: 6450},
		},
		{
			A: 		Point{X: 69, Y: 23},
			B: 		Point{X: 27, Y: 71},
			Prize: 	Point{X: 18641, Y: 10279},
		},
	}

	t.Run("ReadFile", func (t *testing.T){
		if !slices.Equal(challanges, want) {
			t.Errorf("got %+v want %+v", challanges, want)
		}
	})
	
	t.Run("MarkAllPoints", func (t *testing.T){
		
		got := 0
		want := 0
		m := map[Point][]int{}

		for _,ch := range challanges {
			m = MarkAllPoints(ch)
			got = len(m)
			want = 35
		}

		if got != want {
			// t.Errorf("%+v", m)
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func TestLowestPriceForMaxPrizes(t *testing.T) {
	got, _ := LowestPriceForMaxPrizes("./test.txt")
	want := 480

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}


func TestLowestPriceForMaxPrizesPt2(t *testing.T) {
	got, err := LowestPriceForMaxPrizesPt2("./test.txt")
	if err != nil {
		t.Fatalf("Unexpected error")
	}
	want := 480

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
