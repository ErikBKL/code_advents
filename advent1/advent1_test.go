package advent1

import (
	"slices"
	"testing"
)

func TestTotalDistance(t *testing.T) {

	got1, got2, err := PrepareData("./test.txt")
	if err != nil {
		t.Fatal(err)
	}

	want1 := []int{11111,11111}
	want2 := []int{22222,22222}

	if slices.Equal(got1, want1) != true {
		t.Errorf("got %v want %v", got1, want1)
	}
	
	if slices.Equal(got2, want2) != true {
		t.Errorf("got %v want %v", got2, want2)
	}
}
