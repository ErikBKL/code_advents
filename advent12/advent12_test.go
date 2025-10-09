package advent12

import (
	"testing"
)

func Test (t *testing.T) {
	got, _ := TotalFencePrice("./test.txt")
	want := 140

	if got != want {
		t.Errorf("got %d want %d", got ,want)
	}
}