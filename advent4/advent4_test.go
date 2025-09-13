package advent4

import (
	"testing"
)

func TestReadFileToMatrix(t *testing.T) {
	mtx, err := ReadFileToMatrix("./test.txt")
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Stage 1", func (t *testing.T) {
		got := ProcessBatchStage1(mtx)
		want := 2
	
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}

	})
	t.Run("Stage 2", func (t *testing.T){
		got := ProcessBatchStage2(mtx)
		want := 2
	
		if got != want {
			t.Errorf("stage 2 got %d want %d", got, want)
		}
	})
	
}

