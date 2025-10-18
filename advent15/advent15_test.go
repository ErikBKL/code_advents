package advent15

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	file, err := os.Open("./test.txt")
	if err != nil {
		t.Fatalf("unexpected error")
	}

	ch, err := ReadFile(file)
	if err != nil {
		t.Fatalf("unexpected error")
	}

	t.Errorf("%q\n, %q", *ch.mtx, ch.moves)
}

func TestSumBoxGps(t *testing.T) {

	got, err := SumBoxGps("./test.txt")
	if err != nil {
		t.Fatalf("unexpected error")
	}
	
	want := 2028

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}