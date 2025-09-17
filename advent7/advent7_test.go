package advent7

import (
	"testing"
)

func TestSplitTargetOperands(t *testing.T) {

	target, operands, err := SplitTargetOperands("127:31 20 13 15")
	if err != nil {
		t.Fatalf("Unexpected error parsing numbers")
	}
	t.Errorf("target: %d, operands: %v", target, operands)
}

func TestSumValidEquations(t *testing.T) {
	got, _ := SumValidEquations("./test.txt")
	want := 3749

	if got != want {
		t.Errorf("got: %d want: %d", got, want)
	}
}