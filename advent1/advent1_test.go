package advent1

import (
	"slices"
	"testing"
)

func TestTotalDistance(t *testing.T) {

	got, err := PrepareData("./test.txt")
	if err != nil {
		t.Fatal(err)
	}

	want := []int{11111,11111}

	if slices.Equal(got, want) != true {
		t.Errorf("got %v want %v", got, want)
	}
}
