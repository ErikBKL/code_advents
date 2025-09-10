package advent2

import (
	"testing"
)

func TestStage1(t *testing.T) {

	got, _ := AmountSafeReports("./test.txt")
	want := 1
	
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestStage2(t *testing.T) {

	got, _ :=	AmountSafeReports("./stage2test.txt")
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

