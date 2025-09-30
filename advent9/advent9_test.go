
package advent9

import (
	"testing"
	"slices"
)

func TestDiskChecksum(t *testing.T) {
	got, _ := DiskChecksum("./test.txt")
	want := 1928

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestReadFile(t *testing.T) {
	got,_ := ReadFile("./test.txt")
	want := []rune{'1','2','3','4','5'}

	if ! slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestMapToImg(t *testing.T) {
	diskMap, _ := ReadFile("./test.txt")

	diskImg := MapToImg(diskMap)

	want := []rune{'0','.','.','1','1','1','.','.','.','.','2','2','2','2','2',}

	if ! slices.Equal(diskImg, want) {
		t.Errorf("got %q want %q", diskImg, want)
	}
}

func TestCheckSum(t *testing.T) {
	diskMap, _ := ReadFile("./test.txt")

	diskImg := CompressDiskImg(MapToImg(diskMap))

	got := CheckSum(diskImg)

	want := 192

	if got != want {
		t.Errorf("got %d want %d", got, want)
		t.Errorf("%q", diskImg)
	}
}

func TestStatsFreeChunk( t *testing.T) {
	diskMap, _ := ReadFile("./test.txt")

	diskImg := MapToImg(diskMap)

	end, len := StatsFreeChunk(diskImg, 2)
	wantEnd ,wantLen := 5, 3

	if end != wantEnd {
		t.Errorf ("got %d want %d", end, wantEnd)
	}
	if len != wantLen {
		t.Errorf ("got %d want %d", len, wantLen)
		t.Errorf("%q", diskImg)
	}
}