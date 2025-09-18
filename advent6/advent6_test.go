package advent6

import (
	"testing"
)


func TestCountDistinctPositions(t *testing.T) {

	got, err := CountDistinctPositions("./test.txt")
	if err != nil {
		t.Fatalf("Unexpected error")
	}

	want := 41

	if got != want {
		t.Errorf("want: %d, got :%d", want, got)
	}
}

