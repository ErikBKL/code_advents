package advent2

import (
	"testing"
)

func TestStage1(t *testing.T) {

	got, _ := AmountSafeReports("./sample_part1.txt")
	want := 2

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestStage2(t *testing.T) {

	got, _ := AmountSafeReportsWithDampener("test.txt")
	want := 4

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
