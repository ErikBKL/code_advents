package matrix

import (
	"testing"
	"slices"
)

func TestNew(t *testing.T) {
	mtx := New[int]()
	t.Run("TestNew", func (t *testing.T){
		
		if mtx == nil {
			t.Errorf("got %p and want not nil", mtx)
		}
	})

	t.Run("TestResize", func (t *testing.T){
		mtx.Resize(3,3)
		mtxSize := cap(mtx.Data)
		want := 9

		if mtxSize != want {
			t.Errorf("got %d want %d", mtxSize, want)
			t.Errorf("slice: %v", mtx.Data)

		}


	})

	t.Run("TestSize", func (t *testing.T){
		got := mtx.Size()
		want := 9

		if got != want {
			t.Errorf("got %d want %d", got, want)
			t.Errorf("slice: %v", mtx.Data)


		}
	})

	t.Run("TestSetAt", func (t *testing.T){
		ctr := 0 
		for r := 0 ; r < mtx.Rows ; r++ {
			for c := 0 ; c < mtx.Cols ; c++ {
				mtx.Set(r, c, ctr)
				ctr++
			}
		}
		
		want := []int{0,1,2,3,4,5,6,7,8}
		got :=  []int{}
		for r := 0 ; r < mtx.Rows ; r++ {
			for c := 0 ; c < mtx.Cols ; c++ {
				got = append(got, mtx.At(r, c))
			}
		}

		if ! slices.Equal(got, want) {
			t.Errorf("got: %v, want: %v", got, want)
		}
	})
}


func TestIsNextValid(t *testing.T) {
	var tests = []struct{
		point 	Point
		dir		Direction
		want	bool
	}{
		{point: Point{0,0}, dir: LEFT_UP, want: false},
		{point: Point{1,0}, dir: DOWN, want: false},
		{point: Point{2,0}, dir: DOWN, want: false},
		{point: Point{0,1}, dir: LEFT, want: false},
		{point: Point{1,1}, dir: LEFT, want: true},
		{point: Point{2,1}, dir: RIGHT_UP, want: false},
		{point: Point{0,2}, dir: DOWN, want: true},
		{point: Point{1,2}, dir: LEFT_DOWN, want: true},
		{point: Point{2,2}, dir: LEFT_DOWN, want: true},
	}

	// setup
	mtx := New[int]()
	mtx.Resize(3,3)

	for _, testCase := range tests {
		got := mtx.IsNextValid(testCase.dir, testCase.point)

		if got != testCase.want {
			t.Errorf("For point %v got %t want %t", testCase.point, got, testCase.want)
		}
	}
}