package advent14

import (
	"os"
	"testing"
)

func TestReadFile (t *testing.T) {

	file, err := os.Open("./test.txt")
	if err != nil {
		t.Fatalf("Unexpected error")
	}
	defer file.Close()

	got, _ := ReadFile(file)
	// want := 1
	t.Errorf("%+v", got)

	// if got != want {
	// }
}

func TestSecurityFactor(t *testing.T) {

	got := SecurityFactor("./test.txt")
	want := 12
	
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
