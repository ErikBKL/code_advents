package advent12

import (
	"testing"
)

func Test (t *testing.T) {
	got, _ := TotalFencePricePt2("./test.txt")
	want := 236

	if got != want {
		t.Errorf("got %d want %d", got ,want)
	}
}