package advent2

import (
	"testing"
)

func Test(t *testing.T) {

	got, _ := AmountSafeReports("./test.txt")
	want := 1
	
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
