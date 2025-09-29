package advent8

import(
	"testing"
)
func TestUniqueAntinodes(t *testing.T) {
	got, _ := UniqueAntinodes("./test.txt")
	want := 13

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}