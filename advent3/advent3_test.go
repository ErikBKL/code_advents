package advent3

import (
	"fmt"
	"testing"
)


func TestSumOfAllMul (t *testing.T) {
	got, err := SumOfAllMul("./test.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	want := 10000

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}