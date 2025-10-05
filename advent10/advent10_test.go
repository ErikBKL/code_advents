package advent10

import (
	"testing"
)

func TestTotalTrailheadScore(t *testing.T) {
	got, _ := TotalTrailheadScore("./test.txt")
	want := 36

	if got != want {
		t.Errorf("got %d want %d", got , want)
	}
}